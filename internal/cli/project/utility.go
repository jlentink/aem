package project

import (
	"fmt"
	"github.com/spf13/afero"
	"io"
	"os"
	"regexp"
)

const (
	configFilename                 = "aem.toml"
	configAemJar                   = "AEM-%s.jar"
	configAemRunDir                = "crx-quickstart"
	configAemInstallDir            = "install"
	configAemInstancesDir          = "instances"
	configPackageDir               = "packages"
	configAppDir                   = "app"
	configBinDir                   = "bin"
	configInstanceGitIgnore        = ".gitignore"
	configInstanceGitIgnoreContent = "# Ignore everything in this directory\n*\n# Except this file\n!.gitignore"
	configAemLogDir                = "logs"
	configAemPidFile               = "conf/cq.pid"
	configLicenseFile              = "license.properties"
)

var (
	fs    = afero.NewOsFs()
	fsErr = false
)

// Mock replace fs with memFS for mocking and testing
func Mock(ro bool) afero.Fs {
	fs = afero.NewMemMapFs()
	if ro {
		fs = afero.NewReadOnlyFs(fs)
	}
	return fs
}

// FilesystemError trigger filesystem error for unit testing
func FilesystemError(err bool) {
	fsErr = err
}

func appendSlash(path string) string {
	if len(path) == 0 {
		return ``
	}
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}
	return path
}

// Remove slash from end of path
func RemoveSlash(path string) string {
	if len(path) == 1 && path == `/` || len(path) == 0 {
		return ``
	}

	if path[len(path)-1:] == `/` {
		path = path[:len(path)-1]
	}
	return path
}

func normalizeString(input string) string {
	r, _ := regexp.Compile(`(\s|\W)`)
	return r.ReplaceAllString(input, `-`)
}

// Exists tell if a file exists
func Exists(path string) bool {
	if exists, _ := afero.Exists(fs, path); exists {
		return true
	}
	return false
}

// Rename file
func Rename(source, destination string) error {
	err := fs.Rename(source, destination)
	if err != nil {
		err = fmt.Errorf("could not rename directory from %s to %s", source, destination)
	}
	return err
}

// Copy file
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := fs.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() && !sourceFileStat.IsDir() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := fs.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := fs.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// Open opens a file, returning it or an error, if any happens.
func Open(path string) (afero.File, error) {
	return fs.Open(path)
}

// Create file
func Create(path string) (afero.File, error) {
	return fs.Create(path)
}

// ReadDir Read dir from filesystem
func ReadDir(path string) ([]os.FileInfo, error) {
	return afero.ReadDir(fs, path)
}

// Remove removes a file identified by name, returning an error, if any
// happens.
func Remove(path string) error {
	return fs.Remove(path)
}
