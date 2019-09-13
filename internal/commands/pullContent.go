package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandPullContent struct {
	verbose        bool
	instanceName   string
	toInstanceName string
	build          bool
}

func (c *commandPullContent) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pull-content",
		Short:   "Pull content in from instance via packages",
		Aliases: []string{"cpull"},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "from", "f", ``, "Instance to copy from")
	cmd.Flags().StringVarP(&c.toInstanceName, "to", "t", aem.GetDefaultInstanceName(), "Destination Instance")
	cmd.Flags().BoolVarP(&c.build, "build", "b", false, "Build before download")
	return cmd
}

func (c *commandPullContent) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPullContent) run(cmd *cobra.Command, args []string) {
	cnf, f, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	_, t, errorString, err := getConfigAndInstance(c.toInstanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	output.Printf(output.NORMAL, "\U0001F69A %s => %s\n", f.Name, t.Name)
	for _, cPkg := range cnf.ContentPackages {
		if c.build {
			rebuildPackage(f, cPkg)
		}
		pd, err := pkg.DownloadWithName(f, cPkg)
		if err != nil {
			output.Printf(output.NORMAL, "\U0000274C Issue while fetching content page: %s\n", err.Error())
		}
		path, err := project.GetLocationForPackage(pd)
		if err != nil {
			output.Printf(output.NORMAL, errorString, err.Error())
			os.Exit(ExitError)
		}

		crx, err := pkg.Upload(*t, path, true, true)
		if err != nil {
			output.Printf(output.NORMAL, errorString, err.Error())
			os.Exit(ExitError)
		}
		output.Printf(output.VERBOSE, "%s", crx.Response.Data.Log.Text)
	}

}
