package dispatcher

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/http"
	"github.com/jlentink/aem/internal/output"
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
		{Key: "Content-Type", Value: "application/octet-stream"},
		{Key: "CQ-Path", Value: path},
		{Key: "CQ-Handle", Value: path},
	}
	_, err := http.GetPlainWithHeaders(i.URLString()+invalidateURL, i.Username, i.Password, header)
	if err != nil {
		return fmt.Errorf("could not invalidate path: %s", err.Error())
	}

	return nil
}

// InvalidateAll sent invalidate to all dispatchers
func InvalidateAll(is []objects.Instance, p []string) bool {
	var status = true
	for idx, i := range is {
		for _, path := range p {
			output.Printf(output.NORMAL, "\U0001F5D1 Invalidating: %s (%s)\n", path, i.Name)
			err := Invalidate(&is[idx], path)
			if err != nil {
				output.Printf(output.NORMAL, "Could not invalidate path: %s\n", err.Error())
				status = false
			}
		}
	}
	return status
}
