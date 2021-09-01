package commands

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/cachedir"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/sliceutil"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

// Exit codes for AEM
const (
	ExitNormal  = 0
	ExitError   = 1
	HomeDirFile = ".aem"
)

// Command internal interface for commands
type Command interface {
	setup() *cobra.Command
	preRun(cmd *cobra.Command, args []string)
	run(cmd *cobra.Command, args []string)
}

func getConfig() (*objects.Config, error) {
	return aem.GetConfig()
}
func getConfigAndInstance(i string) (*objects.Config, *objects.Instance, string, error) {
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, nil, "Could not load config file. (%s)", err
	}

	currentInstance, err := aem.GetByName(i, cnf.Instances)
	if err != nil {
		return cnf, nil, "Could not find instance. (%s)", err
	}

	aem.Cnf = cnf
	return cnf, currentInstance, ``, nil
}

//nolint
func getConfigAndInstanceOrGroup(i, g string) (*objects.Config, []objects.Instance, string, error) {
	if len(g) > 0 {
		return getConfigAndGroup(g)
	}
	c, in, s, e := getConfigAndInstance(i)
	return c, []objects.Instance{*in}, s, e
}

func getConfigAndInstanceOrGroupWithRoles(i, g string, r []string) (*objects.Config, []objects.Instance, string, error) {
	if len(g) > 0 {
		return getConfigAndGroupWithRoles(g, r)
	}

	c, in, s, e := getConfigAndInstance(i)
	if e != nil {
		return c, nil, s, e
	}

	if !sliceutil.InSliceString(r, in.Type) {
		return c, []objects.Instance{*in}, s, fmt.Errorf("instance is not of type %s", in.Type)
	}

	return c, []objects.Instance{*in}, s, e
}

func getCdnConfig(name string) (*objects.CDN, error){
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("Could not load config file. (%s)", err)
	}

	for _, cdn := range cnf.CDNs {
		if  strings.EqualFold(cdn.Name, name) {
			return &cdn, nil
		}
	}

	return nil, fmt.Errorf("Could not find cdn: %s.", name)
}

func getConfigAndGroup(i string) (*objects.Config, []objects.Instance, string, error) {
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, nil, "Could not load config file. (%s)", err
	}

	currentInstance, err := aem.GetByGroup(i, cnf.Instances)
	if err != nil {
		return cnf, nil, "Could not find instance. (%s)", err
	}
	aem.Cnf = cnf
	return cnf, currentInstance, ``, nil
}

func getConfigAndGroupWithRole(i, r string) (*objects.Config, []objects.Instance, string, error) {
	return getConfigAndGroupWithRoles(i, []string{r})
}

func getConfigAndGroupWithRoles(i string, r []string) (*objects.Config, []objects.Instance, string, error) {
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, nil, "Could not load config file. (%s)", err
	}

	currentInstance, err := aem.GetByGroupAndRoles(i, cnf.Instances, r)
	if err != nil {
		return cnf, nil, "Could not find instance. (%s)", err
	}
	aem.Cnf = cnf
	return cnf, currentInstance, ``, nil
}

// CheckConfigExists is the configuration file available
func CheckConfigExists() (bool, error) {
	p, err := project.GetConfigFileLocation()
	if err != nil {
		return false, fmt.Errorf("could not get config file location: %s", err)
	}

	if !project.Exists(p) {
		return false, nil
	}

	return true, nil
}

// ReadRegisteredProjects reads aem file to find registered projects
func ReadRegisteredProjects(homedir string) objects.Projects {
	projects := objects.Projects{}
	if project.Exists(homedir + "/" + HomeDirFile) {
		toml.DecodeFile(homedir+"/"+HomeDirFile, &projects) // nolint: errcheck
	}
	return projects
}

func changeProjectDir(projectName string) {
	projects := cachedir.RegisteredProjects()
	for _, cProject := range projects {
		if strings.EqualFold(cProject.Name, projectName) {
			err := os.Chdir(cProject.Path)
			if err != nil {
				output.Printf(output.NORMAL, "Could not change to  project folder: %s.\n", err.Error())
				os.Exit(ExitError)
			}
			return
		}
	}
	output.Printf(output.NORMAL, "Could not find project: %s.\n", projectName)
	os.Exit(ExitError)
}

func lookupByPath() error {
	cwd, err := project.GetWorkDir()
	projects := cachedir.RegisteredProjects()
	if err != nil {
		return err
	}

	for _, p := range projects {
		if strings.HasPrefix(cwd, p.Path) {
			err := os.Chdir(p.Path)
			return err
		}
	}
	return fmt.Errorf("not found")
}

// ConfigCheckListProjects Check for config and list projects if needed
func ConfigCheckListProjects() {

	if Project != "" {
		changeProjectDir(Project)
	}

	b, err := CheckConfigExists()
	if err != nil {
		output.Print(output.NORMAL, "Error while searching for config file.\n")
		output.Printf(output.VERBOSE, "error: %s", err.Error())
		os.Exit(ExitError)
	}

	if !b {
		err := lookupByPath()
		if err == nil {
			return
		}

		output.Print(output.NORMAL, "No config file in the current directory.\n")
		p, err := project.HomeDir()
		if err != nil {
			output.Print(output.NORMAL, "Could not get the home dir.\n")
			os.Exit(ExitError)
		}

		projects := ReadRegisteredProjects(p)
		if len(p) > 0 {
			output.Print(output.NORMAL, "You have the following projects on registered.\n")
			output.Print(output.NORMAL, "Switch to the project location to start using the tool\n\n")

			for _, project := range projects.Project {
				fmt.Printf(" * %s - %s\n", project.Name, project.Path)
			}
			output.Print(output.NORMAL, "\n")
		}
		os.Exit(ExitError)
	}
}

// RegisterProject in homedir
func RegisterProject() {
	cnf, err := getConfig()
	if err != nil {
		return
	}
	cwd, err := project.GetWorkDir()
	if err != nil {
		return
	}
	cachedir.RegisterProject(cnf.ProjectName, cwd)
}

// WriteRegisterFile writes project in project registry file
func WriteRegisterFile(projects objects.Projects, homedir string) {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(projects)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(homedir+"/"+HomeDirFile, buf.Bytes(), 0644)
	if err != nil {
		return
	}

}

// GetInstancesAndConfig gets config and configuration for instance or group
func GetInstancesAndConfig(i, g string) (*objects.Config, []objects.Instance, string, error) {
	if len(g) > 0 {
		return getConfigAndGroup(g)
	}
	cnf, instance, errString, err := getConfigAndInstance(i)
	aem.Cnf = cnf
	return cnf, []objects.Instance{*instance}, errString, err

}
