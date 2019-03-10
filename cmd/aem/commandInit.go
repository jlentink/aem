package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/pborman/getopt/v2"
	"github.com/spf13/afero"
)

func newInitCommand() commandInit {
	return commandInit{
		u:         new(utility),
		p:         new(projectStructure),
		fs:        afero.NewOsFs(),
		dump:      false,
		overwrite: false,
	}
}


type commandInit struct {
	u         *utility
	p         *projectStructure
	fs        afero.Fs
	dump      bool
	overwrite bool
}

func (p *commandInit) survey() string {
	answers := newConfigAnswers()
	answers.AdditionalPackages = []string{}

	survey.Ask(surveyInitialQuestionsQuestions, &answers)

	if answers.JarLocationType == "filesystem" {
		survey.Ask(surveyJarFileQuestions, &answers)
	} else {
		survey.Ask(surveyJarHTTPQuestions, &answers)
	}

	for {
		survey.Ask(surveyAdditionalPackagesQuestions, &answers)
		answers.AdditionalPackages = append(answers.AdditionalPackages, answers.AdditionalPackage)
		answers.AdditionalPackage = ""
		if !answers.MorePackages {
			break
		}
	}

	return answers.getConfig()

}

func (p *commandInit) Execute(args []string) {
	p.getOpt(args)
	configTemplateStr := ""

	if !p.u.Exists(p.p.getConfigFileLocation()) || p.overwrite {
		if !p.dump {
			configTemplateStr = p.survey()
		} else {
			answers := newConfigAnswers()
			configTemplateStr = answers.getConfig()
		}

		err := afero.WriteFile(p.fs, p.p.getConfigFileLocation(), []byte(configTemplateStr), 0644)
		exitFatal(err, "Could not write config file.")
		fmt.Printf("Written sample config file. please edit .aem\n")

	} else {
		exitProgram("\".aem\" file found; please edit to update the values.")
	}

}

func (p *commandInit) getOpt(args []string) {
	getopt.FlagLong(&p.dump, "dump", 'd', "Write default config file without setup questions")
	getopt.FlagLong(&p.overwrite, "force-overwrite", 'f', "Overwrite current configuration")
	getopt.CommandLine.Parse(args)
}
