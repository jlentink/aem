package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/spf13/afero"
	"io/ioutil"
	"regexp"
	"strings"
)

// Labels to retrieve data from the manifest
//nolint
const (
	ManifestLabelManifestVersion           = "Manifest-Version"
	ManifestLabelImplementationTitle       = "Implementation-Title"
	ManifestLabelImplementationVersion     = "Implementation-Version"
	ManifestLabelArchiverVersion           = "Archiver-Version"
	ManifestLabelBuiltBy                   = "Built-By"
	ManifestLabelImplementationVendorID    = "Implementation-Vendor-Id"
	ManifestLabelImportPackage             = "Import-Package"
	ManifestLabelContentPackageType        = "Content-Package-Type"
	ManifestLabelContentPackageDescription = "Content-Package-Description"
	ManifestLabelContentPackageRoots       = "Content-Package-Roots"
	ManifestLabelCreatedBy                 = "Created-By"
	ManifestLabelBuildJdk                  = "Build-Jdk"
	ManifestLabelContentPackageID          = "Content-Package-Id"
	ManifestLabelPackageGroup              = "Package-Group"
	ManifestLabelPackageName               = "Package-Name"
	ManifestLabelPackageVersion            = "Package-Version"

	manifestKeySeperator = ":"
	manifestReturn       = "\r"
	manifestPath         = "META-INF/MANIFEST.MF"
)

func newManifestPackage() manifestPackage {
	return manifestPackage{
		fs:        afero.NewOsFs(),
		u:         new(utility),
		keyValues: make(map[string]string),
	}
}

type manifestPackage struct {
	fs         afero.Fs
	u          *utility
	currentKey string
	keyValues  map[string]string
}

func (m *manifestPackage) Get(manifestLabel string) (string, error) {
	if _, ok := m.keyValues[manifestLabel]; ok {
		return m.keyValues[manifestLabel], nil
	}

	return "", errors.New("could not find the key")
}

func (m *manifestPackage) labelValue(line string) (string, string) {
	regLabel, _ := regexp.Compile(`^([A-Z])([A-Za-z\-]*) ?: ?(.*)$`)
	regValue, _ := regexp.Compile(`^ (.*)$`)

	if regLabel.MatchString(line) {
		line = strings.TrimSuffix(line, manifestReturn)
		matches := regLabel.FindAllStringSubmatch(line, -1)
		return fmt.Sprintf("%s%s", matches[0][1], matches[0][2]), fmt.Sprintf("%s", matches[0][3])
	}

	if regValue.MatchString(line) {
		line = strings.TrimSuffix(line, manifestReturn)

		return "", line[1:]

	}

	return "", ""
}

func (m *manifestPackage) store(label, value string) {
	if len(label) > 0 {
		m.currentKey = label
		m.keyValues[label] = value
	} else {
		m.keyValues[m.currentKey] = m.keyValues[m.currentKey] + value
	}
}

func (m *manifestPackage) explodeContentPackageID() {
	splits := strings.Split(m.keyValues[ManifestLabelContentPackageID], manifestKeySeperator)
	m.keyValues[ManifestLabelPackageGroup] = splits[0]
	m.keyValues[ManifestLabelPackageName] = splits[1]
	m.keyValues[ManifestLabelPackageVersion] = splits[2]

}

func (m *manifestPackage) parse(manifest string) map[string]string {

	lines := strings.Split(manifest, "\n")
	for _, line := range lines {
		tLabel, tValue := m.labelValue(line)
		m.store(tLabel, tValue)
	}

	if len(m.keyValues[ManifestLabelContentPackageID]) > 0 {
		m.explodeContentPackageID()
	}
	return m.keyValues
}

func (m *manifestPackage) readZip(path string) (string, error) {
	reader, err := zip.OpenReader(path)
	exitFatal(err, "Error while reading the package: ")

	for _, file := range reader.File {
		if file.Name == manifestPath {
			zippedManifest, _ := file.Open()
			defer zippedManifest.Close()

			mData, _ := ioutil.ReadAll(zippedManifest)
			return string(mData), nil
		}
	}

	return "", errors.New("could not find manifest in zip")
}

func (m *manifestPackage) fromZip(path string) map[string]string {
	if m.u.Exists(path) {
		manifest, err := m.readZip(path)
		exitFatal(err, "Could not find manifest.")

		return m.parse(manifest)
	}
	exitProgram("Could not find package at %s", path)
	return m.keyValues
}
