package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandPackageInstall struct {
	verbose      bool
	instanceName string
	packageName  []string
}

func (c *commandPackageInstall) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "install",
		Short:  "Install uploaded package",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to rebuild package on")
	cmd.Flags().StringArrayVarP(&c.packageName, "package", "p", []string{}, "Package to install (allowed multiple)")
	return cmd
}

func (c *commandPackageInstall) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageInstall) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if len(c.packageName) <= 0 {
		installPackageSearch(i)
		os.Exit(ExitNormal)
	}

	for _, packageName := range c.packageName {
		installPackage(i, packageName)
	}

}
