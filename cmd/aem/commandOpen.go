package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"os/exec"
	"runtime"
)

func newOpenCommand() commandOpen {
	return commandOpen{
		p:       new(projectStructure),
		utility: new(utility),
		name:    configDefaultInstance,
	}
}

type commandOpen struct {
	p        *projectStructure
	utility  *utility
	name     string
	instance aemInstanceConfig
}

func (o *commandOpen) Execute(args []string) {
	o.getOpt(args)
	o.instance = o.utility.getInstanceByName(o.name)

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", o.instance.URL())
		cmd.Start()
	case "darwin":
		cmd := exec.Command("open", o.instance.URL())
		cmd.Start()
	case "linux":
		cmd := exec.Command("xdg-open", o.instance.URL())
		cmd.Start()

	default:
		fmt.Printf("unsuported operating systen %s", runtime.GOOS)
	}
}

func (o *commandOpen) getOpt(args []string) {
	getopt.FlagLong(&o.name, "name", 'n', "Instance to open. (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
