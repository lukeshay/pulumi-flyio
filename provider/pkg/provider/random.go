package provider

import (
	"math/rand"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
)

// TODO: Add annotations
type Random struct{}

type RandomArgs struct {
	Length int `pulumi:"length"`
}

type RandomState struct {
	Result string `pulumi:"result"`
	RandomArgs
}

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
