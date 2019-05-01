package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandActivatePage struct {
	Group            string
	Type             string
	name             string
	Page             string
	activate         bool
	deactivate       bool
	u                *utility
	projectStructure *projectStructure
	http             *httpRequests
}

func (c *commandActivatePage) Init() {
	c.Group = ""
	c.name = configDefaultInstance
	c.Page = ""
	c.activate = false
	c.deactivate = false
	c.u = new(utility)
	c.projectStructure = new(projectStructure)
	c.http = new(httpRequests)
}

func (c *commandActivatePage) readConfig() bool {
	return true
}

func (c *commandActivatePage) GetCommand() []string {
	return []string{"activate-page"}
}

func (c *commandActivatePage) GetHelp() string {
	return "Activate page on instance."
}

func (c *commandActivatePage) Execute(args []string) {
	c.getOpt(args)
	c.name = c.u.getDefaultInstance(c.name)
	instances := c.u.getInstance(c.name, c.Group)

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

func (c *commandActivatePage) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name",
		'n', "Instance that need page activation (default: "+c.u.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.Group, "group", 'g', "Instances to target based on group")
	getopt.FlagLong(&c.Page, "page", 'p', "Page to activate")
	getopt.FlagLong(&c.activate, "activate", 'a', "Activate")
	getopt.FlagLong(&c.deactivate, "deactivate", 'd', "Deactivate")
	getopt.CommandLine.Parse(args)
}
