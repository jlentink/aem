package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandInit struct {
	verbose bool
	force   bool
}

func (c *commandInit) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "init",
		Long:   "Init new project in current directory by writing aem.toml. Edit it before starting you project. Run this file on the same level as the root pom.xml",
		Short:  "Init new project",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().BoolVarP(&c.force, "force", "f", false, "Force override of current configuration")
	return cmd
}

func (c *commandInit) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)
}

func (c *commandInit) run(cmd *cobra.Command, args []string) {
	if aem.ConfigExists() && !c.force {
		output.Print(output.NORMAL, "Instance file already exists. (use force to override)\n")
		os.Exit(ExitError)
	}

	output.Printf(output.NORMAL, "Writing sample file.")
	aem.WriteConfigFile()
	os.Exit(ExitNormal)

}
