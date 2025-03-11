package provider

import (
	"context"
	"math/rand"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Random struct{}

var _ infer.Annotated = (*Random)(nil)

func (r *Random) Annotate(anno infer.Annotator) {
	anno.Describe(&r, "Generates a random string of the specified length.")
}

type RandomArgs struct {
	Length int `pulumi:"length"`
}

var _ infer.Annotated = (*RandomArgs)(nil)

func (a *RandomArgs) Annotate(anno infer.Annotator) {
	anno.Describe(&a.Length, "The length of the random string to generate.")
}

type RandomState struct {
	Result string `pulumi:"result"`
	RandomArgs
}

var _ infer.Annotated = (*RandomState)(nil)

func (s *RandomState) Annotate(anno infer.Annotator) {
	anno.Describe(&s.Result, "The generated random string.")
	anno.Describe(&s.RandomArgs, "The input arguments used to generate the random string.")
}

func (Random) Create(ctx context.Context, name string, input RandomArgs, preview bool) (string, RandomState, error) {
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
