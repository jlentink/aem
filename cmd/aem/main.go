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
	config   = configStruct{}
	commands = []ICommand{
		&commandStart{},
		&commandStop{},
		&commandOpen{},
		&commandLog{},
		&commandSync{},
		&commandInit{},
		&commandVersion{},
		&commandPassword{},
		&commandPullContent{},

		&commandPackagesList{},
		&commandPackageInstall{},
		&commandPackageDownload{},
		&commandPackageRebuild{},
		&commandPackageCopy{},

		&commandSystemInformation{},
		&commandActivatePage{},
		&commandActivateTree{},

		&commandBundleList{},
		&commandBundleInstall{},
		&commandBundleStart{},
		&commandBundleStop{},

		&commandOakCheckpoints{},
		&commandOakExplore{},
		&commandOakCheck{},
		&commandOakCompact{},
		&commandOakConsole{},
	}

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
	fmt.Print(strings.ToLower(input))
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
	s := new(sliceUtil)

	for _, cmd := range commands {
		if s.inSliceString(cmd.GetCommand(), config.Command) {
			if cmd.readConfig() {
				readConfig()
			}
			cmd.Init()
			cmd.Execute(config.CommandArgs)
			return
		}
	}

	showHelp()
}

func showHelp() {
	fmt.Printf("aem-cli Usage:\n")
	fmt.Printf("\n")
	fmt.Printf("Available commands:\n")
	for _, cmd := range commands {
		fmt.Printf(" - %s", strings.Join(cmd.GetCommand(), ", "))
		fmt.Printf(": %s\n", cmd.GetHelp())
	}
	fmt.Printf("\nversion: %s\n", BuiltVersion)
	fmt.Printf("%s <command> -h for more information per command.\n\n", os.Args[0])
}
