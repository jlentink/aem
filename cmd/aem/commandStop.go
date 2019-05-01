package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type commandStop struct {
	p        *projectStructure
	utility  *utility
	name     string
	instance aemInstanceConfig
}

func (c *commandStop) Init() {
	c.p = new(projectStructure)
	c.utility = new(utility)
	c.name = configDefaultInstance
}

func (c *commandStop) readConfig() bool {
	return true
}

func (c *commandStop) GetCommand() []string {
	return []string{"stop"}
}

func (c *commandStop) GetHelp() string {
	return "Stop an Adobe Experience Manager instance."
}

func (c *commandStop) checkEnvironmentName() {
	envName := os.Getenv(aemEnvName)
	if c.name == configDefaultInstance && len(envName) > 0 {
		fmt.Printf("Found env variable changing instance name to %s.\n", envName)
		c.name = envName
	}
}


func (c *commandStop) Execute(args []string) {
	c.getOpt(args)
	c.checkEnvironmentName()
	c.instance = c.utility.getInstanceByName(c.name)
	rundir := c.p.getRunDirLocation(c.instance)

	if _, err := os.Stat(rundir); os.IsNotExist(err) {
		log.Fatal("Could not find instance dir.")
	}

	if c.utility.Exists(c.p.getPidFileLocation(c.instance)) {
		pid, _ := ioutil.ReadFile(c.p.getPidFileLocation(c.instance))
		cmd := exec.Command("kill", string(pid))
		cmd.Dir = rundir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		exitFatal(err, "Error stoping AEM")
		fmt.Println("AEM stopping...")
		os.Remove(c.p.getPidFileLocation(c.instance))
	} else {
		fmt.Printf("No Pid file found. No running AEM expected.")
	}
}

func (c *commandStop) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Instance to stop. (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
