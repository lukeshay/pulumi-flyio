package provider

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

// TODO: Add annotations
type Machine struct{}

var (
	_ infer.CustomResource[MachineArgs, MachineState] = (*Machine)(nil)
	_ infer.CustomDelete[MachineState]                = (*Machine)(nil)
	_ infer.CustomRead[MachineArgs, MachineState]     = (*Machine)(nil)
	_ infer.CustomUpdate[MachineArgs, MachineState]   = (*Machine)(nil)
	_ infer.CustomDiff[MachineArgs, MachineState]     = (*Machine)(nil)
)

type MachineArgs struct {
	flyio.CreateMachineRequest
	WaitForChecks  *int   `pulumi:"waitForChecks,optional"`
	WaitForUpdate  *int   `pulumi:"waitForUpdate,optional"`
	AppName        string `pulumi:"appName"`
	UpdateStrategy string `pulumi:"updateStrategy,optional"`
	SkipLaunch     *bool  `pulumi:"skipLaunch,optional"`
}

type MachineState struct {
	flyio.Machine
	AppName     string      `pulumi:"appName"`
	Input       MachineArgs `pulumi:"input"`
	MachineName string      `pulumi:"machineName"`
}

func (m Machine) Create(ctx p.Context, name string, input MachineArgs, preview bool) (string, MachineState, error) {
	state := MachineState{
		Input:   input,
		AppName: input.AppName,
	}
	if preview {
		return name, state, nil
	}

	if input.Name == nil {
		name := fmt.Sprintf("%s-%s", name, randSeq(6))
		input.Name = &name
	}

	machine, err := createMachine(ctx, input.AppName, input.CreateMachineRequest)
	if err != nil {
		return "", MachineState{}, err
	}

	if input.SkipLaunch == nil || !*input.SkipLaunch {
		machine, err = waitForState(ctx, state.AppName, *machine.Id, *machine.InstanceId, flyio.Started)
		if err == nil {
			state.Machine = *machine
			_, err = m.waitForChecks(ctx, *machine.Id, input, state, input.WaitForChecks)
		}
	}

	state.Machine = *machine
	state.MachineName = *machine.Name

	return state.MachineName, state, err
}

func (m Machine) Delete(ctx p.Context, reqID string, state MachineState) error {
	return destroyMachine(ctx, state, &state.Machine)
}

func (Machine) Read(ctx p.Context, id string, inputs MachineArgs, state MachineState) (
	canonicalID string, normalizedInputs MachineArgs, normalizedState MachineState, err error,
) {
	client, err := getFlyClient()
	if err != nil {
		return id, inputs, state, err
	}

	res, err := client.MachinesList(ctx, state.Input.AppName, &flyio.MachinesListParams{
		Region: state.Region,
	})
	if err != nil {
		return id, inputs, state, err
	}

	result, err := flyio.ParseMachinesListResponse(res)
	if err != nil {
		return id, inputs, state, err
	}

	if result.JSON200 == nil {
		return id, inputs, state, fmt.Errorf("error showing app: %s", result.Body)
	}

	for _, machine := range *result.JSON200 {
		if *machine.Name == state.MachineName {
			state.Machine = machine
		}
	}

	state.MachineName = *state.Name

	return id, inputs, state, nil
}

