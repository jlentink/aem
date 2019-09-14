package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"

	"github.com/spf13/cobra"
)

type commandBuild struct {
	verbose bool
}

func (c *commandBuild) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "build",
		Short:   "Build application",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	return cmd
}

func (c *commandBuild) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandBuild) run(cmd *cobra.Command, args []string) {
	getConfig() // nolint: errcheck
	aem.GetConfig() // nolint: errcheck
	aem.BuildProject() // nolint: errcheck
}
