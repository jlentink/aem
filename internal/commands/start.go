package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/dispatcher"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandStart struct {
	verbose       bool
	instanceName  string
	groupName     string
	allowRoot     bool
	foreground    bool
	forceDownload bool
	ignorePid     bool
}

func (c *commandStart) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "start",
		Short:  "Start Adobe Experience Manager instance",
		PreRun: c.preRun,
		Run:    c.run,
	}

	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to start")
	cmd.Flags().StringVarP(&c.groupName, "group", "g", "", "Instance to start")
	cmd.Flags().BoolVarP(&c.forceDownload, "download", "d", false, "Force re-download")
	cmd.Flags().BoolVarP(&c.foreground, "foreground", "f", false, "on't detach aem from current tty")
	cmd.Flags().BoolVarP(&c.allowRoot, "allow-root", "r", false, "Allow to start as root user (UID: 0)")
	cmd.Flags().BoolVarP(&c.ignorePid, "ignore-pid", "p", false, "Ignore existing PID file and start AEM")

	return cmd
}

func (c *commandStart) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandStart) run(cmd *cobra.Command, args []string) {
	if !aem.AllowUserStart(c.allowRoot) {
		output.Print(output.NORMAL, "You are starting aem as a root. This is not allowed. override with: --allow-root\n")
		os.Exit(ExitError)
	}

	cnf, instances, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName, c.groupName, []string{aem.RoleAuthor, aem.RolePublisher, aem.RoleDispatcher})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	for _, currentInstance := range instances {
		if currentInstance.InstanceOf([]string{aem.RoleAuthor, aem.RolePublisher}) {
			err := aem.FullStart(currentInstance, c.ignorePid, c.forceDownload, c.foreground, cnf, nil)
			if err != nil {
				os.Exit(ExitError)
			}

 		} else if currentInstance.InstanceOf([]string{aem.RoleDispatcher}){
			err := dispatcher.Start(currentInstance, cnf, c.foreground)
			if err != nil {
				fmt.Printf("%s", err.Error())
				os.Exit(ExitError)
			}
		}
	}
}
