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

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Machine struct{}

// verify Machine complies with resource.CustomCreate interface
var (
	_ infer.CustomResource[MachineArgs, MachineState] = (*Machine)(nil)
	_ infer.CustomDelete[MachineState]                = (*Machine)(nil)
	_ infer.CustomRead[MachineArgs, MachineState]     = (*Machine)(nil)
	_ infer.CustomRead[MachineArgs, MachineState]     = (*Machine)(nil)
)

type MachineArgs struct {
	flyio.MachinesCreateJSONRequestBody
	// WaitForChecks specifies how long to wait in milliseconds for all checks to pass before returning.
	WaitForChecks *int   `pulumi:"waitForChecks,optional"`
	AppName       string `pulumi:"appName"`
	// UpdateStrategy specifies what happens when a machine is updated. Options: "replace", "bluegreen". Default: "replace"
	UpdateStrategy string `pulumi:"updateStrategy,optional"`
}

type MachineState struct {
	flyio.Machine
	AppName string      `pulumi:"appName"`
	Input   MachineArgs `pulumi:"input"`
}

func (m Machine) Create(ctx p.Context, name string, input MachineArgs, preview bool) (string, MachineState, error) {
	state := MachineState{Input: input, AppName: input.AppName}
	if preview {
		return name, state, nil
	}

	client, err := getFlyClient()
	if err != nil {
		return "", MachineState{}, err
	}

	res, err := client.MachinesCreate(ctx, input.AppName, input.MachinesCreateJSONRequestBody)
	if err != nil {
		return "", MachineState{}, err
	}

	result, err := flyio.ParseMachinesCreateResponse(res)
	if err != nil {
		return "", MachineState{}, err
	}

	if result.JSON200 == nil {
		return "", MachineState{}, fmt.Errorf("error creating machine: %s", result.Body)
	}

	machine := result.JSON200

	if input.SkipLaunch != nil && !*input.SkipLaunch {
		machine, err = m.waitForState(ctx, input.AppName, *result.JSON200.Id, flyio.Started)
	}

	state.Machine = *machine

	return *result.JSON200.Id, state, err
}

func (m Machine) Delete(ctx p.Context, reqID string, state MachineState) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	signal := "SIGKILL"
	timeout := "30s"

	res, err := client.MachinesStop(ctx, state.AppName, *state.Id, flyio.StopRequest{
		Signal:  &signal,
		Timeout: &timeout,
	})
	if err != nil {
		return err
	}

	result2, err := flyio.ParseMachinesStopResponse(res)
	if err != nil {
		return err
	}

	if result2.JSON400 != nil {
		return fmt.Errorf("error stopping machine: %s", result2.Body)
	}

	_, err = m.waitForState(ctx, state.AppName, *state.Id, flyio.Stopped)
	if err != nil {
		return err
	}

	res, err = client.MachinesDelete(ctx, state.AppName, *state.Id, &flyio.MachinesDeleteParams{})
	if err != nil {
		return err
	}

	result, err := flyio.ParseMachinesDeleteResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() != 200 {
		return fmt.Errorf("error deleting machine: %s", result.Body)
	}

	_, err = m.waitForState(ctx, state.AppName, *state.Id, flyio.Destroyed)
	if err != nil {
		return err
	}

	return nil
}

func (Machine) Read(ctx p.Context, id string, inputs MachineArgs, state MachineState) (
	canonicalID string, normalizedInputs MachineArgs, normalizedState MachineState, err error,
) {
	client, err := getFlyClient()
	if err != nil {
		return id, inputs, state, err
	}

	res, err := client.MachinesShow(ctx, state.Input.AppName, *state.Id)
	if err != nil {
		return id, inputs, state, err
	}

	result, err := flyio.ParseMachinesShowResponse(res)
	if err != nil {
		return id, inputs, state, err
	}

	if result.JSON200 == nil {
		return id, inputs, state, fmt.Errorf("error showing app: %s", result.Body)
	}

	state.Machine = *result.JSON200

	return *result.JSON200.Id, inputs, state, nil
}

func (m Machine) Update(ctx p.Context, id string, state MachineState, input MachineArgs, preview bool) (MachineState, error) {
	if preview {
		return state, nil
	}

	client, err := getFlyClient()
	if err != nil {
		return state, err
	}

	ttl := 30

	if input.WaitForChecks != nil {
		ttl += *input.WaitForChecks / 1000
	}

	res, err := client.MachinesUpdate(ctx, state.AppName, *state.Id, flyio.UpdateMachineRequest{
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

	if input.SkipLaunch != nil && !*input.SkipLaunch {
		var machine *flyio.Machine
		machine, err = m.waitForState(ctx, input.AppName, *result.JSON200.Id, flyio.Started)
		if err == nil {
			state.Machine = *machine
		}
	}

	return state, err
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

func (m Machine) waitForState(ctx context.Context, appName, machineId string, state flyio.MachinesWaitParamsState) (*flyio.Machine, error) {
	client, err := getFlyClient()
	if err != nil {
		return nil, err
	}

	res, err := client.MachinesWait(ctx, appName, machineId, &flyio.MachinesWaitParams{
		State: &state,
	})
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
