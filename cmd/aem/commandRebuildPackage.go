package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewRebuildPackageCommand() commandRebuildPackage {
	return commandRebuildPackage{
		From:             CONFIG_DEFAULT_INSTANCE,
		utility:          new(Utility),
		projectStructure: new(projectStructure),
		forceDownload:    false,
		Package:          "",
		http:             new(HttpRequests),
	}
}

type commandRebuildPackage struct {
	From             string
	Package          string
	utility          *Utility
	projectStructure *projectStructure
	http             *HttpRequests
	forceDownload    bool
}

func (p *commandRebuildPackage) Execute(args []string) {
	p.getOpt(args)

	fromInstance := p.utility.getInstanceByName(p.From)

	pkgs := make([]PackageDescription, 0)
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

func (p *commandRebuildPackage) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from-name", 'f', "Rebuild package on instance (Default: "+CONFIG_DEFAULT_INSTANCE+")")
	getopt.FlagLong(&p.Package, "package", 'p', "Define package package:version (no interactive mode)")
	getopt.CommandLine.Parse(args)
}
