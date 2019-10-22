package commands

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/generate"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandGenerate struct {
	verbose bool
}

func (c *commandGenerate) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "Generate code block",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	dump := &commandGenerateDump{}
	cmd.AddCommand(dump.setup())
	return cmd
}

func (c *commandGenerate) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandGenerate) run(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("templates"); os.IsNotExist(err) {
		output.Printf(output.NORMAL, "Could not find template dir. Please run aem generate init first.")
		os.Exit(ExitError)
	}

	components, err := generate.ListComponents()
	if err != nil {
		output.Printf(output.NORMAL, "Unexpected error while reading components: %s", err.Error())
		os.Exit(ExitError)
	}
	componentName := ""
	options := []string{}
	for _, component := range components {
		options = append(options, component.Name+" - "+component.Version)
	}

	promptComponent := &survey.Select{
		Message: "Which component to generate",
		Options: options,
	}

	err = survey.AskOne(promptComponent, &componentName)
	validateSurveyError(err)
	component, err := componentSelector(componentName, components)
	if err != nil {
		output.Printf(output.NORMAL, "Unexpected error while searching components: %s", err.Error())
		os.Exit(ExitError)
	}

	nameComponent := &survey.Input{
		Message: "Name for component",
	}

	generateName := ""
	err = survey.AskOne(nameComponent, &generateName)
	validateSurveyError(err)
	if err != nil {
		output.Printf(output.NORMAL, "Unexpected error while searching components: %s", err.Error())
		os.Exit(ExitError)
	}
	component.Name = generateName

	destination, err := generate.ParseTemplate(component.Destination, map[string]string{"project": aem.Cnf.ProjectName})
	if err != nil {
		output.Printf(output.NORMAL, "Template error in destination: %s", err.Error())
		os.Exit(ExitError)
	}
	destinationPath := ""
	promptDestination := &survey.Input{
		Message: "destination",
		Default: destination,
	}
	err = survey.AskOne(promptDestination, &destinationPath)
	validateSurveyError(err)
	component.Destination = destinationPath
	err = component.Generate()
	if err != nil {
		output.Printf(output.NORMAL, "\U0001F622 Error while generating component %s\n", err.Error())
		os.Exit(ExitError)
	}
}

func validateSurveyError(err error) {
	if err != nil {
		if err.Error() == "interrupt" {
			output.Printf(output.NORMAL, "Program interrupted. (CTRL+C)")
			os.Exit(ExitError)
		}
		output.Printf(output.NORMAL, "Unexpected error: %s", err.Error())
		os.Exit(ExitError)
	}
}

func componentSelector(selectedComponent string, components []*generate.Component) (*generate.Component, error) {
	parts := strings.Split(selectedComponent, "-")
	var name, version string
	if len(parts) > 2 {
		for i := 0; i < len(parts)-1; i++ {
			if i > 0 {
				name += "-"
			}
			name += parts[i]
		}
		name = strings.TrimSpace(name)
		version = strings.TrimSpace(parts[len(parts)-1])
	} else {
		name = strings.TrimSpace(parts[0])
		version = strings.TrimSpace(parts[1])
	}

	for _, component := range components {
		if strings.TrimSpace(component.Name) == name && strings.TrimSpace(component.Version) == version {
			return component, nil
		}
	}
	return nil, fmt.Errorf("could not find the component: %s", selectedComponent)
}
