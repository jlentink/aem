package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandActivation struct {
	verbose bool
	cmd     *cobra.Command
}

func (c *commandActivation) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "activation",
		Aliases: []string{},
		Short:   "Activation commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	commands = []Command{
		&commandReplicationPage{},
		&commandActivateTree{},
		&commandInvalidate{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandActivation) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandActivation) run(cmd *cobra.Command, args []string) {
	c.cmd.Help() // nolint: errcheck
}
