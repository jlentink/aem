package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"os"
	"os/exec"
	"strings"
	"time"
)

func NewSyncCommand() commandSync {
	return commandSync{
		instanceName:     CONFIG_DEFAULT_INSTANCE,
		instanceGroup:    "",
		utility:          new(Utility),
		projectStructure: new(projectStructure),
		disableLog:       false,
		bin:              "aemsync",
	}
}

type commandSync struct {
	instanceName     string
	instanceGroup    string
	utility          *Utility
	disableLog       bool
	bin              string
	projectStructure *projectStructure
}

func (p *commandSync) Execute(args []string) {
	p.getOpt(args)
	instances := p.utility.getInstance(p.instanceName, p.instanceGroup)

	target := p.getTargetsString(instances)

	for _, watchPath := range config.WatchPath {
		p.sync(target, watchPath)
	}

	fmt.Printf("Press CTRL + c to stop.\n")
	for {
		time.Sleep(1 * time.Second)
	}
}

func (p *commandSync) getTargetsString(instances []AEMInstanceConfig) string {
	instancesStr := make([]string, 0)
	for _, instance := range instances {
		instancesStr = append(instancesStr, instance.PasswordURL())
	}

	return strings.Join(instancesStr, ",")
}

func (p *commandSync) sync(instance string, path string) {
	cmd := exec.Command(p.bin, "-t", instance, "-w", path)
	if !p.disableLog {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	cmd.Start()
	go func() {
		err := cmd.Wait()
		fmt.Printf("Command finished with error: %v\n", err)
	}()
}

func (p *commandSync) getOpt(args []string) {
	getopt.FlagLong(&p.instanceName, "instance-name", 'i', "Instance to sync to. (default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.FlagLong(&p.instanceGroup, "instance-group", 'g', "Instance group to sync to.")
	getopt.FlagLong(&p.disableLog, "disable-log", 'l', "Disable AEM log output")
	getopt.FlagLong(&p.bin, "aemsync", 's', "Path to AEM sync")
	getopt.CommandLine.Parse(args)
}
