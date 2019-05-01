package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"regexp"
)

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

func (c *commandPackageInstall) Init() {
	c.From = ""
	c.ToGroup = ""
	c.ToName = configDefaultInstance
	c.u = new(utility)
	c.projectStructure = newProjectStructure()
	c.showLog = false
	c.forceDownload = false
	c.yes = false
	c.noInstall = false
	c.http = new(httpRequests)
}

func (c *commandPackageInstall) readConfig() bool {
	return true
}

func (c *commandPackageInstall) GetCommand() []string {
	return []string{"package-install"}
}

func (c *commandPackageInstall) GetHelp() string {
	return "Install packages from local zip."
}

func (c *commandPackageInstall) Execute(args []string) {
	c.getOpt(args)
	toInstances := make([]aemInstanceConfig, 0)

	if len(c.ToName) > 0 {
		c.ToName = c.u.getDefaultInstance(c.ToName)
		toInstances = append(toInstances, c.u.getInstanceByName(c.ToName))
	} else if len(c.ToGroup) > 0 {
		toInstances = c.u.getInstanceByGroup(c.ToGroup)
	}

	if c.u.Exists(c.Package) {
		r, _ := regexp.Compile(regexZipFile)
		if r.MatchString(c.Package) {
			description := c.u.zipToPackage(c.Package)
			if confirm("Do you want to install %s (%s)\n", c.yes, description.Name, description.Version) {
				for _, i := range toInstances {
					fmt.Printf("Installing to %s\n", i.Name)
					c.http.uploadPackage(i, description, true, !c.noInstall)
				}
			}
		} else {
			exitProgram("Package must be zipfile. (%s)", c.Package)
		}
	} else {
		exitProgram("Could not find package (%s)", c.Package)
	}
}

func (c *commandPackageInstall) getOpt(args []string) {
	getopt.FlagLong(&c.From, "to-name",
		't', "Install package to instance (default: "+c.u.getDefaultInstance(configDefaultInstance)+")")
	getopt.FlagLong(&c.ToGroup, "to-group", 'g', "Install package to group")
	getopt.FlagLong(&c.Package, "package", 'p', "Package to install (path to file)")
	getopt.FlagLong(&c.yes, "yes", 'y', "Skip confirmation")
	getopt.FlagLong(&c.noInstall, "no-install", 'n', "Do not install package")
	getopt.CommandLine.Parse(args)
}
