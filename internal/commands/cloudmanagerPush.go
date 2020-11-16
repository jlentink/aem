package commands

import (
	"github.com/jlentink/aem/internal/aem/cloudmanager"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandCloudManagerPush struct {
	verbose        bool
	instanceName   string
	toInstanceName string
	toGroup        string
	install        bool
	force          bool
	cPackage       []string
}

func (c *commandCloudManagerPush) setup() *cobra.Command {
	c.install = false
	c.force = false
	cmd := &cobra.Command{
		Use:    "push",
		Short:  "Push current branch to Adobe Cloud manager GIT",
		PreRun: c.preRun,
		Run:    c.run,
	}

	return cmd
}

func (c *commandCloudManagerPush) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCloudManagerPush) run(cmd *cobra.Command, args []string) {
	cnf, err := getConfig()
	if err != nil {
		output.Printf(output.NORMAL, "Error getting config", err.Error())
		os.Exit(ExitError)
	}

	err = cloudmanager.GitPush(cnf)
	if err != nil {
		output.Printf(output.NORMAL, "Error pushing to adobe: %s.", err.Error())
		os.Exit(ExitError)

	}

}
