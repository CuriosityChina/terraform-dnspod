package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/magicshui/terraform-dnspod"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dnspod.Provider,
	})
}
