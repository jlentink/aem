package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"regexp"
)

func newPackageInstallCommand() commandPackageInstall {
	return commandPackageInstall{
		From:             "",
		ToGroup:          "",
		ToName:           configDefaultInstance,
		u:                new(utility),
		projectStructure: newProjectStructure(),
		showLog:          false,
		forceDownload:    false,
		yes:              false,
		noInstall:        false,
		http:             new(httpRequests),
	}
}

type commandPackageInstall struct {
	From             string
	ToName           string
	ToGroup          string
	Type             string
	Role             string
	Name             string
	Package          string
	yes              bool
	noInstall        bool
	u                *utility
	showLog          bool
	projectStructure projectStructure
	forceDownload    bool
	http             *httpRequests
}

func (p *commandPackageInstall) Execute(args []string) {

	p.getOpt(args)
	toInstances := make([]aemInstanceConfig, 0)

	if len(p.ToName) > 0 {
		toInstances = append(toInstances, p.u.getInstanceByName(p.ToName))
	} else if len(p.ToGroup) > 0 {
		toInstances = p.u.getInstanceByGroup(p.ToGroup)
	}

	if p.u.Exists(p.Package) {
		r, _ := regexp.Compile(regexZipFile)
		if r.MatchString(p.Package) {
			description := p.u.zipToPackage(p.Package)
			if confirm("Do you want to install %s (%s)\n", p.yes, description.Name, description.Version) {
				for _, i := range toInstances {
					fmt.Printf("Installing to %s\n", i.Name)
					p.http.uploadPackage(i, description, true, !p.noInstall)
				}
			}
		} else {
			exitProgram("Package must be zipfile. (%s)", p.Package)
		}
	} else {
		exitProgram("Could not find package (%s)", p.Package)
	}
}

func (p *commandPackageInstall) getOpt(args []string) {
	getopt.FlagLong(&p.ToName, "to-name", 't', "Push package to instance (default: "+configDefaultInstance+")")
	getopt.FlagLong(&p.ToGroup, "to-group", 'g', "Push package to group")
	getopt.FlagLong(&p.Package, "package", 'p', "Package to install (path to file)")
	getopt.FlagLong(&p.yes, "yes", 'y', "Skip confirmation")
	getopt.FlagLong(&p.noInstall, "no-install", 'n', "Do not install package")
	getopt.CommandLine.Parse(args)
}
