package main

import (
	"fmt"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func NewProjectStructure() projectStructure{
	return projectStructure{
		fs: afero.NewOsFs(),
	}
}

type projectStructure struct {
	fs afero.Fs
}

func (p *projectStructure) mock(){
	p.fs = afero.NewMemMapFs()
}

func (p *projectStructure) getWorkDir() string {
	dir, err := os.Getwd()
	exitFatal(err, "Could not get current working directory.")
	return dir
}



func (p *projectStructure) getConfigFileLocation() string {
	return p.appendSlash(p.getWorkDir()) + CONFIG_FILENAME;
}

func (p *projectStructure) getInstanceDirLocation() string {
	dir := p.getWorkDir()
	return p.appendSlash(dir) + p.appendSlash(CONFIG_INSTANCE_DIR)
}

func (p *projectStructure) appendSlash(path string) string {
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}
	return path
}

func (p *projectStructure) removeSlash(path string) string {
	if len(path) == 1 && path == "/" {
		return ""
	}

	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}
	return path
}

func (p *projectStructure) getIgnoreFileLocation() string {
	return p.appendSlash(p.getInstanceDirLocation()) + CONFIG_INSTANCE_GIT_IGNORE
}

func (p *projectStructure) getUnpackDirLocation() string {
	return p.appendSlash(p.getInstanceDirLocation()) + CONFIG_AEM_RUN_DIR
}

func (p *projectStructure) normalizeString(input string) string {
	r, _ := regexp.Compile("(\\s|\\W)")
	return r.ReplaceAllString(input, "-")
}

func (p *projectStructure) getRunDirLocation(instance AEMInstanceConfig) string {
	return p.appendSlash(p.getInstanceDirLocation()) + p.normalizeString(instance.Name)
}

func (p *projectStructure) getPidFileLocation(instance AEMInstanceConfig) string {
	return p.appendSlash(p.getRunDirLocation(instance)) + CONFIG_AEM_PID
}

func (p *projectStructure) getAppDirLocation(instance AEMInstanceConfig) string {
	return p.appendSlash(p.getRunDirLocation(instance)) + CONFIG_APP_DIR
}

func (p *projectStructure) getAemInstallDirLocation(instance AEMInstanceConfig) string {
	return p.appendSlash(p.getRunDirLocation(instance)) + CONFIG_AEM_INSTALL_DIR
}

func (p *projectStructure) getJarFileLocation() string {
	return p.appendSlash(p.getInstanceDirLocation()) + CONFIG_AEM_JAR_NAME
}

func (p *projectStructure) getLogFileLocation(instance AEMInstanceConfig) string {
	return p.appendSlash(p.getRunDirLocation(instance)) + CONFIG_AEM_LOG
}

func (p *projectStructure) getPackagesDirLocation() string {
	return p.appendSlash(p.getInstanceDirLocation()) + CONFIG_PACKAGES_DIR
}

func (p *projectStructure) getDirForPackage(aemPackage PackageDescription) string {
	location := p.appendSlash(p.getPackagesDirLocation()) + p.appendSlash(aemPackage.Name) + aemPackage.Version
	return location
}

func (p *projectStructure) getLocationForPackage(aemPackage PackageDescription) string {
	location := p.appendSlash(p.getDirForPackage(aemPackage)) + aemPackage.DownloadName
	return location
}

func (p *projectStructure) createAemInstallDir(instance AEMInstanceConfig) string {
	installDir := p.getAemInstallDirLocation(instance)
	if exists, _ := afero.Exists(p.fs, installDir); !exists {
		err := p.fs.MkdirAll(installDir, 0755)
		exitFatal(err, "Could not create install dir (%s)", installDir)
	}
	return installDir
}

func (p *projectStructure) rename(source, destination string) {
	err := p.fs.Rename(source, destination)
	exitFatal(err, "Could not rename directory from %s to %s", source, destination)
}

func (p *projectStructure) createInstanceDir() string {
	instanceDir := p.getInstanceDirLocation()
	fmt.Printf(instanceDir)
	if exists, _ := afero.Exists(p.fs, instanceDir); !exists {
		err := p.fs.MkdirAll(instanceDir, 0755)
		exitFatal(err, "Could not create instance dir (%s)", instanceDir)
	}

	return instanceDir
}

func (p *projectStructure) exists(path string) bool {
	if exists, _ := afero.Exists(p.fs, path); exists {
		return true
	}
	return false
}

func (p *projectStructure) createDirForPackage(aemPackage PackageDescription) string {
	packageDir := p.getDirForPackage(aemPackage)
	if exists, _ := afero.Exists(p.fs, packageDir); !exists {
		err := p.fs.MkdirAll(packageDir, 0755)
		exitFatal(err, "Could not create package dir (%s)", packageDir)
	}

	return packageDir
}

func (p *projectStructure) writeGitIgnoreFile() string {
	gitIgnorePath := p.getIgnoreFileLocation()
	if _, err := os.Stat(gitIgnorePath); os.IsNotExist(err) {
		content := []byte(CONFIG_INSTANCE_GIT_IGNORE_CONTENT)
		err := ioutil.WriteFile(gitIgnorePath, content, 0644)
		exitFatal(err, "Could not create ignore file")
	}
	return gitIgnorePath
}

func (p *projectStructure) copy(src, dst string) (int64, error) {
	sourceFileStat, err := p.fs.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := p.fs.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := p.fs.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
