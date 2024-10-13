package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// TODO: Add annotations
type Secrets struct{}

var (
	_ infer.CustomResource[SecretsArgs, SecretsState] = (*Secrets)(nil)
	_ infer.CustomDelete[SecretsState]                = (*Secrets)(nil)
)

type SecretsArgs struct {
	App    string            `json:"app" pulumi:"app"`
	Values map[string]string `pulumi:"values" provider:"secret"`
}

type SecretsState struct {
	App        string   `json:"app" pulumi:"app"`
	SecretKeys []string `pulumi:"secretKeys"`
}

func (Secrets) Create(ctx context.Context, name string, input SecretsArgs, preview bool) (string, SecretsState, error) {
	state := SecretsState{
		App:        input.App,
		SecretKeys: make([]string, 0, len(input.Values)),
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

func (Secrets) Delete(ctx context.Context, reqID string, state SecretsState) error {
	cfg := infer.GetConfig[Config](ctx)

	_, err := cfg.flyClient.UnsetSecrets(ctx, state.App, state.SecretKeys)
	return err
}
