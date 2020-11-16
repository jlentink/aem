package commands

import (
	"bufio"
	"fmt"
	"github.com/jlentink/aem/internal/aem/cloudmanager"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandCloudManagerGitAuthenticate struct {
	verbose        bool
	instanceName   string
	toInstanceName string
	toGroup        string
	install        bool
	force          bool
	cPackage       []string
}

func (c *commandCloudManagerGitAuthenticate) setup() *cobra.Command {
	c.install = false
	c.force = false
	cmd := &cobra.Command{
		Use:    "git-auth",
		Short:  "Add git credentials.",
		PreRun: c.preRun,
		Run:    c.run,
	}

	return cmd
}

func (c *commandCloudManagerGitAuthenticate) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCloudManagerGitAuthenticate) run(cmd *cobra.Command, args []string) {
	cnf, err := getConfig()
	if err != nil {
		output.Printf(output.NORMAL, "Error getting config", err.Error())
		os.Exit(ExitError)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	usr, _ := reader.ReadString('\n')

	fmt.Print("Enter password: ")
	pwd, _ := reader.ReadString('\n')

	err = cloudmanager.GitSetAuthentication(usr, pwd, cnf)
	if err != nil {
		output.Printf(output.NORMAL, "Error setting the credentials. %s", err.Error())
		os.Exit(ExitError)
	}

}
