package aem

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/zalando/go-keyring"
)

const (
	serviceNameID = "aem-cli"
)

func serviceName(i objects.Instance) string {
	return serviceNameID + "-" + i.Name + "-" + i.Hostname
}

// KeyRingSetPassword Store password in keyring.
func KeyRingSetPassword(i objects.Instance, password string) error {
	return keyring.Set(serviceName(i), i.Username, password)
}

// KeyRingGetPassword Store password from keyring.
func KeyRingGetPassword(i objects.Instance) (string, error) {
	return keyring.Get(serviceName(i), i.Username)
}

// GetPasswordForInstance get password for instance.
func GetPasswordForInstance(i objects.Instance, useKeyring bool) (string, error) {
	if useKeyring {
		pw, err := keyring.Get(serviceName(i), i.Username)
		if err != nil {
			return ``, fmt.Errorf("could not get password from keychain for %s", i.Name)
		}
		return pw, nil
	}
	return i.Password, nil
}
