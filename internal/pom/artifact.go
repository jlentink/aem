package pom

import "fmt"

// Artifact types
const (
	Unknown = iota
	Bundle
	Package
	All
)

// Artifact holding struct
type Artifact struct {
	Name      string
	BasePath  string
	Path      string
	ID        string
	Version   string
	Packaging string
	Parent    string
}

// CompletePath complete path to artifact
func (a *Artifact) CompletePath() string {
	return fmt.Sprintf("%s%s", a.Path, a.Filename())
}

// Kind What is the kind of the artifact
func (a *Artifact) Kind() int {
	switch a.Packaging {
	case "content-package":
		return Package
	case "bundle":
		return Bundle
	default:
		return Unknown
	}
}

// Filename get the filename for the artifact
func (a *Artifact) Filename() string {
	switch a.Kind() {
	case Package:
		return fmt.Sprintf("%s-%s.zip", a.ID, a.Version)
	case Bundle:
		return fmt.Sprintf("%s-%s.jar", a.ID, a.Version)
	default:
		return fmt.Sprintf("%s-%s.unkown", a.ID, a.Version)
	}
}

// PakageName returns the package name
func (a *Artifact) PakageName() string {
	return a.ID
}
