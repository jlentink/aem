package packageproperties

import (
	"encoding/xml"
	"fmt"
)

// Label descriptors
const (
	CreatedBy             = "createdBy"
	AllowIndexDefinitions = "allowIndexDefinitions"
	Name                  = "name"
	GroupId               = "groupId"
	Version               = "version"
	PackageType           = "packageType"
	RequiresRoot          = "requiresRoot"
	Group                 = "group"
	Description           = "description"
	ArtifactId            = "artifactId"
)

type propertiesXML struct {
	XMLName xml.Name `xml:"properties"`
	Text    string   `xml:",chardata"`
	Comment struct {
		Text string `xml:",chardata"`
	} `xml:"comment"`
	Entry []struct {
		Text string `xml:",chardata"`
		Key  string `xml:"key,attr"`
	} `xml:"entry"`
}

// Properties description struct
type Properties struct {
	Comment  string
	Property map[string]string
}

// Get gets the value from the package properties
func (p *Properties) Get(key string) (string, error) {
	if _, ok := p.Property[key]; ok {
		return p.Property[key], nil
	}
	return "", fmt.Errorf("could not find key: %s", key)
}
