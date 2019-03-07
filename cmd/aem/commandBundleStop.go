package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newBundleStopCommand() bundleStopCommand {
	return bundleStopCommand{
		name:    configDefaultInstance,
		http:    new(httpRequests),
		utility: new(utility),
		bundle:  "",
	}
}

type bundleStopCommand struct {
	name    string
	http    *httpRequests
	utility *utility
	bundle  string
}

func (c *bundleStopCommand) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)
	bundles := make([]bundle, 0)
	bundlePicker := newBundlePicker()

	if len(c.bundle) > 0 {
		bundles = append(bundles, bundle{SymbolicName: c.bundle})
	} else {
		bundles = bundlePicker.picker(instance)
	}

	for _, bundle := range bundles {
		fmt.Printf("Stopping %s\n", bundle.Name)
		resp := c.http.bundleStopStart(instance, bundle, BundleStatusStop)
		fmt.Printf("%s (%s) | Status %s -> %s\n", bundle.Name, bundle.SymbolicName, bundle.State, bundleRawState[resp.StateRaw])
	}

}

func (c *bundleStopCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "bundle to start (Symbolic name)")
	getopt.CommandLine.Parse(args)
}
