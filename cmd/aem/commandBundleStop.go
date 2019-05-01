package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandBundleStop struct {
	name    string
	http    *httpRequests
	utility *utility
	bundle  string
}

func (c *commandBundleStop) Init() {
	c.name = configDefaultInstance
	c.http = new(httpRequests)
	c.utility = new(utility)
	c.bundle = ""

}

func (c *commandBundleStop) readConfig() bool {
	return true
}

func (c *commandBundleStop) GetCommand() []string {
	return []string{"bundle-stop"}
}

func (c *commandBundleStop) GetHelp() string {
	return "Stop bundle on instance."
}

func (c *commandBundleStop) Execute(args []string) {
	c.getOpt(args)
	c.name = c.utility.getDefaultInstance(c.name)
	instance := c.utility.getInstanceByName(c.name)
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

func (c *commandBundleStop) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name",
		'n', "Stop bundle on instance (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "bundle to start (Symbolic name)")
	getopt.CommandLine.Parse(args)
}
