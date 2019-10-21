package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandShellCompletion struct {
	verbose      bool
	instanceName string
	aemVersion   string
	oakVersion   string
	cmd          *cobra.Command
}

func (c *commandShellCompletion) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "shell",
		Aliases: []string{},
		Short:   "Shell completion commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}


	commands = []Command{
		&commandZsh{},
		&commandBash{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandShellCompletion) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandShellCompletion) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
