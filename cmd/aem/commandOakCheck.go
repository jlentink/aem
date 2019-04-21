package main

import (
	"github.com/pborman/getopt/v2"
)

type commandOakCheck struct {
	name       string
	aemVersion string
	oakVersion string
	oak        oak
	pStructure projectStructure
	utility    *utility
}

func (c *commandOakCheck) Init() {
	c.name = configDefaultInstance
	c.aemVersion = ""
	c.oakVersion = config.OakVersion
	c.oak = newOak()
	c.pStructure = newProjectStructure()
	c.utility = new(utility)

}

func (c *commandOakCheck) readConfig() bool {
	return true
}

func (c *commandOakCheck) GetCommand() []string {
	return []string{"oak-check"}
}

func (c *commandOakCheck) GetHelp() string {
	return "Run oak-run check on instance. Check the FileStore for inconsistencies."
}

func (c *commandOakCheck) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	instancePath := c.pStructure.getRunDirLocation(instance)
	oakPath := c.oak.getVersion(c.aemVersion, c.oakVersion)

	oakArgs := []string{"check", instancePath + oakRepoPath}

	c.oak.execute(oakPath, oakArgs)

}

func (c *commandOakCheck) getOpt(args []string) {
	getopt.FlagLong(&c.aemVersion, "aem", 'a', "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	getopt.FlagLong(&c.oakVersion, "oak", 'o', "Define version of oak-run to use")
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to use oak-run on (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
