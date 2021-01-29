package project

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/spf13/afero"
)

//CreateBinDir Create bin dir
func CreateBinDir() (string, error) {
	binDir, err := GetBinDir()
	if err != nil {
		return ``, err
	}
	if exists, _ := afero.Exists(fs, binDir); !exists {
		err := fs.MkdirAll(binDir, 0755)
		if err != nil {
			return ``, fmt.Errorf("could not create install dir (%s)", binDir)
		}
	}
	return binDir, nil
}

//CreateInstanceDir Create dir for instance
func CreateInstanceDir(instance objects.Instance) (string, error) {
	instanceDir, err := GetInstanceDirLocation(instance)
	if err != nil {
		return ``, err
	}
	if exists, _ := afero.Exists(fs, instanceDir); !exists {
		err := fs.MkdirAll(instanceDir, 0755)
		if err != nil {
			return ``, fmt.Errorf("could not create install instance (%s)", instanceDir)
		}
	}
	return instanceDir, nil
}

// CreateAemInstallDir creates aem install dir
func CreateAemInstallDir(instance objects.Instance) (string, error) {
	path, err := GetAemInstallDirLocation(instance)
	if err != nil {
		return ``, err
	}
	if exists, _ := afero.Exists(fs, path); !exists {
		err := fs.MkdirAll(path, 0755)
		if err != nil {
			return ``, fmt.Errorf("could not create install dir (%s)", path)
		}
	}
	return path, nil
}

// CreateInstancesDir creates instances dir for all instances to be created under
func CreateInstancesDir() (string, error) {
	path, err := GetInstancesDirLocation()
	if err != nil {
		return ``, err
	}
	if exists, _ := afero.Exists(fs, path); !exists {
		err := fs.MkdirAll(path, 0755)
		if err != nil {
			return ``, fmt.Errorf("could not create instance dir (%s)", path)
		}
	}
	return path, nil
}

// CreateDirForPackage creates a dir for package struct
func CreateDirForPackage(aemPackage *objects.Package) (string, error) {
	path, err := GetDirForPackage(aemPackage)
	if err != nil {
		return ``, err
	}

	if exists, _ := afero.Exists(fs, path); !exists {
		err := fs.MkdirAll(path, 0755)
		if err != nil {
			return ``, fmt.Errorf("could not create pkg dir (%s)", path)
		}
	}
	return path, nil
}

// CreateDir creates a dir for package struct
func CreateDir(path string) (string, error) {
	if exists, _ := afero.Exists(fs, path); !exists {
		err := fs.MkdirAll(path, 0755)
		if err != nil {
			return ``, fmt.Errorf("could not create pkg dir (%s)", path)
		}
	}
	return path, nil
}
