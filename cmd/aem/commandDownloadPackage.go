package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewDownloadPackageCommand() commandDownloadPackage {
	return commandDownloadPackage{
		From:             CONFIG_DEFAULT_INSTANCE,
		utility:          new(Utility),
		projectStructure: NewProjectStructure(),
		forceDownload:    false,
		Package:          "",
		http:             new(HttpRequests),
	}
}

type commandDownloadPackage struct {
	From             string
	To               string
	Type             string
	Role             string
	Name             string
	Package          string
	utility          *Utility
	projectStructure projectStructure
	http             *HttpRequests
	forceDownload    bool
}

func (p *commandDownloadPackage) Execute(args []string) {
	p.getOpt(args)

	fromInstance := p.utility.getInstanceByName(p.From)

	pkgs := make([]PackageDescription, 0)
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

func (p *commandDownloadPackage) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from", 'f', "Pull content from (default: "+CONFIG_DEFAULT_INSTANCE + ")")
	getopt.FlagLong(&p.Package, "package", 'p', "Define package package:version (no interactive mode)")
	getopt.FlagLong(&p.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
