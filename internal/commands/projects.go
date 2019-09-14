package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandProjects struct {
	verbose      bool
	instanceName string
}

func (c *commandProjects) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "projects",
		Short:   "List know projects",
		PreRun:  c.preRun,
		Run:     c.run,
	}
	return cmd
}

func (c *commandProjects) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandProjects) run(cmd *cobra.Command, args []string) {
	p, err := project.HomeDir()
	if err != nil {
		output.Printf(output.NORMAL, "Could not find homedir: %s", err.Error())
		os.Exit(ExitError)
	}
	projects := ReadRegisteredProjects(p)
	for _, project := range projects.Project {
		fmt.Printf(" * %s - %s\n", project.Name, project.Path)
	}

}
