package cachedir

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/jlentink/aem/internal/cli/project"
	"io/ioutil"
)

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
	if project.Exists(getCacheRoot() + "/projects.toml") {
		toml.DecodeFile(getCacheRoot() + "/projects.toml", &projects) // nolint: errcheck
	}
	return projects.Project
}

func RegisterPorject(name, path string){
	mutated := false
	projects := RegisteredProjects()
	for _, project := range projects {
		if project.Name == name && project.Path == path {
			return
		}
		if project.Name == name {
			project.Path = path
			mutated = true
		}
	}

	if !mutated {
		projects = append(projects, ProjectRegistered{Name: name, Path: path})
		mutated = true
	}

	if mutated {
		writeRegisterFile(projects)
	}
}

func writeRegisterFile(projects []ProjectRegistered) {
	data := Projects{Project: projects}
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(data)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(getCacheRoot() + "/projects.toml", buf.Bytes(), 0644)
	if err != nil {
		return
	}

}




