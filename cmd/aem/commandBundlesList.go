package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pborman/getopt/v2"
	"os"
)

type commandBundleList struct {
	name string
	http *httpRequests
}

func (c *commandBundleList) Init() {
	c.name = configDefaultInstance
	c.http = new(httpRequests)
}

func (c *commandBundleList) readConfig() bool {
	return true
}

func (c *commandBundleList) GetCommand() []string {
	return []string{"bundle-list", "bundles-list"}
}

func (c *commandBundleList) GetHelp() string {
	return "List bundle on instance."
}

func (c *commandBundleList) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)

	bundles := c.http.listBundles(instance)

	fmt.Printf("%s\n", bundles.Status)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Name", "Symbolic Name", "Version", "Category", "Status"})
	for _, bundle := range bundles.Data {
		t.AppendRow(table.Row{bundle.ID, bundle.Name, bundle.SymbolicName, bundle.Version, bundle.Category, bundle.State})
	}
	t.Render()
}

func (c *commandBundleList) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
