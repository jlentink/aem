package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandPackageRebuild struct {
	From             string
	Package          string
	utility          *utility
	projectStructure *projectStructure
	http             *httpRequests
	forceDownload    bool
}

func (c *commandPackageRebuild) Init() {
	c.From = configDefaultInstance
	c.utility = new(utility)
	c.projectStructure = new(projectStructure)
	c.forceDownload = false
	c.Package = ""
	c.http = new(httpRequests)
}

func (c *commandPackageRebuild) readConfig() bool {
	return true
}

func (c *commandPackageRebuild) GetCommand() []string {
	return []string{"package-rebuild"}
}

func (c *commandPackageRebuild) GetHelp() string {
	return "Rebuild package on AEM instance."
}

func (c *commandPackageRebuild) Execute(args []string) {
	c.getOpt(args)
	c.From = c.utility.getDefaultInstance(c.From)
	fromInstance := c.utility.getInstanceByName(c.From)

	if len(c.Package) > 0 {
		pkgs := c.utility.pkgsFromString(fromInstance, c.Package)
		c.buildPackage(fromInstance, pkgs)
	} else {
		pkgPicker := newPackagePicker()
		pkgs := pkgPicker.picker(fromInstance)
		c.buildPackage(fromInstance, pkgs)
	}

}

func (c *commandPackageRebuild) buildPackage(instance aemInstanceConfig, packages []packageDescription) {
	for _, pkg := range packages {
		fmt.Printf("Build: %s\n", pkg.Name)
		c.http.buildPackage(instance, pkg)
	}
}

func (c *commandPackageRebuild) getOpt(args []string) {
	getopt.FlagLong(&c.From, "from-name",
		'f', "Rebuild package on instance (default: "+c.utility.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.Package, "package", 'p', "Define package package:version (no interactive mode)")
	getopt.CommandLine.Parse(args)
}
