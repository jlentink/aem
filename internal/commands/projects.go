package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/cli/cachedir"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandProjects struct {
	verbose bool
}

func (c *commandProjects) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "projects",
		Short:  "List know projects",
		PreRun: c.preRun,
		Run:    c.run,
	}
	return cmd
}

func (c *commandProjects) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
	RegisterProject()
}

func (c *commandProjects) run(cmd *cobra.Command, args []string) {
	projects := cachedir.RegisteredProjects()
	for _, project := range projects {
		fmt.Printf(" * %s - %s\n", project.Name, project.Path)
	}
	if len(projects) == 0 {
		output.Printf(output.NORMAL, "\U00002049 No registered projects found.")
		os.Exit(ExitError)
	}
}
