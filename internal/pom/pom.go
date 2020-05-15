package pom

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/jlentink/aem/internal/cli/project"
	"os"
	"path/filepath"
	"strings"
)

// NewPom creates a pom object from path
func NewPom(path string) Pom {
	p := Pom{}
	//nolint - for later use
	p.Open(path)
	return p
}

// Pom is a java pom file parser
type Pom struct {
	name     string
	path     string
	children []*Pom
	doc      *xmlquery.Node
	err      error
}

// Open a pomfile based on path
func (p *Pom) Open(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		p.err = err
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

//nolint - for later use
func (p *Pom) hasParent() bool {
	if g := xmlquery.FindOne(p.doc, "/project/parent"); g != nil {
		return true
	}
	return false
}

func (p *Pom) basePath() string {
	return filepath.Dir(p.path)
}

// GetAllPomModules gets the modules available in the pom file
func (p *Pom) GetAllPomModules() {
	m := xmlquery.Find(p.doc, "/project/modules/module")
	for _, module := range m {
		mPom := NewPom(fmt.Sprintf("%s/%s/pom.xml", p.basePath(), module.InnerText()))
		p.children = append(p.children, &mPom)
		mPom.GetAllPomModules()
	}
}

// GetChildren Get children of this pom
func (p *Pom) GetChildren() []*Pom {
	return p.children
}

func (p *Pom) getInnerTextForPath(node *xmlquery.Node, xpath string) string {
	artifactNode := xmlquery.FindOne(node, xpath)
	if artifactNode != nil {
		return artifactNode.InnerText()
	}
	return ""
}

// GetArtifact from the pom
func (p *Pom) GetArtifact() Artifact {
	artifact := Artifact{}
	artifact.ID = p.getInnerTextForPath(p.doc, "//project/artifactId")
	artifact.Parent = p.getInnerTextForPath(p.doc, "//project/parent/artifactId")
	artifact.Path = p.basePath() + "/target/"
	artifact.BasePath = p.basePath()
	artifact.Version = p.getInnerTextForPath(p.doc, "//project/parent/version")
	artifact.Packaging = p.getInnerTextForPath(p.doc, "//project/packaging")
	if artifact.Packaging == "" {
		plugins := xmlquery.Find(p.doc, "//project/build/plugins/plugin")
		for _, plugin := range plugins {
			artifactId := p.getInnerTextForPath(plugin, "/artifactId")
			bdn := p.getInnerTextForPath(plugin, "executions/execution/configuration/bnd")
			if artifactId == "bnd-maven-plugin" && strings.Contains(bdn, "Import-Package") {
				artifact.Packaging = "bundle"
			}
		}
	}
	return artifact
}

// GetAllArtifacts gets all child artifacts
func (p *Pom) GetAllArtifacts(filter int) ([]Artifact, error) {
	artifacts := make([]Artifact, 0)
	p.GetAllPomModules()
	poms := p.flatPomList()
	for _, p := range poms {
		a := p.GetArtifact()
		if a.Kind() == Package && filter == Package {
			artifacts = append(artifacts, a)
		} else if a.Kind() == Bundle && filter == Bundle {
			artifacts = append(artifacts, a)
		} else if filter == All {
			artifacts = append(artifacts, a)
		}
	}
	return artifacts, nil
}

// GetArtifactByName find artifact by name
func (p *Pom) GetArtifactByName(name string) (*Artifact, error) {
	poms := p.flatPomList()
	for _, pom := range poms {
		if pom.name == name {
			artifact := pom.GetArtifact()
			return &artifact, nil
		}
	}
	return nil, fmt.Errorf("could not find artifact")
}

func (p *Pom) flatPomList() []*Pom {
	poms := make([]*Pom, 0)
	poms = append(poms, p)

	for _, c := range p.children {
		poms = append(poms, c.flatPomList()...)
	}
	return poms
}

/*
// GetModules gets the modules available in the pom file
func (p *Pom) GetModules() []string {
	m := xmlquery.Find(p.doc, "/project/modules/module")
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
		if len(tp.GetModules()) > 0 {
			p.GetModulePOMs()
		}
		poms = append(poms, &tp)
	}
	return poms, nil
}
*/
/*


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
*/
/*
// GetAllArtifacts Get all artifacts:59
func (p *Pom) GetAllArtifacts(filter int) ([]*Artifact, error) {
	artifacts := make([]*Artifact, 0)
	poms := p.GetAllModules()
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
*/
