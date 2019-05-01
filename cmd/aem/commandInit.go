package main

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/pborman/getopt/v2"
	"github.com/spf13/afero"
)

type commandInit struct {
	u         *utility
	p         *projectStructure
	fs        afero.Fs
	dump      bool
	overwrite bool
}

func (c *commandInit) Init() {
	c.u = new(utility)
	c.p = new(projectStructure)
	c.fs = afero.NewOsFs()
	c.dump = false
	c.overwrite = false

}

func (c *commandInit) readConfig() bool {
	return true
}

func (c *commandInit) GetCommand() []string {
	return []string{"init"}
}

func (c *commandInit) GetHelp() string {
	return "Create aem cli config file."
}

func (c *commandInit) survey() string {
	answers := newConfigAnswers()
	answers.AdditionalPackages = []string{}

	err := survey.Ask(surveyInitialQuestionsQuestions, &answers)

	validateSurveyInput(err)

	if answers.JarLocationType == "filesystem" {
		err = survey.Ask(surveyJarFileQuestions, &answers)
	} else {
		err = survey.Ask(surveyJarHTTPQuestions, &answers)
	}

	validateSurveyInput(err)

	err = survey.Ask(surveyLicenseQuestions, &answers)

	validateSurveyInput(err)

	for {
		err = survey.Ask(surveyAdditionalPackagesQuestions, &answers)
		answers.AdditionalPackages = append(answers.AdditionalPackages, answers.AdditionalPackage)
		answers.AdditionalPackage = ""

		validateSurveyInput(err)
		if !answers.MorePackages {
			break
		}
	}

	return answers.getConfig()

}

//validateSurveyInput validates the returned error object from survey.Ask()
func validateSurveyInput(err error) {
	if nil != err {
		if err.Error() == "interrupt" {
			exitProgram("Interrupted: no config file created\n")
		}
		// exit with regular error (validation)
		exitProgram(err.Error() + "\n")
	}
}

func (c *commandInit) Execute(args []string) {
	c.getOpt(args)
	configTemplateStr := ""

	if !c.u.Exists(c.p.getConfigFileLocation()) || c.overwrite {
		if !c.dump {
			configTemplateStr = c.survey()
		} else {
			answers := newConfigAnswers()
			configTemplateStr = answers.getConfig()
		}

		err := afero.WriteFile(c.fs, c.p.getConfigFileLocation(), []byte(configTemplateStr), 0644)
		exitFatal(err, "Could not write config file.")
		fmt.Printf("Written sample config file. please edit .aem\n")

	} else {
		exitProgram("\".aem\" file found; please edit to update the values.\n")
	}

}

func (c *commandInit) getOpt(args []string) {
	getopt.FlagLong(&c.dump, "dump", 'd', "Write default config file without setup questions")
	getopt.FlagLong(&c.overwrite, "force-overwrite", 'f', "Overwrite current configuration")
	getopt.CommandLine.Parse(args)
}
