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
	config = configStruct{}

	// BuiltVersion Holds the build version
	BuiltVersion = ""

	// BuiltHash Holds the git hash of the build
	BuiltHash = ""
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
	}
	fmt.Printf(strings.ToLower(input))
	return false
}

func readConfig() {
	dir, err := os.Getwd()
	exitFatal(err, "Could not get current working directory.")

	if _, err := os.Stat(dir + "/" + configFilename); err == nil {
		log.Debug("Found config file.")
		_, err := toml.DecodeFile(dir+"/"+configFilename, &config)

		exitFatal(err, "Config file error: ")
	}
}

func setupLog() {
	if config.Verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	parseParameters()
	setupLog()

	switch config.Command {
	case "package-list":
		readConfig()
		command := newPackageListCommand()
		command.Execute(config.CommandArgs)
	case "package-copy":
		readConfig()
		command := newPackageCopyCommand()
		command.Execute(config.CommandArgs)
	case "package-install":
		readConfig()
		command := newPackageInstallCommand()
		command.Execute(config.CommandArgs)
	case "package-rebuild":
		readConfig()
		command := newPackageRebuildCommand()
		command.Execute(config.CommandArgs)
	case "package-download":
		readConfig()
		command := newPackageDownloadCommand()
		command.Execute(config.CommandArgs)
	case "pull-content":
		readConfig()
		command := newPullContentCommand()
		command.Execute(config.CommandArgs)
	case "start":
		readConfig()
		command := newStartCommand()
		command.Execute(config.CommandArgs)
	case "stop":
		readConfig()
		command := newStopCommand()
		command.Execute(config.CommandArgs)
	case "open":
		readConfig()
		command := newOpenCommand()
		command.Execute(config.CommandArgs)
	case "sync":
		readConfig()
		command := newSyncCommand()
		command.Execute(config.CommandArgs)
	case "log":
		readConfig()
		command := newLogCommand()
		command.Execute(config.CommandArgs)
	case "system-information", "sysinfo":
		readConfig()
		command := newSystemInformationCommand()
		command.Execute(config.CommandArgs)
	case "password", "passwords":
		readConfig()
		command := newPasswordCommand()
		command.Execute(config.CommandArgs)
	case "init":
		command := newInitCommand()
		command.Execute(config.CommandArgs)
	case "version":
		command := newVersionCommand()
		command.Execute(config.CommandArgs)
	case "page-replicate":
		readConfig()
		command := newPageActivateCommand()
		command.Execute(config.CommandArgs)
	case "activate-tree":
		readConfig()
		command := newActivateTreeCommand()
		command.Execute(config.CommandArgs)
	case "bundle-list", "bundles-list":
		readConfig()
		command := newBundleListCommand()
		command.Execute(config.CommandArgs)
	case "bundle-stop":
		readConfig()
		command := newBundleStopCommand()
		command.Execute(config.CommandArgs)
	case "bundle-start":
		readConfig()
		command := newBundleStartCommand()
		command.Execute(config.CommandArgs)
	case "bundle-install":
		readConfig()
		command := newBundleInstallCommand()
		command.Execute(config.CommandArgs)
	case "oak-checkpoints":
		readConfig()
		command := newcommandOakCheckpoints()
		command.Execute(config.CommandArgs)
	case "oak-explore":
		readConfig()
		command := newCommandOakExplore()
		command.Execute(config.CommandArgs)
	case "oak-check":
		readConfig()
		command := newCommandOakCheck()
		command.Execute(config.CommandArgs)
	case "oak-console":
		readConfig()
		command := newCommandOakConsole()
		command.Execute(config.CommandArgs)
	case "oak-compact":
		readConfig()
		command := newCommandOakCompact()
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
	fmt.Printf("- oak-check\n")
	fmt.Printf("  Check the FileStore for inconsistencies\n")
	fmt.Printf("- oak-checkpoints\n")
	fmt.Printf("  Manage checkpoints\n")
	fmt.Printf("- oak-compact\n")
	fmt.Printf("  Segment compaction on a TarMK repository.\n")
	fmt.Printf("- oak-console\n")
	fmt.Printf("  Start an interactive console.\n")
	fmt.Printf("- oak-explorer\n")
	fmt.Printf("  \n")
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
