package provider

import (
	"context"
	"fmt"

	"github.com/lukeshay/pulumi-flyio/provider/pkg/flyio"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type GetApp struct{}

func (c *GetApp) Annotate(a infer.Annotator) {
	a.Describe(c,
		"Gets the state of a Fly.io app.",
	)
}

type GetAppArgs struct {
	Name string `json:"name" pulumi:"name"`
}

func (c *GetAppArgs) Annotate(a infer.Annotator) {
	a.Describe(&c.Name, "The name of the app.")
}

type GetAppState struct {
	Id     string `pulumi:"flyId"`
	Name   string `pulumi:"name"`
	Org    string `pulumi:"org"`
	Status string `pulumi:"status"`
}

func (GetApp) Call(ctx context.Context, args GetAppArgs) (GetAppState, error) {
	cfg := infer.GetConfig[Config](ctx)

	appShowRes, err := cfg.flyioClient.AppsShow(ctx, args.Name)
	if err != nil {
		return GetAppState{}, err
	}

	appShowResult, err := flyio.ParseAppsShowResponse(appShowRes)
	if err != nil {
		return GetAppState{}, err
	}

	if appShowResult.JSON200 == nil {
		return GetAppState{}, fmt.Errorf("unexpected response from API: %s", appShowResult.Body)
	}

	if appShowRes.StatusCode > 299 {
		return GetAppState{}, fmt.Errorf("unexpected status code from API: %s", appShowRes.Body)
	}

	return GetAppState{
		Id:     *appShowResult.JSON200.Id,
		Name:   *appShowResult.JSON200.Name,
		Org:    *appShowResult.JSON200.Organization.Slug,
		Status: *appShowResult.JSON200.Status,
	}, nil
}
