package main

import (
	"github.com/pborman/getopt/v2"
)

type commandVltCopy struct {
	follow           bool
	projectStructure *projectStructure
	utility          *utility
	from             string
	to               string
	path             string
	v                vlt
}

func (c *commandVltCopy) Init() {
	c.follow = false
	c.projectStructure = new(projectStructure)
	c.utility = new(utility)
	c.v = newVlt()
	c.to = configDefaultInstance
	c.from = ""
}

func (c *commandVltCopy) readConfig() bool {
	return true
}

func (c *commandVltCopy) GetCommand() []string {
	return []string{"vlt-copy"}
}

func (c *commandVltCopy) GetHelp() string {
	return "Copy data via VLT from one to another AEM instance"
}

func (c *commandVltCopy) Execute(args []string) {
	c.getOpt(args)
	c.to = c.utility.getDefaultInstance(configDefaultInstance)
	to := c.utility.getInstanceByName(c.to)
	from := c.utility.getInstanceByName(c.from)

	logger.Infof("Copy content from %s to %s", c.from, c.to)
	paths := c.v.getPaths(c.path, config.VltPaths)
	for _, path := range paths {
		args := []string{"rcp"}
		if path.recursive {
			args = append(args, "--recursive")
		}
		args = append(args, from.PasswordURL()+"/crx/-/jcr:root"+path.path)
		args = append(args, to.PasswordURL()+"/crx/-/jcr:root"+path.path)
		c.v.execute(args...)
	}
}

func (c *commandVltCopy) getOpt(args []string) {
	getopt.FlagLong(&c.to, "to",
		't', "Instance to data to (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.from, "from", 'f', "Instance to data from")
	getopt.FlagLong(&c.path, "path", 'p', "Use path instead of configuration paths")

	getopt.CommandLine.Parse(args)
}
