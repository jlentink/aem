package project

import (
	"errors"
	"github.com/spf13/afero"
	"os"
)

// WriteTextFile write textfile to disk
func WriteTextFile(path, content string) (int, error) {
	bytes := 0
	f, err := Create(path)

	if err == nil {
		defer f.Close() // nolint: errcheck
		bytes, err = f.WriteString(content)
	}
	return bytes, err
}

// WriteGitIgnoreFile Write ignore file to disk
func WriteGitIgnoreFile() (string, error) {
	path, err := getIgnoreFileLocation()
	if err != nil {
		return ``, err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		content := []byte(configInstanceGitIgnoreContent)
		err := afero.WriteFile(fs, path, content, 0644)
		if err != nil {
			return ``, errors.New("could not create ignore file")
		}
	}
	return path, nil
}
