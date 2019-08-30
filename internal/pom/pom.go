package pom

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/jlentink/aem/internal/cli/project"
	"os"
	"path/filepath"
)

// Pom is a java pom file parser
type Pom struct {
	name     string
	path     string
	doc      *xmlquery.Node
	children []*xmlquery.Node
}

// Open a pomfile based on path
func (p *Pom) Open(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("could not find: %s", path)
	}

	p.path = path
	f, err := project.Open(path)
	if err != nil {
		return err
	}
	p.doc, err = xmlquery.Parse(f)
	return err
}

func (p *Pom) hasParent() bool {
	if g := xmlquery.FindOne(p.doc, "//project//parent"); g != nil {
		return true
	}
	return false
}

func (p *Pom) basePath() string {
	return filepath.Dir(p.path)
}

// GetModules gets the modules available in the pom file
func (p *Pom) GetModules() []string {
	m := xmlquery.Find(p.doc, "//project//modules//module")
	var modules []string
	if m == nil {
		return modules
	}

	for _, module := range m {
		modules = append(modules, module.InnerText())
	}
	return modules
}

// GetName Get name from pom
func (p *Pom) GetName() string {
	return p.name
}

// GetModulePOMs gets the poms of the modules
func (p *Pom) GetModulePOMs() ([]*Pom, error) {
	var poms []*Pom
	for _, module := range p.GetModules() {
		tp := Pom{name: module}
		err := tp.Open(p.basePath() + "/" + module + "/pom.xml")
		if err != nil {
			return nil, fmt.Errorf("could not find pom for module '%s'", module)
		}
		poms = append(poms, &tp)
	}
	return poms, nil
}

func (p *Pom) getInnerTextForPath(xpath string) string {
	artifactNode := xmlquery.FindOne(p.doc, xpath)
	if artifactNode != nil {
		return artifactNode.InnerText()
	}
	return ""
}

// GetArtifact from the pom
func (p *Pom) GetArtifact() Artifact {
	artifact := Artifact{}
	artifact.ID = p.getInnerTextForPath("//project/artifactId")
	artifact.Path = p.basePath() + "/target/"
	artifact.BasePath = p.basePath()
	artifact.Version = p.getInnerTextForPath("//project/parent/version")
	artifact.Packaging = p.getInnerTextForPath("//project/packaging")
	return artifact
}

// GetArtifactByName from pom by name
func (p *Pom) GetArtifactByName(name string) (*Artifact, error) {
	poms, err := p.GetModulePOMs()
	if err != nil {
		return nil, err
	}
	for _, pom := range poms {
		if pom.name == name {
			artifact := pom.GetArtifact()
			return &artifact, nil
		}
	}
	return nil, fmt.Errorf("could not find artifact")
}

// GetAllArtifacts Get all artifacts:59
func (p *Pom) GetAllArtifacts(filter int) ([]*Artifact, error) {
	artifacts := make([]*Artifact, 0)
	poms, err := p.GetModulePOMs()
	if err != nil {
		return nil, err
	}
	for _, pom := range poms {
		a := pom.GetArtifact()
		if a.Kind() == Package && filter == Package {
			artifacts = append(artifacts, &a)
		} else if a.Kind() == Bundle && filter == Bundle {
			artifacts = append(artifacts, &a)
		} else if filter == All {
			artifacts = append(artifacts, &a)
		}

	}
	return artifacts, nil
}
