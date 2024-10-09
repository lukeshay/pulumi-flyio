package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	csgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	tsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pygen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

const Name string = "flyio"

func Provider() p.Provider {
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
				"fly",
				"flyio",
				"category/cloud",
				"kind/native",
			},
			License:           "Apache-2.0",
			Repository:        "https://github.com/lukeshay/pulumi-flyio",
			Publisher:         "Luke Shay",
			PluginDownloadURL: "https://api.github.com/lukeshay",
			LanguageMap: map[string]any{
				"nodejs": tsgen.NodePackageInfo{
					PackageName: "pulumi-flyio",
				},
				"go": gogen.GoPackageInfo{
					RespectSchemaVersion:           true,
					GenerateResourceContainerTypes: true,
					PackageImportAliases: map[string]string{
						"github.com/lukeshay/pulumi-flyio/sdk/go/flyio": "flyio",
					},
					ImportBasePath: "github.com/lukeshay/pulumi-flyio/sdk/go/flyio",
				},
				"python": pygen.PackageInfo{
					RespectSchemaVersion: true,
				},
				"csharp": csgen.CSharpPackageInfo{
					RespectSchemaVersion: true,
					PackageReferences: map[string]string{
						"Pulumi": "3.*",
					},
				},
			},
		},
	})
}
