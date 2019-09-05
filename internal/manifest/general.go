package manifest

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Manifest labels
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

	manifestPath         = "META-INF/MANIFEST.MF"
	manifestReturn       = "\r"
)

var (
	currentKey string
	keyValues  map[string]string
)

// OpenPackage Open package to read manifest
func OpenPackage(location string) (*Manifest, error) {
	reader, err := zip.OpenReader(location)
	if err != nil {
		return nil, err
	}

	for _, file := range reader.File {
		if file.Name == manifestPath {
			manifest, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer manifest.Close()
			return openManifest(manifest)
		}
	}
	return nil, fmt.Errorf("no manifest found in zip")
}

// OpenManifest Open manifest for reading
func OpenManifest(location string) (*Manifest, error) {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return nil, fmt.Errorf("could not find file at: %s", err)
	}

	manifest, err := os.Open(location)
	if err != nil {
		return nil, err
	}

	return openManifest(manifest)
}

func openManifest(r io.Reader) (*Manifest, error) {
	currentKey = ""
	//@todo: clear map
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(d), "\n")
	for _, line := range lines {
		tLabel, tValue := labelValue(line)
		store(tLabel, tValue)
	}
	return &Manifest{keyValues: keyValues}, nil
}

func labelValue(line string) (string, string) {
	regLabel, _ := regexp.Compile(`^([A-Z])([A-Za-z\-]*) ?: ?(.*)$`)
	regValue, _ := regexp.Compile(`^ (.*)$`)

	if regLabel.MatchString(line) {
		line = strings.TrimSuffix(line, manifestReturn)
		matches := regLabel.FindAllStringSubmatch(line, -1)
		//nolint
		return fmt.Sprintf("%s%s", matches[0][1], matches[0][2]), fmt.Sprintf("%s", matches[0][3])
	}

	if regValue.MatchString(line) {
		line = strings.TrimSuffix(line, manifestReturn)

		return "", line[1:]

	}

	return "", ""
}

func store(label, value string) {
	if len(label) > 0 {
		currentKey = label
		keyValues[label] = value
	} else {
		keyValues[currentKey] = keyValues[currentKey] + value
	}
}
