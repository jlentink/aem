package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/replication"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandReplicationPage struct {
	verbose       bool
	instanceName  string
	instanceGroup string
	path          string
	activate      bool
	deactivate    bool
}

func (c *commandReplicationPage) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "page",
		Short:  "Activate / Deactivate page",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to (de)activate page on")
	cmd.Flags().StringVarP(&c.instanceGroup, "group", "g", ``, "Instance group to (de)activate page on")
	cmd.Flags().StringVarP(&c.path, "path", "p", ``, "Path to (de)activate")
	cmd.Flags().BoolVarP(&c.activate, "activate", "a", false, "Activate page")
	cmd.Flags().BoolVarP(&c.deactivate, "deactivate", "d", false, "Deactivate")
	return cmd
}

func (c *commandReplicationPage) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandReplicationPage) run(cmd *cobra.Command, args []string) {
	_, is, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName, c.instanceGroup, []string{aem.RoleAuthor})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	for _, i := range is {
		i := i
		if c.activate {
			_, err := replication.Activate(&i, c.path)
			if err != nil {
				output.Printf(output.NORMAL, "Could not activate page: %s", err.Error())
				os.Exit(ExitError)
			}
			output.Printf(output.NORMAL, "\U00002705 Page activated: %s\n", c.path)
		} else if c.deactivate {
			_, err := replication.Deactivate(&i, c.path)
			if err != nil {
				output.Printf(output.NORMAL, "Could not activate page: %s", err.Error())
				os.Exit(ExitError)
			}
			output.Printf(output.NORMAL, "\U00002705 Page deactivated: %s\n", c.path)
		}
	}
}
