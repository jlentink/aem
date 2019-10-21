package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandBundle struct {
	verbose      bool
	instanceName string
	aemVersion   string
	oakVersion   string
	cmd          *cobra.Command
}

func (c *commandBundle) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "bundle",
		Aliases: []string{},
		Short:   "Bundle commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}


	commands = []Command{
		&commandBundleList{},
		&commandBundleInstall{},
		&commandBundleStart{},
		&commandBundelStop{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandBundle) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandBundle) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
