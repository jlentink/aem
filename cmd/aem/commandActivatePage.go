package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newPageActivateCommand() commandPageActivate {
	return commandPageActivate{
		Group:            "",
		Name:             configDefaultInstance,
		Page:             "",
		activate:         false,
		deactivate:       false,
		u:                new(utility),
		projectStructure: new(projectStructure),
		http:             new(httpRequests),
	}
}

type commandPageActivate struct {
	Group            string
	Type             string
	Name             string
	Page             string
	activate         bool
	deactivate       bool
	u                *utility
	projectStructure *projectStructure
	http             *httpRequests
}

func (c *commandPageActivate) Execute(args []string) {
	c.getOpt(args)
	instances := c.u.getInstance(c.Name, c.Group)

	for _, instance := range instances {
		if c.activate {
			status := c.http.activatePage(instance, c.Page)
			fmt.Printf("Send action status received: %d\n", status)
		} else if c.deactivate {
			status := c.http.deactivatePage(instance, c.Page)
			fmt.Printf("Send action status received: %d\n", status)
		} else {
			exitProgram("Use --activate or --deactivate")
		}
	}
	fmt.Printf("Action(s) performed..")
}

func (c *commandPageActivate) getOpt(args []string) {
	getopt.FlagLong(&c.Name, "name", 'n', "Instance to target based on name")
	getopt.FlagLong(&c.Group, "group", 'g', "Instances to target based on group")
	getopt.FlagLong(&c.Page, "page", 'p', "Page to activate")
	getopt.FlagLong(&c.activate, "activate", 'a', "Activate")
	getopt.FlagLong(&c.deactivate, "deactivate", 'd', "Deactivate")
	getopt.CommandLine.Parse(args)
}
