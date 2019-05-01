package main

import (
	"github.com/pborman/getopt/v2"
)

type commandOakExplore struct {
	name       string
	aemVersion string
	oakVersion string
	oak        oak
	pStructure projectStructure
	utility    *utility
}

func (c *commandOakExplore) Init() {
	c.name = configDefaultInstance
	c.aemVersion = ""
	c.oakVersion = config.OakVersion
	c.oak = newOak()
	c.pStructure = newProjectStructure()
	c.utility = new(utility)

}

func (c *commandOakExplore) readConfig() bool {
	return true
}

func (c *commandOakExplore) GetCommand() []string {
	return []string{"oak-explore"}
}

func (c *commandOakExplore) GetHelp() string {
	return "Open oak explorer"
}

func (c *commandOakExplore) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	instancePath := c.pStructure.getRunDirLocation(instance)
	oakPath := c.oak.getVersion(c.aemVersion, c.oakVersion)

	oakArgs := []string{"explore", instancePath + oakRepoPath}

	c.oak.execute(oakPath, oakArgs)

}

func (c *commandOakExplore) getOpt(args []string) {
	getopt.FlagLong(&c.aemVersion, "aem", 'a', "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	getopt.FlagLong(&c.oakVersion, "oak", 'o', "Define version of oak-run to use")
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to use oak-run on (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
