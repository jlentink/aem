package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type commandPullContent struct {
	From             string
	To               string
	Type             string
	Role             string
	Name             string
	utility          *utility
	projectStructure projectStructure
	forceDownload    bool
	http             *httpRequests
}

func (c *commandPullContent) Init() {
	c.From = ""
	c.To = configDefaultInstance
	c.utility = new(utility)
	c.projectStructure = newProjectStructure()
	c.forceDownload = false
	c.http = new(httpRequests)
}

func (c *commandPullContent) readConfig() bool {
	return true
}

func (c *commandPullContent) GetCommand() []string {
	return []string{"pull-content"}
}

func (c *commandPullContent) GetHelp() string {
	return "Pull content packages from instance to..."
}

func (c *commandPullContent) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	fromInstance := u.getInstanceByName(c.From)
	toInstance := u.getInstanceByName(c.To)

	packages := c.http.getListForInstance(fromInstance)
	for _, currentPackage := range config.ContentPackages {
		fmt.Printf("Content package: %s\n", currentPackage)
		packageName, packageVersion := u.packageNameVersion(currentPackage)
		contentPackages := u.filterPackages(packages, packageName)

		for _, contentPackage := range contentPackages {
			if packageVersion == contentPackage.Version {
				_, err := c.http.downloadPackage(fromInstance, contentPackage, c.forceDownload)
				if nil == err {
					fmt.Printf("Uploading package...")
					crx, err := c.http.uploadPackage(toInstance, contentPackage, true, true)
					fmt.Printf("%s\n", crx.Response.Data.Log)
					exitFatal(err, "Error installing package.")
				}
			}
		}
	}
}

func (c *commandPullContent) getOpt(args []string) {
	getopt.FlagLong(&c.From, "from-name", 'f', "Pull content from")
	getopt.FlagLong(&c.To, "to-name", 't', "Push content to")
	getopt.FlagLong(&c.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