func (m Machine) Update(ctx p.Context, id string, state MachineState, input MachineArgs, preview bool) (MachineState, error) {
	if preview {
		state.Input = input
		return state, nil
	}

	client, err := getFlyClient()
	if err != nil {
		return state, err
	}

	ttl := 30

	diff, _ := m.Diff(ctx, id, state, input)

	if input.UpdateStrategy == "bluegreen" || diff.DeleteBeforeReplace {
		ctx.Log(diag.Info, "Redeploying machine using the bluegreen strategy")

		// 1. Create a new machine with SkipServiceRegistration: true
		skipServiceRegistration := true
		request := input.CreateMachineRequest
		request.SkipServiceRegistration = &skipServiceRegistration

		if *input.Name == *state.Input.Name {
			name := fmt.Sprintf("%s-%s", *state.Input.Name, time.Now().Format("20060102150405"))
			request.Name = &name
		}

		newMachine, err := createMachine(ctx, state.AppName, request)
		if err != nil {
			return state, err
		}

		if input.SkipLaunch == nil || !*input.SkipLaunch {
			// 2. Wait for checks on new machine
			ctx.Log(diag.Info, "Waiting for new machine to start")
			newMachine, err = waitForState(ctx, state.AppName, *newMachine.Id, *newMachine.InstanceId, flyio.Started)
			if err != nil {
				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			ctx.Log(diag.Info, "Waiting for checks on new machine")
			_, err := m.waitForChecks(ctx, *newMachine.Id, input, state, input.WaitForChecks)
			if err != nil {
				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			// 3. Uncordon new machine
			ctx.Log(diag.Info, "Uncordoning new machine")
			err = uncordonMachine(ctx, state.AppName, *newMachine.Id)
			if err != nil {
				destroyOldMachine(ctx, state, newMachine)

				return state, err
			}

			// 4. Cordorn old machine
			ctx.Log(diag.Info, "Cordoning old machine")
			err = cordonMachine(ctx, state.AppName, *state.Id)
			if err != nil {
				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			if input.WaitForUpdate != nil {
				// 5. Pause for WaitForUpdate
				ctx.Log(diag.Info, "Waiting for update")
				time.Sleep(time.Duration(*input.WaitForUpdate) * time.Millisecond)
			}

			// 6. Wait for checks on new machine
			//    - If checks fail, uncordon old machine, destroy new machine, and return error
			ctx.Log(diag.Info, "Waiting for checks on new machine")
			newState, err := m.waitForChecks(ctx, *newMachine.Id, input, state, input.WaitForChecks)
			if err != nil {
				ctx.Log(diag.Info, "Checks failed on new machine, uncordoning old machine and destroying new machine")
				err2 := uncordonMachine(ctx, state.AppName, *state.Id)
				if err2 != nil {
					return state, err2
				}

				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			ctx.Log(diag.Info, "Checks passed on new machine, destroying old machine")
			destroyOldMachine(ctx, state, newMachine)

			newState.MachineName = *request.Name
			newState.Input = input

			return newState, nil
		}

		// 7. Destroy old machine
		err = destroyOldMachine(ctx, state, newMachine)
		if err != nil {
			return MachineState{}, err
		}

		state.Machine = *newMachine
		state.Input = input

		return state, nil
	}

	ctx.Log(diag.Info, "Redeploying machine using the replacement strategy")

	if input.WaitForChecks != nil {
		ttl += *input.WaitForChecks / 1000
	}

	res, err := client.MachinesList(ctx, state.AppName, &flyio.MachinesListParams{
		Region: state.Region,
	})
	if err != nil {
		return state, err
	}

	machinesList, err := flyio.ParseMachinesListResponse(res)
	if err != nil {
		return state, err
	}

	var machine flyio.Machine
	for _, m := range *machinesList.JSON200 {
		if *m.Name == state.MachineName {
			machine = m
		}
	}

	res, err = client.MachinesUpdate(ctx, state.AppName, *machine.Id, flyio.UpdateMachineRequest{
		Config:                  input.Config,
		LeaseTtl:                input.LeaseTtl,
		Lsvd:                    input.Lsvd,
		Name:                    state.Name,
		Region:                  input.Region,
		SkipLaunch:              input.SkipLaunch,
		SkipServiceRegistration: input.SkipServiceRegistration,
	})
	if err != nil {
		return state, err
	}

	result, err := flyio.ParseMachinesUpdateResponse(res)
	if err != nil {
		return state, err
	}

	if result.JSON400 != nil {
		return state, fmt.Errorf("error updating machine: %s", *result.JSON400.Error)
	}
	if result.JSON200 == nil {
		return state, fmt.Errorf("error updating machine: %s", result.Body)
	}

	state.Machine = *result.JSON200
	state.Input = input

	if input.SkipLaunch == nil || !*input.SkipLaunch {
		var machine *flyio.Machine
		machine, err = waitForState(ctx, state.AppName, *result.JSON200.Id, *result.JSON200.InstanceId, flyio.Started)
		if err == nil {
			state.Machine = *machine

			_, err = m.waitForChecks(ctx, *result.JSON200.Id, input, state, input.WaitForChecks)
		}
	}

	return state, err
}

var machineDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{},
	DeleteBeforeReplaceProps: []string{"Region", "Name", "AppName"},
}

func (Machine) Diff(ctx p.Context, id string, state MachineState, input MachineArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, machineDiffOpts)
}

func (m Machine) waitForChecks(ctx p.Context, id string, input MachineArgs, state MachineState, waitTime *int) (
	MachineState, error,
) {
	if input.WaitForChecks != nil {
		waitUntil := time.Now().Add(time.Duration(*waitTime) * time.Millisecond)
		done := []string{}
		var err error

		for time.Now().Before(waitUntil) && len(done) < len(*state.Checks) {
			time.Sleep(1 * time.Second)
			_, _, state, err = m.Read(ctx, id, input, state)
			if err != nil {
				continue
			}

			for _, check := range *state.Checks {
				if *check.Status == "passing" && !slices.Contains(done, *check.Name) {
					done = append(done, *check.Name)
				}
			}
		}

		if len(done) != len(*state.Checks) {
			result := map[string]string{}

			for _, check := range *state.Checks {
				result[*check.Name] = *check.Status
				_, _, state, err = m.Read(ctx, id, input, state)
			}

			return state, fmt.Errorf("not all checks succeeded: %#v", result)
		}
	}

	return state, nil
}

func (m Machine) waitForStopped(ctx p.Context, id string, input MachineArgs, state MachineState, waitTime int) (
	MachineState, error,
) {
	waitUntil := time.Now().Add(time.Duration(waitTime) * time.Millisecond)

	for time.Now().Before(waitUntil) {
		time.Sleep(1 * time.Second)
		_, _, state, err := m.Read(ctx, id, input, state)
		if err != nil {
			continue
		}

		if *state.State == "stopped" {
			return state, nil
		}
	}

	return state, fmt.Errorf("machine is not stopped")
}

type machineState string

var (
	created    machineState = "created"
	starting   machineState = "starting"
	started    machineState = "started"
	stopping   machineState = "stopping"
	stopped    machineState = "stopped"
	replacing  machineState = "replacing"
	destroying machineState = "destroying"
	destroyed  machineState = "destroyed"
)

func waitForState(ctx context.Context, appName, machineId, instanceId string, state flyio.MachinesWaitParamsState) (*flyio.Machine, error) {
	client, err := getFlyClient()
	if err != nil {
		return nil, err
	}

	params := &flyio.MachinesWaitParams{
		State: &state,
	}

	if instanceId != "" {
		params.InstanceId = &instanceId
	}

	res, err := client.MachinesWait(ctx, appName, machineId, params)
	if err != nil {
		return nil, err
	}

	result, err := flyio.ParseMachinesWaitResponse(res)
	if err != nil {
		return nil, err
	}

	if result.JSON400 != nil {
		return nil, fmt.Errorf("error waiting for machine state: %s", *result.JSON400.Error)
	}

	res, err = client.MachinesShow(ctx, appName, machineId)
	if err != nil {
		return nil, err
	}

	result2, err := flyio.ParseMachinesShowResponse(res)
	if err != nil {
		return nil, err
	}

	if result2.JSON200 == nil {
		return nil, fmt.Errorf("error showing app: %s", result.Body)
	}

	return result2.JSON200, nil
}

func createMachine(ctx context.Context, appName string, input flyio.CreateMachineRequest) (*flyio.Machine, error) {
	client, err := getFlyClient()
	if err != nil {
		return nil, err
	}

	res, err := client.MachinesCreate(ctx, appName, input)
	if err != nil {
		return nil, err
	}

	result, err := flyio.ParseMachinesCreateResponse(res)
	if err != nil {
		return nil, err
	}

	if result.JSON200 == nil {
		return nil, fmt.Errorf("error creating machine: %s", result.Body)
	}

	return result.JSON200, err
}

func destroyOldMachine(ctx p.Context, state MachineState, newMachin *flyio.Machine) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	res, err := client.MachinesList(ctx, state.AppName, &flyio.MachinesListParams{
		Region: state.Region,
	})

	ctx.Logf(diag.Info, "Listing machines: %v", err)
	if err == nil {
		result, err := flyio.ParseMachinesListResponse(res)
		ctx.Logf(diag.Info, "Listing machines: %v", err)
		if err != nil {
			return err
		}

		var machine flyio.Machine

		for _, m := range *result.JSON200 {
			if *m.Name == state.MachineName && (newMachin != nil && *m.InstanceId != *newMachin.InstanceId) {
				machine = m
			}
		}

		ctx.Logf(diag.Info, "Destroying machine: %s", *machine.Id)

		return destroyMachine(ctx, state, &machine)
	}

	return nil
}

func destroyMachine(ctx p.Context, state MachineState, machine *flyio.Machine) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	signal := "SIGKILL"
	timeout := "30s"

	res, err := client.MachinesStop(ctx, state.AppName, *machine.Id, flyio.StopRequest{
		Signal:  &signal,
		Timeout: &timeout,
	})
	if err != nil {
		ctx.Logf(diag.Info, "Error stopping machine: %v", err)
		return err
	}

	result2, err := flyio.ParseMachinesStopResponse(res)
	if err != nil {
		return err
	}

	if result2.JSON400 != nil {
		return fmt.Errorf("error stopping machine: %s", result2.Body)
	}

	_, err = waitForState(ctx, state.AppName, *machine.Id, *machine.InstanceId, flyio.Stopped)
	if err != nil {
		return err
	}

	force := true

	res, err = client.MachinesDelete(ctx, state.AppName, *machine.Id, &flyio.MachinesDeleteParams{
		Force: &force,
	})
	if err != nil {
		return err
	}

	delRes, err := flyio.ParseMachinesDeleteResponse(res)
	if err != nil {
		return err
	}

	if delRes.StatusCode() != 200 {
		result, _ := flyio.ParseMachinesDeleteResponse(res)
		return fmt.Errorf("error deleting machine: %s", result.Body)
	}

	_, err = waitForState(ctx, state.AppName, *machine.Id, *machine.InstanceId, flyio.Destroyed)
	if err != nil {
		return err
	}

	return nil
}

func cordonMachine(ctx context.Context, appName, id string) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	res, err := client.MachinesCordon(ctx, appName, id)
	if err != nil {
		return err
	}

	result, err := flyio.ParseMachinesCordonResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() != 200 {
		return fmt.Errorf("error cordoning machine: %s", result.Body)
	}

	return nil
}

func uncordonMachine(ctx context.Context, appName, id string) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	res, err := client.MachinesUncordon(ctx, appName, id)
	if err != nil {
		return err
	}

	result, err := flyio.ParseMachinesUncordonResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() != 200 {
		return fmt.Errorf("error uncordoning machine: %s", result.Body)
	}

	return nil
}
