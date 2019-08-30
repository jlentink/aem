package commands

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/bundle"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandBundleList struct {
	verbose      bool
	instanceName string
}

func (c *commandBundleList) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bundle-list",
		Short:   "List bundles",
		Aliases: []string{"blist"},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to list bundles off")
	return cmd
}

func (c *commandBundleList) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)
}

func (c *commandBundleList) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(EXIT_ERROR)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Symbolic name", "Version", "State"})
	bundles, err := bundle.List(i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not get list from instance: %s", err.Error())
		os.Exit(EXIT_ERROR)

	}
	for _, cBundle := range bundles {
		t.AppendRow([]interface{}{cBundle.ID, cBundle.Name, cBundle.SymbolicName, cBundle.Version, cBundle.State})
	}
	t.Render()
}
