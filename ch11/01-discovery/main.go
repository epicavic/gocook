package main

import (
	"main/discovery"

	"github.com/hashicorp/consul/api"
)

func main() {
	// return consul default client config (localhost:8500)
	config := api.DefaultConfig()

	// fake service registration
	cli, err := discovery.NewClient(config, "localhost", "discovery", "fake", 8080)
	if err != nil {
		panic(err)
	}

	if err := discovery.Exec(cli); err != nil {
		panic(err)
	}
}

/*
$ ./run.sh
&api.AgentService{
	Kind:"fake",
	ID:"discovery",
	Service:"discovery",
	Tags:[]string{"Go", "Awesome"},
	Meta:map[string]string(nil),
	Port:8080,
	Address:"localhost",
	TaggedAddresses:map[string]api.ServiceAddress(nil),
	Weights:api.AgentWeights{Passing:1, Warning:1},
	EnableTagOverride:false,
	CreateIndex:0xe,
	ModifyIndex:0xe,
	ContentHash:"",
	Proxy:(*api.AgentServiceConnectProxyConfig)(0xc0002007e0),
	Connect:(*api.AgentServiceConnect)(0xc000212300),
	Namespace:"",
	Datacenter:""
}
*/
