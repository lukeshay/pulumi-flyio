package provider

import (
	"math/rand"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "flyio"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Random, RandomArgs, RandomState](),
			infer.Resource[App, AppArgs, AppState](),
			infer.Resource[Machine, MachineArgs, MachineState](),
		},
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
type Random struct{}

// Each resource has an input struct, defining what arguments it accepts.
type RandomArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but it's generally a
	// good idea.
	Length int `pulumi:"length"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type RandomState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	RandomArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minimum.
func (Random) Create(ctx p.Context, name string, input RandomArgs, preview bool) (string, RandomState, error) {
	state := RandomState{RandomArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = makeRandom(input.Length)
	return name, state, nil
}

func makeRandom(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	result := make([]rune, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
