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
	"github.com/jlentink/aem/internal/sliceutil"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	packageListURL      = "/crx/packmgr/list.jsp"
	packageUploadURL    = "/crx/packmgr/service.jsp"
	packageDeletedURL   = "/crx/packmgr/service/script.html/etc/packages/%s/%s"
	packageRebuildURL   = "/crx/packmgr/service/.json%s?cmd=build"
	packageInstallURL   = "/crx/packmgr/service/.json%s?cmd=install"
	packageCreateURL    = "/crx/packmgr/service/exec.json?cmd=create"
	packageUpdateURL    = "/crx/packmgr/update.jsp"
	packagePathTemplate = "{\"root\":\"%s\",\"rules\":[]},"
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

// GetPackageByNameAndVersion finds a package based on name and version
func GetPackageByNameAndVersion(i objects.Instance, name, version string) (*objects.Package, error) {
	pkgList, err := list(i)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgList {
		if strings.ToLower(pkg.Name) == strings.ToLower(name) &&
			pkg.Version == version {
			return &pkg, nil
		}
	}
	return nil, fmt.Errorf("could not find package")
}

// GetPackageByNameAndGroupAndVersion finds a package based on name,group and version
func GetPackageByNameAndGroupAndVersion(i objects.Instance, name, version, group string) (*objects.Package, error) {
	pkgList, err := list(i)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgList {
		if  strings.EqualFold(pkg.Name, name) &&
			strings.EqualFold(pkg.Version, version) &&
			strings.EqualFold(pkg.Group, group) {
			return &pkg, nil
		}
	}
	return nil, fmt.Errorf("could not find package")
}

// GetPackageByName finds a package based on name and version
func GetPackageByName(i objects.Instance, name string) (*objects.Package, error) {
	pkgList, err := list(i)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgList {
		if strings.ToLower(pkg.Name) == strings.ToLower(name) {
			return &pkg, nil
		}
	}
	emptyPkg := objects.Package{}
	return &emptyPkg, fmt.Errorf("could not find package")
}

// PackageList return list of packages on instance
func PackageList(i objects.Instance) ([]objects.Package, error) {
	return list(i)
}

// FilteredByGroupPackageList Filter the packages based on group
func FilteredByGroupPackageList(i objects.Instance, group string) ([]objects.Package, error) {
	filtered := make([]objects.Package, 0)
	pkgs, err := list(i)
	if err != nil || group == "" {
		return pkgs, err
	}
	groups := strings.Split(group, ",")
	for _, pkg := range pkgs {
		if sliceutil.StringInSliceEqualFold(pkg.Group, groups) {
			filtered = append(filtered, pkg)
		}
	}
	return filtered, err
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

	part.Write(fileContent)                                   // nolint: errcheck
	writer.WriteField("name", pkgName)                        // nolint: errcheck
	writer.WriteField("force", strconv.FormatBool(force))     // nolint: errcheck
	writer.WriteField("install", strconv.FormatBool(install)) // nolint: errcheck
	writer.Close()                                            // nolint: errcheck

	return body, writer.FormDataContentType(), nil
}

func initCreate(i objects.Instance, name, group, version string) error {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	pw, err := i.GetPassword()
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Add("_charset_", "utf-8")
	data.Add("packageName", name)
	data.Add("packageVersion", version)
	data.Add("groupName", group)
	_, _, err = http.PostFormEncode(i.URLString()+packageCreateURL, i.Username, pw, data)

	if err != nil {
		return err
	}
	return nil
}

func updateCreate(i objects.Instance, pkg *objects.Package, name, group, version string, paths []string) error {
	pw, err := i.GetPassword()
	if err != nil {
		return err
	}

	filter := "["
	for _, path := range paths {
		filter += fmt.Sprintf(packagePathTemplate, path)
	}
	filter = filter[0:len(filter)-1] + "]"

	data := map[string]string{}
	data["path"] = fmt.Sprintf(pkg.Path)
	data["packageName"] = name
	data["groupName"] = group
	data["description"] = "AEM-CLI create content backup package."
	data["filter"] = filter
	data["version"] = version

	_, _, err = http.PostMultiPart(i.URLString()+packageUpdateURL, i.Username, pw, data)
	if err != nil {
		return err
	}

	return nil
}

