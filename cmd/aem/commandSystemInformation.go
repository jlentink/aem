package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"strings"
)

func newSystemInformationCommand() commandSystemInformation {
	return commandSystemInformation{
		p:       new(projectStructure),
		a:       new(httpRequests),
		i:       new(instance),
		utility: new(utility),
		name:    configDefaultInstance,
	}
}

type commandSystemInformation struct {
	p        *projectStructure
	a        *httpRequests
	i        *instance
	utility  *utility
	name     string
	instance aemInstanceConfig
}

func (s *commandSystemInformation) Execute(args []string) {
	s.getOpt(args)
	s.instance = s.i.getByName(s.name)

	sysInfo, err := s.a.getSystemInformation(s.instance)
	exitFatal(err, "Could not get information about instance. Check version AEM >= 6.4 or credentials")
	s.printInstanceInformation(sysInfo)
	s.printRepositoryInformation(sysInfo)
	s.printSystemInformation(sysInfo)
	s.printInstanceInformation(sysInfo)
	s.printHealthCheck(sysInfo)
	s.printMaintenanceTasks(sysInfo)
	s.printReplicationAgents(sysInfo)
	s.printDistributionAgents(sysInfo)
}

func (s *commandSystemInformation) printInstanceInformation(information *systemInformation) {
	fmt.Printf("Adobe Experience manager:\n")
	fmt.Printf("- Version: %s\n", information.Instance.AdobeExperienceManager)
	fmt.Printf("- Run mode: %s\n", information.Instance.RunModes)
	fmt.Printf("- Up since: %s\n\n", information.Instance.InstanceUpSince)
}

func (s *commandSystemInformation) printRepositoryInformation(information *systemInformation) {
	fmt.Printf("Repostitory\n")
	fmt.Printf("- Version: %s\n", information.Repository.ApacheJackrabbitOak)
	fmt.Printf("- Repository size: %s\n", information.Repository.RepositorySize)
	fmt.Printf("- Storage location: %s\n", information.Repository.FileDataStore)
	fmt.Printf("- Nodes total: %s\n", information.EstimatedNodeCounts.Total)
	fmt.Printf("- Tags total: %s\n", information.EstimatedNodeCounts.Tags)
	fmt.Printf("- Authorizables: %s\n", information.EstimatedNodeCounts.Authorizables)
	fmt.Printf("- Pages: %s\n\n", information.EstimatedNodeCounts.Pages)
}

func (s *commandSystemInformation) printSystemInformation(information *systemInformation) {
	fmt.Printf("Operating system:\n")
	if len(information.SystemInformation.Windows) > 0 {
		fmt.Printf("- Operating system: Windows %s\n", information.SystemInformation.Windows)
	}
	if len(information.SystemInformation.Linux) > 0 {
		fmt.Printf("- Operating system: Linux %s\n", information.SystemInformation.Linux)
	}
	if len(information.SystemInformation.MacOSX) > 0 {
		fmt.Printf("- Operating system: Mac OSX %s\n", information.SystemInformation.MacOSX)
	}
	fmt.Printf("- System Load Average: %s\n", information.SystemInformation.SystemLoadAverage)
	fmt.Printf("- Usable Disk Space %s\n", information.SystemInformation.UsableDiskSpace)
	fmt.Printf("- Maximum Heap %s\n\n", information.SystemInformation.MaximumHeap)
}
func (s *commandSystemInformation) printMaintenanceTasks(information *systemInformation) {
	fmt.Printf("Maintenance Tasks:\n")
	tasks := information.MaintenanceTasks.(map[string]interface{})
	for key, value := range tasks {
		s.printListing(key, fmt.Sprintf("%s", value))
	}
	fmt.Printf("\n")
}

func (s *commandSystemInformation) printHealthCheck(information *systemInformation) {
	warns := information.HealthChecks.(map[string]interface{})

	fmt.Printf("Current health checks:\n")
	if len(warns) > 0 {
		for key, check := range warns {
			s.printListing(key, fmt.Sprintf("%s", check))
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("No Warnings or errors.\n\n")
	}
}

func (s *commandSystemInformation) printReplicationAgents(information *systemInformation) {
	agents := information.ReplicationAgents.(map[string]interface{})
	fmt.Printf("Replication Agents:\n")
	if len(agents) > 0 {
		for key, check := range agents {
			s.printListing(key, fmt.Sprintf("%s", check))
		}
		fmt.Printf("\n")
	}
}

func (s *commandSystemInformation) printDistributionAgents(information *systemInformation) {
	agents := information.DistributionAgents.(map[string]interface{})
	fmt.Printf("Distribution Agents:\n")
	if len(agents) > 0 {
		for key, check := range agents {
			s.printListing(key, fmt.Sprintf("%s", check))
		}
		fmt.Printf("\n")
	}
}

func (s *commandSystemInformation) printListing(key, valuesString string) {
	values := strings.Split(valuesString, ",")
	fmt.Printf("- %s:\n", key)
	for i, value := range values {
		fmt.Printf("\t %d) %s\n", i+1, strings.TrimSpace(value))
	}
}

func (s *commandSystemInformation) getOpt(args []string) {
	getopt.FlagLong(&s.name, "name", 'n', "Instance to start. (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
