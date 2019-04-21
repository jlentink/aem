package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-http-utils/headers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

//
// AEM URLs
//
const (
	URLSystemInformation = "/libs/granite/operations/content/systemoverview/export.json"
	URLActivateTree      = "/etc/replication/treeactivation.html"
	URLBundles           = "/system/console/bundles"
	URLRebuildPackage    = "/crx/packmgr/service/.json%s?cmd=build"
	URLBundleInstall     = "/system/console/bundles/%s"
	URLBundlePage        = "/system/console/bundles/%s"
	URLReplication       = "/bin/replicate.json"
	URLPackageList       = "/crx/packmgr/list.jsp"
	URLPackageEndpoint   = "/crx/packmgr/service.jsp"

	ServiceName = "aem-cli"

	JarContentType = "application/java-archive"

	bundleResponseErrorMsg = "Could not parse bundle response from Adobe Experience manager. (%s)"
	bundleStartLevel       = 20
	bundleFormActionField  = "action"
	bundleInstall          = "install"
	bundleRefresh          = "refreshPackages"
	bundleUninstall        = "uninstall"
)

type httpRequests struct {
}

func (a *httpRequests) getPassword(aeminstance aemInstanceConfig) string {
	i := new(instance)
	return i.getPasswordForInstance(aeminstance)
}

func (a *httpRequests) buildPackage(instance aemInstanceConfig, pkg packageDescription) {

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s"+URLRebuildPackage, instance.URL(), pkg.Path), nil)

	a.addAuthentication(instance, req)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Body : ", string(respBody))
}

func (a *httpRequests) downloadPackage(instance aemInstanceConfig, aemPackage packageDescription, forceDownload bool) (string, error) {
	projectStructure := newProjectStructure()
	projectStructure.createDirForPackage(aemPackage)
	destination := projectStructure.getLocationForPackage(aemPackage)
	url := instance.URL() + aemPackage.Path

	err := a.downloadFile(destination, url, instance.Username, a.getPassword(instance), forceDownload)
	exitFatal(err, "Download issue")

	return destination, err
}

func (a *httpRequests) uploadPackage(instance aemInstanceConfig, aemPackage packageDescription, force bool, install bool) (*crxResponse, error) {
	projectStructure := newProjectStructure()

	fileLocation := projectStructure.getLocationForPackage(aemPackage)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", aemPackage.DownloadName)
	fileContent, err := ioutil.ReadFile(fileLocation)
	exitFatal(err, "Could not read package for upload")

	part.Write(fileContent)
	writer.WriteField("name", aemPackage.Name)
	writer.WriteField("force", strconv.FormatBool(force))
	writer.WriteField("install", strconv.FormatBool(install))
	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, instance.URL()+URLPackageEndpoint, &ProgressReporter{r: body, totalSize: uint64(body.Len()), label: "Uploading"})

	// Headers
	// Set Authentication
	req.SetBasicAuth(instance.Username, a.getPassword(instance))
	req.Header.Add(headers.ContentType, writer.FormDataContentType())

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Print("\n")
	crxResponse := new(crxResponse)

	err = xml.Unmarshal(respBody, crxResponse)

	return crxResponse, err
}

func (a *httpRequests) getSystemInformation(instance aemInstanceConfig) (*systemInformation, error) {

	systemInformation := &systemInformation{}

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", instance.URL(), URLSystemInformation), nil)

	a.addAuthentication(instance, req)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		return systemInformation, err
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(respBody), systemInformation)

	return systemInformation, err

}
func (a *httpRequests) activatePage(instance aemInstanceConfig, path string) int {
	return a.activateDeactivePage(instance, PageStatusActivate, path)
}

func (a *httpRequests) deactivatePage(instance aemInstanceConfig, path string) int {
	return a.activateDeactivePage(instance, PageStatusDeactivate, path)
}

func (a *httpRequests) activateDeactivePage(instance aemInstanceConfig, mode string, path string) int {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("path", path)
	writer.WriteField("cmd", mode)
	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", instance.URL(), URLReplication), body)
	req.Header.Add(headers.ContentType, writer.FormDataContentType())

	a.addAuthentication(instance, req)

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, "Activate page failed")

	// Display Results
	return resp.StatusCode
}

func (a *httpRequests) getListForInstance(instance aemInstanceConfig) []packageDescription {

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodGet, instance.URL()+URLPackageList, nil)

	// Headers
	req.Header.Add(headers.CacheControl, configNoCache)

	a.addAuthentication(instance, req)

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, "Could not retrieve list from Adobe Experience manager.")

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	packageFeed := new(packagesFeed)

	// Parse JSON feed
	err = json.Unmarshal(respBody, packageFeed)
	exitFatal(err, "Could not parse package feed from Adobe Experience manager. (%s)", instance.URL())

	return packageFeed.Package
}

func (a *httpRequests) downloadFile(filepath string, url string, username string, password string, forceDownload bool) error {
	u := new(utility)
	if u.Exists(filepath) && !forceDownload {
		fmt.Printf("Found \"%s\" file skipping...\n", filepath)
		return nil
	}
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Headers
	req.Header.Add(headers.CacheControl, configNoCache)

	if len(username) > 0 || len(password) > 0 {
		req.SetBasicAuth(username, password)
	}

	filesize := a.downloadSize(req)

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, "Could not retrieve list from Adobe Experience manager.")
	defer resp.Body.Close()

	fmt.Printf("%d", resp.StatusCode)

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &ProgressReporter{r: resp.Body, totalSize: filesize, label: "Downloading"}

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		fmt.Print("\n")
		return err
	}

	defer out.Close()

	//_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	_, err = io.Copy(out, counter)

	if err != nil {
		fmt.Print("\n")
		return err
	}

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		fmt.Print("\n")
		return err
	}

	fmt.Print("\n")
	return nil
}