// Create a package
func Create(i objects.Instance, name, group, version string, paths []string, build bool) (*objects.Package, error) {
	err := initCreate(i, name, group, version)
	if err != nil {
		return nil, err
	}

	pkg, err := GetPackageByNameAndVersion(i, name, version)
	if err != nil {
		return nil, err
	}

	err = updateCreate(i, pkg, name, group, version, paths)
	if err != nil {
		return nil, err
	}

	if build {
		Rebuild(&i, pkg)
	}

	return pkg, nil
}

// AwaitBuild to be done
func AwaitBuild(i *objects.Instance, pkg *objects.Package) error {
	pkgs, err := list(*i)
	if err != nil {
		return err
	}
	for _, cPkg := range pkgs {
		if cPkg.Equals(pkg) {
			if cPkg.BuildCount < 1 {
				fmt.Print(".")
				time.Sleep(1 * time.Second)
				AwaitBuild(i, pkg)
			}
		}
	}

	return nil
}

// Upload package to instance
func Upload(i objects.Instance, pkgLocation string, install, force bool) (*objects.CrxResponse, string, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	if _, err := os.Stat(pkgLocation); os.IsNotExist(err) {
		return nil, "", fmt.Errorf("could not find package: %s", pkgLocation)
	}
	if m, _ := regexp.MatchString(`(?i)(.*).zip$`, pkgLocation); !m {
		return nil, "", fmt.Errorf("package should be a zipfile")
	}

	body, formHeader, err := constructInstallBody(pkgLocation, install, force)
	if err != nil {
		return nil, "", err
	}
	pw, err := i.GetPassword()
	if err != nil {
		return nil, "", fmt.Errorf("could not get password")
	}
	resp, err := http.Upload(i.URLString()+packageUploadURL, i.Username, pw, body, []http.Header{{Key: headers.ContentType, Value: formHeader}})
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode >= 400 {
		return nil, "", fmt.Errorf("received invalid http status %d - %s", resp.StatusCode, resp.Status)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Print("\n")
	crxResp := objects.CrxResponse{}
	err = xml.Unmarshal(respBody, &crxResp)
	if err != nil {
		return nil, string(respBody), err
	}
	return &crxResp, "", nil
}

// Download package from instance
func Download(i *objects.Instance, pkg *objects.Package) (*objects.Package, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}

	_, err := project.CreateDirForPackage(pkg)
	if err != nil {
		return nil, err
	}
	p, err := project.GetLocationForPackage(pkg)
	if err != nil {
		return nil, err
	}

	_, err = http.DownloadFile(p, i.URLString()+pkg.Path, i.Username, i.GetPasswordSimple(), true)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}

// Delete package from instance
func Delete(i *objects.Instance, pkg *objects.Package) (*objects.Package, error) {
	if !aem.Cnf.ValidateSSL {
		http.DisableSSLValidation()
	}


	data := map[string]string{"cmd" : "delete",
								"pid": "next-gen2913",
								"callback": "window.parent.Ext.Ajax.Stream.callback"}

	pw, err := i.GetPassword()
	if err != nil {
		return nil, err
	}
	_, _, err = http.PostMultiPart(i.URLString()+fmt.Sprintf(packageDeletedURL, pkg.Group, pkg.DownloadName), i.Username, pw, data)
	return pkg, err
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
		if strings.EqualFold(cPkg.Name, pkgName) && (pkgVersion == "" || strings.EqualFold(pkgVersion, cPkg.Version)) {
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
		if strings.EqualFold(cPkg.Name, pkgName) && (pkgVersion == "" || strings.EqualFold(pkgVersion, cPkg.Version)) {
			return Rebuild(i, &cPkg)
		}
	}
	return nil, fmt.Errorf("could not find package: %s", n)

}

// GetTimeVersion get version based on time
func GetTimeVersion() string {
	now := time.Now()
	return fmt.Sprintf("%s.%d", now.Format("20060102"), now.UnixNano())
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
	_, _, err = http.PostPlain(i.URLString()+fmt.Sprintf(packageRebuildURL, pkg.Path), i.Username, pw, nil)

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

	_, _, err = http.PostPlain(i.URLString()+fmt.Sprintf(packageInstallURL, pkg.Path), i.Username, pw, nil)

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
		if strings.EqualFold(cPkg.Name, pkgName) && (pkgVersion == "" || strings.EqualFold(pkgVersion, cPkg.Version)) {
			return Install(i, &cPkg)
		}
	}
	return nil, fmt.Errorf("could not find package: %s", n)

}
