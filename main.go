package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/ondrejsika/terraform-provider-linodex/linodex"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: linodex.Provider,
	})
}
