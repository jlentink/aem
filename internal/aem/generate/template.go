package generate

import (
	"bytes"
	"github.com/jlentink/aem/internal/casetypes"
	"text/template"
)

func templateFuncs() map[string]interface{} {
	return template.FuncMap{
		"CC": casetypes.CamelCase,
		"UC": casetypes.UpperCase,
		"LC": casetypes.LowerCase,
		"PC": casetypes.PascalCase,
		"SC": casetypes.SnakeCase,
		"KC": casetypes.KababCase,
		"TC": casetypes.TitleCase,
	}
}

// ParseTemplate parses template string with map of variables
func ParseTemplate(tpl string, vars map[string]string) (string, error) {
	t, err := template.New("toTemplate").Funcs(templateFuncs()).Parse(tpl)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = t.Execute(&b, vars)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
