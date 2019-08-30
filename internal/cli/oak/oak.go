package oak

import (
	"fmt"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/http"
	"os"
	"os/exec"
)

const (
	oakRunRepository = "https://repo1.maven.org/maven2/org/apache/jackrabbit/oak-run/%s/oak-run-%s.jar"
	// RepoPath Path to repo
	RepoPath = "/repository/segmentstore"
	oakJar   = "oak-run-%s.jar"
)

var (
	oakDefaultVersion = "1.8.12"
)

// SetDefaultVersion sets the default oak value to be used
func SetDefaultVersion(version string) {
	if len(version) > 0 {
		oakDefaultVersion = version
	}
}

func versionForAem(version string) string {
	switch version {
	case "6.0":
		return "1.0.42"
	case "6.1":
		return "1.2.31"
	case "6.2":
		return "1.4.24"
	case "6.3":
		return "1.6.16"
	case "6.4":
		return "1.8.12"
	case "6.5":
		return "1.8.2"
	default:
		return oakDefaultVersion
	}
}

// Get download the correct oak version
func Get(aemVersion, version string) (string, error) {

	if len(aemVersion) > 0 {
		version = versionForAem(aemVersion)
	}

	if len(version) == 0 {
		return ``, fmt.Errorf("no oak version selected. Use --aem or --oak to define version")
	}

	binPath, err := project.CreateBinDir()
	if err != nil {
		return ``, err
	}

	oakJar := binPath + fmt.Sprintf(oakJar, version)

	if !project.Exists(oakJar) {
		fmt.Printf("Downloading oak version %s...\n", version)
		_, err = http.DownloadFile(oakJar, fmt.Sprintf(oakRunRepository, version, version), ``, ``, true)
		if err != nil {
			return ``, err
		}
	}
	return oakJar, nil
}

// Execute 's command
func Execute(oakPath string, opts, args []string) error {

	path, err := project.GetInstancesDirLocation()
	if err != nil {
		return err
	}

	param := opts
	param = append(param, []string{"-jar", oakPath}...)
	param = append(param, args...)
	cmd := exec.Command("java", param...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
