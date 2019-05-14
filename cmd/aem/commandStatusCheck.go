package main

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/pborman/getopt/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type binReq = int

const (
	binRequirementNone = binReq(iota)
	binRequirementRequired
	binRequirementOptional
)

type binaryDescription struct {
	bin         string
	description string
	required    int
}

type commandSetupCheck struct {
	binaries   []binaryDescription
	returnCode int
}

func (c *commandSetupCheck) Init() {
	c.returnCode = 0
	c.binaries = []binaryDescription{
		{bin: "java", description: "needed to run AEM and run-oak. (Crucial)", required: binRequirementRequired},
		{bin: "vlt", description: "needed for all vlt actions", required: binRequirementOptional},
		{bin: "aemsync", description: "needed to sync frontend code with AEM instance", required: binRequirementOptional},
		{bin: "lazybones", description: "Lazybones templates for Adobe Experience Manager", required: binRequirementOptional},
	}

	switch strings.ToLower(runtime.GOOS) {
	case "darwin":
		c.binaries = append(c.binaries, []binaryDescription{
			{bin: "kill", description: "needed to run stop. (Crucial)", required: binRequirementRequired},
			{bin: "open", description: "needed to open the browser with the correct URL.", required: binRequirementOptional},
		}...)
	case "linux":
		c.binaries = append(c.binaries, []binaryDescription{
			{bin: "kill", description: "needed to run stop. (Crucial)", required: binRequirementRequired},
		}...)
	case "windows":
		c.binaries = append(c.binaries, []binaryDescription{
			{bin: "rundll32", description: "needed to open the browser with the correct URL.", required: binRequirementOptional},
		}...)
	default:
		exitProgram("Unknown operating system")
	}
}

func (c *commandSetupCheck) readConfig() bool {
	return false
}

func (c *commandSetupCheck) GetCommand() []string {
	return []string{"setup-check", "check"}
}

func (c *commandSetupCheck) GetHelp() string {
	return "Check if all needed binaries are available for all functionality"
}

func (c *commandSetupCheck) Execute(args []string) {
	c.getOpt(args)
	c.checkBinaries()
	os.Exit(c.returnCode)
}

func (c *commandSetupCheck) checkBinaries() {
	for _, bin := range c.binaries {
		_, err := exec.LookPath(bin.bin)
		c.printStatus(err == nil, bin)
	}
}

func (c *commandSetupCheck) printStatus(status bool, bin binaryDescription) {
	color := ct.Green
	statusMsg := " FOUND "
	if !status {
		statusMsg = "MISSING"
		switch bin.required {
		case binRequirementRequired:
			c.returnCode = 1
			color = ct.Red
		case binRequirementOptional:
			color = ct.Yellow
		case binRequirementNone:
			color = ct.White
		}
	}
	fmt.Print("[ ")
	ct.ChangeColor(color, false, ct.None, false)
	fmt.Print(statusMsg)
	ct.ResetColor()
	fmt.Printf(" ] %s -  %s\n", bin.bin, bin.description)
}

func (c *commandSetupCheck) getOpt(args []string) {
	getopt.CommandLine.Parse(args)
}
