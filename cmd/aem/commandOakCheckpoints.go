package main

import (
	"github.com/pborman/getopt/v2"
)

type commandOakCheckpoints struct {
	name       string
	aemVersion string
	oakVersion string
	rm         bool
	oak        oak
	pStructure projectStructure
	utility    *utility
}

func (c *commandOakCheckpoints) Init() {
	c.name = configDefaultInstance
	c.aemVersion = ""
	c.oakVersion = config.OakVersion
	c.oak = newOak()
	c.pStructure = newProjectStructure()
	c.rm = false
	c.utility = new(utility)

}

func (c *commandOakCheckpoints) readConfig() bool {
	return true
}

func (c *commandOakCheckpoints) GetCommand() []string {
	return []string{"oak-checkpoints"}
}

func (c *commandOakCheckpoints) GetHelp() string {
	return "Run oak-run checkpoints on instance."
}

func (c *commandOakCheckpoints) Execute(args []string) {
	c.getOpt(args)
	instance := c.utility.getInstanceByName(c.name)
	instancePath := c.pStructure.getRunDirLocation(instance)
	oakPath := c.oak.getVersion(c.aemVersion, c.oakVersion)

	oakArgs := []string{"checkpoints", instancePath + oakRepoPath}

	if c.rm {
		oakArgs = append(oakArgs, "rm-all")
	}

	c.oak.execute(oakPath, oakArgs)

}

func (c *commandOakCheckpoints) getOpt(args []string) {
	getopt.FlagLong(&c.aemVersion, "aem", 'a', "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	getopt.FlagLong(&c.oakVersion, "oak", 'o', "Define version of oak-run to use")
	getopt.FlagLong(&c.rm, "rm", 'd', "Define version of oak-run to use")
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to use oak-run on (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
