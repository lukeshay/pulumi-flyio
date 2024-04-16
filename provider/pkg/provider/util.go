package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
)

func requireEnv(key string) (string, error) {
	value, found := os.LookupEnv(key)
	if found {
		return value, nil
	}

	return "", fmt.Errorf("missing required env var: %s", key)
}

func addFlyApiToken(req *http.Request) error {
	flyApiToken, err := requireEnv("FLY_API_TOKEN")
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", flyApiToken))

	return nil
}

func getFlyClient() (*flyio.Client, error) {
	client, err := flyio.NewClient("https://api.machines.dev/v1")
	if err != nil {
		return nil, err
	}

	client.RequestEditors = append(client.RequestEditors, func(ctx context.Context, req *http.Request) error {
		return addFlyApiToken(req)
	})

	return client, nil
}
