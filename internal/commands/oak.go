package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandOak struct {
	verbose      bool
	instanceName string
	aemVersion   string
	oakVersion   string
	cmd          *cobra.Command
}

func (c *commandOak) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "oak",
		Aliases: []string{},
		Short:   "Oak commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}


	commands = []Command{
		&commandOakCheck{},
		&commandOakCheckpoints{},
		&commandOakCompact{},
		&commandOakConsole{},
		&commandOakExplore{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandOak) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandOak) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
