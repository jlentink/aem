package main

import (
	"errors"
	"fmt"
	"github.com/pborman/getopt/v2"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	user2 "os/user"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func newStartCommand() commandStart {
	return commandStart{
		projectStructure: newProjectStructure(),
		utility:          new(utility),
		httpClient:       new(httpRequests),
		forceDownload:    false,
		forGround:        false,
		root:             false,
		name:             configDefaultInstance,
	}
}

type commandStart struct {
	projectStructure projectStructure
	utility          *utility
	httpClient       *httpRequests
	forceDownload    bool
	forGround        bool
	root             bool
	name             string
	instance         aemInstanceConfig
}

func (s *commandStart) Execute(args []string) {
	s.getOpt(args)
	s.instance = s.utility.getInstanceByName(s.name)

	s.checkRoot()

	s.projectStructure.createInstanceDir()
	s.projectStructure.writeGitIgnoreFile()
	runDir := s.projectStructure.getRunDirLocation(s.instance)
	s.getJarFile()

	if !s.utility.Exists(runDir) {
		s.executeUnpack()
		s.renameUnpackFolder()
	}
	s.getAdditionPackages()
	s.cleanupDeprecated()

	s.executeStart(s.instance)
}

func (s *commandStart) checkRoot() {
	if s.root {
		return
	}
	user, err := user2.Current()
	exitFatal(err, "Could not get current user.")

	if (strings.ToLower(runtime.GOOS) == "linux" || strings.ToLower(runtime.GOOS) == "darwin") && user.Uid == "0" {
		exitProgram("You are running as root. Please use an other non-root user.\n")
	}
}

func (s *commandStart) getJarFile() {
	jarFile := s.projectStructure.getJarFileLocation()
	if !s.utility.Exists(jarFile) || s.forceDownload {
		if len(config.AemJar) > 8 && ("https://" == config.AemJar[0:8] || "http://" == config.AemJar[0:7]) {

			url, err := url.Parse(config.AemJar)
			exitFatal(err, "Could not parse url (%s). Please check url format.", config.AemJar)
			password, _ := url.User.Password()

			fmt.Printf("Downloading AEM jar...\n")
			err = s.httpClient.downloadFile(s.projectStructure.getJarFileLocation(), s.utility.returnURLString(url), url.User.Username(), password, s.forceDownload)
			exitFatal(err, "Error occurred during downloading aem jar.")
		} else {
			if _, err := os.Stat(config.AemJar); os.IsNotExist(err) {
				fmt.Printf("Could not find file at %s\n", config.AemJar)
				os.Exit(1)
			}
			s.utility.copy(config.AemJar, jarFile)
		}
	} else {
		fmt.Printf("Found AEM jar. Skipping download...\n")
	}

}

func (s *commandStart) executeUnpack() {
	cmd := exec.Command("java", "-jar", s.projectStructure.getJarFileLocation(), "-unpack")
	cmd.Dir = s.projectStructure.getInstanceDirLocation()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	exitFatal(err, "Error while unpacking AEM jar.")
}

func (s *commandStart) renameUnpackFolder() {
	s.projectStructure.rename(s.projectStructure.getUnpackDirLocation(), s.projectStructure.getRunDirLocation(s.instance))
}

func (s *commandStart) findJar() string {
	files, err := ioutil.ReadDir(s.projectStructure.getAppDirLocation(s.instance))
	exitFatal(err, "Could not found app dir.")
	r, _ := regexp.Compile("(.*)\\.jar")
	for _, file := range files {
		if r.MatchString(file.Name()) {
			return file.Name()
		}
	}
	exitProgram("Could not find main jar in application")
	return ""
}

func (s *commandStart) executeStart(instance aemInstanceConfig) {
	javaOptions := config.JVMOpts
	javaOptions = append(javaOptions, instance.JVMOptions...)
	if instance.Debug {
		if len(config.JVMDebugOptions) > 0 {
			javaOptions = append(javaOptions, config.JVMDebugOptions...)
		}

		if len(instance.JVMDebugOptions) > 0 {
			javaOptions = append(javaOptions, instance.JVMDebugOptions...)
		}
	}

	javaOptions = append(javaOptions,
		"-Dsling.run.modes="+instance.Type+",crx3,crx3tar",
		"-jar", "app/"+s.findJar(), "start",
		"-c", s.projectStructure.getRunDirLocation(instance),
		"-i", "launchpad",
		"-p", fmt.Sprintf("%d", instance.Port),
		"-Dsling.properties=conf/sling.properties")

	cmd := exec.Command("java", javaOptions...)
	cmd.Dir = s.projectStructure.getRunDirLocation(instance)
	err := errors.New("")

	if !s.forGround {
		err = cmd.Start()
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	}

	exitFatal(err, "Error starting AEM ")
	fmt.Printf("Starting AEM with (pid: %d)\n", cmd.Process.Pid)
	ioutil.WriteFile(s.projectStructure.getPidFileLocation(s.instance), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
}

func (s *commandStart) getAdditionPackages() {
	installDir := s.projectStructure.createAemInstallDir(s.instance)

	for _, additional := range config.AdditionalPackages {
		packagename := path.Base(additional)
		if !s.utility.Exists(installDir+"/"+packagename) || s.forceDownload {
			url, err := url.Parse(additional)
			exitFatal(err, "Could not parse url (%s). Please check url format.", additional)

			password, _ := url.User.Password()
			fmt.Printf("Downloading additional (%s)\n", additional)
			s.httpClient.downloadFile(installDir+"/"+packagename, s.utility.returnURLString(url), url.User.Username(), password, s.forceDownload)
		} else {
			logrus.Debugf("Found package %s. Skipping download...", packagename)
		}
	}
}

func (s *commandStart) cleanupDeprecated() {
	files, err := ioutil.ReadDir(s.projectStructure.getAemInstallDirLocation(s.instance))
	packages := make([]string, 0)
	exitFatal(err, "Could not find install dir")

	for _, pkg := range config.AdditionalPackages {
		packages = append(packages, path.Base(pkg))
	}
	for _, file := range files {
		if !s.utility.inSliceString(packages, file.Name()) {
			fmt.Printf("Removing package not in config anymore %s\n", file.Name())
			os.Remove(s.projectStructure.appendSlash(s.projectStructure.getAemInstallDirLocation(s.instance)) + file.Name())
		}
	}
}

func (s *commandStart) getOpt(args []string) {
	getopt.FlagLong(&s.forceDownload, "download", 'd', "Force new download")
	getopt.FlagLong(&s.forGround, "for-ground", 'f', "Don't detach aem")
	getopt.FlagLong(&s.name, "name", 'n', "Instance to start. (default: "+configDefaultInstance+")")
	getopt.FlagLong(&s.root, "root", 'r', "Allow root")
	getopt.CommandLine.Parse(args)
}
