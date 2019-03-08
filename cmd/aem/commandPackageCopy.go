package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"strings"
)

func newPackageCopyCommand() commandPackageCopy {
	return commandPackageCopy{
		From:             configDefaultInstance,
		ToGroup:          "",
		ToName:           "",
		utility:          new(utility),
		projectStructure: new(projectStructure),
		showLog:          false,
		forceDownload:    false,
		http:             new(httpRequests),
	}
}

type commandPackageCopy struct {
	From             string
	ToName           string
	ToGroup          string
	Type             string
	Role             string
	Name             string
	Packages         string
	utility          *utility
	showLog          bool
	projectStructure *projectStructure
	http             *httpRequests
	forceDownload    bool
}

func (p *commandPackageCopy) matchPackage(pkgString string, pkg packageDescription) bool {
	packageName, packageVersion := p.utility.packageNameVersion(pkgString)
	if packageName == pkg.Name && packageVersion == pkg.Version {
		return true
	}
	return false
}

func (p *commandPackageCopy) matchInstancePackages(packageString string, fromInstance aemInstanceConfig) []packageDescription {
	selectedPkg := make([]packageDescription, 0)
	authorPackages := p.http.getListForInstance(fromInstance)
	packagesStrings := strings.Split(p.Packages, ",")
	for _, pkgStr := range packagesStrings {
		for _, currentAuthorPackage := range authorPackages {
			if p.matchPackage(pkgStr, currentAuthorPackage) {
				selectedPkg = append(selectedPkg, currentAuthorPackage)
			}
		}
	}

	return selectedPkg
}

func (p *commandPackageCopy) getPackages(packageString string, fromInstance aemInstanceConfig) []packageDescription {
	if len(packageString) > 0 {
		return p.matchInstancePackages(packageString, fromInstance)
	}
	pkgPicker := newPackagePicker()
	return pkgPicker.picker(fromInstance)
}

func (p *commandPackageCopy) Execute(args []string) {
	u := utility{}
	p.getOpt(args)

	fromInstance := u.getInstanceByName(p.From)
	toInstances := u.getInstance(p.ToName, p.ToGroup)

	copyPackages := p.getPackages(p.Packages, fromInstance)

	for _, pkg := range copyPackages {
		fmt.Printf("Downloading %s from %s\n", pkg.Name, fromInstance.URL())
		_, err := p.http.downloadPackage(fromInstance, pkg, p.forceDownload)
		exitFatal(err, "Could not download %s from %s\n", pkg.Name, fromInstance.URL())
		for _, toInstance := range toInstances {
			fmt.Printf("Uploading %s to %s\n", pkg.Name, toInstance.URL())
			crx, err := p.http.uploadPackage(toInstance, pkg, true, true)
			exitFatal(err, "Error installing package %s on %s.", pkg.Name, fromInstance.URL())
			if p.showLog {
				fmt.Printf("%s\n", crx.Response.Data.Log)
			}
		}

	}
}

func (p *commandPackageCopy) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from-name", 'f', "Pull content from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&p.ToName, "to-name", 't', "Push package to instance")
	getopt.FlagLong(&p.ToGroup, "to-group", 'g', "Push package to group")
	getopt.FlagLong(&p.Packages, "package", 'p', "Packages (multiple use comma separated list.)")
	getopt.FlagLong(&p.showLog, "log", 'l', "Show AEM log output")
	getopt.FlagLong(&p.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
