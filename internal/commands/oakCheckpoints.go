package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/cli/oak"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandOakCheckpoints struct {
	verbose      bool
	instanceName string
	aemVersion   string
	oakVersion   string
}

func (c *commandOakCheckpoints) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "checkpoints",
		Aliases: []string{},
		Short:   "Run oak checkpoints",
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().StringVarP(&c.aemVersion, "aem", "a", ``, "Version of AEM to use oak-run on. (use matching AEM version of oak-run)")
	cmd.Flags().StringVarP(&c.oakVersion, "oak", "o", ``, "Define version of oak-run to use")
	cmd.MarkFlagRequired("name")
	return cmd
}

func (c *commandOakCheckpoints) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandOakCheckpoints) run(cmd *cobra.Command, args []string) {
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
	oakArgs := []string{"checkpoints", instancePath + oak.RepoPath}
	oak.Execute(path, aem.Cnf.OakOptions, oakArgs) // nolint: errcheck

}
