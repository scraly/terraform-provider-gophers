package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/scraly/terraform-provider-gophers/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug: debugMode,

		// TODO: update this string with the full name of your provider as used in your configs
		// ProviderAddr: "registry.terraform.io/hashicorp/scaffolding",
		// ProviderAddr: "registry.terraform.io/scraly/gophers",
		// ProviderAddr: "terraform.local/local/gophers",

		ProviderFunc: provider.New(version),
	}

	plugin.Serve(opts)
}
