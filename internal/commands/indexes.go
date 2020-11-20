package commands

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/indexes"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandIndexes struct {
	verbose      bool
	instanceName string
}

func (c *commandIndexes) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "indexes",
		Aliases: []string{"list"},
		Short:  "Show indexes on instance",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	return cmd
}

func (c *commandIndexes) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandIndexes) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	searchIndexes, err := indexes.GetIndexes(i)
	if err != nil {
		output.Printf(output.NORMAL, "Error retrieving indexes. (%s)", err.Error())
		os.Exit(ExitError)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Reindex count", "info"})
	for i, index := range searchIndexes {
		t.AppendRow([]interface{}{i, index.Name, index.ReindexCount, index.Info})
	}

	t.Render()

}
