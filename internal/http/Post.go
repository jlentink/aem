package http

import (
	"bytes"
	"fmt"
	"github.com/go-http-utils/headers"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)



// PostPlain Simple HTTP POST with url string
func PostPlain(uri string, username string, password string, body *bytes.Buffer) ([]byte, *http.Response, error) {
	URL, err := setupPlainToURL(uri, username, password)
	if err != nil {
		return nil, nil, err
	}

	return PostWithHeaders(URL, body, []Header{})
}

// PostPlainWithHeaders Simple HTTP POST with url string and define the headers
func PostPlainWithHeaders(uri string, username string, password string, body *bytes.Buffer, header []Header) ([]byte, *http.Response, error) {
	URL, err := setupPlainToURL(uri, username, password)
	if err != nil {
		return nil, nil, err
	}
	return PostWithHeaders(URL, body, header)
}

// Post Perform a HTTP POST
func Post(uri *url.URL, body *bytes.Buffer) ([]byte, *http.Response, error) {
	return PostWithHeaders(uri, body, []Header{})
}

// PostMultiPart post multipart form
func PostMultiPart(uri string, username string, password string, data map[string]string) ([]byte, *http.Response, error) {
	return PostMultiPartWithHeaders(uri, username, password, data, []Header{})
}

// PostMultiPart Post multipart form
// With the ability to add headers
func PostMultiPartWithHeaders(uri string, username string, password string, data map[string]string, header []Header) ([]byte, *http.Response, error) {
	URL, err := setupPlainToURL(uri, username, password)
	if err != nil {
		return nil, nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)  // nolint: errcheck
	for key, val := range data {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, nil, err
		}
	}
	writer.Close()
	reqHeaders := append(header, Header{Key: headers.ContentType, Value: writer.FormDataContentType()})

	return PostWithHeaders(URL, body, reqHeaders)
}

// PostFormEncode post form encoded form
func PostFormEncode(uri string, username string, password string, data url.Values) ([]byte, *http.Response, error) {
	URL, err := setupPlainToURL(uri, username, password)
	if err != nil {
		return nil, nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, uri, strings.NewReader(data.Encode()))
	req.Header.Add(headers.ContentType, ApplicationFormEncode)
	req.Header.Add(headers.ContentLength, strconv.Itoa(len(data.Encode())))
	req.URL = URL

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, resp, err
	}

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, resp, err
	}

	if resp.StatusCode >= 400 {
		return rBody, resp, fmt.Errorf("received unexpected HTTP status. (%d)", resp.StatusCode)
	}

	return rBody, resp, nil
}


// PostWithHeaders Do a HTTP POST and define the headers used
func PostWithHeaders(uri *url.URL, body *bytes.Buffer, header []Header) ([]byte, *http.Response, error) {
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
		return []byte{}, resp, err
	}
	defer resp.Body.Close() // nolint: errcheck

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, resp, err
	}

	if resp.StatusCode >= 400 {
		return rBody, resp, fmt.Errorf("received unexpected HTTP status. (%d)", resp.StatusCode)
	}

	return rBody, resp, nil
}
