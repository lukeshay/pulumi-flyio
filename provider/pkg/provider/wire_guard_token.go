package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Add annotations
type WireGuardToken struct{}

var (
	_ infer.CustomResource[WireGuardTokenArgs, WireGuardTokenState] = (*WireGuardToken)(nil)
	_ infer.CustomDelete[WireGuardTokenState]                       = (*WireGuardToken)(nil)
)

type WireGuardTokenArgs struct {
	Org  string `pulumi:"org"`
	Name string `pulumi:"name,optional"`
}

type WireGuardTokenState struct {
	Org   string `pulumi:"org"`
	Name  string `pulumi:"name,optional"`
	Token string `pulumi:"token" provider:"secret"`
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
