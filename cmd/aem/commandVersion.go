package main

import (
	"fmt"
)

func NewVersionCommand() commandVersion {
	return commandVersion{
	}
}

type commandVersion struct {
}

func (p *commandVersion) Execute(args []string) {
	fmt.Printf("%s (%s)\n", BuiltVersion, BuiltHash)
}
