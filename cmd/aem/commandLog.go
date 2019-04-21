package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/pborman/getopt/v2"
)

type commandLog struct {
	follow           bool
	projectStructure *projectStructure
	utility          *utility
	name             string
}

func (c *commandLog) Init() {
	c.follow = false
	c.projectStructure = new(projectStructure)
	c.utility = new(utility)
	c.name = configDefaultInstance
}

func (c *commandLog) readConfig() bool {
	return true
}

func (c *commandLog) GetCommand() []string {
	return []string{"log"}
}

func (c *commandLog) GetHelp() string {
	return "Show log file content."
}

func (c *commandLog) Execute(args []string) {
	c.getOpt(args)
	logFile := c.projectStructure.getLogFileLocation(c.utility.getInstanceByName(c.name))

	t, err := tail.TailFile(logFile, tail.Config{Follow: c.follow, MustExist: true})
	exitFatal(err, "Could not read error log. %s", logFile)
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func (c *commandLog) getOpt(args []string) {
	getopt.FlagLong(&c.follow, "follow", 'f', "Follow log file. Show new lines if they come in.")
	getopt.FlagLong(&c.name, "name", 'n', "Instance to start. (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
