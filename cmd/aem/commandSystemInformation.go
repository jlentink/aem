package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"strings"
)

type commandSystemInformation struct {
	p        *projectStructure
	a        *httpRequests
	i        *instance
	utility  *utility
	name     string
	instance aemInstanceConfig
}

func (c *commandSystemInformation) Init() {
	c.p = new(projectStructure)
	c.a = new(httpRequests)
	c.i = new(instance)
	c.utility = new(utility)
	c.name = configDefaultInstance
}

func (c *commandSystemInformation) readConfig() bool {
	return true
}

func (c *commandSystemInformation) GetCommand() []string {
	return []string{"system-information", "sysinfo"}
}

func (c *commandSystemInformation) GetHelp() string {
	return "Show system information of instance (>=AEM 6.4)"
}

func (c *commandSystemInformation) Execute(args []string) {
	c.getOpt(args)
	c.instance = c.i.getByName(c.name)

	sysInfo, err := c.a.getSystemInformation(c.instance)
	exitFatal(err, "Could not get information about instance. Check version AEM >= 6.4 or credentials")
	c.printInstanceInformation(sysInfo)
	c.printRepositoryInformation(sysInfo)
	c.printSystemInformation(sysInfo)
	c.printInstanceInformation(sysInfo)
	c.printHealthCheck(sysInfo)
	c.printMaintenanceTasks(sysInfo)
	c.printReplicationAgents(sysInfo)
	c.printDistributionAgents(sysInfo)
}

func (c *commandSystemInformation) printInstanceInformation(information *systemInformation) {
	fmt.Printf("Adobe Experience manager:\n")
	fmt.Printf("- Version: %s\n", information.Instance.AdobeExperienceManager)
	fmt.Printf("- Run mode: %s\n", information.Instance.RunModes)
	fmt.Printf("- Up since: %s\n\n", information.Instance.InstanceUpSince)
}

func (c *commandSystemInformation) printRepositoryInformation(information *systemInformation) {
	fmt.Printf("Repostitory\n")
	fmt.Printf("- Version: %s\n", information.Repository.ApacheJackrabbitOak)
	fmt.Printf("- Repository size: %s\n", information.Repository.RepositorySize)
	fmt.Printf("- Storage location: %s\n", information.Repository.FileDataStore)
	fmt.Printf("- Nodes total: %s\n", information.EstimatedNodeCounts.Total)
	fmt.Printf("- Tags total: %s\n", information.EstimatedNodeCounts.Tags)
	fmt.Printf("- Authorizables: %s\n", information.EstimatedNodeCounts.Authorizables)
	fmt.Printf("- Pages: %s\n\n", information.EstimatedNodeCounts.Pages)
}

func (c *commandSystemInformation) printSystemInformation(information *systemInformation) {
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
func (c *commandSystemInformation) printMaintenanceTasks(information *systemInformation) {
	fmt.Printf("Maintenance Tasks:\n")
	tasks := information.MaintenanceTasks.(map[string]interface{})
	for key, value := range tasks {
		c.printListing(key, fmt.Sprintf("%s", value))
	}
	fmt.Printf("\n")
}

func (c *commandSystemInformation) printHealthCheck(information *systemInformation) {
	warns := information.HealthChecks.(map[string]interface{})

	fmt.Printf("Current health checks:\n")
	if len(warns) > 0 {
		for key, check := range warns {
			c.printListing(key, fmt.Sprintf("%s", check))
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("No Warnings or errors.\n\n")
	}
}

func (c *commandSystemInformation) printReplicationAgents(information *systemInformation) {
	agents := information.ReplicationAgents.(map[string]interface{})
	fmt.Printf("Replication Agents:\n")
	if len(agents) > 0 {
		for key, check := range agents {
			c.printListing(key, fmt.Sprintf("%s", check))
		}
		fmt.Printf("\n")
	}
}

func (c *commandSystemInformation) printDistributionAgents(information *systemInformation) {
	agents := information.DistributionAgents.(map[string]interface{})
	fmt.Printf("Distribution Agents:\n")
	if len(agents) > 0 {
		for key, check := range agents {
			c.printListing(key, fmt.Sprintf("%s", check))
		}
		fmt.Printf("\n")
	}
}

func (c *commandSystemInformation) printListing(key, valuesString string) {
	values := strings.Split(valuesString, ",")
	fmt.Printf("- %s:\n", key)
	for i, value := range values {
		fmt.Printf("\t %d) %s\n", i+1, strings.TrimSpace(value))
	}
}

func (c *commandSystemInformation) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Instance to start. (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
