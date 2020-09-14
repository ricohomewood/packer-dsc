package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/ricohomewood/packer-dsc/provisioner/dsc"
)

func main() {

	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(dsc.Provisioner))
	server.Serve()
}
