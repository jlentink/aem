package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pborman/getopt/v2"
	"os"
)

func NewListBundlesCommand() bundlesListCommand {
	return bundlesListCommand{
		name: CONFIG_DEFAULT_INSTANCE,
		http: new(HttpRequests),
	}
}

type bundlesListCommand struct {
	name   string
	http   *HttpRequests
}

func (c *bundlesListCommand) Execute(args []string) {
	u := Utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)

	bundles := c.http.listBundles(instance)

	fmt.Printf("%s\n", bundles.Status)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Name", "Version", "Category", "Status"})
	for _, bundle := range bundles.Data {
		t.AppendRow(table.Row{bundle.ID, bundle.Name, bundle.Version, bundle.Category, bundle.State})
	}
	t.Render()
}

func (c *bundlesListCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.CommandLine.Parse(args)
}
