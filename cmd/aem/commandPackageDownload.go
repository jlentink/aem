package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newPackageDownloadCommand() commandPackageDownload {
	return commandPackageDownload{
		From:             configDefaultInstance,
		utility:          new(utility),
		projectStructure: newProjectStructure(),
		forceDownload:    false,
		Package:          "",
		http:             new(httpRequests),
	}
}

type commandPackageDownload struct {
	From             string
	To               string
	Type             string
	Role             string
	Name             string
	Package          string
	utility          *utility
	projectStructure projectStructure
	http             *httpRequests
	forceDownload    bool
}

func (p *commandPackageDownload) Execute(args []string) {
	p.getOpt(args)

	fromInstance := p.utility.getInstanceByName(p.From)

	pkgs := make([]packageDescription, 0)
	if len(p.Package) > 0 {
		pkgs = p.utility.pkgsFromString(fromInstance, p.Package)
	} else {
		pkgs = p.utility.pkgPicker(fromInstance)
	}

	for _, pkg := range pkgs {
		fmt.Printf("Download: %s\n", pkg.Name)
		p.http.downloadPackage(fromInstance, pkg, p.forceDownload)
	}
}

func (p *commandPackageDownload) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from", 'f', "Pull content from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&p.Package, "package", 'p', "Define package package:version (no interactive mode)")
	getopt.FlagLong(&p.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
