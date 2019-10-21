package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandZsh struct {
	verbose bool
}

func (c *commandZsh) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "zsh",
		Short:   "Generate zsh completion for aemCLI",
		Aliases: []string{"zsh"},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	return cmd
}

func (c *commandZsh) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandZsh) run(cmd *cobra.Command, args []string) {
	rootCmd.GenZshCompletion(os.Stdout) // nolint: errcheck
}
