package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"strings"
)

func NewCopyPackageCommand() commandCopyPackage {
	return commandCopyPackage{
		From:             "",
		ToGroup:          "",
		ToName:           "",
		utility:          new(Utility),
		projectStructure: new(projectStructure),
		showLog:          false,
		forceDownload:    false,
		http:             new(HttpRequests),
	}
}

type commandCopyPackage struct {
	From             string
	ToName           string
	ToGroup          string
	Type             string
	Role             string
	Name             string
	Packages         string
	utility          *Utility
	showLog          bool
	projectStructure *projectStructure
	http             *HttpRequests
	forceDownload    bool
}

func (p *commandCopyPackage) Execute(args []string) {
	u := Utility{}
	p.getOpt(args)
	toInstances := make([]AEMInstanceConfig, 0)

	fromInstance := u.getInstanceByName(p.From)
	if len(p.ToName) > 0 {
		toInstances = append(toInstances, u.getInstanceByName(p.ToName))
	} else if len(p.ToGroup) > 0 {
		toInstances = u.getInstanceByGroup(p.ToGroup)
	}

	authorPackages := p.http.getListForInstance(fromInstance)
	packages := strings.Split(p.Packages, ",")

	for _, currentPackage := range authorPackages {
		for _, cPackage := range packages {
			packageName, packageVersion := u.packageNameVersion(cPackage)
			if packageName == currentPackage.Name && packageVersion == currentPackage.Version {
				fmt.Printf("\n%s (%s)\n", packageName, packageVersion)
				fmt.Println("downloading...")
				err, _ := p.http.downloadPackage(fromInstance, currentPackage, p.forceDownload)
				if nil == err {
					for _, toInstance := range toInstances {
						fmt.Printf("\nuploading to %s...\n", toInstance.Name)
						crx, err := p.http.uploadPackage(toInstance, currentPackage, true, true)
						exitFatal(err, "Error installing package.")
						if p.showLog {
							fmt.Printf("%s\n", crx.Response.Data.Log)
						}
					}
				}
			}
		}
	}
}

func (p *commandCopyPackage) getOpt(args []string) {
	getopt.FlagLong(&p.From, "from-name", 'f', "Pull content from")
	getopt.FlagLong(&p.ToName, "to-name", 't', "Push package to instance")
	getopt.FlagLong(&p.ToGroup, "to-group", 'g', "Push package to group")
	getopt.FlagLong(&p.Packages, "package", 'p', "Packages (multiple use comma separated list.)")
	getopt.FlagLong(&p.showLog, "log", 'l', "Show AEM log output")
	getopt.FlagLong(&p.forceDownload, "force-download", 'd', "Force new download")
	getopt.CommandLine.Parse(args)
}
