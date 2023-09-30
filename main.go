package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/winebarrel/terraform-provider-cronplan/cronplan"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name cronplan

func main() {
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: cronplan.Provider,
		ProviderAddr: "registry.terraform.io/winebarrel/cronplan",
		Debug:        *debug,
	}

	plugin.Serve(opts)
}
