package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newPullContentCommand() commandPullContent {
	return commandPullContent{
		From:             "",
		To:               configDefaultInstance,
		utility:          new(utility),
		projectStructure: newProjectStructure(),
		forceDownload:    false,
		http:             new(httpRequests),
	}
}

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

func (p *commandPullContent) Execute(args []string) {
	u := utility{}
	p.getOpt(args)

	fromInstance := u.getInstanceByName(p.From)
	toInstance := u.getInstanceByName(p.To)

	packages := p.http.getListForInstance(fromInstance)
	for _, currentPackage := range config.ContentPackages {
		fmt.Printf("Content package: %s\n", currentPackage)
		packageName, packageVersion := u.packageNameVersion(currentPackage)
		contentPackages := u.filterPackages(packages, packageName)

		for _, contentPackage := range contentPackages {
			if packageVersion == contentPackage.Version {
				_, err := p.http.downloadPackage(fromInstance, contentPackage, p.forceDownload)
				if nil == err {
					fmt.Printf("Uploading package...")
					crx, err := p.http.uploadPackage(toInstance, contentPackage, true, true)
					fmt.Printf("%s\n", crx.Response.Data.Log)
					exitFatal(err, "Error installing package.")
				}
			}
		}
	}
}

func (p *commandPullContent) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from-name", 'f', "Pull content from")
	getopt.FlagLong(&p.To, "to-name", 't', "Push content to")
	getopt.FlagLong(&p.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
