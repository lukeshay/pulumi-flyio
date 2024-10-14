package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	csgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	tsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pygen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/superfly/fly-go"
)

var Version string

const Name string = "flyio"

type Config struct {
	FlyApiToken string `pulumi:"token,optional" provider:"secret"`

	flyioClient *flyio.Client
	flyClient   *fly.Client
}

var (
	_ infer.CustomConfigure = (*Config)(nil)
	_ infer.Annotated       = (*Config)(nil)
)

// Annotate provides user-facing descriptions and defaults for Config's fields.
func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.FlyApiToken, "API key for the Fly.io API.")
	a.SetDefault(&c.FlyApiToken, "", "FLY_API_TOKEN", "FLY_TOKEN", "FLY_API_KEY", "FLY_KEY")
}

// Configure validates and processes user-provided configuration values.
func (c *Config) Configure(_ context.Context) error {
	flyioClient, err := flyio.NewClient("https://api.machines.dev/v1")
	if err != nil {
		return err
	}

	flyioClient.RequestEditors = append(flyioClient.RequestEditors, func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.FlyApiToken))

		return nil
	})

	c.flyioClient = flyioClient

	fly.SetBaseURL("https://api.fly.io")
	fly.SetErrorLog(true)
	c.flyClient = fly.NewClientFromOptions(fly.ClientOptions{
		AccessToken: c.FlyApiToken,
		Name:        "pulumi-flyio",
		Version:     Version,
	})

	return nil
}

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Config: infer.Config[*Config](),
		Resources: []infer.InferredResource{
			infer.Resource[Random](),
			infer.Resource[App](),
			infer.Resource[Machine](),
			infer.Resource[Volume](),
			infer.Resource[IP](),
			infer.Resource[Certificate](),
			infer.Resource[Secrets](),
			infer.Resource[WireGuardPeer](),
			infer.Resource[WireGuardToken](),
			infer.Resource[PostgresAttachment](),
		},
		Functions: []infer.InferredFunction{
			infer.Function[GetApp](),
		},
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
			PluginDownloadURL: "github://api.github.com/lukeshay",
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
					RootNamespace:        "PulumiFlyio",
					PackageReferences: map[string]string{
						"Pulumi": "3.*",
					},
				},
			},
		},
	})
}
