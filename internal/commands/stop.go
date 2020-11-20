package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/dispatcher"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandStop struct {
	verbose      bool
	instanceName string
	groupName     string
}

func (c *commandStop) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "stop",
		Short:  "stop Adobe Experience Manager instance",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().StringVarP(&c.groupName, "group", "g", "", "Instance to start")
	return cmd
}

func (c *commandStop) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandStop) run(cmd *cobra.Command, args []string) {

	cnf, instances, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName, c.groupName, []string{aem.RoleAuthor, aem.RolePublisher, aem.RoleDispatcher})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	for _, currentInstance := range instances {
		if currentInstance.InstanceOf([]string{aem.RoleAuthor, aem.RolePublisher}) {
			err = aem.Stop(currentInstance)
			if err != nil {
				output.Printf(output.NORMAL, "Could not stop instance. (%s)", err.Error())
				os.Exit(ExitError)
			}
		} else if currentInstance.InstanceOf([]string{aem.RoleDispatcher}){
			err := dispatcher.Stop(currentInstance, cnf)
			if err != nil {
				output.Printf(output.NORMAL, "Could not stop instance. (%s)", err.Error())
				os.Exit(ExitError)
			}
		}
	}
}
