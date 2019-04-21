package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandBundleStart struct {
	name    string
	http    *httpRequests
	utility *utility
	bundle  string
}

func (c *commandBundleStart) Init() {
	c.name = configDefaultInstance
	c.http = new(httpRequests)
	c.utility = new(utility)
	c.bundle = ""
}

func (c *commandBundleStart) readConfig() bool {
	return true
}

func (c *commandBundleStart) GetCommand() []string {
	return []string{"bundle-start"}
}

func (c *commandBundleStart) GetHelp() string {
	return "Start bundle on AEM instance."
}

func (c *commandBundleStart) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)
	bundlePicker := newBundlePicker()
	bundles := make([]bundle, 0)

	if len(c.bundle) > 0 {
		bundles = append(bundles, bundle{SymbolicName: c.bundle})
	} else {
		bundles = bundlePicker.picker(instance)
	}

	for _, bundle := range bundles {
		fmt.Printf("Starting %s\n", bundle.Name)
		resp := c.http.bundleStopStart(instance, bundle, BundleStatusStart)
		fmt.Printf("%s (%s) | Status %s -> %s\n", bundle.Name, bundle.SymbolicName, bundle.State, bundleRawState[resp.StateRaw])
	}

}

func (c *commandBundleStart) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "bundle to start (Symbolic name)")
	getopt.CommandLine.Parse(args)
}
