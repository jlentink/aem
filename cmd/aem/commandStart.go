package main

import (
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

func (c *commandStart) Init() {
	c.projectStructure = newProjectStructure()
	c.utility = new(utility)
	c.httpClient = new(httpRequests)
	c.forceDownload = false
	c.forGround = false
	c.root = false
	c.name = configDefaultInstance
}

func (c *commandStart) readConfig() bool {
	return true
}

func (c *commandStart) GetCommand() []string {
	return []string{"start"}
}

func (c *commandStart) GetHelp() string {
	return "Start an Adobe Experience Manager instance."
}

func (c *commandStart) checkEnvironmentName() {
	envName := os.Getenv(aemEnvName)
	if c.name == configDefaultInstance && len(envName) > 0 {
		fmt.Printf("Found env variable changing instance name to %s.\n", envName)
		c.name = envName
	}
}

func (c *commandStart) Execute(args []string) {
	c.getOpt(args)
	c.checkEnvironmentName()
	c.instance = c.utility.getInstanceByName(c.name)

	c.checkRoot()

	c.projectStructure.createInstanceDir()
	c.projectStructure.writeGitIgnoreFile()
	runDir := c.projectStructure.getRunDirLocation(c.instance)
	c.getJarFile()
	c.writeLicense()

	if !c.utility.Exists(runDir) {
		c.executeUnpack()
		c.renameUnpackFolder()
	}
	c.getAdditionPackages()
	c.cleanupDeprecated()

	//c.executeStart(c.instance)
}

func (c *commandStart) checkRoot() {
	if c.root {
		return
	}
	user, err := user2.Current()
	exitFatal(err, "Could not get current user.")

	if (strings.ToLower(runtime.GOOS) == "linux" || strings.ToLower(runtime.GOOS) == "darwin") && user.Uid == "0" {
		exitProgram("You are running as root. Please use an other non-root user.\n")
	}
}

func (c *commandStart) getJarFile() {
	jarFile := c.projectStructure.getJarFileLocation()
	if !c.utility.Exists(jarFile) || c.forceDownload {
		if len(config.AemJar) > 8 && ("https://" == config.AemJar[0:8] || "http://" == config.AemJar[0:7]) {
			fmt.Printf("Downloading AEM jar...\n")
			err := c.httpClient.downloadFile(c.projectStructure.getJarFileLocation(), config.AemJar, config.AemJarUsername, config.AemJarPassword, c.forceDownload)
			exitFatal(err, "Error occurred during downloading aem jar\n.")
		} else {
			if _, err := os.Stat(config.AemJar); os.IsNotExist(err) {
				fmt.Printf("Could not find file at %s\n", config.AemJar)
				os.Exit(1)
			}
			c.utility.copy(config.AemJar, jarFile)
		}
	} else {
		fmt.Printf("Found AEM jar. Skipping download...\n")
	}
}

func (c *commandStart) writeLicense() {
	if !c.utility.Exists(c.projectStructure.getLicenseLocation()) {
		license := fmt.Sprintf(
			"#Adobe Granite License Properties\n"+
				"license.product.name=Adobe Experience Manager\n"+
				"license.customer.name=%s\n"+
				"license.product.version=%s\n"+
				"license.downloadID=%s\n", config.LicenseCustomer, config.LicenseVersion, config.LicenseDownloadID)
		c.projectStructure.writeTextFile(c.projectStructure.getLicenseLocation(), license)
	} else {
		fmt.Println("License found skipping...")
	}
}

func (c *commandStart) executeUnpack() {
	cmd := exec.Command("java", "-jar", c.projectStructure.getJarFileLocation(), "-unpack")
	cmd.Dir = c.projectStructure.getInstanceDirLocation()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	exitFatal(err, "Error while unpacking AEM jar.")
}

func (c *commandStart) renameUnpackFolder() {
	c.projectStructure.rename(c.projectStructure.getUnpackDirLocation(), c.projectStructure.getRunDirLocation(c.instance))
}

func (c *commandStart) findJar() string {
	files, err := ioutil.ReadDir(c.projectStructure.getAppDirLocation(c.instance))
	exitFatal(err, "Could not found app dir.")
	r, _ := regexp.Compile(`(.*)\.jar`)
	for _, file := range files {
		if r.MatchString(file.Name()) {
			return file.Name()
		}
	}
	exitProgram("Could not find main jar in application")
	return ""
}

func (c *commandStart) executeStart(instance aemInstanceConfig) {
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
		"-jar", "app/"+c.findJar(), "start",
		"-c", c.projectStructure.getRunDirLocation(instance),
		"-i", "launchpad",
		"-p", fmt.Sprintf("%d", instance.Port),
		"-Dsling.properties=conf/sling.properties")

	cmd := exec.Command("java", javaOptions...)
	cmd.Dir = c.projectStructure.getRunDirLocation(instance)

	if !c.forGround {
		err := cmd.Start()
		exitFatal(err, "Error starting AEM ")
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		exitFatal(err, "Error starting AEM ")
	}

	fmt.Printf("Starting AEM with (pid: %d)\n", cmd.Process.Pid)
	ioutil.WriteFile(c.projectStructure.getPidFileLocation(c.instance), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
}

func (c *commandStart) getAdditionPackages() {
	installDir := c.projectStructure.createAemInstallDir(c.instance)

	for _, additional := range config.AdditionalPackages {
		packagename := path.Base(additional)
		if !c.utility.Exists(installDir+"/"+packagename) || c.forceDownload {
			url, err := url.Parse(additional)
			exitFatal(err, "Could not parse url (%s). Please check url format.", additional)

			password, _ := url.User.Password()
			fmt.Printf("Downloading additional (%s)\n", additional)
			c.httpClient.downloadFile(installDir+"/"+packagename, c.utility.returnURLString(url), url.User.Username(), password, c.forceDownload)
		} else {
			logrus.Debugf("Found package %s. Skipping download...", packagename)
		}
	}
}

func (c *commandStart) cleanupDeprecated() {
	files, err := ioutil.ReadDir(c.projectStructure.getAemInstallDirLocation(c.instance))
	sutil := new(sliceUtil)
	packages := make([]string, 0)
	exitFatal(err, "Could not find install dir")

	for _, pkg := range config.AdditionalPackages {
		packages = append(packages, path.Base(pkg))
	}
	for _, file := range files {
		if !sutil.inSliceString(packages, file.Name()) {
			fmt.Printf("Removing package not in config anymore %s\n", file.Name())
			os.Remove(c.projectStructure.appendSlash(c.projectStructure.getAemInstallDirLocation(c.instance)) + file.Name())
		}
	}
}

func (c *commandStart) getOpt(args []string) {
	getopt.FlagLong(&c.forceDownload, "download", 'd', "Force new download")
	getopt.FlagLong(&c.forGround, "foreground", 'f', "Don't detach aem")
	getopt.FlagLong(&c.name, "name", 'n', "Instance to start. (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.root, "root", 'r', "Allow root")
	getopt.CommandLine.Parse(args)
}
