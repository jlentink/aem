package commands

import (
	"embed"
	"fmt"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

//go:embed _templates/*
var templateFS embed.FS

type commandGenerateDump struct {
	verbose bool
}

func (c *commandGenerateDump) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Write default templates to disk",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	return cmd
}

func (c *commandGenerateDump) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandGenerateDump) run(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("templates"); !os.IsNotExist(err) {
		output.Printf(output.NORMAL, "Found a template dir. No need to run. exiting.")
		os.Exit(ExitNormal)
	}

	entries, err := templateFS.ReadDir("_templates")
	if err != nil {
		output.Printf(output.NORMAL, "Unexpected error while reading templates: %s", err.Error())
		os.Exit(ExitError)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			d := "templates/" + entry.Name()
			err = os.MkdirAll(d, 0777)
			if err != nil {
				output.Printf(output.NORMAL, "Error while writing directory to disk. %s\n", err.Error())
				os.Exit(ExitNormal)
			}
			templateEntries, err := templateFS.ReadDir("_templates/" + entry.Name())
			if err != nil {
				output.Printf(output.NORMAL, "Unexpected error while reading templates: %s", err.Error())
				os.Exit(ExitError)
			}
			for _, templateEntry := range templateEntries {
				if !templateEntry.IsDir() {
					content, err := templateFS.ReadFile("_templates/" + entry.Name() + "/" + templateEntry.Name())
					if err != nil {
						output.Printf(output.NORMAL, "Unexpected error while reading templates: %s", err.Error())
						os.Exit(ExitError)
					}
					err = os.WriteFile("templates/"+entry.Name()+"/"+templateEntry.Name(), content, 0600)
					if err != nil {
						output.Printf(output.NORMAL, "Unexpected error while writing templates: %s", err.Error())
						os.Exit(ExitError)
					}
				}
			}
		}
	}
	fmt.Print("\n Files written to disk. You can now edit or start generating code.\n")
}
