package generate

import (
	"github.com/BurntSushi/toml"
	"github.com/jlentink/aem/internal/output"
	"io/ioutil"
	"os"
)

const templateDir = "templates"
const settingsFile = "settings.toml"

// ListComponents in templates folder
func ListComponents() ([]*Component, error) {
	templateRoot, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	components := make([]*Component, 0)

	templateRoot = templateRoot + "/" + templateDir
	entries, err := ioutil.ReadDir(templateRoot)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if _, err := os.Stat(templateRoot + "/" + entry.Name() + "/" + settingsFile); os.IsNotExist(err) {
				output.Printf(output.VERBOSE, "Did not find settingsfile for %s\n", entry.Name())
				continue
			}

			component := Component{}
			_, err := toml.DecodeFile(templateRoot+"/"+entry.Name()+"/"+settingsFile, &component)
			if err != nil {
				output.Printf(output.VERBOSE, "Error in settings file of %s - %s\n", entry.Name(), err.Error())
				continue
			}
			component.SourcePath = templateRoot + "/" + entry.Name()
			components = append(components, &component)
		} else {
			continue
		}
	}
	return components, nil
}
