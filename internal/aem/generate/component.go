package generate

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"io/ioutil"
	"os"
)

// ComponentField defines a input for the component
type ComponentField struct {
	Name     string   `toml:"name"`
	Question string   `toml:"question"`
	Help     string   `toml:"help"`
	Type     string   `toml:"type"`
	Options  []string `toml:"options"`
	Default  string   `toml:"default"`
	Value    string
}

// Component to render
type Component struct {
	Name        string           `toml:"name"`
	Description string           `toml:"description"`
	Destination string           `toml:"destination"`
	Version     string           `toml:"version"`
	Fields      []ComponentField `toml:"fields"`
	SourcePath  string
	ValueMap    map[string]string
}

func (c *Component) renderField(field *ComponentField) error {

	defaultValue, err := ParseTemplate(field.Default, c.ValueMap)
	if err != nil {
		defaultValue = field.Default
	}

	switch field.Type {
	case "select":
		value := ""
		prompt := &survey.Select{
			Options: field.Options,
			Message: field.Question,
			Help:    field.Help,
			Default: defaultValue,
		}
		err := survey.AskOne(prompt, &value)
		field.Value = value
		return err
	case "input":
		value := ""
		prompt := &survey.Input{
			Message: field.Question,
			Help:    field.Help,
			Default: defaultValue,
		}
		err := survey.AskOne(prompt, &value)
		field.Value = value
		return err
	default:
		return fmt.Errorf("unknown field type %s", field.Type)
	}
}

func (c *Component) requestInput() error {
	if c.ValueMap == nil {
		c.ValueMap = make(map[string]string)
	}
	c.ValueMap["name"] = c.Name
	for i := range c.Fields {
		err := c.renderField(&c.Fields[i])
		if err != nil {
			return err
		}
		c.ValueMap[c.Fields[i].Name] = c.Fields[i].Value
	}
	return nil
}

func (c *Component) render(src, dest string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	e, err := c.exists(dest)
	if err != nil {
		return err
	}

	if !e {
		err := os.MkdirAll(dest, 0755)
		if err != nil {
			return err
		}
	}
	for _, entry := range entries {
		if entry.Name() == "settings.toml" {
			continue
		}
		name, err := ParseTemplate(entry.Name(), c.ValueMap)
		if err != nil {
			cErr := fmt.Errorf("error while processing filename template name %s at location: %s: %s", entry.Name(), dest, err.Error())
			return cErr
		}
		if entry.IsDir() {
			err = c.render(src+"/"+entry.Name(), dest+"/"+name)
			if err != nil {
				return err
			}
		} else {
			e, err := c.exists(dest + "/" + name)
			if err != nil {
				return err
			}

			if e {
				continue
			}
			input, err := ioutil.ReadFile(src + "/" + entry.Name())
			if err != nil {
				cErr := fmt.Errorf("error while reading template name %s at location: %s: %s", entry.Name(), dest, err.Error())
				return cErr
			}
			content, err := ParseTemplate(string(input), c.ValueMap)
			if err != nil {
				cErr := fmt.Errorf("error while processing template %s at location: %s: %s", entry.Name(), src, err.Error())
				return cErr
			}
			err = ioutil.WriteFile(dest+"/"+name, []byte(content), 0644)
			if err != nil {
				cErr := fmt.Errorf("error while writing template %s at location: %s: %s", entry.Name(), src, err.Error())
				return cErr
			}
		}
	}

	return nil
}

func (c *Component) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Generate component and write to disk
func (c *Component) Generate() error {
	err := c.requestInput()
	if err != nil {
		return err
	}
	err = c.render(c.SourcePath, c.Destination)
	if err != nil {
		return err
	}
	return nil
}
