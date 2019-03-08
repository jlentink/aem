package main

import (
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/pborman/getopt/v2"
	"os"
	"strings"
)

func newPackageListCommand() packagesListCommand {
	return packagesListCommand{
		Download:   false,
		Type:       "author",
		Role:       "development",
		Name:       configDefaultInstance,
		Ascending:  true,
		Descending: false,
		SortBy:     "Package",
		http:       new(httpRequests),
	}
}

type packagesListCommand struct {
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

func (p *packagesListCommand) Execute(args []string) {
	u := utility{}
	p.getOpt(args)

	instance := u.getInstanceByName(p.Name)
	packages := p.http.getListForInstance(instance)
	sortFields := make([]string, 0)

	if strings.Index(p.SortBy, ",") == -1 {
		sortFields = append(sortFields, p.SortBy)
	} else {
		sortFields = strings.Split(p.SortBy, ",")
	}

	u.sortPackages(packages, p.Descending, sortFields)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Package", "Version", "Size", "Build", "Created"})
	for _, packageObject := range packages {
		t.AppendRow(table.Row{packageObject.Name, packageObject.Version, humanize.Bytes(packageObject.Size), packageObject.BuildCount, u.unixTime(packageObject.Created)})
	}
	t.Render()
}

func (p *packagesListCommand) getOpt(args []string) {
	getopt.FlagLong(&p.Name, "name", 'n', "Name of instance to download from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&p.Descending, "descending", 'd', "Sort Descending")
	getopt.FlagLong(&p.SortBy, "sort", 's', "Sort comma separated list")
	getopt.CommandLine.Parse(args)

	if p.Descending {
		p.Ascending = false
	}
}

func (p *packagesListCommand) help() string {
	return "ssadsad"
}
