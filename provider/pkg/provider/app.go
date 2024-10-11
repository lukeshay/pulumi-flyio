package provider

import (
	"context"
	"fmt"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Add annotations
type App struct{}

var (
	_ infer.CustomResource[AppArgs, AppState] = (*App)(nil)
	_ infer.CustomDelete[AppState]            = (*App)(nil)
	_ infer.CustomRead[AppArgs, AppState]     = (*App)(nil)
)

type AppArgs struct {
	flyio.CreateAppRequest
}

type AppState struct {
	flyio.App
	Input AppArgs `pulumi:"input"`
}

func (App) Create(ctx context.Context, name string, input AppArgs, preview bool) (string, AppState, error) {
	state := AppState{Input: input}
	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.AppsCreate(ctx, input.CreateAppRequest)
	if err != nil {
		return "", AppState{}, err
	}

	result, err := flyio.ParseAppsCreateResponse(res)
	if err != nil {
		return "", AppState{}, err
	}

	if result.JSON400 != nil {
		return "", AppState{}, fmt.Errorf(*result.JSON400.Error)
	}

	if result.StatusCode() != 201 {
		return "", AppState{}, fmt.Errorf("error creating app: %s", result.Body)
	}

	res, err = cfg.flyioClient.AppsShow(ctx, *input.AppName)
	if err != nil {
		return "", AppState{}, err
	}

	result2, err := flyio.ParseAppsShowResponse(res)
	if err != nil {
		return "", AppState{}, err
	}

	if result2.JSON200 == nil {
		return "", AppState{}, fmt.Errorf("error showing app: %s", result.Body)
	}

	state.App = *result2.JSON200

	return *result2.JSON200.Id, state, nil
}

func (App) Delete(ctx context.Context, reqID string, state AppState) error {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.AppsDelete(ctx, *state.Name)
	if err != nil {
		return err
	}

	result, err := flyio.ParseAppsDeleteResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() != 202 {
		return fmt.Errorf("error deleting app: %s", result.Body)
	}

	return nil
}

func (App) Read(ctx context.Context, id string, inputs AppArgs, state AppState) (
	canonicalID string, normalizedInputs AppArgs, normalizedState AppState, err error,
) {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.AppsShow(ctx, *state.Name)
	if err != nil {
		return id, inputs, state, err
	}

	result, err := flyio.ParseAppsShowResponse(res)
	if err != nil {
		return id, inputs, state, err
	}

	if result.JSON200 == nil {
		return id, inputs, state, fmt.Errorf("error showing app: %s", result.Body)
	}

	state.App = *result.JSON200

	return id, inputs, state, nil
}

var appDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{},
	DeleteBeforeReplaceProps: []string{"AppName", "OrgSlug", "EnableSubdomains", "Network"},
}

func (App) Diff(ctx context.Context, id string, state AppState, input AppArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, appDiffOpts)
}
