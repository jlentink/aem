package project

import (
	"errors"
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/spf13/afero"
	"os"
)

// GetWorkDir get workdir for project
func GetWorkDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil || fsErr {
		return ``, errors.New("could not get current working directory")
	}
	return dir, nil
}

// GetBinDir get global project bin dir
func GetBinDir() (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + appendSlash(configBinDir), nil
}

// GetConfigFileLocation get project config file location
func GetConfigFileLocation() (string, error) {
	cwd, err := GetWorkDir()
	if err != nil {
		return ``, err
	}
	return appendSlash(cwd) + configFilename, nil
}

// GetInstancesDirLocation get location of instances dir
func GetInstancesDirLocation() (string, error) {
	cwd, err := GetWorkDir()
	if err != nil {
		return ``, err
	}
	return appendSlash(cwd) + appendSlash(configAemInstancesDir), nil
}

// GetLicenseLocation get license location for instance
func GetLicenseLocation(instance objects.Instance) (string, error) {
	path, err := GetInstanceDirLocation(instance)
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + configLicenseFile, nil
}

func getIgnoreFileLocation() (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + configInstanceGitIgnore, nil
}

// GetUnpackDirLocation get unpack dir location
func GetUnpackDirLocation() (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + configAemRunDir, nil
}

// GetInstanceDirLocation Get's location for instance
func GetInstanceDirLocation(instance objects.Instance) (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	return appendSlash(path + normalizeString(instance.Name)), nil
}

// GetRunDirLocation Get location of the run dir of aem
func GetRunDirLocation(instance objects.Instance) (string, error) {
	path, err := GetInstanceDirLocation(instance)
	if err != nil {
		return ``, err
	}
	return appendSlash(path + configAemRunDir), nil
}

// GetPidFileLocation gets path of pidfile
func GetPidFileLocation(instance objects.Instance) (string, error) {
	path, err := GetRunDirLocation(instance)
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + configAemPidFile, nil
}

// GetAppDirLocation gets app dir for instance
func GetAppDirLocation(instance objects.Instance) (string, error) {
	path, err := GetRunDirLocation(instance)
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + configAppDir, nil
}

// GetAemInstallDirLocation gets AEM install dir for location
func GetAemInstallDirLocation(instance objects.Instance) (string, error) {
	path, err := GetRunDirLocation(instance)
	if err != nil {
		return ``, err
	}
	return appendSlash(appendSlash(path) + configAemInstallDir), nil
}

// CreateInstallDir create the install dir
func CreateInstallDir(instance objects.Instance) (string, error) {
	path, err := GetAemInstallDirLocation(instance)
	if err != nil {
		return ``, err
	}

	ex, err := afero.DirExists(fs, path)
	if err != nil {
		return path, err
	}

	if ex {
		return path, nil
	}

	return path, fs.MkdirAll(path, 0777)
}

// GetJarFileLocation Get the location of the jar file
func GetJarFileLocation(jar *objects.AemJar) (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + fmt.Sprintf(configAemJar, jar.Version), nil
}

// GetLogDirLocation Get the log file location
func GetLogDirLocation(instance objects.Instance) (string, error) {
	path, err := GetRunDirLocation(instance)
	if err != nil {
		return ``, err
	}
	return appendSlash(appendSlash(path) + configAemLogDir), nil
}

func getPackagesDirLocation() (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	return appendSlash(path) + configPackageDir, nil
}

// GetDirForPackage get the path for the packages dir
func GetDirForPackage(aemPackage *objects.Package) (string, error) {
	path, err := getPackagesDirLocation()
	if err != nil {
		return ``, err
	}
	location := appendSlash(path) + appendSlash(aemPackage.Name) + aemPackage.Version
	return location, nil
}

// GetLocationForPackage location for package
func GetLocationForPackage(aemPackage *objects.Package) (string, error) {
	path, err := GetDirForPackage(aemPackage)
	if err != nil {
		return ``, err
	}
	location := appendSlash(path) + aemPackage.DownloadName
	return location, nil
}
