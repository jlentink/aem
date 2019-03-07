package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func NewStopCommand() commandStop {
	return commandStop{
		p:       new(projectStructure),
		utility: new(Utility),
		name:    CONFIG_DEFAULT_INSTANCE,
	}
}

type commandStop struct {
	p       *projectStructure
	utility *Utility
	name    string
	instance AEMInstanceConfig
}

func (s *commandStop) Execute(args []string) {
	s.getOpt(args)
	s.instance = s.utility.getInstanceByName(s.name)
	rundir := s.p.getRunDirLocation(s.instance)

	if _, err := os.Stat(rundir); os.IsNotExist(err) {
		log.Fatal("Could not find instance dir.")
	}

	if s.utility.Exists(s.p.getPidFileLocation(s.instance)) {
		pid, _ := ioutil.ReadFile(s.p.getPidFileLocation(s.instance))
		cmd := exec.Command("kill", string(pid))
		cmd.Dir = rundir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		exitFatal(err, "Error stoping AEM")
		fmt.Println("AEM stopping...")
		os.Remove(s.p.getPidFileLocation(s.instance))
	} else {
		fmt.Printf("No Pid file found. No running AEM expected.")
	}
}

func (s *commandStop) getOpt(args []string) {
	getopt.FlagLong(&s.name, "name", 'n', "Instance to stop. (default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.CommandLine.Parse(args)
}
