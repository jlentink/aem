package casetypes

import (
	"regexp"
	"strings"
)

func tokenize(input string) []string {
	return strings.Split(input, ` `)
}

func normalize(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z\d]`)
	input = re.ReplaceAllString(input, " ")
	return strings.TrimSpace(input)
}

// LowerCase transform all characters to lower case. Alias of strings.ToLower
func LowerCase(input string) string {
	return strings.ToLower(input)
}

// UpperCase transform all characters to upper case. Alias of strings.UpperCase
func UpperCase(input string) string {
	return strings.ToUpper(input)
}

// CamelCase transform string to camel case (camelCase)
func CamelCase(input string) string {
	output := ""
	for i, val := range tokenize(normalize(input)) {
		tok := LowerCase(val)
		if i > 0 {
			tok = strings.Title(tok)
		}
		output += tok
	}
	return output
}

// PascalCase transform string to pascal case (PascalCase)
func PascalCase(input string) string {
	output := ""
	for _, val := range tokenize(normalize(input)) {
		tok := LowerCase(val)
		tok = strings.Title(tok)
		output += tok
	}
	return output
}

// KababCase transform string to kabab case (kabab-case)
func KababCase(input string) string {
	input = LowerCase(normalize(input))
	input = strings.ReplaceAll(input, " ", "-")
	return input
}

// SnakeCase transform string to snake case (slug_case)
func SnakeCase(input string) string {
	input = LowerCase(normalize(input))
	re := regexp.MustCompile(`[^a-zA-Z\d]`)
	return re.ReplaceAllString(input, "_")
}

// TitleCase returns a copy of the string s with all Unicode letters that begin words
// mapped to their title case.
func TitleCase(input string) string {
	input = LowerCase(input)
	return UpperCase(input[0:1]) + input[1:]
}
