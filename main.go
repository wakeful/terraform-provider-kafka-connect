package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/wakeful/terraform-provider-kafka-connect/kafkaconnect"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kafkaconnect.Provider})
}
