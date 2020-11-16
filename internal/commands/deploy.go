package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/bundle"
	"github.com/jlentink/aem/internal/aem/dispatcher"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/pom"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

type commandDeploy struct {
	verbose         bool
	instanceName    string
	instanceGroup   string
	forceBuild      bool
	username        string
	password        string
	artifact        string
	flush           bool
	productionBuild bool
}

func (c *commandDeploy) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy to server(s)",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	cmd.Flags().StringVarP(&c.instanceName, "name", "n",
		aem.GetDefaultInstanceName(), "Instance to deploy to")
	cmd.Flags().StringVarP(&c.instanceGroup, "group", "g", "",
		"Group to deploy to")
	cmd.Flags().BoolVarP(&c.forceBuild, "build", "b", false,
		"Build SNAPSHOT before deploy")
	cmd.Flags().StringVarP(&c.username, "username", "u", "",
		"Overwrite username to use if not using the one from config file")
	cmd.Flags().StringVarP(&c.password, "password", "p", "",
		"Overwrite password to use if not using the one from config file")
	cmd.Flags().StringVarP(&c.artifact, "artifact", "a", "",
		"Deploy one a single artifact")
	cmd.Flags().BoolVarP(&c.flush, "flush", "f", true,
		"Flush after deploy")
	cmd.Flags().BoolVarP(&c.productionBuild, "production-build", "B", false,
		"Build versioned before deploy")

	return cmd
}

func (c *commandDeploy) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandDeploy) run(cmd *cobra.Command, args []string) {
	var status bool
	_, is, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName,
		c.instanceGroup, []string{aem.RoleAuthor, aem.RolePublisher})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	path, err := project.GetWorkDir()
	if err != nil {
		output.Printf(output.NORMAL, "Could not get Working dir. (%s)", err.Error())
		os.Exit(ExitError)
	}

	p, err := pom.Open(path + "/pom.xml")
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	if c.artifact == "" {
		status = c.deployAllPackages(is, p)
	} else {
		status = c.deployModule(is, p)
	}

	if status && c.flush {
		_, fis, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName,
			c.instanceGroup, []string{aem.RoleDispatcher})
		if err != nil {
			output.Printf(output.NORMAL, errorString)
			os.Exit(ExitError)
		}
		c.invalidateCache(fis)
	}
}

func (c *commandDeploy) invalidateCache(is []objects.Instance) {
	status := dispatcher.InvalidateAll(is, aem.Cnf.InvalidatePaths)
	if !status {
		os.Exit(ExitError)
	}
}
func (c *commandDeploy) deployModule(is []objects.Instance, p *pom.Pom) bool {
	a, err := p.GetArtifactByName(c.artifact)
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		return false
	}

	if c.forceBuild {
		err = aem.BuildModuleProject(a.BasePath, c.productionBuild)
		if err != nil {
			output.Printf(output.NORMAL, "\U0000274C build failed: %s", err.Error())
			return false
		}
	}

	for _, i := range is {
		i := i
		switch a.Kind() {
		case pom.Bundle:
			err = bundle.Install(&i, a.CompletePath(), "20")
			if err != nil {
				output.Printf(output.NORMAL, "%s\n", err)
			}
		case pom.Package:
			resp, htmlBody, err := pkg.Upload(i, a.CompletePath(), true, true)
			if err != nil {
				fmt.Printf("Status: \U0000274C\n")
				output.Printf(output.NORMAL, "%s\n", err)
			}
			if resp != nil {
				fmt.Printf("Status: \U00002705\n")
				output.Printf(output.VERBOSE, "%s\n", resp.Response.Data.Log)
				if len(htmlBody) > 0 {
					output.Printf(output.VERBOSE, "%s\n", htmlBody)
				}
			}
		default:
			output.Printf(output.NORMAL, "Unknown package type. %s", a.Packaging)
			return false
		}
	}
	return true
}
func (c *commandDeploy) stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
		matched, _ := regexp.MatchString(b, a)
		if matched {
			return true
		}
	}
	return false
}
func (c *commandDeploy) deployAllPackages(is []objects.Instance, p *pom.Pom) bool {
	if c.forceBuild {
		err := aem.BuildProject(c.productionBuild) // nolint: errcheck
		if err != nil {
			output.Printf(output.NORMAL, "\U0000274C Build failed...")
			os.Exit(1)
		}
	}

	var success, failed = 0, 0
	artifacts, _ := p.GetAllArtifacts(pom.Package)
	for _, i := range is {
		if len(c.username) > 0 {
			objects.Cnf.KeyRing = false
			i.Username = c.username
		}

		if len(c.password) > 0 {
			objects.Cnf.KeyRing = false
			i.Password = c.password
		}

		fmt.Printf("Deploying to %s\n", i.Name)
		for _, artifact := range artifacts {
			//aem.Cnf.PackagesExcluded
			if !c.stringInSlice(artifact.PakageName(), aem.Cnf.PackagesExcluded) {
				fmt.Printf("\r%s\n", artifact.Filename())
				resp, HtmlBody, err := pkg.Upload(i, artifact.CompletePath(), true, true)
				if resp != nil {
					success++
					fmt.Printf("Status: \U00002705\n")
					output.Printf(output.VERBOSE, "%s\n", resp.Response.Data.Log)
				}
				if err != nil {
					fmt.Printf("Status: \U0000274C\n")
					failed++
					output.Printf(output.NORMAL, "%s\n", err)
					if len(HtmlBody) != 0 {
						output.Printf(output.NORMAL, "%s\n", HtmlBody)
					}
				}
			}
		}
	}
	fmt.Printf("\n\n"+
		"=============================================================\n"+
		"  Install Summary: %d Success, %d Failed\n"+
		"=============================================================\n", success, failed)

	return failed <= 0
}
