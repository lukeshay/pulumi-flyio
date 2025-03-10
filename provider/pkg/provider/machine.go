package provider

import (
	"context"
	"fmt"
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

type MachineDeploymentStrategy string

var (
	ImmediateMachineDeploymentStrategy MachineDeploymentStrategy = "immediate"
	BlueGreenMachineDeploymentStrategy MachineDeploymentStrategy = "bluegreen"
)

type MachineArgs struct {
	flyio.CreateMachineRequest
	App                string                     `pulumi:"app"`
	DeploymentStrategy *MachineDeploymentStrategy `pulumi:"deploymentStrategy,optional"`
}

type MachineState struct {
	flyio.Machine
	App                string                     `pulumi:"app"`
	DeploymentStrategy *MachineDeploymentStrategy `pulumi:"deploymentStrategy,optional"`
	Input              MachineArgs                `pulumi:"input"`
}

func (m Machine) Create(ctx context.Context, name string, input MachineArgs, preview bool) (string, MachineState, error) {
	input.CreateMachineRequest.Name = nil

	if preview {
		return name, MachineState{
			Input:              input,
			App:                input.App,
			DeploymentStrategy: input.DeploymentStrategy,
		}, nil
	}

	machine, err := createMachine(ctx, input.App, input.CreateMachineRequest)
	if err != nil {
		return "", MachineState{}, err
	}

	state := MachineState{
		Input:              input,
		App:                input.App,
		Machine:            *machine,
		DeploymentStrategy: input.DeploymentStrategy,
	}

	return *state.Name, state, err
}

func (m Machine) Delete(ctx context.Context, _id string, state MachineState) error {
	return m.delete(ctx, state.App, *state.Id)
}

func (m Machine) Read(ctx context.Context, id string, inputs MachineArgs, state MachineState) (
	string, MachineArgs, MachineState, error,
) {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.MachinesShow(ctx, state.App, *state.Id)
	if err != nil {
		return id, inputs, state, err
	}

	result, err := flyio.ParseMachinesShowResponse(res)
	if err != nil {
		return id, inputs, state, err
	}

	if result.JSON200 == nil {
		return id, inputs, state, fmt.Errorf("error showing machine: %s", result.Body)
	}

	state.Machine = *result.JSON200

	return id, inputs, state, nil
}

func (c Machine) Update(ctx context.Context, id string, state MachineState, input MachineArgs, preview bool) (MachineState, error) {
	input.Name = nil

	if preview {
		state.Input = input
		return state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	if input.DeploymentStrategy == nil || *input.DeploymentStrategy == "immediate" {
		res, err := cfg.flyioClient.MachinesUpdate(ctx, state.App, *state.Id, flyio.MachinesUpdateJSONRequestBody{
			Config:                  input.Config,
			SkipLaunch:              input.SkipLaunch,
			Region:                  input.Region,
			Lsvd:                    input.Lsvd,
			SkipServiceRegistration: input.SkipServiceRegistration,
			LeaseTtl:                input.LeaseTtl,
		})
		if err != nil {
			return state, err
		}

		result, err := flyio.ParseMachinesUpdateResponse(res)
		if err != nil {
			return state, err
		}

		if result.JSON200 == nil {
			return state, fmt.Errorf("error updating machine: %s", result.Body)
		}

		state.Machine = *result.JSON200
	} else {
		newMachine, err := createMachine(ctx, input.App, input.CreateMachineRequest)
		if err != nil {
			return state, err
		}

		err = c.waitForState(ctx, state.App, newMachine, flyio.MachinesWaitParamsStateStarted)
		if err != nil {
			force := true
			cfg.flyioClient.MachinesDelete(ctx, state.App, *newMachine.Id, &flyio.MachinesDeleteParams{
				Force: &force,
			})
			return state, err
		}

		err = c.waitForChecks(ctx, state.App, newMachine)
		if err != nil {
			force := true
			cfg.flyioClient.MachinesDelete(ctx, state.App, *newMachine.Id, &flyio.MachinesDeleteParams{
				Force: &force,
			})
			return state, err
		}

		err = c.delete(ctx, state.App, *state.Id)
		if err != nil {
			return state, err
		}

		state.Machine = *newMachine
	}
	state.Input = input

	return state, nil
}

func (Machine) delete(ctx context.Context, app string, id string) error {
	cfg := infer.GetConfig[Config](ctx)

	force := true

	res, err := cfg.flyioClient.MachinesDelete(ctx, app, id, &flyio.MachinesDeleteParams{
		Force: &force,
	})
	if err != nil {
		return err
	}

	result, err := flyio.ParseMachinesDeleteResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() > 299 {
		return fmt.Errorf("error deleting machine: %s", result.Body)
	}

	return nil
}

var machineDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{"App", "Config.AutoDestroy"},
	DeleteBeforeReplaceProps: []string{},
}

func (Machine) Diff(ctx context.Context, id string, state MachineState, input MachineArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, machineDiffOpts)
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

func (Machine) waitForState(ctx context.Context, app string, machine *flyio.Machine, state flyio.MachinesWaitParamsState) error {
	cfg := infer.GetConfig[Config](ctx)
	res, err := cfg.flyioClient.MachinesWait(ctx, app, *machine.Id, &flyio.MachinesWaitParams{
		State:      &state,
		InstanceId: machine.InstanceId,
	})
	if err != nil {
		return err
	}

	result, err := flyio.ParseMachinesWaitResponse(res)
	if err != nil {
		return err
	}

	if result.JSON400 != nil {
		return fmt.Errorf("error waiting for machine: %s", *result.JSON400.Error)
	}

	if result.StatusCode() > 299 {
		return fmt.Errorf("error waiting for machine: %s", result.Body)
	}

	return nil
}

func (Machine) waitForChecks(ctx context.Context, app string, machine *flyio.Machine) error {
	cfg := infer.GetConfig[Config](ctx)

	passing := true
	iterations := 1

	for _, check := range *machine.Checks {
		if *check.Status != "passing" {
			passing = false
			continue
		}
	}

	for !passing {
		iterations += 1

		if iterations >= 20 {
			return fmt.Errorf("timed out waiting for machine checks to pass")
		}

		time.Sleep(5 * time.Second)

		res, err := cfg.flyioClient.MachinesShow(ctx, app, *machine.Id)
		if err != nil {
			continue
		}

		result, err := flyio.ParseMachinesShowResponse(res)
		if err != nil {
			continue
		}

		if result.JSON200 == nil {
			continue
		}

		machine = result.JSON200

		passing = true

		for _, check := range *machine.Checks {
			if *check.Status != "passing" {
				p.GetLogger(ctx).Infof("Check %s is not passing: %s", *check.Name, *check.Status)

				passing = false
				break
			}
		}
	}

	return nil
}

func createMachine(ctx context.Context, appName string, input flyio.CreateMachineRequest) (*flyio.Machine, error) {
	cfg := infer.GetConfig[Config](ctx)
	input.Name = nil

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
