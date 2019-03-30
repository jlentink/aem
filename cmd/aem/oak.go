package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	oakRunRepository = "https://repo1.maven.org/maven2/org/apache/jackrabbit/oak-run/%s/oak-run-%s.jar"
	oakRepoPath      = "/repository/segmentstore"
	oakJar           = "oak-run-%s.jar"
)

func newOak() oak {
	return oak{
		pStructure: newProjectStructure(),
		http:       new(httpRequests),
	}
}

type oak struct {
	pStructure projectStructure
	http       *httpRequests
}

func (o *oak) versionForAem(version string) string {
	switch version {
	case "6.0":
		return "1.0.42"
	case "6.1":
		return "1.2.31"
	case "6.2":
		return "1.4.24"
	case "6.3":
		return "1.6.16"
	case "6.4":
		return "1.8.12"
	default:
		return config.OakVersion
	}
}

func (o *oak) getVersion(aemVersion, version string) string {

	if len(aemVersion) > 0 {
		version = o.versionForAem(aemVersion)
	}

	if len(version) == 0 {
		exitProgram("No oak version selected. Use --aem or --oak to define version.")
	}

	oakJar := o.pStructure.createBinDir() + fmt.Sprintf(oakJar, version)
	if !o.pStructure.exists(oakJar) {
		fmt.Printf("Downloading oak version %s...\n", version)
		o.http.downloadFile(oakJar, fmt.Sprintf(oakRunRepository, version, version), "", "", false)
	}
	return oakJar
}

func (o *oak) execute(oakPath string, args []string) {
	param := config.OakOptions
	param = append(param, []string{"-jar", oakPath}...)
	param = append(param, args...)
	cmd := exec.Command("java", param...)
	cmd.Dir = o.pStructure.getInstanceDirLocation()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	exitFatal(err, "Error while unpacking AEM jar.")

}
