package main

import (
	"github.com/pborman/getopt/v2"
)

func newCommandOakConsole() commandOakConsole {
	return commandOakConsole{
		name:       configDefaultInstance,
		aemVersion: "",
		oakVersion: config.OakVersion,
		oak:        newOak(),
		write:      false,
		pStructure: newProjectStructure(),
		utility:    new(utility),
	}
}

type commandOakConsole struct {
	name       string
	aemVersion string
	oakVersion string
	oak        oak
	write      bool
	pStructure projectStructure
	utility    *utility
}

func (c *commandOakConsole) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	instancePath := c.pStructure.getRunDirLocation(instance)
	oakPath := c.oak.getVersion(c.aemVersion, c.oakVersion)

	oakArgs := []string{"console", instancePath + oakRepoPath}

	if c.write {
		oakArgs = append(oakArgs, "--read-write")
	}

	c.oak.execute(oakPath, oakArgs)

}

func (c *commandOakConsole) getOpt(args []string) {
	getopt.FlagLong(&c.aemVersion, "aem", 'a', "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	getopt.FlagLong(&c.oakVersion, "oak", 'o', "Define version of oak-run to use")
	getopt.FlagLong(&c.write, "write", 'w', "Enable write mode")
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to use oak-run on (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
