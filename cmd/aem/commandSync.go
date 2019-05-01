package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"os"
	"os/exec"
	"strings"
	"time"
)

type commandSync struct {
	instanceName     string
	instanceGroup    string
	utility          *utility
	disableLog       bool
	bin              string
	projectStructure *projectStructure
}

func (c *commandSync) Init() {
	c.instanceName = configDefaultInstance
	c.instanceGroup = ""
	c.utility = new(utility)
	c.projectStructure = new(projectStructure)
	c.disableLog = false
	c.bin = "aemsync"

}

func (c *commandSync) readConfig() bool {
	return true
}

func (c *commandSync) GetCommand() []string {
	return []string{"sync"}
}

func (c *commandSync) GetHelp() string {
	return "Show log file content."
}

func (c *commandSync) Execute(args []string) {
	c.getOpt(args)
	instances := c.utility.getInstance(c.instanceName, c.instanceGroup)

	target := c.getTargetsString(instances)

	for _, watchPath := range config.WatchPath {
		c.sync(target, watchPath)
	}

	fmt.Printf("Press CTRL + c to stop.\n")
	for {
		time.Sleep(1 * time.Second)
	}
}

func (c *commandSync) getTargetsString(instances []aemInstanceConfig) string {
	instancesStr := make([]string, 0)
	for _, instance := range instances {
		instancesStr = append(instancesStr, instance.PasswordURL())
	}

	return strings.Join(instancesStr, ",")
}

func (c *commandSync) sync(instance string, path string) {
	cmd := exec.Command(c.bin, "-t", instance, "-w", path)
	if !c.disableLog {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	cmd.Start()
	go func() {
		err := cmd.Wait()
		fmt.Printf("Command finished with error: %v\n", err)
	}()
}

func (c *commandSync) getOpt(args []string) {
	getopt.FlagLong(&c.instanceName, "instance-name", 'i', "Instance to sync to. (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.instanceGroup, "instance-group", 'g', "Instance group to sync to.")
	getopt.FlagLong(&c.disableLog, "disable-log", 'l', "Disable AEM log output")
	getopt.FlagLong(&c.bin, "aemsync", 's', "Path to AEM sync")
	getopt.CommandLine.Parse(args)
}
