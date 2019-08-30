package manifest

import "errors"

// Manifest holds the manifest values
type Manifest struct {
	keyValues map[string]string
}

// Get a value from the manifest based on the label
func (m *Manifest) Get(manifestLabel string) (string, error) {
	if _, ok := m.keyValues[manifestLabel]; ok {
		return m.keyValues[manifestLabel], nil
	}

	return "", errors.New("could not find the key")
}
