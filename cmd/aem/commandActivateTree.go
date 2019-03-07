package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewActivateTreeCommand() commandActivateTree {
	return commandActivateTree{
		Name:             CONFIG_DEFAULT_INSTANCE,
		Path:             "",
		utility:          new(Utility),
		i:                new(Instance),
		projectStructure: new(projectStructure),
		http:             new(HttpRequests),
	}
}

type commandActivateTree struct {
	Name             string
	Path             string
	utility          *Utility
	i                *Instance
	projectStructure *projectStructure
	http             *HttpRequests
}

func (c *commandActivateTree) Execute(args []string) {
	c.getOpt(args)
	instance := c.i.getByName(c.Name)
	if c.http.ActivateTree(instance, c.Path) {
		fmt.Printf("Tree activated.\n")
	} else {
		fmt.Printf("Error while activating tree.\n")
	}
}

func (c *commandActivateTree) getOpt(args []string) {
	getopt.FlagLong(&c.Name, "instance", 'i', "Activate Tree on instance (Default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.FlagLong(&c.Path, "path", 'p', "Path to activate")
	getopt.CommandLine.Parse(args)
}
