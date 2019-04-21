package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"strings"
)

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

func (c *commandPackageCopy) Init() {
	c.From = configDefaultInstance
	c.ToGroup = ""
	c.ToName = ""
	c.utility = new(utility)
	c.projectStructure = new(projectStructure)
	c.showLog = false
	c.forceDownload = false
	c.http = new(httpRequests)
}

func (c *commandPackageCopy) readConfig() bool {
	return true
}

func (c *commandPackageCopy) GetCommand() []string {
	return []string{"package-copy"}
}

func (c *commandPackageCopy) GetHelp() string {
	return "Copy packages from one instance to another."
}

func (c *commandPackageCopy) matchPackage(pkgString string, pkg packageDescription) bool {
	packageName, packageVersion := c.utility.packageNameVersion(pkgString)
	if packageName == pkg.Name && packageVersion == pkg.Version {
		return true
	}
	return false
}

func (c *commandPackageCopy) matchInstancePackages(packageString string, fromInstance aemInstanceConfig) []packageDescription {
	selectedPkg := make([]packageDescription, 0)
	authorPackages := c.http.getListForInstance(fromInstance)
	packagesStrings := strings.Split(c.Packages, ",")
	for _, pkgStr := range packagesStrings {
		for _, currentAuthorPackage := range authorPackages {
			if c.matchPackage(pkgStr, currentAuthorPackage) {
				selectedPkg = append(selectedPkg, currentAuthorPackage)
			}
		}
	}

	return selectedPkg
}

func (c *commandPackageCopy) getPackages(packageString string, fromInstance aemInstanceConfig) []packageDescription {
	if len(packageString) > 0 {
		return c.matchInstancePackages(packageString, fromInstance)
	}
	pkgPicker := newPackagePicker()
	return pkgPicker.picker(fromInstance)
}

func (c *commandPackageCopy) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	fromInstance := u.getInstanceByName(c.From)
	toInstances := u.getInstance(c.ToName, c.ToGroup)

	copyPackages := c.getPackages(c.Packages, fromInstance)

	for _, pkg := range copyPackages {
		fmt.Printf("Downloading %s from %s\n", pkg.Name, fromInstance.URL())
		_, err := c.http.downloadPackage(fromInstance, pkg, c.forceDownload)
		exitFatal(err, "Could not download %s from %s\n", pkg.Name, fromInstance.URL())
		for _, toInstance := range toInstances {
			fmt.Printf("Uploading %s to %s\n", pkg.Name, toInstance.URL())
			crx, err := c.http.uploadPackage(toInstance, pkg, true, true)
			exitFatal(err, "Error installing package %s on %s.", pkg.Name, fromInstance.URL())
			if c.showLog {
				fmt.Printf("%s\n", crx.Response.Data.Log)
			}
		}

	}
}

func (c *commandPackageCopy) getOpt(args []string) {
	getopt.FlagLong(&c.From, "from-name", 'f', "Pull content from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.ToName, "to-name", 't', "Push package to instance")
	getopt.FlagLong(&c.ToGroup, "to-group", 'g', "Push package to group")
	getopt.FlagLong(&c.Packages, "package", 'p', "Packages (multiple use comma separated list.)")
	getopt.FlagLong(&c.showLog, "log", 'l', "Show AEM log output")
	getopt.FlagLong(&c.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
