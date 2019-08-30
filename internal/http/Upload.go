package http

import (
	"bytes"
	"github.com/go-http-utils/headers"
	"github.com/jlentink/aem/internal/version"
	"net/http"
	"net/url"
)

// UploadWithURL upload bytes.buffer with *url.URL
func UploadWithURL(uri *url.URL, body *bytes.Buffer, httpHeaders []Header) (*http.Response, error) {
	pr := &progressReporter{
		r:         body,
		totalSize: uint64(body.Len()),
		label:     "Uploading",
	}

	req, _ := http.NewRequest(http.MethodPost, ``, pr)
	req.URL = uri

	for _, header := range httpHeaders {
		req.Header.Add(header.Key, header.Value)
	}

	req.Header.Add(headers.UserAgent, "aemCLI - "+version.GetVersion())
	if p, ps := uri.User.Password(); ps {
		req.SetBasicAuth(uri.User.Username(), p)
	}

	client := &http.Client{}
	return client.Do(req)
}

// Upload bytes.buffer with string username, password and uri
func Upload(uri, username, password string, body *bytes.Buffer, httpHeaders []Header) (*http.Response, error) {
	URL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)
	}

	return UploadWithURL(URL, body, httpHeaders)
}
