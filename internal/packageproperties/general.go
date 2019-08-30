package packageproperties

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	propertiesPath = "META-INF/vault/properties.xml"
)

var (
	currentKey string
	keyValues  map[string]string
)

// Open properties file of package
func Open(location string) (*Properties, error) {
	reader, err := zip.OpenReader(location)
	if err != nil {
		return nil, err
	}

	for _, file := range reader.File {
		if file.Name == propertiesPath {
			properties, err := file.Open()
			defer properties.Close()
			if err != nil {
				return nil, err
			}
			return parseProperties(properties)
		}
	}
	return nil, fmt.Errorf("could not find properties file in zip")
}

// OpenXML properties file for reading
func OpenXML(location string) (*Properties, error) {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return nil, fmt.Errorf("could not find file at: %s", err)
	}

	properties, err := os.Open(location)
	if err != nil {
		return nil, err
	}

	return parseProperties(properties)
}

func parseProperties(r io.Reader) (*Properties, error) {
	keyValues = make(map[string]string, 0)
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	rawXML := propertiesXML{}
	err = xml.Unmarshal(d, &rawXML)
	if err != nil {
		return nil, err
	}

	for _, entry := range rawXML.Entry {
		keyValues[entry.Key] = entry.Text
	}

	return &Properties{Property: keyValues}, nil
}
