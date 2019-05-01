package main

import (
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pborman/getopt/v2"
	"os"
	"strings"
)

type commandPackagesList struct {
	Download   bool
	Verbose    bool
	Type       string
	Role       string
	Name       string
	Ascending  bool
	Descending bool
	SortBy     string
	http       *httpRequests
}

func (c *commandPackagesList) Init() {
	c.Download = false
	c.Type = "author"
	c.Role = "development"
	c.Name = configDefaultInstance
	c.Ascending = true
	c.Descending = false
	c.SortBy = "Package"
	c.http = new(httpRequests)
}

func (c *commandPackagesList) readConfig() bool {
	return true
}

func (c *commandPackagesList) GetCommand() []string {
	return []string{"package-list", "packages-list"}
}

func (c *commandPackagesList) GetHelp() string {
	return "List packages on server."
}

func (c *commandPackagesList) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.Name)
	packages := c.http.getListForInstance(instance)
	sortFields := make([]string, 0)

	if strings.Contains(c.SortBy, ",") {
		sortFields = append(sortFields, c.SortBy)
	} else {
		sortFields = strings.Split(c.SortBy, ",")
	}

	u.sortPackages(packages, c.Descending, sortFields)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Package", "Version", "Size", "Build", "Created"})
	for _, packageObject := range packages {
		t.AppendRow(table.Row{packageObject.Name, packageObject.Version, humanize.Bytes(packageObject.Size), packageObject.BuildCount, u.unixTime(packageObject.Created)})
	}
	t.Render()
}

func (c *commandPackagesList) getOpt(args []string) {
	getopt.FlagLong(&c.Name, "name", 'n', "Name of instance to download from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.Descending, "descending", 'd', "Sort Descending")
	getopt.FlagLong(&c.SortBy, "sort", 's', "Sort comma separated list")
	getopt.CommandLine.Parse(args)

	if c.Descending {
		c.Ascending = false
	}
}
