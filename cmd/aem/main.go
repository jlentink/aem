package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pborman/getopt/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	config       = Config{}
	BuiltVersion = ""
	BuiltHash    = ""
)

func parseParameters() {
	config.CommandArgs = os.Args

	if len(os.Args) > 1 && os.Args[0][1:1] != "-" {
		config.Command = os.Args[1]
		config.CommandArgs = os.Args[1:]
	}

	getopt.FlagLong(&config.Verbose, "verbose", 'v', "Enable verbose")
	getopt.CommandLine.Getopt(config.CommandArgs, nil)
}

func exitFatal(err error, message string, args ...interface{}) {
	if nil != err {
		if len(args) > 0 {
			callExit(fmt.Sprintf(message, args...))
		} else {
			callExit(fmt.Sprintf("%s%s", message, err.Error()))
		}
	}
}

func callExit(message string) {
	if os.Getenv("UNIT_TEST") == "1" {
		panic("os.exit called")
	} else {
		fmt.Print(message)
		os.Exit(1)
	}
}

func exitProgram(message string, args ...interface{}) {
	callExit(fmt.Sprintf(message, args...))
}

func confirm(format string, force bool, args ...interface{}) bool {

	if force {
		return true
	}

	fmt.Printf(format, args...)
	fmt.Print("[Y/n] ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if strings.ToLower(input) == "y\n" {
		return true
	} else {
		fmt.Printf(strings.ToLower(input))
		return false
	}

}

func readConfig() {
	dir, err := os.Getwd()
	exitFatal(err, "Could not get current working directory.")

	if _, err := os.Stat(dir + "/" + CONFIG_FILENAME); err == nil {
		log.Debug("Found config file.")
		_, err := toml.DecodeFile(dir+"/"+CONFIG_FILENAME, &config)

		exitFatal(err, "Config file error: ")
	}
}

func setupLog() {
	if (config.Verbose) {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	parseParameters()
	setupLog()

	switch config.Command {
	case "package-list":
		readConfig()
		command := NewListPackagesCommand()
		command.Execute(config.CommandArgs)
	case "package-copy":
		readConfig()
		command := NewCopyPackageCommand()
		command.Execute(config.CommandArgs)
	case "package-install":
		readConfig()
		command := NewInstallPackageCommand()
		command.Execute(config.CommandArgs)
	case "package-rebuild":
		readConfig()
		command := NewRebuildPackageCommand()
		command.Execute(config.CommandArgs)
	case "package-download":
		readConfig()
		command := NewDownloadPackageCommand()
		command.Execute(config.CommandArgs)
	case "pull-content":
		readConfig()
		command := NewPullContentCommand()
		command.Execute(config.CommandArgs)
	case "start":
		readConfig()
		command := NewStartCommand()
		command.Execute(config.CommandArgs)
	case "stop":
		readConfig()
		command := NewStopCommand()
		command.Execute(config.CommandArgs)
	case "open":
		readConfig()
		command := NewOpenCommand()
		command.Execute(config.CommandArgs)
	case "sync":
		readConfig()
		command := NewSyncCommand()
		command.Execute(config.CommandArgs)
	case "log":
		readConfig()
		command := NewLogCommand()
		command.Execute(config.CommandArgs)
	case "system-information", "sysinfo":
		readConfig()
		command := NewSystemInformationCommand()
		command.Execute(config.CommandArgs)
	case "password", "passwords":
		readConfig()
		command := NewPasswordCommand()
		command.Execute(config.CommandArgs)
	case "init":
		command := NewInitCommand()
		command.Execute(config.CommandArgs)
	case "version":
		command := NewVersionCommand()
		command.Execute(config.CommandArgs)
	case "page-replicate":
		readConfig()
		command := NewActivatePageCommand()
		command.Execute(config.CommandArgs)
	case "activate-tree":
		readConfig()
		command := NewActivateTreeCommand()
		command.Execute(config.CommandArgs)
	case "bundle-list", "bundles-list":
		readConfig()
		command := NewListBundlesCommand()
		command.Execute(config.CommandArgs)
	case "bundle-stop":
		readConfig()
		command := NewBundleStopCommand()
		command.Execute(config.CommandArgs)
	case "bundle-start":
		readConfig()
		command := NewBundleStartCommand()
		command.Execute(config.CommandArgs)
	case "bundle-install":
		readConfig()
		command := NewBundleInstallCommand()
		command.Execute(config.CommandArgs)
	default:
		showHelp()
	}
}

func showHelp() {
	fmt.Printf("aem-cli Usage:\n")
	fmt.Printf("Built version: " + BuiltVersion + " (" + BuiltHash + ")\n")
	fmt.Printf("\n")
	fmt.Printf("Available commands:\n")
	fmt.Printf("- init\n")
	fmt.Printf("  Create sample config file.\n")
	fmt.Printf("- package-list\n")
	fmt.Printf("  List packages on server.\n")
	fmt.Printf("- package-copy\n")
	fmt.Printf("  Copy packages from one instance to another.\n")
	fmt.Printf("- package-install\n")
	fmt.Printf("  Install packages from local zip.\n")
	fmt.Printf("- package-rebuild\n")
	fmt.Printf("  Rebuild package on instance.\n")
	fmt.Printf("- package-download\n")
	fmt.Printf("  Download package from instance.\n")
	fmt.Printf("- pull-content\n")
	fmt.Printf("  Pull content packages from intance to...\n")
	fmt.Printf("- start\n")
	fmt.Printf("  Start AEM.\n")
	fmt.Printf("- stop\n")
	fmt.Printf("  Stop AEM.\n")
	fmt.Printf("- log\n")
	fmt.Printf("  Show log file.\n")
	fmt.Printf("- Open\n")
	fmt.Printf("  Open instance in browser\n")
	fmt.Printf("- system-information or sysinfo\n")
	fmt.Printf("  Show system information of instance\n")
	fmt.Printf("- replicate-page\n")
	fmt.Printf("  Activate or deactive page\n")
	fmt.Printf("- activate-tree\n")
	fmt.Printf("  Activate tree on instance\n")
	fmt.Printf("- password\n")
	fmt.Printf("  Save passwords to OS keyring\n")
	fmt.Printf("- version\n")
	fmt.Printf("  Show aemcli version.\n")
	fmt.Printf("- sync\n")
	fmt.Printf("  Watch file changes and sync to server. (needs aemsync installed)\n")
	fmt.Printf("  to install: npm install aemsync -g\n")
	fmt.Printf("- help\n")
	fmt.Printf("  Show this help section.\n")
	fmt.Printf("\n")
	fmt.Printf("aem-cli <command> -h for more information per command.\n")
	fmt.Printf("\n")
	fmt.Printf("WIP:\n")
	fmt.Printf("- pull-content-vlt-rcp\n")
	fmt.Printf("- threaddump\n")
	fmt.Printf("\n")
}
