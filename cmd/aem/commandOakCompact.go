package main

import (
	"github.com/pborman/getopt/v2"
)

type commandOakCompact struct {
	name       string
	aemVersion string
	oakVersion string
	oak        oak
	pStructure projectStructure
	utility    *utility
}

func (c *commandOakCompact) Init() {
	c.name = configDefaultInstance
	c.aemVersion = ""
	c.oakVersion = config.OakVersion
	c.oak = newOak()
	c.pStructure = newProjectStructure()
	c.utility = new(utility)

}

func (c *commandOakCompact) readConfig() bool {
	return true
}

func (c *commandOakCompact) GetCommand() []string {
	return []string{"oak-compact"}
}

func (c *commandOakCompact) GetHelp() string {
	return "Run oak-run offline compaction on instance."
}

func (c *commandOakCompact) Execute(args []string) {
	c.getOpt(args)
	c.name = c.utility.getDefaultInstance(c.name)
	instance := c.utility.getInstanceByName(c.name)
	instancePath := c.pStructure.getRunDirLocation(instance)
	oakPath := c.oak.getVersion(c.aemVersion, c.oakVersion)

	oakArgs := []string{"compact", instancePath + oakRepoPath}

	c.oak.execute(oakPath, oakArgs)

}

func (c *commandOakCompact) getOpt(args []string) {
	getopt.FlagLong(&c.aemVersion, "aem", 'a', "Version of AEM to use oak on. (use matching AEM version of oak-run)")
	getopt.FlagLong(&c.oakVersion, "oak", 'o', "Define version of oak-run to use")
	getopt.FlagLong(&c.name, "name",
		'n', "Name of instance to use oak-run on (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.CommandLine.Parse(args)
}
