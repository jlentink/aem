package bundle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-http-utils/headers"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/http"
	"github.com/jlentink/aem/internal/output"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	//"mime/multipart"
	//"mime/quotedprintable"
)

// List bundles on instance
func List(i *objects.Instance) ([]*Bundle, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	data := url.Values{}
	data.Set("name", "foo")
	data.Add(bundleFormActionField, bundleRefresh)

	pw, err := i.GetPassword()
	if err != nil {
		return nil, fmt.Errorf("could not get password")
	}

	buf := bytes.NewBufferString(data.Encode())
	resp, err := http.PostPlainWithHeaders(i.URLString()+bundlesURL, i.Username, pw, buf, []http.Header{
		{Key: "Content-Type", Value: "application/x-www-form-urlencoded; charset=UTF-8"},
	})
	if err != nil {
		return nil, err
	}

	list := bundlesFeed{}
	err = json.Unmarshal(resp, &list)
	if err != nil {
		return nil, err
	}

	bundleList := make([]*Bundle, 0)
	for i := range list.Data {
		bundleList = append(bundleList, &list.Data[i])
	}

	return bundleList, nil
}

// Stop bundle on instance
func Stop(i *objects.Instance, bundle *Bundle) (*Bundle, error) {
	return bundleAction(i, bundle, "stop")
}

// Get bundle from instance
func Get(i *objects.Instance, name string) (*Bundle, error) {
	bundles, err := List(i)
	if err != nil {
		return nil, err
	}
	for _, bundle := range bundles {
		if bundle.SymbolicName == name {
			return bundle, nil
		}
	}
	return nil, fmt.Errorf("could not find bundle with name: %s", name)
}

// Start a bundle
func Start(i *objects.Instance, bundle *Bundle) (*Bundle, error) {
	return bundleAction(i, bundle, "start")
}

// Build instance on server
func Build(i *objects.Instance) {

}

func bundleAction(i *objects.Instance, bundle *Bundle, action string) (*Bundle, error) {

	data := url.Values{}
	data.Set("action", action)
	data.Add(bundleFormActionField, bundleRefresh)

	pw, err := i.GetPassword()
	if err != nil {
		return nil, fmt.Errorf("could not get password")
	}

	buf := bytes.NewBufferString(data.Encode())
	resp, err := http.PostPlainWithHeaders(i.URLString()+fmt.Sprintf(bundlePageURL, fmt.Sprintf("%d", bundle.ID)), i.Username, pw, buf, []http.Header{
		{Key: headers.ContentType, Value: "application/x-www-form-urlencoded; charset=UTF-8"},
	})

	if err != nil {
		return nil, err
	}

	bResponse := bundleResponse{}
	json.Unmarshal(resp, &bResponse)

	if err != nil {
		return nil, err
	}

	bundle.StateRaw = bResponse.StateRaw
	bundle.Fragment = bResponse.Fragment

	return bundle, nil
}

// Install instance on server
func Install(i *objects.Instance, bundlePath string, level string) error {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}
	output.Printf(output.NORMAL, "\U0001F69A %s - %s\n", i.Name, path.Base(bundlePath))
	if _, err := os.Stat(bundlePath); os.IsNotExist(err) {
		return fmt.Errorf("could not find bundle: %s", bundlePath)
	}
	if m, _ := regexp.MatchString(`(?i)(.*).jar$`, bundlePath); !m {
		return fmt.Errorf("bundle should be a  jar file")
	}

	body, formHeader, err := constructInstallBody(bundlePath, level)
	if err != nil {
		return err
	}
	pw, err := i.GetPassword()
	if err != nil {
		return fmt.Errorf("could not get password")
	}
	resp, err := http.Upload(i.URLString()+bundlesURL, i.Username, pw, body, []http.Header{
		{Key: headers.ContentLength, Value: fmt.Sprintf("%d", body.Len())},
		{Key: headers.ContentType, Value: formHeader},
	})
	fmt.Print("\n")
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("received invalid http status %d - %s", resp.StatusCode, resp.Status)
	}

	return nil
}

// Uninstall bundle on instance
func Uninstall(i *objects.Instance) {

}

// Delete bundle on instance
func Delete(i *objects.Instance) {

}

func constructInstallBody(bundlePath, level string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("action", "install")
	writer.WriteField("bundlestartlevel", level)

	part, _ := writer.CreateFormFile("bundlefile", filepath.Base(bundlePath))
	fileContent, err := ioutil.ReadFile(bundlePath)
	if err != nil {
		return nil, "", err
	}

	part.Write(fileContent)
	writer.Close()

	return body, writer.FormDataContentType(), nil
}

// BundleSearch searches through bundles on instance
func BundleSearch(i *objects.Instance, label string) (*Bundle, error) {
	bundles, err := List(i)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve bundle list from server %s", err.Error())
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ .Name | cyan }} {{ .SymbolicName | cyan }} ({{ .Version | red }})",
		Inactive: "  {{ .Name | cyan }} {{ .SymbolicName | cyan }} ({{ .Version | red }})",
		Selected: "\U0001F48A  " + label + "... {{ .Name | red | cyan }} {{ .SymbolicName | cyan }}  ({{ .Version | red }})",
		Details: `--------- Package ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Version:" | faint }}	{{ .Version }}
{{ "Symbolic name:" | faint }}	{{ .SymbolicName }}
{{ "State:" | faint }}	{{ .State }}
{{ "Category:" | faint }}	{{ .Category }}
`,
	}

	searcher := func(input string, index int) bool {
		bndl := bundles[index]
		name := strings.Replace(strings.ToLower(bndl.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select Bundle:",
		Items:     bundles,
		Templates: templates,
		Size:      20,
		Searcher:  searcher,
	}

	in, _, err := prompt.Run()

	if err != nil {
		return nil, fmt.Errorf("prompt failed %v\n", err)
	}

	return bundles[in], nil
}
