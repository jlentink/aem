package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newBundleInstallCommand() bundleInstallCommand {
	return bundleInstallCommand{
		name:             configDefaultInstance,
		http:             new(httpRequests),
		utility:          new(utility),
		projectStructure: newProjectStructure(),
		bundle:           "",
		bundleStartLevel: bundleStartLevel,
	}
}

type bundleInstallCommand struct {
	name             string
	http             *httpRequests
	utility          *utility
	projectStructure projectStructure
	bundle           string
	bundleStartLevel int
}

func (c *bundleInstallCommand) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	fmt.Printf("%+v", instance)
	if c.projectStructure.exists(c.bundle) {
		if c.http.bundleInstall(instance, c.bundle, bundleInstall, bundleStartLevel) {
			fmt.Printf("bundle installed.")
		} else {
			fmt.Printf("bundle response was unexprected")
		}
	} else {
		exitProgram("Could not find bundle. (%s)\n", c.bundle)
	}
}

func (c *bundleInstallCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "Path to bundle (.jar)")
	getopt.FlagLong(&c.bundleStartLevel, "startlevel", 's', "bundle start level (default: "+string(bundleStartLevel)+")")
	getopt.CommandLine.Parse(args)
}
