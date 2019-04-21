package main

import (
	"fmt"
)

type commandVersion struct {
}

func (c *commandVersion) Init() {
}

func (c *commandVersion) readConfig() bool {
	return false
}

func (c *commandVersion) GetCommand() []string {
	return []string{"version"}
}

func (c *commandVersion) GetHelp() string {
	return "Version of aem cli"
}

func (c *commandVersion) Execute(args []string) {
	fmt.Printf("%s (%s)\n", BuiltVersion, BuiltHash)
}

func (c *commandVersion) getOpt(args []string) {}
