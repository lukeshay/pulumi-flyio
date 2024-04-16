package main

import (
	p "github.com/pulumi/pulumi-go-provider"

	flyio "github.com/lukeshay/pulumi-flyio/provider/pkg/provider"
)

//go:generate go run ./generate.go

// Serve the provider against Pulumi's Provider protocol.
func main() { p.RunProvider(flyio.Name, flyio.Version, flyio.Provider()) }
