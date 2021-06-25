package objects

import (
	"strings"
)

// PackageList description struct
type PackageList struct {
	Results []Package `json:"results"`
	Total   int       `json:"total"`
}

// Package description struct
type Package struct {
	Pid               string              `json:"pid"`
	Path              string              `json:"path"`
	Name              string              `json:"name"`
	DownloadName      string              `json:"downloadName"`
	Group             string              `json:"group"`
	GroupTitle        string              `json:"groupTitle"`
	Version           string              `json:"version"`
	Description       string              `json:"description"`
	Thumbnail         string              `json:"thumbnail"`
	BuildCount        int                 `json:"buildCount"`
	Created           int64               `json:"created,omitempty"`
	CreatedStr        string              ``
	CreatedBy         string              `json:"createdBy,omitempty"`
	LastUnpacked      int64               `json:"lastUnpacked"`
	LastUnpackedBy    string              `json:"lastUnpackedBy"`
	LastUnwrapped     int64               `json:"lastUnwrapped"`
	SizeHuman         string              ``
	Size              int                 `json:"size"`
	HasSnapshot       bool                `json:"hasSnapshot"`
	NeedsRewrap       bool                `json:"needsRewrap"`
	RequiresRoot      bool                `json:"requiresRoot"`
	RequiresRestart   bool                `json:"requiresRestart"`
	AcHandling        string              `json:"acHandling"`
	Dependencies      []PackageDependency `json:"dependencies"`
	Resolved          bool                `json:"resolved"`
	Filter            []Filter            `json:"filter"`
	Screenshots       []string            `json:"screenshots"`
	ProviderName      string              `json:"providerName,omitempty"`
	ProviderURL       string              `json:"providerUrl,omitempty"`
	ProviderLink      string              `json:"providerLink,omitempty"`
	BuiltWith         string              `json:"builtWith,omitempty"`
	TestedWith        string              `json:"testedWith,omitempty"`
	LastUnwrappedBy   string              `json:"lastUnwrappedBy,omitempty"`
	LastModified      int64               `json:"lastModified,omitempty"`
	LastModifiedBy    string              `json:"lastModifiedBy,omitempty"`
	LastModifiedByStr string              ``
	LastWrapped       int64               `json:"lastWrapped,omitempty"`
	LastWrappedStr    string              ``
	LastWrappedBy     string              `json:"lastWrappedBy,omitempty"`
}

// Equals check if packages are the same
func (p *Package) Equals(comparePackage *Package) bool {
	if p.Name == comparePackage.Name && p.Version == comparePackage.Version {
		return true
	}
	return false
}

// FromString Get package from String
func (p *Package) FromString(pkgStr string) bool {
	parts := strings.Split(pkgStr, ",")
	if len(parts) != 4 {
		return false
	}
	p.Name = parts[0]
	p.Version = parts[1]
	p.Group = parts[2]
	p.DownloadName = parts[3]
	return true
}

// PackageDependency description struct
type PackageDependency struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// PackageFilter description struct
type PackageFilter struct {
	Root  string        `json:"root"`
	Rules []PackageRule `json:"rules"`
}

// PackageRule description struct
type PackageRule struct {
	Modifier string `json:"modifier"`
	Pattern  string `json:"pattern"`
}
