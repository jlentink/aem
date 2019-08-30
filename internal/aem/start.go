package aem

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/http"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/sliceutil"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// AllowUserStart Allow AEM to start for the current user
func AllowUserStart(allow bool) bool {
	if allow {
		return true
	}

	usr, err := getCurrentUser()
	if err != nil {
		return false
	}

	output.Printf(output.VERBOSE, "UserID detected %s\n", usr.Uid)
	if (strings.ToLower(runtime.GOOS) == "linux" || strings.ToLower(runtime.GOOS) == "darwin") && usr.Uid == "0" {
		return false
	}
	return true
}

// FindJarVersion find jar configuration in config for version
func FindJarVersion(configVersion, instanceVersion string, config *objects.Config) (*objects.AemJar, error) {
	var version string
	if len(instanceVersion) > 0 {
		version = instanceVersion
	} else {
		version = configVersion
	}

	for _, jar := range config.AemJar {
		if jar.Version == version {
			return &jar, nil
		}
	}
	return nil, fmt.Errorf("could not find aem jar settings for version: %s", version)
}

//GetJar Download jar to disk
func GetJar(forceDownload bool, aemJar *objects.AemJar) (uint64, error) {
	location, err := project.GetJarFileLocation(aemJar)
	if err != nil {
		return 0, fmt.Errorf("Could not get location to set jar to. ")
	}

	_, err = project.CreateInstancesDir()
	if err != nil {
		return 0, err
	}

	if !project.Exists(location) || forceDownload {
		output.Printf(output.VERBOSE, "Force download: %v\n", forceDownload)
		c, _ := GetConfig()
		if isURL(aemJar.Location) {
			output.Printf(output.VERBOSE, "Download from : %v\n", c.AemJar)
			return http.DownloadFile(location, aemJar.Location, aemJar.Username, aemJar.Password, forceDownload)
		}
		if project.Exists(aemJar.Location) {
			output.Printf(output.VERBOSE, "copying from : %v\n", c.AemJar)
			size, err := project.Copy(aemJar.Location, location)
			if err != nil {
				return 0, err
			}
			return uint64(size), nil
		}
		return 0, fmt.Errorf("could not find jar at %s", c.AemJar)
	}

	output.Printf(output.VERBOSE, "Found jar at: \"%s\". Skipping download...", location)
	return 0, nil
}

func findJar(i objects.Instance) (string, error) {
	appPath, _ := project.GetAppDirLocation(i)
	files, err := ioutil.ReadDir(appPath)
	if err != nil {
		return ``, err
	}

	r, _ := regexp.Compile(`(.*)\.jar`)
	for _, file := range files {
		if r.MatchString(file.Name()) {
			return file.Name(), nil
		}
	}
	return ``, fmt.Errorf("could not find main jar in application")
}

// WriteIgnoreFile Write ignore file to disk
func WriteIgnoreFile() (string, error) {
	return project.WriteGitIgnoreFile()
}

// TCPPortOpen is the port open on the server
func TCPPortOpen(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return true
	}
	defer l.Close()
	return false
}

// Unpack the jar instance on the correct location
func Unpack(i objects.Instance, version *objects.AemJar) error {
	jarPath, _ := project.GetJarFileLocation(version)
	workPath, _ := project.GetInstancesDirLocation()
	unpackPath, _ := project.GetUnpackDirLocation()
	runPath, _ := project.GetRunDirLocation(i)

	if !project.Exists(runPath) {
		if !project.Exists(unpackPath) {
			cmd := exec.Command("java", "-jar", jarPath, "-unpack")
			cmd.Dir = workPath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
		}
		project.CreateInstanceDir(i)
		return project.Rename(unpackPath, runPath)
	}

	return nil
}

// Start start a new instance
func Start(i objects.Instance, forGround bool) error {
	cnf, _ := GetConfig()
	runDir, _ := project.GetRunDirLocation(i)
	jarPath, _ := findJar(i)
	javaOptions := cnf.JVMOpts
	javaOptions = append(javaOptions, i.JVMOptions...)
	if i.Debug {
		if len(cnf.JVMDebugOptions) > 0 {
			javaOptions = append(javaOptions, cnf.JVMDebugOptions...)
		}

		if len(i.JVMDebugOptions) > 0 {
			javaOptions = append(javaOptions, i.JVMDebugOptions...)
		}
	}

	javaOptions = append(javaOptions,
		"-Dsling.run.modes="+i.Type+","+i.RunMode,
		"-jar", "app/"+jarPath, "start",
		"-c", runDir,
		"-i", "launchpad",
		"-p", fmt.Sprintf("%d", i.Port),
		"-Dsling.properties=conf/sling.properties")

	cmd := exec.Command("java", javaOptions...)
	cmd.Dir = runDir

	if !forGround {
		err := cmd.Start()
		if err != nil {
			return err
		}
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	pidFile, _ := project.GetPidFileLocation(i)
	fmt.Printf("Starting AEM with (pid: %d)\n", cmd.Process.Pid)
	ioutil.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	return nil
}

// SyncPackages Sync the packages in the install dir
func SyncPackages(i objects.Instance, c objects.Config, forceDownload bool) error {
	installDir, err := project.CreateInstallDir(i)
	if err != nil {
		return err
	}
	for _, cPkg := range c.AdditionalPackages {
		_, err := http.DownloadFile(installDir+path.Base(cPkg), cPkg, ``, ``, forceDownload)
		if err != nil {
			return err
		}
	}

	return cleanupPackages(installDir, c)
}

func getAdditionalPackageList(c objects.Config) []string {
	list := make([]string, 0)
	for _, value := range c.AdditionalPackages {
		list = append(list, path.Base(value))
	}

	return list
}

func cleanupPackages(p string, c objects.Config) error {
	files, err := project.ReadDir(p)
	pkgList := getAdditionalPackageList(c)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() && sliceutil.InSliceString(pkgList, f.Name()) {
			output.Printf(output.VERBOSE, "pkg still required (%s)\n", f.Name())
		} else {
			output.Printf(output.NORMAL, "pkg cleanup (%s)\n", f.Name())
			err := project.Remove(p + f.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// PidExists Does the pid exists?
func PidExists(i objects.Instance) bool {
	path, _ := project.GetPidFileLocation(i)
	if project.Exists(path) {
		return true
	}
	return false
}
