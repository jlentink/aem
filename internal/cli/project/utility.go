package project

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/afero"
	"io"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
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

func pathSeparator(path string) string {
	if strings.ToLower(runtime.GOOS) != "windows" {
		return path
	}

	return strings.ReplaceAll(path, "/", "\\")
}

// RemoveSlash Removes slash from end of path
func RemoveSlash(path string) string {
	if len(path) == 1 && path == `/` || len(path) == 0 {
		return ``
	}

	if path[len(path)-1:] == `/` {
		path = path[:len(path)-1]
	}
	return path
}

// nolint
func filesystem() afero.Fs {
	return fs
}

func normalizeString(input string) string {
	r, _ := regexp.Compile(`(\s|\W)`)
	return r.ReplaceAllString(input, `-`)
}

// HomeDir returns user home's dir
func HomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
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
	source = pathSeparator(source)
	destination = pathSeparator(destination)
	err := fs.Rename(source, destination)
	if err != nil {
		err = fmt.Errorf("could not rename directory from %s to %s (%s)", source, destination, err)
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
	defer source.Close() // nolint: errcheck

	destination, err := fs.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close() // nolint: errcheck
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

// RemoveAll removes a directory path and any children it contains. It
// does not fail if the path does not exist (return nil).
func RemoveAll(path string) error {
	return fs.RemoveAll(path)
}

// IsFile tells you if it is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Confirm ask user for confirmation.
func Confirm(message, help string, force bool) bool {
	confirmed := false
	if force {
		return true
	}
	prompt := &survey.Confirm{
		Message: message,
		Help: help,
	}
	err := survey.AskOne(prompt, &confirmed)
	if err != nil {
		if err.Error() == "interrupt" {
			output.Printf(output.NORMAL, "Program interrupted. (CTRL+C)")
			return false
		}
		output.Printf(output.NORMAL, "Unexpected error: %s", err.Error())
		return false
	}

	return confirmed
}