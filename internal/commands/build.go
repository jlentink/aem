package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"os"

	"github.com/spf13/cobra"
)

type commandBuild struct {
	verbose         bool
	versionOnly     bool
	productionBuild bool
	skipTests       bool
	skipFrontend	bool
	skipCheckStyle	bool
}

func (c *commandBuild) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "build",
		Short:   "Build application",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	cmd.Flags().BoolVarP(&c.productionBuild, "production-build", "B", false,
		"Flush after deploy")
	cmd.Flags().BoolVarP(&c.skipTests, "skip-tests", "t", false,
		"Skip tests")
	cmd.Flags().BoolVarP(&c.skipFrontend, "skip-frontend", "F", false,
		"Skip frontend build")
	cmd.Flags().BoolVarP(&c.skipCheckStyle, "skip-checkstyle", "c", false,
		"Skip checkstyle")
	cmd.Flags().BoolVarP(&c.versionOnly, "version", "V", false,
		"Don't build version only.")

	return cmd
}

func (c *commandBuild) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")

	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandBuild) run(cmd *cobra.Command, args []string) {
	getConfig()     // nolint: errcheck
	aem.GetConfig() // nolint: errcheck

	if c.versionOnly {
		err := aem.SetBuildVersion(c.productionBuild)
		if err != nil {
			output.Printf(output.NORMAL, "\U0000274C Version failed...")
			os.Exit(1)
		}
		os.Exit(0)
	}

	err := aem.BuildProject(c.productionBuild, c.skipTests, c.skipCheckStyle, c.skipFrontend)
	if err != nil {
		output.Printf(output.NORMAL, "\U0000274C Build failed...")
		os.Exit(1)
	}
}
