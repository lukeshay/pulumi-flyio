// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flyio

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "flyio:index:App":
		r = &App{}
	case "flyio:index:Certificate":
		r = &Certificate{}
	case "flyio:index:IP":
		r = &IP{}
	case "flyio:index:Machine":
		r = &Machine{}
	case "flyio:index:Random":
		r = &Random{}
	case "flyio:index:Secrets":
		r = &Secrets{}
	case "flyio:index:Volume":
		r = &Volume{}
	case "flyio:index:WireGuardPeer":
		r = &WireGuardPeer{}
	case "flyio:index:WireGuardToken":
		r = &WireGuardToken{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:flyio" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"flyio",
		"index",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"flyio",
		&pkg{version},
	)
}
