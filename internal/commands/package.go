package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
)

type commandPackage struct {
	verbose      bool
	instanceName string
	aemVersion   string
	oakVersion   string
	cmd          *cobra.Command
}

func (c *commandPackage) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "package",
		Aliases: []string{},
		Short:   "Package commands",
		PreRun:  c.preRun,
		Run:     c.run,
	}


	commands = []Command{
		&commandPackageDownload{},
		&commandPackageCopy{},
		&commandPackageInstall{},
		&commandPackageUpload{},
		&commandPackageRebuild{},
		&commandPackageList{},
	}
	for _, cmd := range commands {
		c.cmd.AddCommand(cmd.setup())
	}
	return c.cmd
}

func (c *commandPackage) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackage) run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
