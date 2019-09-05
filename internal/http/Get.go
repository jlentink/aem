package http

import (
	"fmt"
	"github.com/go-http-utils/headers"
	"github.com/jlentink/aem/internal/version"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetPlain does a plain get request to an URL
func GetPlain(uri string, username string, password string) ([]byte, error) {
	URL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return Get(URL)
}

// GetPlainWithHeaders Do a plain request with specific headers
func GetPlainWithHeaders(uri string, username string, password string, header []Header) ([]byte, error) {
	URL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return GetWithHeaders(URL, header)
}

// GetWithHeaders Do a get request with url.URL
func GetWithHeaders(uri *url.URL, header []Header) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, URLToURLString(uri), nil)
	for _, h := range header {
		req.Header.Add(h.Key, h.Value)
	}
	req.Header.Add(headers.UserAgent, "aemCLI - "+version.GetVersion())
	req.URL = uri

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode >= 400 {
		return body, fmt.Errorf("received unexpected HTTP status. (%d)", resp.StatusCode)
	}

	return body, nil
}

// Get Perform a simple get request
func Get(uri *url.URL) ([]byte, error) {
	return GetWithHeaders(uri, []Header{{Key: headers.CacheControl, Value: configNoCache}})
}
