package commands

import (
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

var (
	verbose bool
	// Project - Global set project
	Project  string
	commands = []Command{
		&commandVersion{},
		&commandInit{},
		&commandStart{},
		&commandStop{},
		&commandOpen{},
		&commandLog{},
		&commandShellCompletion{},
		&commandSetupCheck{},
		&commandSystemInformation{},
		&commandDeploy{},
		&commandBuild{},
		&commandPullContent{},
		&commandPassword{},
		&commandBundle{},
		&commandPackage{},
		&commandOak{},
		&commandIndex{},
		&commandActivation{},
		&commandProjects{},
		&commandGenerate{},
		&commandInvalidate{},
		&commandDestroy{},
		//&commandCloudManager{},
	}
	rootCmd = &cobra.Command{Use: "aem"}
)

// Execute init commands
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&Project, "project", "P", "", "Run command for project. (if not current working directory)")
	for _, cmd := range commands {
		rootCmd.AddCommand(cmd.setup())
	}
	err := rootCmd.Execute()
	if err != nil {
		output.Printf(output.NORMAL, "Could not execute root command.\n")
		os.Exit(ExitError)
	}
}
