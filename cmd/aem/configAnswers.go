package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"regexp"
	"strings"
)

var surveyInitialQuestionsQuestions = []*survey.Question{
	{
		Name: "UseKeyRing",
		Prompt: &survey.Confirm{
			Message: "Use the password manager of your operating system.",
			Help:    "Where is the location of the AEM jar.",
		},
		Validate: survey.Required,
	}, {
		Name: "JarLocationType",
		Prompt: &survey.Select{
			Message: "Location of AEM jar",
			Options: []string{"http(s)", "filesystem"},
			Help:    "Where is the location of the AEM jar.",
		},
		Validate: survey.Required,
	},
}

var surveyJarHTTPQuestions = []*survey.Question{
	{
		Name: "JarLocation",
		Prompt: &survey.Input{
			Message: "Location of AEM jar",
			Help:    "The URL to download the AEM jar from.",
			Default: "https://somedomain.tld/some-4502.jar",
		},
		Validate: survey.Required,
	},
	{
		Name: "JarUsername",
		Prompt: &survey.Input{
			Message: "Http username to use (if needed)",
			Default: "some-username",
			Help:    "If authentication what would be the username",
		},
	},
	{
		Name: "JarPassword",
		Prompt: &survey.Password{
			Message: "Http password to use. (if needed)",
			Help:    "If authentication what would be the password",
		},
	},
}

var surveyJarFileQuestions = []*survey.Question{
	{
		Name: "JarLocation",
		Prompt: &survey.Input{
			Message: "Path to the AEM jar",
			Help:    "Where on your filesystem did you store your Jar",
			Default: "/foo/bar/some-4502.jar",
		},
		Validate: survey.Required,
	},
}

var surveyLicenseQuestions = []*survey.Question{
	{
		Name: "LicenseCustomer",
		Prompt: &survey.Input{
			Message: "License customer",
			Help:    "What is the license customer name. (only use this in private projects; keep license a secret!)",
			Default: "foo-bar",
		},
	},
	{
		Name: "LicenseVersion",
		Prompt: &survey.Input{
			Message: "License version",
			Help:    "What is the AEM version.",
			Default: "6.x",
		},
	},
	{
		Name: "LicenseDownloadID",
		Prompt: &survey.Input{
			Message: "License download ID",
			Help:    "What is the AEM download id. (only use this in private projects; keep license a secret!)",
			Default: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		},
	},
}

var surveyAdditionalPackagesQuestions = []*survey.Question{
	{
		Name: "AdditionalPackage",
		Prompt: &survey.Input{
			Message: "Additional packages to install (url with or without password.",
			Default: "https://github.com/Adobe-Consulting-Services/acs-aem-commons/releases/download/acs-aem-commons-3.19.0/acs-aem-commons-content-3.19.0.zip",
			Help:    "http(s)://username:password@somedomain.tld/foobar.zip",
		},
		Validate: func(val interface{}) error {
			r, _ := regexp.Compile("(?i)^http(s?)://((.*:.*)@)?([a-z0-9-./]*).zip$")
			str, _ := val.(string)
			URLs := strings.Split(str, "\n")
			for _, URL := range URLs {
				fmt.Printf("%s", URL)
				if len(str) != 0 && !r.MatchString(URL) {
					return fmt.Errorf("invalid url found: %s", URL)
				}
			}

			return nil
		},
	},
	{
		Name: "MorePackages",
		Prompt: &survey.Confirm{
			Message: "More additional packages?",
			Help:    "More additional packages to add?",
		},
		Validate: survey.Required,
	},
}

func newConfigAnswers() configAnswers {
	return configAnswers{
		UseKeyRing:         true,
		JarLocation:        "",
		JarUsername:        "admin",
		JarPassword:        "admin",
		LicenseCustomer:    "",
		LicenseVersion:     "",
		LicenseDownloadID:  "",
		AdditionalPackages: []string{"https://github.com/Adobe-Consulting-Services/acs-aem-commons/releases/download/acs-aem-commons-3.19.0/acs-aem-commons-content-3.19.0.zip"},
	}
}

