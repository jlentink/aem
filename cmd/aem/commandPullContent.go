package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func NewPullContentCommand() commandPullContent {
	return commandPullContent{
		From:             "",
		To:               CONFIG_DEFAULT_INSTANCE,
		utility:          new(Utility),
		projectStructure: NewProjectStructure(),
		forceDownload:    false,
		http:             new(HttpRequests),
	}
}

type commandPullContent struct {
	From             string
	To               string
	Type             string
	Role             string
	Name             string
	utility          *Utility
	projectStructure projectStructure
	forceDownload    bool
	http             *HttpRequests
}

func (p *commandPullContent) Execute(args []string) {
	u := Utility{}
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
				err, _ := p.http.downloadPackage(fromInstance, contentPackage, p.forceDownload)
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
