package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/version"
	"github.com/spf13/cobra"
	"os"
)

type commandVersion struct {
	verbose bool
	minimal bool
}

func (c *commandVersion) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Short:  "Show version of aemcli",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().BoolVarP(&c.minimal, "minimal", "m", false, "Show the minimal version information")
	return cmd
}

func (c *commandVersion) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
}

func (c *commandVersion) run(cmd *cobra.Command, args []string) {
	output.Print(output.NORMAL, version.DisplayVersion(c.verbose, c.minimal))
	os.Exit(ExitNormal)
}
