package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandBash struct {
	verbose bool
}

func (c *commandBash) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "bash",
		Short:  "Generate bash completion for aemCLI",
		PreRun: c.preRun,
		Run:    c.run,
	}
	return cmd
}

func (c *commandBash) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandBash) run(cmd *cobra.Command, args []string) {
	rootCmd.GenBashCompletion(os.Stdout) // nolint: errcheck
}
