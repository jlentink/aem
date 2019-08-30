package dispatcher

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/http"
)

const (
	invalidateURL = "/dispatcher/invalidate.cache"
)

// Invalidate sent invalidate to dispatcher
func Invalidate(i *objects.Instance, path string) error {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	header := []http.Header{
		{Key: "CQ-Action", Value: "Activate"},
		{Key: "Content-Length", Value: "0"},
		{Key: "Content-Type", Value: " application/octet-stream"},
	}

	_, err := http.GetPlainWithHeaders(i.URLString()+invalidateURL, i.Username, i.Password, header)
	if err != nil {
		return fmt.Errorf("could not invalidate path: %s", err.Error())
	}

	return nil
}