func (a *httpRequests) downloadSize(req *http.Request) uint64 {
	oldMethod := req.Method
	req.Method = http.MethodHead

	client := &http.Client{}

	resp, err := client.Do(req)
	exitFatal(err, "Unable to create Http Client")

	if resp.StatusCode != http.StatusOK {
		req.Method = oldMethod
		return 0
	}

	size, err := strconv.Atoi(resp.Header.Get(headers.ContentLength))
	if err != nil {
		req.Method = oldMethod
		return 0
	}
	req.Method = oldMethod
	return uint64(size)
}

func (a *httpRequests) activateTree(instance aemInstanceConfig, path string) bool {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("cmd", "activate")
	writer.WriteField("ignoredeactivated", "true")
	writer.WriteField("onlymodified", "true")
	writer.WriteField("path", path)
	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, instance.URL()+URLActivateTree, body)

	a.addAuthentication(instance, req)

	req.Header.Add(headers.ContentType, writer.FormDataContentType())

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, "Error during tree activation")

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	}

	return false
}

func (a *httpRequests) listBundles(instance aemInstanceConfig) *bundlesFeed {

	params := url.Values{}
	params.Set(bundleFormActionField, bundleRefresh)
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, instance.URL()+URLBundles, body)

	a.addAuthentication(instance, req)

	// Headers
	req.Header.Add(headers.ContentType, "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add(headers.CacheControl, configNoCache)

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, "Could not retrieve bundle list")

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	bundleFeed := new(bundlesFeed)

	// Parse JSON feed
	err = json.Unmarshal(respBody, bundleFeed)
	exitFatal(err, bundleResponseErrorMsg, instance.URL())

	return bundleFeed
}

func (a *httpRequests) bundleStopStart(instance aemInstanceConfig, bundle bundle, status string) *bundleResponse {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField(bundleFormActionField, status)
	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, instance.URL()+fmt.Sprintf(URLBundlePage, bundle.SymbolicName), body)

	a.addAuthentication(instance, req)

	req.Header.Add(headers.ContentType, writer.FormDataContentType())

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, "Error stopping bundle")

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	bundleResp := new(bundleResponse)

	// Parse JSON feed
	err = json.Unmarshal(respBody, bundleResp)
	exitFatal(err, bundleResponseErrorMsg, instance.URL())

	return bundleResp
}

func (a *httpRequests) bundleUninstall(instance aemInstanceConfig, bundle bundle) *bundleResponse {
	// cURL (POST http://localhost:4505/system/console/bundles/name%20of%20bundle)

	params := url.Values{}
	params.Set(bundleFormActionField, bundleUninstall)
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:4505/system/console/bundles/%s", body)

	a.addAuthentication(instance, req)

	req.Header.Add(headers.ContentType, "application/x-www-form-urlencoded; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	exitFatal(err, bundleResponseErrorMsg, instance.URL())

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

	bundleResp := new(bundleResponse)

	// Parse JSON feed
	err = json.Unmarshal(respBody, bundleResp)
	exitFatal(err, bundleResponseErrorMsg, instance.URL())

	return bundleResp
}

func (a *httpRequests) createFilePart(w *multipart.Writer, path string, fieldname, mimeType string) (*multipart.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set(headers.ContentDisposition, fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filepath.Base(path)))
	h.Set(headers.ContentType, mimeType)
	part, err := w.CreatePart(h)
	if err != nil {
		return w, err
	}

	fileContent, err := ioutil.ReadFile(path)
	part.Write(fileContent)

	return w, err
}

func (a *httpRequests) bundleInstall(instance aemInstanceConfig, bundleFile string, command string, bundleStartLevel int) bool {

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	writer.WriteField(bundleFormActionField, command)
	writer.WriteField("bundlestartlevel", string(bundleStartLevel))
	_, err := a.createFilePart(writer, bundleFile, "bundlefile", JarContentType)
	exitFatal(err, "Could not read package for upload")
	writer.Close()

	// Create client
	client := &http.Client{}

	// Create request
	//req, err := http.NewRequest(http.MethodPost, instance.URL()+URLBundles, body)
	req, _ := http.NewRequest(http.MethodPost, instance.URL()+URLBundles, body)

	a.addAuthentication(instance, req)

	// Headers
	req.Header.Add(headers.ContentType, writer.FormDataContentType())
	bodySize := fmt.Sprintf("%d", len(body.String()))
	req.Header.Add(headers.ContentLength, bodySize)

	// Fetch Request
	resp, err := client.Do(req)
	exitFatal(err, bundleResponseErrorMsg, instance.URL())

	if resp.StatusCode == 200 {
		return true
	}

	return false
}

func (a *httpRequests) addAuthentication(instance aemInstanceConfig, req *http.Request) {
	if len(instance.Username) > 0 || len(a.getPassword(instance)) > 0 {
		req.SetBasicAuth(instance.Username, a.getPassword(instance))
	}
}
