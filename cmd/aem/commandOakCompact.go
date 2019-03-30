package main

import (
	"github.com/pborman/getopt/v2"
)

func newCommandOakCompact() commandOakCompact {
	return commandOakCompact{
		name:       configDefaultInstance,
		aemVersion: "",
		oakVersion: config.OakVersion,
		oak:        newOak(),
		pStructure: newProjectStructure(),
		utility:    new(utility),
	}
}

type commandOakCompact struct {
	name       string
	aemVersion string
	oakVersion string
	oak        oak
	pStructure projectStructure
	utility    *utility
}

func (c *commandOakCompact) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	instancePath := c.pStructure.getRunDirLocation(instance)
	oakPath := c.oak.getVersion(c.aemVersion, c.oakVersion)

	oakArgs := []string{"compact", instancePath + oakRepoPath}

	c.oak.execute(oakPath, oakArgs)

}

func (c *commandOakCompact) getOpt(args []string) {
	getopt.FlagLong(&c.aemVersion, "aem", 'a', "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	getopt.FlagLong(&c.oakVersion, "oak", 'o', "Define version of oak-run to use")
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to use oak-run on (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
