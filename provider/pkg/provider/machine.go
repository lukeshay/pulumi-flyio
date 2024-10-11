package provider

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
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
	SkipLaunch     *bool  `pulumi:"skipLaunch,optional"`
	AppName        string `pulumi:"appName"`
	UpdateStrategy string `pulumi:"updateStrategy,optional"`
}

type MachineState struct {
	flyio.Machine
	AppName     string      `pulumi:"appName"`
	Input       MachineArgs `pulumi:"input"`
	MachineName string      `pulumi:"machineName"`
}

func (m Machine) Create(ctx context.Context, name string, input MachineArgs, preview bool) (string, MachineState, error) {
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

func (m Machine) Delete(ctx context.Context, reqID string, state MachineState) error {
	cfg := infer.GetConfig[Config](ctx)

	machine, err := m.listMachinesWithName(ctx, cfg, state)
	if err != nil {
		return err
	}

	return destroyMachine(ctx, state, &machine)
}

func (Machine) listMachinesWithName(ctx context.Context, cfg Config, state MachineState) (flyio.Machine, error) {
	res, err := cfg.flyioClient.MachinesList(ctx, state.Input.AppName, &flyio.MachinesListParams{
		Region: state.Region,
	})
	if err != nil {
		return flyio.Machine{}, err
	}

	result, err := flyio.ParseMachinesListResponse(res)
	if err != nil {
		return flyio.Machine{}, err
	}

	if result.JSON200 == nil {
		return flyio.Machine{}, fmt.Errorf("error showing app: %s", result.Body)
	}

	var machine flyio.Machine

	for _, m := range *result.JSON200 {
		if *m.Name == state.MachineName {
			machine = m
		}
	}

	return machine, nil
}

func (m Machine) Read(ctx context.Context, id string, inputs MachineArgs, state MachineState) (
	string, MachineArgs, MachineState, error,
) {
	cfg := infer.GetConfig[Config](ctx)

	machine, err := m.listMachinesWithName(ctx, cfg, state)
	if err != nil {
		return id, inputs, state, err
	}

	state.Machine = machine
	state.MachineName = *state.Name

	return id, inputs, state, nil
}

func (m Machine) Update(ctx context.Context, id string, state MachineState, input MachineArgs, preview bool) (MachineState, error) {
	if preview {
		state.Input = input
		return state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	ttl := 30

	diff, _ := m.Diff(ctx, id, state, input)

	if input.UpdateStrategy == "bluegreen" || diff.DeleteBeforeReplace {
		p.GetLogger(ctx).Info("Redeploying machine using the bluegreen strategy")

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
			p.GetLogger(ctx).Info("Waiting for new machine to start")
			newMachine, err = waitForState(ctx, state.AppName, *newMachine.Id, *newMachine.InstanceId, flyio.Started)
			if err != nil {
				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			p.GetLogger(ctx).Info("Waiting for checks on new machine")
			_, err := m.waitForChecks(ctx, *newMachine.Id, input, state, input.WaitForChecks)
			if err != nil {
				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			// 3. Uncordon new machine
			p.GetLogger(ctx).Info("Uncordoning new machine")
			err = uncordonMachine(ctx, state.AppName, *newMachine.Id)
			if err != nil {
				destroyOldMachine(ctx, state, newMachine)

				return state, err
			}

			// 4. Cordorn old machine
			p.GetLogger(ctx).Info("Cordoning old machine")
			err = cordonMachine(ctx, state.AppName, *state.Id)
			if err != nil {
				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			if input.WaitForUpdate != nil {
				// 5. Pause for WaitForUpdate
				p.GetLogger(ctx).Info("Waiting for update")
				time.Sleep(time.Duration(*input.WaitForUpdate) * time.Millisecond)
			}

			// 6. Wait for checks on new machine
			//    - If checks fail, uncordon old machine, destroy new machine, and return error
			p.GetLogger(ctx).Info("Waiting for checks on new machine")
			newState, err := m.waitForChecks(ctx, *newMachine.Id, input, state, input.WaitForChecks)
			if err != nil {
				p.GetLogger(ctx).Info("Checks failed on new machine, uncordoning old machine and destroying new machine")
				err2 := uncordonMachine(ctx, state.AppName, *state.Id)
				if err2 != nil {
					return state, err2
				}

				destroyMachine(ctx, state, newMachine)

				return state, err
			}

			p.GetLogger(ctx).Info("Checks passed on new machine, destroying old machine")
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

	p.GetLogger(ctx).Info("Redeploying machine using the replacement strategy")

	if input.WaitForChecks != nil {
		ttl += *input.WaitForChecks / 1000
	}

	machine, err := m.listMachinesWithName(ctx, cfg, state)
	if err != nil {
		return state, err
	}

	res, err := cfg.flyioClient.MachinesUpdate(ctx, state.AppName, *machine.Id, flyio.UpdateMachineRequest{
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

func (Machine) Diff(ctx context.Context, id string, state MachineState, input MachineArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, machineDiffOpts)
}

func (m Machine) waitForChecks(ctx context.Context, id string, input MachineArgs, state MachineState, waitTime *int) (
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
				_, _, state, _ = m.Read(ctx, id, input, state)
			}

			return state, fmt.Errorf("not all checks succeeded: %#v", result)
		}
	}

	return state, nil
}

func (m Machine) waitForStopped(ctx context.Context, id string, input MachineArgs, state MachineState, waitTime int) (
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
	cfg := infer.GetConfig[Config](ctx)

	params := &flyio.MachinesWaitParams{
		State: &state,
	}

	if instanceId != "" {
		params.InstanceId = &instanceId
	}

	res, err := cfg.flyioClient.MachinesWait(ctx, appName, machineId, params)
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

	res, err = cfg.flyioClient.MachinesShow(ctx, appName, machineId)
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
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.MachinesCreate(ctx, appName, input)
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

func destroyOldMachine(ctx context.Context, state MachineState, newMachin *flyio.Machine) error {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.MachinesList(ctx, state.AppName, &flyio.MachinesListParams{
		Region: state.Region,
	})

	p.GetLogger(ctx).Infof("Listing machines: %v", err)
	if err == nil {
		result, err := flyio.ParseMachinesListResponse(res)
		p.GetLogger(ctx).Infof("Listing machines: %v", err)
		if err != nil {
			return err
		}

		var machine flyio.Machine

		for _, m := range *result.JSON200 {
			if *m.Name == state.MachineName && (newMachin != nil && *m.InstanceId != *newMachin.InstanceId) {
				machine = m
			}
		}

		p.GetLogger(ctx).Infof("Destroying machine: %s", *machine.Id)

		return destroyMachine(ctx, state, &machine)
	}

	return nil
}

func destroyMachine(ctx context.Context, state MachineState, machine *flyio.Machine) error {
	cfg := infer.GetConfig[Config](ctx)

	signal := "SIGKILL"
	timeout := "30s"

	res, err := cfg.flyioClient.MachinesStop(ctx, state.AppName, *machine.Id, flyio.StopRequest{
		Signal:  &signal,
		Timeout: &timeout,
	})
	if err != nil {
		p.GetLogger(ctx).Infof("Error stopping machine: %v", err)
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

	res, err = cfg.flyioClient.MachinesDelete(ctx, state.AppName, *machine.Id, &flyio.MachinesDeleteParams{
		Force: &force,
	})
	if err != nil {
		return err
	}

	_, err = flyio.ParseMachinesDeleteResponse(res)
	if err != nil {
		return err
	}

	// if delRes.StatusCode() != 200 {
	// 	result, _ := flyio.ParseMachinesDeleteResponse(res)
	// 	return fmt.Errorf("error deleting machine: %s", result.Body)
	// }

	_, err = waitForState(ctx, state.AppName, *machine.Id, *machine.InstanceId, flyio.Destroyed)
	if err != nil {
		return err
	}

	return nil
}

func cordonMachine(ctx context.Context, appName, id string) error {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.MachinesCordon(ctx, appName, id)
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
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.MachinesUncordon(ctx, appName, id)
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
