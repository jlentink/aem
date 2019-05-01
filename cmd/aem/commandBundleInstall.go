package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandBundleInstall struct {
	name             string
	http             *httpRequests
	utility          *utility
	projectStructure projectStructure
	bundle           string
	bundleStartLevel int
}

func (c *commandBundleInstall) Init() {
	c.name = configDefaultInstance
	c.http = new(httpRequests)
	c.utility = new(utility)
	c.projectStructure = newProjectStructure()
	c.bundle = ""
	c.bundleStartLevel = bundleStartLevel
}

func (c *commandBundleInstall) readConfig() bool {
	return true
}

func (c *commandBundleInstall) GetCommand() []string {
	return []string{"bundle-install"}
}

func (c *commandBundleInstall) GetHelp() string {
	return "Install bundle on instance."
}

func (c *commandBundleInstall) Execute(args []string) {
	c.getOpt(args)
	c.name = c.utility.getDefaultInstance(c.name)
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

func (c *commandBundleInstall) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name",
		'n', "Install bundle on instance (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "Path to bundle (.jar)")
	getopt.FlagLong(&c.bundleStartLevel, "startlevel", 's',
		fmt.Sprintf("bundle start level (default: %d)", bundleStartLevel))
	getopt.CommandLine.Parse(args)
}
