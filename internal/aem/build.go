package aem

import (
	"bytes"
	"fmt"
	"github.com/jlentink/aem/internal/cli/project"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
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

func replaceVersionPlaceholders(suffix string) (string, error) {
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
	if m, _ := regexp.MatchString(`(.*)DATE(.*)`, suffix); m {
		now := time.Now()
		str := fmt.Sprintf("%s.%d", now.Format("20060102"), now.UnixNano())
		suffix = strings.ReplaceAll(suffix, `DATE`, str)
	}
	if m, _ := regexp.MatchString(`(.*)ADOBE(.*)`, suffix); m {
		now := time.Now()
		str := fmt.Sprintf("%s.%d.%d", now.Format("2006.12"), now.UnixNano(), now.UnixNano())
		suffix = strings.ReplaceAll(suffix, `ADOBE`, str)
	}
	return suffix, nil
}

// SetBuildVersion sets
func SetBuildVersion(productionBuild bool) error {
	versionSuffix := ""
	if productionBuild {
		var err error
		versionSuffix, err = replaceVersionPlaceholders(Cnf.VersionSuffix)
		if err != nil {
			return err
		}
	} else {
		versionSuffix = "-SNAPSHOT"
	}

	version := Cnf.Version

	if m, _ := regexp.MatchString(`(.*)ADOBE(.*)`, Cnf.VersionSuffix); m {
		version = ""
	}

	workPath, _ := project.GetWorkDir()
	cmd := exec.Command("mvn", "versions:set", "-DnewVersion="+version+versionSuffix, "-DallowSnapshots=true")
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
func BuildProject(productionBuild, skipTests, skipCheckStyle, skipFrontend bool) error {
	workPath, _ := project.GetWorkDir()

	err := SetBuildVersion(productionBuild)
	if err != nil {
		return err
	}

	buildOptions := strings.Split(Cnf.BuildCommands, " ")

	if skipTests {
		buildOptions = append(buildOptions, "-DskipTests")
	}

	if skipCheckStyle {
		buildOptions = append(buildOptions, "-Dcheckstyle.skip")
	}

	if skipFrontend {
		buildOptions = append(buildOptions, "-DskipFrontend=true")
	}

	cmd := exec.Command("mvn", buildOptions...)
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
