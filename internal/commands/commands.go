package commands

import (
	"github.com/spf13/cobra"
)

var (
	verbose  bool
	commands = []Command{
		&commandVersion{},
		&commandInit{},
		&commandStart{},
		&commandStop{},
		&commandOpen{},
		&commandBash{},
		&commandLog{},
		&commandSetupCheck{},
		&commandSystemInformation{},
		&commandPackageList{},
		&commandDeploy{},
		&commandBuild{},
		&commandPackageDownload{},
		&commandPullContent{},
		&commandCopyPackage{},
		&commandPackageUpload{},
		&commandPackageRebuild{},
		&commandInvalidate{},
		&commandPassword{},
		&commandPackageInstall{},
		&commandOakExplore{},
		&commandBundleList{},
		&commandBundleInstall{},
		&commandBundleStart{},
		&commandBundelStop{},
		&commandOakCheckpoints{},
		&commandOakCheck{},
		&commandOakCompact{},
		&commandOakConsole{},
		&commandReplicationPage{},
		&commandActivateTree{},
		&commandWatch{},
	}
	rootCmd = &cobra.Command{Use: "aem"}
)

// Execute init commands
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	for _, cmd := range commands {
		rootCmd.AddCommand(cmd.setup())
	}
	rootCmd.Execute()
}
