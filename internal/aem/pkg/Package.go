package pkg

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/go-http-utils/headers"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/http"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/packageproperties"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	packageListURL    = "/crx/packmgr/list.jsp"
	packageUploadURL  = "/crx/packmgr/service.jsp"
	packageRebuildURL = "/crx/packmgr/service/.json%s?cmd=build"
	packageInstallURL = "/crx/packmgr/service/.json%s?cmd=install"
)

func list(i objects.Instance) ([]objects.Package, error) {

	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	body, err := aem.GetFromInstance(&i, packageListURL)
	if err != nil {
		return make([]objects.Package, 0), err
	}

	packageResultList := &objects.PackageList{}
	err = json.Unmarshal(body, packageResultList)
	if err != nil {
		return make([]objects.Package, 0), err
	}
	for i, cPkg := range packageResultList.Results {
		if cPkg.Created > 0 {
			tt := output.UnixTime(cPkg.Created)
			//nolint
			packageResultList.Results[i].CreatedStr = fmt.Sprintf("%s", tt.UTC())
		}
		if cPkg.LastModified > 0 {
			tt := output.UnixTime(cPkg.LastModified)
			//nolint
			packageResultList.Results[i].LastModifiedByStr = fmt.Sprintf("%s", tt.UTC())
		}

		packageResultList.Results[i].SizeHuman = humanize.Bytes(uint64(cPkg.Size))
	}

	return packageResultList.Results, nil
}

func nameVersion(p string) (string, string) {

	if strings.Contains(p, `:`) {
		r := strings.Split(p, `:`)
		return r[0], r[1]
	}
	return p, ``
}

// PackageList return list of packages on instance
func PackageList(i objects.Instance) ([]objects.Package, error) {
	return list(i)
}

func constructInstallBody(pkgLocation string, install, force bool) (*bytes.Buffer, string, error) {
	pkgDesc, err := packageproperties.Open(pkgLocation)
	if err != nil {
		return nil, "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(pkgLocation))
	fileContent, err := ioutil.ReadFile(pkgLocation)
	if err != nil {
		return nil, "", err
	}

	pkgName, err := pkgDesc.Get(packageproperties.Name)
	if err != nil {
		return nil, "", err
	}

	part.Write(fileContent)
	writer.WriteField("name", pkgName)
	writer.WriteField("force", strconv.FormatBool(force))
	writer.WriteField("install", strconv.FormatBool(install))
	writer.Close()

	return body, writer.FormDataContentType(), nil
}

// Upload package to instance
func Upload(i objects.Instance, pkgLocation string, install, force bool) (*objects.CrxResponse, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	if _, err := os.Stat(pkgLocation); os.IsNotExist(err) {
		return nil, fmt.Errorf("could not find package: %s", pkgLocation)
	}
	if m, _ := regexp.MatchString(`(?i)(.*).zip$`, pkgLocation); !m {
		return nil, fmt.Errorf("package should be a zipfile")
	}

	body, formHeader, err := constructInstallBody(pkgLocation, install, force)
	if err != nil {
		return nil, err
	}
	pw, err := i.GetPassword()
	if err != nil {
		return nil, fmt.Errorf("could not get password")
	}
	resp, err := http.Upload(i.URLString()+packageUploadURL, i.Username, pw, body, []http.Header{{Key: headers.ContentType, Value: formHeader}})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("received invalid http status %d - %s", resp.StatusCode, resp.Status)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Print("\n")
	crxResp := objects.CrxResponse{}
	err = xml.Unmarshal(respBody, &crxResp)
	if err != nil {
		return nil, err
	}
	return &crxResp, nil
}

// Download package from instance
func Download(i *objects.Instance, pkg *objects.Package) (*objects.Package, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	project.CreateDirForPackage(pkg)
	p, err := project.GetLocationForPackage(pkg)
	if err != nil {
		return nil, err
	}

	http.DownloadFile(p, i.URLString()+pkg.Path, i.Username, i.GetPasswordSimple(), true)

	return pkg, nil
}

//DownloadWithName Download package based on name
func DownloadWithName(i *objects.Instance, n string) (*objects.Package, error) {
	pkgName, pkgVersion := nameVersion(n)
	pkgs, err := list(*i)
	if err != nil {
		return nil, err
	}
	for _, cPkg := range pkgs {
		cPkg := cPkg
		if strings.ToLower(cPkg.Name) == strings.ToLower(pkgName) && (pkgVersion == "" || pkgVersion == cPkg.Version) {
			return Download(i, &cPkg)
		}
	}
	return nil, fmt.Errorf("could not find package: %s", n)
}

// RebuildbyName rebuild package by name on instance
func RebuildbyName(i *objects.Instance, n string) (*objects.Package, error) {
	pkgName, pkgVersion := nameVersion(n)
	pkgs, err := list(*i)
	if err != nil {
		return nil, err
	}
	for _, cPkg := range pkgs {
		cPkg := cPkg
		if strings.ToLower(cPkg.Name) == strings.ToLower(pkgName) && (pkgVersion == "" || pkgVersion == cPkg.Version) {
			return Rebuild(i, &cPkg)
		}
	}
	return nil, fmt.Errorf("could not find package: %s", n)

}

// Rebuild package on instance
func Rebuild(i *objects.Instance, pkg *objects.Package) (*objects.Package, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	pw, err := i.GetPassword()
	if err != nil {
		return nil, err
	}
	_, err = http.PostPlain(i.URLString()+fmt.Sprintf(packageRebuildURL, pkg.Path), i.Username, pw, nil)

	return pkg, err
}

// Install package on instance
func Install(i *objects.Instance, pkg *objects.Package) (*objects.Package, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	pw, err := i.GetPassword()
	if err != nil {
		return nil, err
	}

	_, err = http.PostPlain(i.URLString()+fmt.Sprintf(packageInstallURL, pkg.Path), i.Username, pw, nil)

	return pkg, err
}

// InstallByName install package on instance by name
func InstallByName(i *objects.Instance, n string) (*objects.Package, error) {
	pkgName, pkgVersion := nameVersion(n)
	pkgs, err := list(*i)
	if err != nil {
		return nil, err
	}
	for _, cPkg := range pkgs {
		cPkg := cPkg
		if strings.ToLower(cPkg.Name) == strings.ToLower(pkgName) && (pkgVersion == "" || pkgVersion == cPkg.Version) {
			return Install(i, &cPkg)
		}
	}
	return nil, fmt.Errorf("could not find package: %s", n)

}
