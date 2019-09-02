package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/bundle"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/pom"
	"github.com/spf13/cobra"
	"os"
)

type commandDeploy struct {
	verbose       bool
	instanceName  string
	instanceGroup string
	forceBuild    bool
	username      string
	password      string
	artifact      string
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
		"Build before deploy")
	cmd.Flags().StringVarP(&c.username, "username", "u", "",
		"Overwrite username to use if not using the one from config file")
	cmd.Flags().StringVarP(&c.password, "password", "p", "",
		"Overwrite password to use if not using the one from config file")
	cmd.Flags().StringVarP(&c.artifact, "artifact", "a", "",
		"Deploy one a single artifact")

	return cmd
}

func (c *commandDeploy) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandDeploy) run(cmd *cobra.Command, args []string) {
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
		c.deployAllPackages(is, p)
	}

	c.deployModule(is, p)
}

func (c *commandDeploy) deployModule(is []objects.Instance, p *pom.Pom) {
	a, err := p.GetArtifactByName(c.artifact)
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	if c.forceBuild {
		err = aem.BuildModuleProject(a.BasePath)
		if err != nil {
			output.Printf(output.NORMAL, "build failed: %s", err.Error())
			os.Exit(ExitError)
		}
	}

	for _, i := range is {
		switch a.Kind() {
		case pom.Bundle:
			bundle.Install(&i, a.CompletePath(), "20")
			break
		case pom.Package:
			resp, err := pkg.Upload(i, a.CompletePath(), true, true)
			if err != nil {
				fmt.Printf("Status: \U0000274C\n")
				output.Printf(output.NORMAL, "%s\n", err)
			}
			if resp != nil {
				fmt.Printf("Status: \U00002705\n")
				output.Printf(output.VERBOSE, "%s\n", resp.Response.Data.Log)
			}
			break
		default:
			output.Printf(output.NORMAL, "Unknown package type. %s", a.Packaging)
			os.Exit(ExitError)
		}
	}
}

func (c *commandDeploy) deployAllPackages(is []objects.Instance, p *pom.Pom) {
	if c.forceBuild {
		aem.BuildProject()
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
			fmt.Printf("\r%s\n", artifact.Filename())
			resp, err := pkg.Upload(i, artifact.CompletePath(), true, true)
			if resp != nil {
				success++
				fmt.Printf("Status: \U00002705\n")
				output.Printf(output.VERBOSE, "%s\n", resp.Response.Data.Log)
			}
			if err != nil {
				fmt.Printf("Status: \U0000274C\n")
				failed++
				output.Printf(output.NORMAL, "%s\n", err)
			}
		}
	}
	fmt.Printf("\n\n"+
		"=============================================================\n"+
		"  Install Summary: %d Success, %d Failed\n"+
		"=============================================================\n", success, failed)
	if failed > 0 {
		os.Exit(ExitError)
	}
	os.Exit(ExitNormal)
}
