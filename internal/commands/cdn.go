package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandCdn struct {
	verbose bool
	cmd     *cobra.Command
}

func (c *commandCdn) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "cdn",
		Aliases: []string{},
		Short:   "cdn commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	commands = []Command{
		&commandCdnPurgeURL{},
		&commandCdnPurgeService{},
		&commandCdnPurgeTag{},
		&commandCdnCredentials{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandCdn) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCdn) run(cmd *cobra.Command, args []string) {
	cmd.Help() // nolint: errcheck
}
