package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandPackageRebuild struct {
	verbose      bool
	instanceName string
	packageName  string
}

func (c *commandPackageRebuild) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "package-rebuild",
		Short:  "package rebuild",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to rebuild package on")
	cmd.Flags().StringVarP(&c.packageName, "package", "p", ``, "Package to rebuild")
	return cmd
}

func (c *commandPackageRebuild) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageRebuild) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if len(c.packageName) <= 0 {
		rebuildPackageSearch(i)
		os.Exit(ExitNormal)
	}

	rebuildPackage(i, c.packageName)

}
