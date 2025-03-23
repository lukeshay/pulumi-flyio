package provider

import (
	"context"
	"fmt"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type IP struct{}

var _ infer.Annotated = (*IP)(nil)

func (i *IP) Annotate(anno infer.Annotator) {
	anno.Describe(&i, "A Fly.io IP address allocation for an application.")
}

var (
	_ infer.CustomResource[IPArgs, IPState] = (*IP)(nil)
	_ infer.CustomDelete[IPState]           = (*IP)(nil)
)

type IPArgs struct {
	App      string  `json:"app" pulumi:"app"`
	AddrType string  `json:"addrType" pulumi:"addrType"`
	Region   *string `json:"region,omitempty" pulumi:"region,optional"`
	Network  *string `json:"network,omitempty" pulumi:"network,optional"`
}

var _ infer.Annotated = (*IPArgs)(nil)

func (a *IPArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.App, "The name of the Fly.io application to allocate the IP address for.")
	anno.Describe(&a.AddrType, "The type of IP address (v4, v6, shared_v4, or private_v6).")
	anno.Describe(&a.Region, "The region to allocate the IP address in. This is required for non-shared IP addresses.")
	anno.Describe(&a.Network, "The network to allocate the IP address in.")
}

type IPStateNetwork struct {
	Name         string `json:"name" pulumi:"name"`
	Organization string `json:"organization" pulumi:"organization"`
}

var _ infer.Annotated = (*IPStateNetwork)(nil)

func (n *IPStateNetwork) Annotate(anno infer.Annotator) {
	anno.Describe(&n.Name, "The name of the network.")
	anno.Describe(&n.Organization, "The organization the network belongs to.")
}

type IPState struct {
	Input     IPArgs    `pulumi:"input"`
	FlyID     *string   `json:"flyId,omitempty" pulumi:"flyId,optional"`
	Address   string    `json:"address" pulumi:"address"`
	Type      string    `json:"type" pulumi:"type"`
	Region    *string   `json:"region,omitempty" pulumi:"region,optional"`
	CreatedAt time.Time `json:"createdAt" pulumi:"createdAt"`
	App       string    `json:"app" pulumi:"app"`
	Network   *string   `json:"network,omitempty" pulumi:"network,optional"`
}

var _ infer.Annotated = (*IPState)(nil)

func (s *IPState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.Input, "The input arguments used to allocate the IP address.")
	anno.Describe(&s.FlyID, "The Fly.io IP address ID.")
	anno.Describe(&s.Address, "The allocated IP address.")
	anno.Describe(&s.Type, "The type of IP address (v4 or v6).")
	anno.Describe(&s.Region, "The region the IP address is allocated in.")
	anno.Describe(&s.CreatedAt, "When the IP address was allocated.")
	anno.Describe(&s.App, "The application the IP address is allocated for.")
	anno.Describe(&s.Network, "The network the IP address belongs to.")
}

func (IP) Create(ctx context.Context, name string, input IPArgs, preview bool) (string, IPState, error) {
	if preview {
		state := IPState{
			Input:     input,
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

	if input.AddrType == "shared_v4" {
		sharedIP, err := cfg.flyClient.AllocateSharedIPAddress(ctx, input.App)
		if err != nil {
			return name, IPState{}, err
		}

		state := IPState{
			Input:     input,
			Address:   sharedIP.String(),
			Type:      input.AddrType,
			CreatedAt: time.Now(),
			App:       input.App,
		}

		return name, state, nil
	}

	if input.Region == nil {
		return name, IPState{}, fmt.Errorf("region is required for non-shared IP addresses")
	}

	ipAddress, err := cfg.flyClient.AllocateIPAddress(ctx, input.App, input.AddrType, *input.Region, org, network)
	if err != nil {
		return name, IPState{}, err
	}

	state := IPState{
		Input:     input,
		FlyID:     &ipAddress.ID,
		Address:   ipAddress.Address,
		Type:      ipAddress.Type,
		Region:    &ipAddress.Region,
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
