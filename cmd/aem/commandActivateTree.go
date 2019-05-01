package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandActivateTree struct {
	name             string
	Path             string
	utility          *utility
	i                *instance
	projectStructure *projectStructure
	http             *httpRequests
}

func (c *commandActivateTree) Init() {
	c.name = configDefaultInstance
	c.Path = ""
	c.utility = new(utility)
	c.i = new(instance)
	c.projectStructure = new(projectStructure)
	c.http = new(httpRequests)
}

func (c *commandActivateTree) readConfig() bool {
	return true
}

func (c *commandActivateTree) GetCommand() []string {
	return []string{"replicate-tree"}
}

func (c *commandActivateTree) GetHelp() string {
	return "Replicate tree on instance."
}

func (c *commandActivateTree) Execute(args []string) {
	c.getOpt(args)
	c.name = c.utility.getDefaultInstance(c.name)
	instance := c.i.getByName(c.name)
	if c.http.activateTree(instance, c.Path) {
		fmt.Printf("Tree activated.\n")
	} else {
		fmt.Printf("Error while activating tree.\n")
	}
}

func (c *commandActivateTree) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name",
		'n', "Activate Tree on instance (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.Path, "path", 'p', "Path to activate")
	getopt.CommandLine.Parse(args)
}
