package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/replication"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandActivateTree struct {
	verbose          bool
	instanceName     string
	instanceGroup    string
	path             string
	ignoreDeactivate bool
	onlyModified     bool
}

func (c *commandActivateTree) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "activate-tree",
		Short:  "Activate Tree",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to (de)activate page on")
	cmd.Flags().StringVarP(&c.instanceGroup, "group", "g", ``, "Instance group to (de)activate page on")
	cmd.Flags().StringVarP(&c.path, "path", "p", ``, "Path to (de)activate")
	cmd.Flags().BoolVarP(&c.ignoreDeactivate, "ignore-deactivated", "d", false, "Ignore Deactivated")
	cmd.Flags().BoolVarP(&c.onlyModified, "only-modified", "o", false, "Only Modified")
	return cmd
}

func (c *commandActivateTree) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)
}

func (c *commandActivateTree) run(cmd *cobra.Command, args []string) {
	_, is, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName, c.instanceGroup, []string{aem.RoleAuthor})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(EXIT_ERROR)
	}

	if len(c.path) == 0 {
		output.Printf(output.NORMAL, "no path provided")
		os.Exit(EXIT_ERROR)
	}

	for _, i := range is {
		body, err := replication.ActivateTree(&i, c.path, c.ignoreDeactivate, c.onlyModified)
		if err != nil {
			output.Printf(output.NORMAL, "Could not activate tree: %s", err.Error())
			os.Exit(EXIT_ERROR)
		}
		output.Printf(output.NORMAL, "\U00002705 tree activated: %s\n", c.path)
		output.Printf(output.VERBOSE, "%s", body)
	}
}
