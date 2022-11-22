package configvariables

import (
	"fmt"
	"os"
	"strings"
)

type envVariable struct {
	raw          string
	name         string
	defaultValue string
	varType      string
}

func (v *envVariable) Replace(content string) string {
	return strings.Replace(content, fmt.Sprintf("${%s}", v.raw), v.Value(), -1)
}

func (v *envVariable) Name() string {
	return v.name
}

func (v *envVariable) DefaultValue() string {
	return v.defaultValue
}

func (v *envVariable) Variable(content string) {
	content = content[2 : len(content)-1]
	typeSeperator := strings.Index(content, ":")
	sepLocation := strings.Index(content, "~")
	v.raw = content
	if typeSeperator > 0 {
		v.varType = strings.ToUpper(content[0:typeSeperator])
	} else {
		v.varType = "ENV"
	}
	if sepLocation > 0 {
		v.name = content[typeSeperator+1 : sepLocation]
		v.defaultValue = content[sepLocation+1:]
	} else {
		v.name = content[typeSeperator+1:]
	}
}

func (v *envVariable) Value() string {
	value := os.Getenv(v.name)
	if value == "" && v.defaultValue != "" {
		return v.defaultValue
	}
	return value
}

func newEnvVariable(content string) envVariable {
	variable := envVariable{}
	variable.Variable(content)
	return variable
}
