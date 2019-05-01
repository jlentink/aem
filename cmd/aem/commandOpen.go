package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"os/exec"
	"runtime"
)

type commandOpen struct {
	p        *projectStructure
	utility  *utility
	name     string
	instance aemInstanceConfig
}

func (c *commandOpen) Init() {
	c.p = new(projectStructure)
	c.utility = new(utility)
	c.name = configDefaultInstance
}

func (c *commandOpen) readConfig() bool {
	return true
}

func (c *commandOpen) GetCommand() []string {
	return []string{"open"}
}

func (c *commandOpen) GetHelp() string {
	return "Open instance in browser."
}

func (c *commandOpen) Execute(args []string) {
	c.getOpt(args)
	c.name = c.utility.getDefaultInstance(c.name)
	c.instance = c.utility.getInstanceByName(c.name)

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", c.instance.URL())
		cmd.Start()
	case "darwin":
		cmd := exec.Command("open", c.instance.URL())
		cmd.Start()
	case "linux":
		cmd := exec.Command("xdg-open", c.instance.URL())
		cmd.Start()

	default:
		fmt.Printf("unsuported operating systen %s", runtime.GOOS)
	}
}

func (c *commandOpen) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name",
		'n', "Open browser to instance (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.CommandLine.Parse(args)
}
