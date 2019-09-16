package http

import (
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/go-http-utils/headers"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/version"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func downloadSize(req *http.Request) (uint64, error) {
	method := req.Method
	req.Method = http.MethodHead

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		output.Print(output.VERBOSE, "unable to create http client\n")
		return 0, errors.New("unable to create http client")
	}
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		req.Method = method
		output.Printf(output.VERBOSE, "received wrong http status %d", resp.StatusCode)
		return 0, fmt.Errorf("received wrong http status %d", resp.StatusCode)
	}

	size, err := strconv.Atoi(resp.Header.Get(headers.ContentLength))
	if err != nil {
		req.Method = method
		return 0, fmt.Errorf("could not create int from response %s", resp.Header.Get(headers.ContentLength))
	}
	output.Printf(output.VERBOSE, "Download size found: %d bytes\n", size)
	req.Method = method
	return uint64(size), nil
}

// DownloadFileWithURL file with url.URL
func DownloadFileWithURL(destination string, uri *url.URL, forceDownload bool) (uint64, error) {
	if project.Exists(destination) && !forceDownload {
		return 0, nil
	}

	req, _ := http.NewRequest(http.MethodGet, "", nil)
	req.Header.Add(headers.CacheControl, configNoCache)
	req.Header.Add(headers.UserAgent, "aemCLI - "+version.GetVersion())
	req.URL = uri

	fs, _ := downloadSize(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() // nolint: errcheck

	fmt.Printf("Downloading: %s (%s)\n", uri.Path, humanize.Bytes(fs))
	counter := &progressReporter{r: resp.Body, totalSize: fs, label: "Downloading"}
	out, err := project.Create(destination + ".tmp")
	if err != nil {
		output.Printf(output.VERBOSE, "unable to create tmp file %s", destination+".tmp")
		return 0, err
	}

	_, err = io.Copy(out, counter)
	if err != nil {
		return 0, err
	}

	err = out.Close()
	if err != nil {
		return 0, err
	}
	err = project.Rename(destination+".tmp", destination)
	if err != nil {
		fmt.Print("\n")
		return 0, err
	}

	fmt.Print("\n")
	return fs, nil
}

// DownloadFile download file from string url
func DownloadFile(destination string, uri string, username string, password string, forceDownload bool) (uint64, error) {
	URL, err := url.Parse(uri)
	if err != nil {
		return 0, err
	}

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)

	}
	return DownloadFileWithURL(destination, URL, forceDownload)
}
