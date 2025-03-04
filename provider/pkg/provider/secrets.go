package provider

import (
	"context"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Add annotations
type Secrets struct{}

var (
	_ infer.CustomResource[SecretsArgs, SecretsState] = (*Secrets)(nil)
	_ infer.CustomUpdate[SecretsArgs, SecretsState]   = (*Secrets)(nil)
	_ infer.CustomDelete[SecretsState]                = (*Secrets)(nil)
	_ infer.CustomDiff[SecretsArgs, SecretsState]     = (*Secrets)(nil)
)

type SecretsArgs struct {
	App    string            `json:"app" pulumi:"app"`
	Values map[string]string `pulumi:"values" provider:"secret"`
}

type SecretsState struct {
	App        string   `json:"app" pulumi:"app"`
	SecretKeys []string `pulumi:"secretKeys"`
	Input      SecretsArgs
}

func (Secrets) Create(ctx context.Context, name string, input SecretsArgs, preview bool) (string, SecretsState, error) {
	state := SecretsState{
		App:        input.App,
		SecretKeys: make([]string, 0, len(input.Values)),
		Input:      input,
	}

	for k := range input.Values {
		state.SecretKeys = append(state.SecretKeys, k)
	}

	if preview {
		return name, state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	_, err := cfg.flyClient.SetSecrets(ctx, input.App, input.Values)
	if err != nil {
		return name, SecretsState{}, err
	}

	return name, state, nil
}

func (Secrets) Update(ctx context.Context, name string, state SecretsState, input SecretsArgs, preview bool) (SecretsState, error) {
	state.SecretKeys = make([]string, 0, len(input.Values))
	state.Input = input

	for k := range input.Values {
		state.SecretKeys = append(state.SecretKeys, k)
	}

	if preview {
		return state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	_, err := cfg.flyClient.SetSecrets(ctx, input.App, input.Values)
	return state, err
}

func (Secrets) Delete(ctx context.Context, reqID string, state SecretsState) error {
	return nil
}

var secretsDiffProps = generateDiffResponseOpts{
	ReplaceProps:             []string{"App"},
	DeleteBeforeReplaceProps: []string{},
}

func (Secrets) Diff(ctx context.Context, id string, state SecretsState, input SecretsArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state.Input, input, secretsDiffProps)
}
