package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

const Name string = "flyio"

func Provider() p.Provider {
	// TODO: Implement Volumne resource
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Random, RandomArgs, RandomState](),
			infer.Resource[App, AppArgs, AppState](),
			infer.Resource[Machine, MachineArgs, MachineState](),
			infer.Resource[Volume, VolumeArgs, VolumeState](),
		},
		Functions: []infer.InferredFunction{},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Metadata: schema.Metadata{
			DisplayName: "Flyio",
			Description: "A native Pulumi provider for Fly.io.",
			Keywords: []string{
				"pulumi",
				"flyio",
				"category/cloud",
				"kind/native",
			},
			License:    "Apache-2.0",
			Repository: "https://github.com/lukeshay/pulumi-flyio",
			Publisher:  "Luke Shay",
			LanguageMap: map[string]any{
				"nodejs": map[string]any{
					"name": "pulumi-flyio",
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/lukeshay/pulumi-flyio/sdk/go/flyio",
				},
			},
		},
	})
}
