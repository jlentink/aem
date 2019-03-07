package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"github.com/spf13/afero"
)

func NewInitCommand() commandInit {
	return commandInit{
		u:  new(Utility),
		p:  new(projectStructure),
		fs: afero.NewOsFs(),
	}
}

type commandInit struct {
	u  *Utility
	p  *projectStructure
	fs afero.Fs
}

func (p *commandInit) Execute(args []string) {
	p.getOpt(args)

	if !p.u.Exists(p.p.getConfigFileLocation()) {
		err := afero.WriteFile(p.fs, p.p.getConfigFileLocation(), []byte(configTemplate), 0644)
		exitFatal(err, "Could not write config file.")
		fmt.Printf("Written sample config file. please edit .aem\n")

	} else {
		exitProgram("\".aem\" file found; please edit to update the values.")
	}

}

func (p *commandInit) getOpt(args []string) {
	getopt.CommandLine.Parse(args)
}
