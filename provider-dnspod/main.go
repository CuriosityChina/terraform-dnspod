package main

import (
	"github.com/CuriosityChina/terraform-dnspod"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dnspod.Provider,
	})
}
