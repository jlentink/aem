package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/systeminformation"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandSystemInformation struct {
	verbose       bool
	instanceName  string
	allowRoot     bool
	foreground    bool
	forceDownload bool
	ignorePid     bool
}

func (c *commandSystemInformation) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "system-information",
		Short:   "Get system information from Adobe Experience Manager instance",
		Aliases: []string{"sys", "sysinfo"},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to start")
	cmd.Flags().BoolVarP(&c.forceDownload, "download", "d", false, "Force re-download")
	cmd.Flags().BoolVarP(&c.foreground, "foreground", "f", false, "on't detach aem from current tty")
	cmd.Flags().BoolVarP(&c.allowRoot, "allow-root", "r", false, "Allow to start as root user (UID: 0)")
	cmd.Flags().BoolVarP(&c.ignorePid, "ignore-pid", "", false, "Ignore existing PID file and start AEM")

	return cmd
}

func (c *commandSystemInformation) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandSystemInformation) run(cmd *cobra.Command, args []string) {

	_, currentInstance, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	sysInfo, err := systeminformation.GetSystemInformation(currentInstance)
	if err != nil {
		output.Printf(output.NORMAL, "Error while fetching system information (%s)", err.Error())
		os.Exit(ExitError)
	}

	fmt.Printf("Adobe Experience manager:\n")
	fmt.Printf("- Version: %s\n", sysInfo.Instance.AdobeExperienceManager)
	output.PrintListing("Run mode", sysInfo.Instance.RunModes, false)
	fmt.Printf("- Up since: %s\n\n", sysInfo.Instance.InstanceUpSince)

	fmt.Printf("Repostitory\n")
	fmt.Printf("- Version: %s\n", sysInfo.Repository.ApacheJackrabbitOak)
	fmt.Printf("- Repository size: %s\n", sysInfo.Repository.RepositorySize)
	fmt.Printf("- Storage location: %s\n", sysInfo.Repository.FileDataStore)
	fmt.Printf("- Nodes total: %s\n", sysInfo.EstimatedNodeCounts.Total)
	fmt.Printf("- Tags total: %s\n", sysInfo.EstimatedNodeCounts.Tags)
	fmt.Printf("- Authorizables: %s\n", sysInfo.EstimatedNodeCounts.Authorizables)
	fmt.Printf("- Pages: %s\n\n", sysInfo.EstimatedNodeCounts.Pages)

	fmt.Printf("Operating system:\n")
	fmt.Printf("- %s\n\n", sysInfo.SystemInformation.CurrentOS)

	fmt.Printf("Maintenance Tasks:\n")
	tasks := sysInfo.MaintenanceTasks.(map[string]interface{})
	for key, value := range tasks {
		output.PrintListing(key, value.(string), true)
	}

	warns := sysInfo.HealthChecks.(map[string]interface{})

	fmt.Printf("Current health checks:\n")
	if len(warns) > 0 {
		for key, check := range warns {
			output.PrintListing(key, check.(string), true)
		}
	} else {
		fmt.Printf("No Warnings or errors.\n\n")
	}

	agents := sysInfo.ReplicationAgents.(map[string]interface{})
	fmt.Printf("Replication Agents:\n")
	if len(agents) > 0 {
		for key, check := range agents {
			output.PrintListing(key, check.(string), true)
		}
	}
	agents = sysInfo.DistributionAgents.(map[string]interface{})
	fmt.Printf("Distribution Agents:\n")
	if len(agents) > 0 {
		for key, check := range agents {
			output.PrintListing(key, check.(string), false)
		}
	}

}
