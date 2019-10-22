package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandIndex struct {
	verbose bool
	cmd     *cobra.Command
}

func (c *commandIndex) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "indexes",
		Aliases: []string{},
		Short:   "index commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	commands = []Command{
		&commandIndexes{},
		&commandReindex{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandIndex) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandIndex) run(cmd *cobra.Command, args []string) {
	cmd.Help() // nolint: errcheck
}
