package provider

import (
	"context"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Add annotations
type IP struct{}

var (
	_ infer.CustomResource[IPArgs, IPState] = (*IP)(nil)
	_ infer.CustomDelete[IPState]           = (*IP)(nil)
)

type IPArgs struct {
	App      string  `json:"app" pulumi:"app"`
	AddrType string  `json:"addrType" pulumi:"addrType"`
	Region   string  `json:"region" pulumi:"region"`
	Network  *string `json:"network,omitempty" pulumi:"network,optional"`
}

type IPStateNetwork struct {
	Name         string `json:"name" pulumi:"name"`
	Organization string `json:"organization" pulumi:"organization"`
}

type IPState struct {
	Input     IPArgs    `pulumi:"input"`
	FlyID     string    `json:"flyId" pulumi:"flyId"`
	Address   string    `json:"address" pulumi:"address"`
	Type      string    `json:"type" pulumi:"type"`
	Region    string    `json:"region" pulumi:"region"`
	CreatedAt time.Time `json:"createdAt" pulumi:"createdAt"`
	App       string    `json:"app" pulumi:"app"`
	Network   *string   `json:"network,omitempty" pulumi:"network,optional"`
}

func (IP) Create(ctx context.Context, name string, input IPArgs, preview bool) (string, IPState, error) {
	if preview {
		state := IPState{
			Input:     input,
			FlyID:     "",
			Address:   "",
			Type:      input.AddrType,
			Region:    input.Region,
			CreatedAt: time.Now(),
			App:       input.App,
		}

		if input.Network != nil {
			state.Network = input.Network
		}

		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	org, err := cfg.flyClient.GetOrganizationByApp(ctx, input.App)
	if err != nil {
		return name, IPState{}, err
	}

	network := ""

	if input.Network != nil {
		network = *input.Network
	}

	ipAddress, err := cfg.flyClient.AllocateIPAddress(ctx, input.App, input.AddrType, input.Region, org, network)
	if err != nil {
		return name, IPState{}, err
	}

	state := IPState{
		Input:     input,
		FlyID:     ipAddress.ID,
		Address:   ipAddress.Address,
		Type:      ipAddress.Type,
		Region:    ipAddress.Region,
		CreatedAt: ipAddress.CreatedAt,
		App:       input.App,
	}

	if ipAddress.Network != nil {
		state.Network = &ipAddress.Network.Name
	}

	return name, state, nil
}

func (IP) Delete(ctx context.Context, reqID string, state IPState) error {
	cfg := infer.GetConfig[Config](ctx)

	p.GetLogger(ctx).Infof("Deleting IP address %s from %s", state.Address, state.App)

	return cfg.flyClient.ReleaseIPAddress(ctx, state.App, state.Address)
}

var ipDiffOpts = generateDiffResponseOpts{
	ReplaceProps:             []string{"App", "AddrType", "Region", "Network"},
	DeleteBeforeReplaceProps: []string{},
}

func (IP) Diff(ctx context.Context, id string, state IPState, input IPArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, ipDiffOpts)
}
