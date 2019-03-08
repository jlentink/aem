package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newPasswordCommand() commandPassword {
	return commandPassword{
		p:       new(projectStructure),
		utility: new(utility),
		i:       new(instance),
		all:     false,
		yes:     false,
	}
}

type commandPassword struct {
	p        *projectStructure
	utility  *utility
	name     string
	group    string
	all      bool
	yes      bool
	instance aemInstanceConfig
	i        *instance
}

func (c *commandPassword) Execute(args []string) {
	c.getOpt(args)

	if !config.KeyRing {
		exitProgram("Keyring is disabled. Insert the passwords in the config file.")
	}
	instances := c.getInstances()
	for _, instance := range instances {
		pw, err := c.i.keyRingGetPassword(instance)
		exitFatal(err, "Could not pull passwords from keychain")

		if len(pw) > 0 {
			if confirm("Password found for %s. Do you want to update the current password?\n", c.yes, instance.Name) {
				c.setPassword(instance)
			}
		} else {
			c.setPassword(instance)
		}
	}

}

func (c *commandPassword) getInstances() []aemInstanceConfig {
	if c.all {
		return config.Instances
	}
	return c.utility.getInstance(c.name, c.group)
}

func (c *commandPassword) setPassword(instance aemInstanceConfig) {
	fmt.Printf("Provide password for: %s\nPassword: ", instance.Name)
	pw := c.utility.readCmdLineInput()
	c.i.keyRingSetPassword(instance, pw)

}

func (c *commandPassword) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Instance to update.")
	getopt.FlagLong(&c.group, "group", 'g', "Instance group to update.")
	getopt.FlagLong(&c.all, "all", 'a', "Update all")
	getopt.FlagLong(&c.yes, "yes", 'y', "Confirm all questions with yes")
	getopt.CommandLine.Parse(args)
}
