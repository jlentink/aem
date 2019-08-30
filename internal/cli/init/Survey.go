package init

//var surveyInitialQuestionsQuestions = []*survey.Question{
//	{
//		Name: "UseKeyRing",
//		Prompt: &survey.Confirm{
//			Message: "Use the password manager of your operating system.",
//			Help:    "Where is the location of the AEM jar.",
//		},
//		Validate: survey.Required,
//	}, {
//		Name: "JarLocationType",
//		Prompt: &survey.Select{
//			Message: "Location of AEM jar",
//			Options: []string{"http(s)", "filesystem"},
//			Help:    "Where is the location of the AEM jar.",
//		},
//		Validate: survey.Required,
//	},
//}
//
//var surveyJarHTTPQuestions = []*survey.Question{
//	{
//		Name: "JarLocation",
//		Prompt: &survey.Input{
//			Message: "Location of AEM jar",
//			Help:    "The URL to download the AEM jar from.",
//			Default: "https://somedomain.tld/some-4502.jar",
//		},
//		Validate: survey.Required,
//	},
//	{
//		Name: "JarUsername",
//		Prompt: &survey.Input{
//			Message: "Http username to use (if needed)",
//			Default: "some-username",
//			Help:    "If authentication what would be the username",
//		},
//	},
//	{
//		Name: "JarPassword",
//		Prompt: &survey.Password{
//			Message: "Http password to use. (if needed)",
//			Help:    "If authentication what would be the password",
//		},
//	},
//}
//
//var surveyJarFileQuestions = []*survey.Question{
//	{
//		Name: "JarLocation",
//		Prompt: &survey.Input{
//			Message: "Path to the AEM jar",
//			Help:    "Where on your filesystem did you store your Jar",
//			Default: "/foo/bar/some-4502.jar",
//		},
//		Validate: survey.Required,
//	},
//}
//
//var surveyLicenseQuestions = []*survey.Question{
//	{
//		Name: "LicenseCustomer",
//		Prompt: &survey.Input{
//			Message: "License customer",
//			Help:    "What is the license customer name. (only use this in private projects; keep license a secret!)",
//			Default: "foo-bar",
//		},
//	},
//	{
//		Name: "LicenseVersion",
//		Prompt: &survey.Input{
//			Message: "License version",
//			Help:    "What is the AEM version.",
//			Default: "6.x",
//		},
//	},
//	{
//		Name: "LicenseDownloadID",
//		Prompt: &survey.Input{
//			Message: "License download ID",
//			Help:    "What is the AEM download id. (only use this in private projects; keep license a secret!)",
//			Default: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
//		},
//	},
//}
//
//var surveyLazybonesQuestions = []*survey.Question{
//	{
//		Name: "Lazybones",
//		Prompt: &survey.Confirm{
//			Message: "Want to generate a project with lazybones?",
//			Help:    "Use Lazybones (https://github.com/Adobe-Consulting-Services/lazybones-aem-templates) to generate new project?",
//		},
//		Validate: survey.Required,
//	},
//}
//
//var surveyLazybonesTemplateQuestions = []*survey.Question{
//	{
//		Name: "LazybonesTemplate",
//		Prompt: &survey.Input{
//			Message: "What template do you want to use for lazybones?",
//			Default: "aem-multimodule-project",
//			Help:    "Default one provided by ACS is: aem-multimodule-project.",
//		},
//		Validate: survey.Required,
//	},
//}
//
//var surveyAdditionalPackagesQuestions = []*survey.Question{
//	{
//		Name: "AdditionalPackage",
//		Prompt: &survey.Input{
//			Message: "Additional packages to install (url with or without password.",
//			Default: "https://github.com/Adobe-Consulting-Services/acs-aem-commons/releases/download/acs-aem-commons-3.19.0/acs-aem-commons-content-3.19.0.zip",
//			Help:    "http(s)://username:password@somedomain.tld/foobar.zip",
//		},
//		Validate: func(val interface{}) error {
//			r, _ := regexp.Compile("(?i)^http(s?)://((.*:.*)@)?([a-z0-9-./]*).zip$")
//			str, _ := val.(string)
//			URLs := strings.Split(str, "\n")
//			for _, URL := range URLs {
//				fmt.Printf("%s", URL)
//				if len(str) != 0 && !r.MatchString(URL) {
//					return fmt.Errorf("invalid url found: %s", URL)
//				}
//			}
//
//			return nil
//		},
//	},
//	{
//		Name: "MorePackages",
//		Prompt: &survey.Confirm{
//			Message: "More additional packages?",
//			Help:    "More additional packages to add?",
//		},
//		Validate: survey.Required,
//	},
//}

//func newConfigAnswers() configAnswers {
//	return configAnswers{
//		UseKeyRing:         true,
//		Lazybones:          false,
//		LazybonesTemplate:  "aem-multimodule-project",
//		JarLocation:        "",
//		JarUsername:        "admin",
//		JarPassword:        "admin",
//		LicenseCustomer:    "",
//		LicenseVersion:     "",
//		LicenseDownloadID:  "",
//		AdditionalPackages: []string{"https://github.com/Adobe-Consulting-Services/acs-aem-commons/releases/download/acs-aem-commons-3.19.0/acs-aem-commons-content-3.19.0.zip"},
//	}
//}
