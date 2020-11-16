package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandCloudManager struct {
	verbose bool
	cmd     *cobra.Command
}

func (c *commandCloudManager) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "cloudmanager",
		Aliases: []string{"cm"},
		Short:   "Adobe Cloud manager commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	commands = []Command{
		&commandCloudManagerPush{},
		&commandCloudManagerGitAuthenticate{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandCloudManager) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCloudManager) run(cmd *cobra.Command, args []string) {
	cmd.Help() // nolint: errcheck
}
