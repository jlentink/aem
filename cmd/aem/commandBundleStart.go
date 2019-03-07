package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewBundleStartCommand() bundleStartCommand {
	return bundleStartCommand{
		name:    CONFIG_DEFAULT_INSTANCE,
		http:    new(HttpRequests),
		utility: new(Utility),
		bundle:  "",
	}
}

type bundleStartCommand struct {
	name    string
	http    *HttpRequests
	utility *Utility
	bundle  string
}

func (c *bundleStartCommand) Execute(args []string) {
	u := Utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)
	bundlePicker := NewBundlePicker()
	bundles := make([]Bundle, 0)

	if len(c.bundle) > 0 {
		bundles = append(bundles, Bundle{SymbolicName: c.bundle})
	} else {
		bundles = bundlePicker.picker(instance)
	}

	for _, bundle := range bundles {
		fmt.Printf("Starting %s\n", bundle.Name)
		resp := c.http.bundleStopStart(instance, bundle, BundleStatusStart)
		fmt.Printf("%s (%s) | Status %s -> %s\n", bundle.Name, bundle.SymbolicName, bundle.State, BundleRawState[resp.StateRaw])
	}

}

func (c *bundleStartCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "Bundle to start (Symbolic name)")
	getopt.CommandLine.Parse(args)
}
