package main

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Utility struct {
}

func (u *Utility) pkgsFromString(instance AEMInstanceConfig, pkgString string) []PackageDescription {
	http := new(HttpRequests)
	pkgs := http.getListForInstance(instance)
	wPkgs := strings.Split(pkgString, ",")
	selectedPkgs := make([]PackageDescription, 0)

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

func (u *Utility) readCmdLineInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[0 : len(input)-1]

}

func (u *Utility) pkgPicker(instance AEMInstanceConfig) []PackageDescription {
	http := new(HttpRequests)
	pkgs := http.getListForInstance(instance)
	pageSize := 20
	selected := make([]int64, 0)
	selectedPkgs := make([]PackageDescription, 0)
	writer := new(TableWriter)

	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Package", "Version"})
	t.SetPageSize(pageSize)
	t.SetOutputMirror(writer)

	for i, pkg := range pkgs {
		t.AppendRow(table.Row{i + 1, pkg.Name, pkg.Version})
	}

	t.Render()
	tables := writer.getTables()

	for i := 0; i < len(tables); i++ {
		fmt.Print(tables[i])

	choose:
		fmt.Printf("Selected %d\n", selected)
		fmt.Print("d: done selecting, q: quit, c: continue, package id ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = input[0 : len(input)-1]
		switch input {
		case "c":
			continue
		case "q":
			return make([]PackageDescription, 0)
		case "d":
			for _, selectedPkg := range selected {
				selectedPkgs = append(selectedPkgs, pkgs[selectedPkg])
			}
			return selectedPkgs
		default:
			r, _ := regexp.Compile("\\d")
			if r.MatchString(input) {
				id, _ := strconv.ParseInt(input, 10, 32)
				if int(id) < len(pkgs)-1 && id > 0 {
					if !u.inSliceInt64(selected, id) {
						selected = append(selected, id)
					}
				} else {
					fmt.Printf("Invalid id: %s\n", input)
					goto choose
				}
			} else {
				fmt.Printf("Unkown option: %s\n", input)
				goto choose
			}
			i = i - 1
		}
	}
	return pkgs
}

func (u *Utility) inSliceInt64(slice []int64, needle int64) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}

func (u *Utility) inSliceString(slice []string, needle string) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}

func (u *Utility) zipToPackage(filepath string) PackageDescription {
	projectStructure := NewProjectStructure()
	manifest := newManifestPackage()
	manifestValues := manifest.fromZip(filepath)
	description := PackageDescription{}
	name := manifestValues[ManifestLabelPackageName]
	version := manifestValues[ManifestLabelPackageVersion]
	
	if len(name) > 0 {
		description = PackageDescription{Name: name, Version: version, DownloadName: name + "-" + version + ".zip"}

		projectStructure.createDirForPackage(description)
		projectStructure.copy(filepath, projectStructure.getLocationForPackage(description))

	} else {
		exitProgram("Could not find name for package")
	}

	return description
}

func (u *Utility) Exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func (u *Utility) returnUrlString(url *url.URL) string {
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

func (u *Utility) getInstanceByName(instanceName string) AEMInstanceConfig {
	for _, instance := range config.Instances {
		if instanceName == instance.Name {
			return instance
		}
	}
	fmt.Printf("Instance %s is not defined.\n", instanceName)
	os.Exit(1)
	return AEMInstanceConfig{}
}

func (u *Utility) getInstanceByGroup(groupName string) []AEMInstanceConfig {
	instances := make([]AEMInstanceConfig, 0)
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

func (u *Utility) copy(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	exitFatal(err, "Could not read file at %s", sourceFile)

	err = ioutil.WriteFile(destinationFile, input, 0644)
	exitFatal(err, "Could not write file to %s", destinationFile)
}

func (u *Utility) filterPackages(packages []PackageDescription, filter string) []PackageDescription {
	filteredList := make([]PackageDescription, 0)

	for _, currentPackage := range packages {
		if currentPackage.Name == filter {
			filteredList = append(filteredList, currentPackage)
		}
	}

	return filteredList
}

func (u *Utility) sortPackages(packages []PackageDescription, ascending bool, additionalSearchFields []string) []PackageDescription {
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
			if compare > 0 {
				return true
			} else {
				return false
			}
		} else {
			if compare <= 0 {
				return true
			} else {
				return false
			}

		}
	})

	return packages
}

func (u *Utility) unixTime(timestamp int64) time.Time {
	tm := time.Unix(timestamp/1000, 0)
	return tm.UTC()
}

func (u *Utility) packageNameVersion(packageName string) (string, string) {
	packageVersion := ""

	packageName = strings.TrimSpace(packageName)
	if strings.Index(packageName, ":") != -1 {
		splits := strings.Split(packageName, ":")
		packageName = splits[0]
		packageVersion = splits[1]
	}

	return packageName, packageVersion
}

func (u *Utility) getInstance(name, group string) []AEMInstanceConfig {
	instances := make([]AEMInstanceConfig, 0)

	if len(group) > 0 {
		instances = u.getInstanceByGroup(group)
	} else if len(name) > 0 {
		instances = append(instances, u.getInstanceByName(name))
	}

	return instances
}
