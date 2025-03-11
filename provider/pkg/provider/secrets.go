package provider

import (
	"context"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Secrets struct{}

var _ infer.Annotated = (*Secrets)(nil)

func (s *Secrets) Annotate(anno infer.Annotator) {
	anno.Describe(&s, "Manages application secrets for a Fly.io application.")
}

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

var _ infer.Annotated = (*SecretsArgs)(nil)

func (a *SecretsArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.App, "The Fly.io application to set secrets for.")
	anno.Describe(&a.Values, "The secret values to set, as key-value pairs.")
}

type SecretsState struct {
	App    string            `json:"app" pulumi:"app"`
	Values map[string]string `pulumi:"values" provider:"secret"`
}

var _ infer.Annotated = (*SecretsState)(nil)

func (s *SecretsState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.App, "The Fly.io application the secrets are set for.")
	anno.Describe(&s.Values, "The secret values, as key-value pairs.")
}

func (Secrets) Create(ctx context.Context, name string, input SecretsArgs, preview bool) (string, SecretsState, error) {
	state := SecretsState(input)

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
	state.Values = input.Values

	if preview {
		return state, nil
	}

	cfg := infer.GetConfig[Config](ctx)

	unsetKeys := []string{}
	for k := range state.Values {
		if _, ok := input.Values[k]; !ok {
			unsetKeys = append(unsetKeys, k)
		}
	}

	_, err := cfg.flyClient.SetSecrets(ctx, input.App, input.Values)
	if err != nil {
		return state, err
	}

	_, err = cfg.flyClient.UnsetSecrets(ctx, input.App, unsetKeys)
	return state, err
}

func (Secrets) Delete(ctx context.Context, reqID string, state SecretsState) error {
	cfg := infer.GetConfig[Config](ctx)

	secretKeys := make([]string, 0, len(state.Values))
	for k := range state.Values {
		secretKeys = append(secretKeys, k)
	}

	_, err := cfg.flyClient.UnsetSecrets(ctx, state.App, secretKeys)
	return err
}

var secretsDiffProps = generateDiffResponseOpts{
	ReplaceProps:             []string{"App"},
	DeleteBeforeReplaceProps: []string{},
}

func (Secrets) Diff(ctx context.Context, id string, state SecretsState, input SecretsArgs) (p.DiffResponse, error) {
	return generateDiffResponse(state, input, secretsDiffProps)
}
