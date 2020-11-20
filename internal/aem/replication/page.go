package replication

import (
	"bytes"
	"fmt"
	"github.com/go-http-utils/headers"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/http"
	"mime/multipart"
)

const (
	replicationURL     = "/bin/replicate.json"
	treeReplicationURL = "/etc/replication/treeactivation.html"
)

// Activate page on instance
func Activate(i *objects.Instance, path string) ([]byte, error) {
	return pageAction(i, path, "activate")
}

// Deactivate page on instance
func Deactivate(i *objects.Instance, path string) ([]byte, error) {
	return pageAction(i, path, "deactivate")
}

func pageAction(i *objects.Instance, path, command string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("path", path)   // nolint: errcheck
	writer.WriteField("cmd", command) // nolint: errcheck
	writer.Close()                    // nolint: errcheck

	pw, err := i.GetPassword()
	if err != nil {
		return nil, err
	}

	bodyContent, _, err := http.PostPlainWithHeaders(i.URLString()+replicationURL, i.Username, pw, body,
		[]http.Header{{Key: headers.ContentType, Value: writer.FormDataContentType()}})

	return bodyContent, err
}

// ActivateTree on instance
func ActivateTree(i *objects.Instance, path string, ignoreDeactivated, onlyModified bool) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("cmd", "activate")                                         // nolint: errcheck
	writer.WriteField("ignoredeactivated", fmt.Sprintf("%t", ignoreDeactivated)) // nolint: errcheck
	writer.WriteField("onlymodified", fmt.Sprintf("%t", onlyModified))           // nolint: errcheck
	writer.WriteField("path", path)                                              // nolint: errcheck
	writer.Close()                                                               // nolint: errcheck

	pw, err := i.GetPassword()
	if err != nil {
		return nil, err
	}
	bodyContent, _, err := http.PostPlainWithHeaders(i.URLString()+treeReplicationURL, i.Username, pw, body,
		[]http.Header{{Key: headers.ContentType, Value: writer.FormDataContentType()}})

	return bodyContent, err
}
