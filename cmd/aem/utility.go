package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"
)

type utility struct {
}

func (u *utility) pkgsFromString(instance aemInstanceConfig, pkgString string) []packageDescription {
	http := new(httpRequests)
	pkgs := http.getListForInstance(instance)
	wPkgs := strings.Split(pkgString, ",")
	selectedPkgs := make([]packageDescription, 0)

	for _, pkg := range pkgs {
		for _, wPkg := range wPkgs {
			name, version := u.packageNameVersion(wPkg)
			if strings.ToLower(pkg.Name) == strings.ToLower(name) && pkg.Version == version {
				selectedPkgs = append(selectedPkgs, pkg)
			}
		}
	}

	return selectedPkgs
}

func (u *utility) readCmdLineInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[0 : len(input)-1]

}

func (u *utility) zipToPackage(filepath string) packageDescription {
	projectStructure := newProjectStructure()
	manifest := newManifestPackage()
	manifestValues := manifest.fromZip(filepath)
	description := packageDescription{}
	name := manifestValues[ManifestLabelPackageName]
	version := manifestValues[ManifestLabelPackageVersion]

	if len(name) <= 0 {
		fmt.Println("Could not find name in manifest using zip to identify name.")

		zipName := path.Base(filepath)
		r, _ := regexp.Compile(regexPackageZip)
		if r.MatchString(zipName) {
			matches := r.FindAllStringSubmatch(zipName, -1)
			name = matches[0][1]
			version = matches[0][2]
		}
	}

	if len(name) > 0 {
		description = packageDescription{Name: name, Version: version, DownloadName: name + "-" + version + ".zip"}

		projectStructure.createDirForPackage(description)
		projectStructure.copy(filepath, projectStructure.getLocationForPackage(description))

	} else {
		exitProgram("Could not find name for package")
	}

	return description
}

func (u *utility) Exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func (u *utility) returnURLString(url *url.URL) string {
	query := url.RawQuery
	fragment := url.Fragment

	if len(query) > 0 {
		query = "?" + query
	}

	if len(fragment) > 0 {
		fragment = "#" + fragment
	}

	return url.Scheme + "://" + url.Host + url.Path + query + fragment
}

func (u *utility) getInstanceByName(instanceName string) aemInstanceConfig {
	for _, instance := range config.Instances {
		if instanceName == instance.Name {
			return instance
		}
	}
	fmt.Printf("Instance %s is not defined.\n", instanceName)
	os.Exit(1)
	return aemInstanceConfig{}
}

func (u *utility) getInstanceByGroup(groupName string) []aemInstanceConfig {
	instances := make([]aemInstanceConfig, 0)
	for _, instance := range config.Instances {
		if groupName == instance.Group {
			instances = append(instances, instance)
		}
	}
	if len(instances) == 0 {
		fmt.Printf("Instance group %s is not defined.\n", groupName)
		os.Exit(1)
	}
	return instances
}

func (u *utility) copy(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	exitFatal(err, "Could not read file at %s", sourceFile)

	err = ioutil.WriteFile(destinationFile, input, 0644)
	exitFatal(err, "Could not write file to %s", destinationFile)
}

func (u *utility) filterPackages(packages []packageDescription, filter string) []packageDescription {
	filteredList := make([]packageDescription, 0)

	for _, currentPackage := range packages {
		if currentPackage.Name == filter {
			filteredList = append(filteredList, currentPackage)
		}
	}

	return filteredList
}

func (u *utility) sortPackages(packages []packageDescription, ascending bool, additionalSearchFields []string) []packageDescription {
	sort.Slice(packages, func(i, j int) bool {
		from := ""
		to := ""

		for _, field := range additionalSearchFields {
			switch field {
			case "name", "package":
				from = from + packages[i].Name
				to = to + packages[j].Name
			case "version":
				from = from + packages[i].Version
				to = to + packages[j].Version
			default:
				logrus.Debugf("Unknown field %s. Skipping", field)
			}
		}

		from = strings.ToLower(from)
		to = strings.ToLower(to)

		compare := strings.Compare(from, to)

		if ascending {
			//nolint
			if compare > 0 {
				return true
			}
			return false
		}
		if compare <= 0 {
			return true
		}
		return false
	})

	return packages
}

func (u *utility) unixTime(timestamp int64) time.Time {
	tm := time.Unix(timestamp/1000, 0)
	return tm.UTC()
}

func (u *utility) packageNameVersion(packageName string) (string, string) {
	packageVersion := ""

	packageName = strings.TrimSpace(packageName)

	if strings.Contains(packageName, ":") {
		splits := strings.Split(packageName, ":")
		packageName = splits[0]
		packageVersion = splits[1]
	}

	return packageName, packageVersion
}

func (u *utility) getInstance(name, group string) []aemInstanceConfig {
	instances := make([]aemInstanceConfig, 0)

	if len(group) > 0 {
		instances = u.getInstanceByGroup(group)
	} else if len(name) > 0 {
		instances = append(instances, u.getInstanceByName(name))
	}

	return instances
}
