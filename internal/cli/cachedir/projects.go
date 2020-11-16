package cachedir

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"io/ioutil"
	"strings"
)

const projectsFile = "projects.toml"

// Projects Registered projects
type Projects struct {
	Project []ProjectRegistered
}

// ProjectRegistered registered project
type ProjectRegistered struct {
	Name string
	Path string
}

func RegisteredProjects() []ProjectRegistered {
	projects := Projects{}
	if project.Exists(getCacheRoot() + "/" + projectsFile) {
		toml.DecodeFile(getCacheRoot() + "/" + projectsFile, &projects) // nolint: errcheck
	}
	return projects.Project
}

func RegisterProject(name, path string){
	Init()
	mutated := false
	projects := RegisteredProjects()
	for index, project := range projects {
		if strings.ToLower(project.Name) == strings.ToLower(name) && project.Path == path {
			return
		}
		if strings.ToLower(project.Name) == strings.ToLower(name) {
			projects[index].Path = path
			mutated = true
		}
	}

	if !mutated {
		projects = append(projects, ProjectRegistered{Name: name, Path: path})
		mutated = true
	}

	writeRegisterFile(projects)
}

func SetProjectMetaData(project ProjectRegistered, key, value string) {

}

func GetProjectMetaData(project ProjectRegistered, keystring string) {

}

func writeRegisterFile(projects []ProjectRegistered) {
	data := Projects{Project: projects}
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(data)
	if err != nil {
		output.Printf(output.VERBOSE, "Error encoding projects file: %s", err.Error())
		return
	}

	err = ioutil.WriteFile(getCacheRoot() + "/" + projectsFile, buf.Bytes(), 0644)
	if err != nil {
		output.Printf(output.VERBOSE, "Error writing projects file: %s", err.Error())
		return
	}
}




