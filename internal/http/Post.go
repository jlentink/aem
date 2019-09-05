package http

import (
	"bytes"
	"fmt"
	"github.com/go-http-utils/headers"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PostPlain Simple HTTP POST with url string
func PostPlain(uri string, username string, password string, body *bytes.Buffer) ([]byte, error) {
	URL, err := url.Parse(uri)

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)

	}
	if err != nil {
		return nil, err
	}

	return PostWithHeaders(URL, body, []Header{})
}

// PostPlainWithHeaders Simple HTTP POST with url string and define the headers
func PostPlainWithHeaders(uri string, username string, password string, body *bytes.Buffer, header []Header) ([]byte, error) {
	URL, err := url.Parse(uri)

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)

	}
	if err != nil {
		return nil, err
	}

	return PostWithHeaders(URL, body, header)
}

// Post Perform a HTTP POST
func Post(uri *url.URL, body *bytes.Buffer) ([]byte, error) {
	return PostWithHeaders(uri, body, []Header{})
}

// PostWithHeaders Do a HTTP POST and define the headers used
func PostWithHeaders(uri *url.URL, body *bytes.Buffer, header []Header) ([]byte, error) {
	if nil == body {
		body = &bytes.Buffer{}
	}
	req, _ := http.NewRequest(http.MethodPost, URLToURLString(uri), body)
	req.Header.Add(headers.CacheControl, configNoCache)
	for _, h := range header {
		req.Header.Add(h.Key, h.Value)
	}
	req.URL = uri

	if p, ps := uri.User.Password(); ps {
		req.SetBasicAuth(uri.User.Username(), p)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode >= 400 {
		return rBody, fmt.Errorf("received unexpected HTTP status. (%d)", resp.StatusCode)
	}

	return rBody, nil
}
