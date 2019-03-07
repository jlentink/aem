package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewBundleInstallCommand() bundleInstallCommand {
	return bundleInstallCommand{
		name:             CONFIG_DEFAULT_INSTANCE,
		http:             new(HttpRequests),
		utility:          new(Utility),
		projectStructure: NewProjectStructure(),
		bundle:           "",
	}
}

type bundleInstallCommand struct {
	name             string
	http             *HttpRequests
	utility          *Utility
	projectStructure projectStructure
	bundle           string
}

func (c *bundleInstallCommand) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	if c.projectStructure.exists(c.bundle) {
		if c.http.bundleInstall(instance, c.bundle) {
			fmt.Printf("Bundle installed.")
		} else {
			fmt.Printf("Bundle response was unexprected")
		}

	} else {
		exitProgram("Could not find bundle. (%s)\n", c.bundle)
	}

	//
	//bundlePicker := NewBundlePicker()
	//bundles := make([]Bundle, 0)

}

func (c *bundleInstallCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "Path to bundle (.jar)")
	getopt.CommandLine.Parse(args)
}
