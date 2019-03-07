package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"regexp"
)

func NewInstallPackageCommand() commandInstallPackage {
	return commandInstallPackage{
		From:             "",
		ToGroup:          "",
		ToName:           "",
		u:                new(Utility),
		projectStructure: new(projectStructure),
		showLog:          false,
		forceDownload:    false,
		yes:              false,
		noInstall:        false,
		http:             new(HttpRequests),
	}
}

type commandInstallPackage struct {
	From             string
	ToName           string
	ToGroup          string
	Type             string
	Role             string
	Name             string
	Package          string
	yes              bool
	noInstall        bool
	u                *Utility
	showLog          bool
	projectStructure *projectStructure
	forceDownload    bool
	http             *HttpRequests
}

func (p *commandInstallPackage) Execute(args []string) {
	p.getOpt(args)
	toInstances := make([]AEMInstanceConfig, 0)

	if len(p.ToName) > 0 {
		toInstances = append(toInstances, p.u.getInstanceByName(p.ToName))
	} else if len(p.ToGroup) > 0 {
		toInstances = p.u.getInstanceByGroup(p.ToGroup)
	}

	if p.u.Exists(p.Package) {
		r, _ := regexp.Compile(REGEX_ZIP)
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

func (p *commandInstallPackage) getOpt(args []string) {
	getopt.FlagLong(&p.ToName, "to-name", 't', "Push package to instance")
	getopt.FlagLong(&p.ToGroup, "to-group", 'g', "Push package to group")
	getopt.FlagLong(&p.Package, "package", 'p', "Package to install")
	getopt.FlagLong(&p.yes, "yes", 'y', "Skip confirmation")
	getopt.FlagLong(&p.noInstall, "no-install", 'n', "Do not install package")
	getopt.CommandLine.Parse(args)
}
