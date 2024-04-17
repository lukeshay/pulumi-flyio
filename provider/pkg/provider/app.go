package provider

import (
	"fmt"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
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
type App struct{}

// verify App complies with resource.CustomCreate interface
var (
	_ infer.CustomResource[AppArgs, AppState] = (*App)(nil)
	_ infer.CustomDelete[AppState]            = (*App)(nil)
	_ infer.CustomRead[AppArgs, AppState]     = (*App)(nil)
	_ infer.CustomUpdate[AppArgs, AppState]   = (*App)(nil)
)

type AppArgs struct {
	flyio.CreateAppRequest
}

type AppState struct {
	flyio.App
	Input AppArgs `pulumi:"input"`
}

func (App) Create(ctx p.Context, name string, input AppArgs, preview bool) (string, AppState, error) {
	state := AppState{Input: input}
	if preview {
		return name, state, nil
	}

	client, err := getFlyClient()
	if err != nil {
		return "", AppState{}, err
	}

	res, err := client.AppsCreate(ctx, flyio.AppsCreateJSONRequestBody{
		AppName:          input.AppName,
		Network:          input.Network,
		EnableSubdomains: input.EnableSubdomains,
		OrgSlug:          input.OrgSlug,
	})
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

	res, err = client.AppsShow(ctx, *input.AppName)
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

func (App) Delete(ctx p.Context, reqID string, state AppState) error {
	client, err := getFlyClient()
	if err != nil {
		return err
	}

	res, err := client.AppsDelete(ctx, *state.Name)
	if err != nil {
		return err
	}

	result, err := flyio.ParseAppsShowResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() != 202 {
		return fmt.Errorf("error deleting app: %s", result.Body)
	}

	return nil
}

func (App) Read(ctx p.Context, id string, inputs AppArgs, state AppState) (
	canonicalID string, normalizedInputs AppArgs, normalizedState AppState, err error,
) {
	client, err := getFlyClient()
	if err != nil {
		return id, inputs, state, err
	}

	res, err := client.AppsShow(ctx, *state.Name)
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

func (a App) Update(ctx p.Context, id string, state AppState, input AppArgs, preview bool) (AppState, error) {
	ctx.LogStatusf(diag.Warning, "Fly Apps cannot be updated. The app you are attempting to update is %s", *state.Name)
	_, _, newState, err := a.Read(ctx, id, input, state)
	return newState, err
}
