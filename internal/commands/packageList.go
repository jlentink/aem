package commands

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jlentink/aem/internal/aem"
	_package "github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandPackageList struct {
	verbose      bool
	instanceName string
}

func (c *commandPackageList) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List packages",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.MarkFlagRequired("name")
	return cmd
}

func (c *commandPackageList) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageList) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Version", "Group", "Size", "Last modified"})
	pkgs, err := _package.PackageList(*i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not get list from instance: %s", err.Error())
		os.Exit(ExitError)

	}
	for i, cP := range pkgs {
		e := output.UnixTime(cP.LastModified)
		tt := ""
		if e != nil {
			//nolint
			tt = fmt.Sprintf("%s", e.UTC())
		}
		t.AppendRow([]interface{}{i, cP.Name, cP.Version, cP.Group, humanize.Bytes(uint64(cP.Size)), tt})
	}
	t.Render()
}
