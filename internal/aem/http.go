package aem

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/http"
	"net/url"
)

const (
	// URLSystemInformation system overview URL
	URLSystemInformation = "/libs/granite/operations/content/systemoverview/export.json"
)

// URL for instance (*url.URL)
func URL(i *objects.Instance, uri string) (*url.URL, error) {
	u, err := url.Parse(URLString(i) + uri)
	if err != nil {
		return nil, err
	}

	pwd, err := i.GetPassword()
	if err != nil {
		return nil, err
	}

	u.User = url.UserPassword(i.Username, pwd)
	return u, nil
}

// URLString for instance
func URLString(i *objects.Instance) string {
	return fmt.Sprintf("%s://%s:%d", i.Protocol, i.Hostname, i.Port)
}

// PasswordURL with credentials for instance
func PasswordURL(i objects.Instance, useKeyring bool) (string, error) {
	password, err := GetPasswordForInstance(i, useKeyring)
	if err != nil {
		return ``, nil
	}
	return fmt.Sprintf("%s://%s:%s@%s:%d", i.Protocol, url.QueryEscape(i.Username), url.QueryEscape(password), i.Hostname, i.Port), nil
}

// GetFromInstance Perform a get request towards an instance
func GetFromInstance(i *objects.Instance, uri string) ([]byte, error) {
	if !Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	u, err := URL(i, uri)
	if err != nil {
		return []byte{}, err
	}

	return http.Get(u)
}
