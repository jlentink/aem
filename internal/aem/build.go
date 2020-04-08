package aem

import (
	"bytes"
	"github.com/jlentink/aem/internal/cli/project"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getGitHash(short bool) (string, error) {
	param := []string{"rev-parse", "HEAD"}
	if short {
		param = []string{"rev-parse", "--short", "HEAD"}
	}
	var outbuf bytes.Buffer
	workPath, _ := project.GetWorkDir()
	cmd := exec.Command("git", param...)
	cmd.Dir = workPath
	cmd.Stdout = &outbuf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return strings.TrimRight(outbuf.String(), "\n"), err
}

func replaceGitHash(suffix string) (string, error) {
	if m, _ := regexp.MatchString(`(.*)GIT_SHORT(.*)`, suffix); m {
		str, err := getGitHash(true)
		if err != nil {
			return ``, err
		}
		suffix = strings.ReplaceAll(suffix, `GIT_SHORT`, str)
	}
	if m, _ := regexp.MatchString(`(.*)GIT_LONG(.*)`, suffix); m {
		str, err := getGitHash(false)
		suffix = strings.ReplaceAll(suffix, `GIT_LONG`, str)
		if err != nil {
			return ``, err
		}
	}
	return suffix, nil
}

// SetBuildVersion sets
func SetBuildVersion(productionBuild bool) error {
	versionSuffix := ""
	if productionBuild {
		var err error
		versionSuffix, err = replaceGitHash(Cnf.VersionSuffix)
		if err != nil {
			return err
		}
	} else {
		versionSuffix = "-SNAPSHOT"
	}

	workPath, _ := project.GetWorkDir()
	cmd := exec.Command("mvn", "versions:set", "-DnewVersion="+Cnf.Version+versionSuffix, "-DallowSnapshots=true")
	cmd.Dir = workPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// BuildProject kicks of the project build
func BuildProject(productionBuild bool) error {
	workPath, _ := project.GetWorkDir()

	err := SetBuildVersion(productionBuild)
	if err != nil {
		return err
	}

	cmd := exec.Command("mvn", strings.Split(Cnf.BuildCommands, " ")...)
	cmd.Dir = workPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// BuildModuleProject build project module
func BuildModuleProject(p string, productionBuild bool) error {
	err := SetBuildVersion(productionBuild)
	if err != nil {
		return err
	}
	cmd := exec.Command("mvn", strings.Split(Cnf.BuildCommands, " ")...)
	cmd.Dir = p
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