type configAnswers struct {
	UseKeyRing         bool
	MorePackages       bool
	JarURL             string
	JarLocationType    string
	JarLocation        string
	JarUsername        string
	JarPassword        string
	LicenseCustomer    string
	LicenseVersion     string
	LicenseDownloadID  string
	AdditionalPackage  string
	AdditionalPackages []string
}

func (c *configAnswers) joinStrings(stringsArr []string) string {
	returnString := ""
	for _, str := range stringsArr {
		returnString = returnString + "\t\"" + str + "\",\n"
	}

	return returnString
}

func (c *configAnswers) getConfig() string {
	return fmt.Sprintf(configTemplate,
		c.UseKeyRing,
		c.JarLocation,
		c.JarUsername,
		c.JarPassword,
		c.LicenseCustomer,
		c.LicenseVersion,
		c.LicenseDownloadID,
		c.joinStrings(c.AdditionalPackages))
}

const (
	configTemplate = `
#
# Example .aem config file.
# Update based on your need
#
verbose = false

#
# When set to true passwords will be stored in your local OS
# keyring. When disabled the passwords from this config file will be used.
#
use-keyring = %t

#
# AEM jar. The AEM jar to start AEM.
#
# Can be an user with username and password or without.
# you can also define a path to a local jar on disk.
#
# Urls where to find the AEM jar
# Example: https://somedomain.tld/aem-author-4502.jar
# 
# The jarUsername and jarPassword are the 
# HTTP basic credentials that can be used.
#
# Beware to not store the AEM jar in a public unprotected path!
#
jarLocation = "%s"
jarUsername = "%s"
jarPassword = "%s"

#
# Project license information.
# licenses are sensitive material and should not be shared outside of the project
# or miss used. So be very careful when using this functonality!!!!
#
licenseCustomer = "%s"
licenseVersion = "%s"
licenseDownloadID = "%s"

#
# paths to watch and sync during development.
# aemsync needs to be installed to use this function
# npm install aemsync -g
#
watchPath = [
    "ui.apps/src/main/content/jcr_root",
    "ui.content/src/main/content/jcr_root"
]

#
# Content packages to use for this project
#
contentPackages = [
	"content:1.0.0",
	"assets:1.0.0",
]


#
# Additional packages to auto install.
# Example given is ACS commons to auto install
# Urls can have username and password added for basic authentication
# Example: https://username@password:somedomain.tld/....
#
additionalPackages = [
%s
]

#
# Which extra arguments to use for oak-run
#
oakOptions = [
    "-mx8G",
    "-Dtar.memoryMapped=true",
]

oakDefaultVersion = "1.8.12"

#
# General JVM options to use when starting AEM
#
jvm-options = [
    "-server",
    "-Xmx4048m",
    "-XX:MaxPermSize=512M",
    "-Djava.awt.headless=true",
    "-Dcom.sun.management.jmxremote.port=3233",
    "-Dcom.sun.management.jmxremote.ssl=false",
    "-Dcom.sun.management.jmxremote.authenticate=false"
]

#
# General JVM debug options to use when starting AEM
#
jvm-debug-options = [
    "-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005",
]

#
# Instances for your project.
# Make sure you have local-author and local-publisher available
# and you can define as many as you want name needs to be unique.
#
[[instance]]
name = "local-author"
group = "local"
debug = true
proto = "http"
hostname = "127.0.0.1"
port = 4502
username = "admin"
password = ""
type = "author"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "local-publish"
group = "local"
debug = false
proto = "http"
hostname = "127.0.0.1"
port = 4503
username = "admin"
password = ""
type = "publish"
jvm-options = []
jvm-debug-options = []


#
# Not used yet
#
packages = [
]
`
)
