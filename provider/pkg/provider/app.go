package provider

import (
	"context"
	"fmt"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type App struct{}

var _ infer.Annotated = (*App)(nil)

func (a *App) Annotate(anno infer.Annotator) {
	anno.Describe(&a, "A Fly.io application.")
}

var (
	_ infer.CustomResource[AppArgs, AppState] = (*App)(nil)
	_ infer.CustomDelete[AppState]            = (*App)(nil)
	_ infer.CustomRead[AppArgs, AppState]     = (*App)(nil)
)

type AppArgs struct {
	Name             string  `json:"name" pulumi:"name"`
	Org              string  `json:"org" pulumi:"org"`
	EnableSubdomains *bool   `json:"enableSubdomains,omitempty" pulumi:"enableSubdomains,optional"`
	Network          *string `json:"network,omitempty" pulumi:"network,optional"`
}

var _ infer.Annotated = (*AppArgs)(nil)

func (a *AppArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.Name, "The name of the Fly.io application.")
	anno.Describe(&a.Org, "The organization the application belongs to.")
	anno.Describe(&a.EnableSubdomains, "Whether to enable subdomains for the application.")
	anno.Describe(&a.Network, "The network the application belongs to.")
}

type AppState struct {
	Name             string  `json:"name" pulumi:"name"`
	Org              string  `json:"org" pulumi:"org"`
	Status           *string `json:"status,omitempty" pulumi:"status,optional"`
	EnableSubdomains *bool   `json:"enableSubdomains,omitempty" pulumi:"enableSubdomains,optional"`
	Network          *string `json:"network,omitempty" pulumi:"network,optional"`
	Input            AppArgs `pulumi:"input"`
}

var _ infer.Annotated = (*AppState)(nil)

func (s *AppState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.Input, "The input arguments used to create the application.")
	anno.Describe(&s.Name, "The name of the Fly.io application.")
	anno.Describe(&s.Org, "The organization the application belongs to.")
	anno.Describe(&s.Status, "The current status of the application.")
	anno.Describe(&s.EnableSubdomains, "Whether subdomains are enabled for the application.")
	anno.Describe(&s.Network, "The network the application belongs to.")
}

func (c App) Create(ctx context.Context, name string, input AppArgs, preview bool) (string, AppState, error) {
	if preview {
		return name, AppState{
			Input:            input,
			Name:             input.Name,
			Org:              input.Org,
			EnableSubdomains: input.EnableSubdomains,
			Network:          input.Network,
		}, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.AppsCreate(ctx, flyio.CreateAppRequest{
		AppName:          &input.Name,
		OrgSlug:          &input.Org,
		EnableSubdomains: input.EnableSubdomains,
		Network:          input.Network,
	})
	if err != nil {
		return "", AppState{}, err
	}

	createResult, err := flyio.ParseAppsCreateResponse(res)
	if err != nil {
		return "", AppState{}, err
	}

	if createResult.JSON400 != nil {
		return "", AppState{}, fmt.Errorf(*createResult.JSON400.Error)
	}

	if createResult.StatusCode() != 201 {
		return "", AppState{}, fmt.Errorf("error creating app: %s", createResult.Body)
	}

	app, state, err := c.show(ctx, cfg, input)
	if err != nil {
		return "", AppState{}, err
	}

	return *app.Id, state, err
}

func (App) Delete(ctx context.Context, reqID string, state AppState) error {
	cfg := infer.GetConfig[Config](ctx)

	res, err := cfg.flyioClient.AppsDelete(ctx, state.Name)
	if err != nil {
		return err
	}

	result, err := flyio.ParseAppsDeleteResponse(res)
	if err != nil {
		return err
	}

	if result.StatusCode() > 299 {
		return fmt.Errorf("error deleting app: %s", result.Body)
	}

	return nil
}

func (c App) Read(ctx context.Context, id string, inputs AppArgs, state AppState) (
	string, AppArgs, AppState, error,
) {
	cfg := infer.GetConfig[Config](ctx)

	_, newState, err := c.show(ctx, cfg, inputs)
	if err != nil {
		return id, inputs, state, err
	}

	return id, inputs, newState, nil
}

func (App) show(ctx context.Context, cfg Config, inputs AppArgs) (*flyio.App, AppState, error) {
	res, err := cfg.flyioClient.AppsShow(ctx, inputs.Name)
	if err != nil {
		return nil, AppState{}, err
	}

	result, err := flyio.ParseAppsShowResponse(res)
	if err != nil {
		return nil, AppState{}, err
	}

	if result.JSON200 == nil {
		return nil, AppState{}, fmt.Errorf("error showing app: %s", result.Body)
	}

	return result.JSON200, AppState{
		Input:            inputs,
		EnableSubdomains: inputs.EnableSubdomains,
		Name:             *result.JSON200.Name,
		Network:          inputs.Network,
		Org:              *result.JSON200.Organization.Slug,
		Status:           result.JSON200.Status,
	}, nil
}

var appDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{"Org", "Name"},
	DeleteBeforeReplaceProps: []string{},
}

func (App) Diff(ctx context.Context, id string, state AppState, input AppArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, appDiffOpts)
}
