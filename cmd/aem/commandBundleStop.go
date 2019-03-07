package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewBundleStopCommand() bundleStopCommand {
	return bundleStopCommand{
		name:    CONFIG_DEFAULT_INSTANCE,
		http:    new(HttpRequests),
		utility: new(Utility),
		bundle:  "",
	}
}

type bundleStopCommand struct {
	name    string
	http    *HttpRequests
	utility *Utility
	bundle  string
}

func (c *bundleStopCommand) Execute(args []string) {
	u := Utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)
	bundles := make([]Bundle, 0)
	bundlePicker := NewBundlePicker()

	if len(c.bundle) > 0 {
		bundles = append(bundles, Bundle{SymbolicName: c.bundle})
	} else {
		bundles = bundlePicker.picker(instance)
	}

	for _, bundle := range bundles {
		fmt.Printf("Stopping %s\n", bundle.Name)
		resp := c.http.bundleStopStart(instance, bundle, BundleStatusStop)
		fmt.Printf("%s (%s) | Status %s -> %s\n", bundle.Name, bundle.SymbolicName, bundle.State, BundleRawState[resp.StateRaw])
	}

}

func (c *bundleStopCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "Bundle to start (Symbolic name)")
	getopt.CommandLine.Parse(args)
}
