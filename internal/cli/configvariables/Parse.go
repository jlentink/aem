package configvariables

import (
	"fmt"
	"regexp"
)

func getVariables(content string) []string {
	r, _ := regexp.Compile("\\${(.*?)}")
	matches := r.FindAllString(content, -1)
	fmt.Printf("%s\n", matches)
	return matches
}

func ReplaceVariables(content string) string {
	for _, v := range getVariables(content) {
		variable := newEnvVariable(v)
		content = variable.Replace(content)
	}

	return content
}
