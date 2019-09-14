package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/cli/oak"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandOakExplore struct {
	verbose      bool
	instanceName string
	aemVersion   string
	oakVersion   string
}

func (c *commandOakExplore) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "oak-explorer",
		Aliases: []string{},
		Short:   "Run oak explorer",
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().StringVarP(&c.aemVersion, "aem", "a", ``, "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	cmd.Flags().StringVarP(&c.oakVersion, "oak", "o", ``, "Define version of oak-run to use")

	return cmd
}

func (c *commandOakExplore) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandOakExplore) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	instancePath, err := project.GetRunDirLocation(*i)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	oak.SetDefaultVersion(aem.Cnf.OakVersion)
	path, _ := oak.Get(i.GetVersion(), aem.Cnf.OakVersion)
	oakArgs := []string{"explore", instancePath + oak.RepoPath}
	oak.Execute(path, aem.Cnf.OakOptions, oakArgs) // nolint: errcheck

}
