package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type WireGuardToken struct{}

var _ infer.Annotated = (*WireGuardToken)(nil)

func (w *WireGuardToken) Annotate(anno infer.Annotator) {
	anno.Describe(&w, "A Fly.io WireGuard token for authenticating WireGuard peers.")
}

var (
	_ infer.CustomResource[WireGuardTokenArgs, WireGuardTokenState] = (*WireGuardToken)(nil)
	_ infer.CustomDelete[WireGuardTokenState]                       = (*WireGuardToken)(nil)
)

type WireGuardTokenArgs struct {
	Org  string `pulumi:"org"`
	Name string `pulumi:"name,optional"`
}

var _ infer.Annotated = (*WireGuardTokenArgs)(nil)

func (a *WireGuardTokenArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.Org, "The organization to create the WireGuard token in.")
	anno.Describe(&a.Name, "The name of the WireGuard token.")
}

type WireGuardTokenState struct {
	Org   string `pulumi:"org"`
	Name  string `pulumi:"name,optional"`
	Token string `pulumi:"token" provider:"secret"`
}

var _ infer.Annotated = (*WireGuardTokenState)(nil)

func (s *WireGuardTokenState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.Org, "The organization the WireGuard token belongs to.")
	anno.Describe(&s.Name, "The name of the WireGuard token.")
	anno.Describe(&s.Token, "The WireGuard token value.")
}

func (WireGuardToken) Create(ctx context.Context, name string, input WireGuardTokenArgs, preview bool) (string, WireGuardTokenState, error) {
	state := WireGuardTokenState{
		Name: input.Name,
		Org:  input.Org,
	}

	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	org, err := cfg.flyClient.GetOrganizationBySlug(ctx, input.Org)
	if err != nil {
		return "", WireGuardTokenState{}, err
	}

	peer, err := cfg.flyClient.CreateDelegatedWireGuardToken(ctx, org, state.Name)
	if err != nil {
		return "", WireGuardTokenState{}, err
	}

	state.Token = peer.Token

	return name, state, nil
}

func (WireGuardToken) Delete(ctx context.Context, reqID string, state WireGuardTokenState) error {
	cfg := infer.GetConfig[Config](ctx)

	org, err := cfg.flyClient.GetOrganizationBySlug(ctx, state.Org)
	if err != nil {
		return err
	}

	var name *string

	if state.Name != "" {
		name = &state.Name
	}

	return cfg.flyClient.DeleteDelegatedWireGuardToken(ctx, org, name, &state.Token)
}
