package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newPackageRebuildCommand() commandPackageRebuild {
	return commandPackageRebuild{
		From:             configDefaultInstance,
		utility:          new(utility),
		projectStructure: new(projectStructure),
		forceDownload:    false,
		Package:          "",
		http:             new(httpRequests),
	}
}

type commandPackageRebuild struct {
	From             string
	Package          string
	utility          *utility
	projectStructure *projectStructure
	http             *httpRequests
	forceDownload    bool
}

func (p *commandPackageRebuild) Execute(args []string) {
	p.getOpt(args)

	fromInstance := p.utility.getInstanceByName(p.From)

	pkgs := make([]packageDescription, 0)
	if len(p.Package) > 0 {
		pkgs = p.utility.pkgsFromString(fromInstance, p.Package)
	} else {
		pkgs = p.utility.pkgPicker(fromInstance)
	}

	for _, pkg := range pkgs {
		fmt.Printf("Build: %s\n", pkg.Name)
		p.http.buildPackage(fromInstance, pkg)
	}
}

func (p *commandPackageRebuild) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from-name", 'f', "Rebuild package on instance (Default: "+configDefaultInstance+")")
	getopt.FlagLong(&p.Package, "package", 'p', "Define package package:version (no interactive mode)")
	getopt.CommandLine.Parse(args)
}
