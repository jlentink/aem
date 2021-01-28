package aem

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
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
			jar := jar
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
	defer l.Close() // nolint: errcheck
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
		_, err := project.CreateInstanceDir(i)
		if err != nil {
			return err
		}
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
		if len(i.JVMDebugOptions) > 0 {
			javaOptions = append(javaOptions, i.JVMDebugOptions...)
		} else {
			javaOptions = append(javaOptions, cnf.JVMDebugOptions...)
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
	output.Printf(output.VERBOSE, "run java %s", javaOptions)
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
	return ioutil.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
}

func FullStart(i objects.Instance, ignorePid, forceDownload, foreground bool, cnf *objects.Config, extraPackages []*objects.Package) error{
	if PidExists(i) && !ignorePid {
		if !PidHandler(ignorePid, i) {
			return fmt.Errorf("could not ignore pid")
		}
	}

	if TCPPortOpen(i.Port) {
		output.Printf(output.NORMAL, "Port already taken by other application (%d)", i.Port)
		return fmt.Errorf("port is already taken")
	}

	version, err := FindJarVersion(cnf.DefaultVersion, i.Version, cnf)
	if err != nil {
		output.Printf(output.NORMAL, "Could not find Jar", err.Error())
		return fmt.Errorf("could not find Jar")
	}

	_, err = GetJar(forceDownload, version)
	if err != nil {
		output.Printf(output.NORMAL, err.Error())
		return fmt.Errorf("could not download Jar")
	}

	err = Unpack(i, version)
	if err != nil {
		output.Printf(output.NORMAL, "Could not unpack AEM jar (%s)", err.Error())
		return fmt.Errorf("could not unpack AEM jar (%s)", err.Error())
	}

	_, err = WriteLicense(&i, cnf)
	if err != nil {
		output.Printf(output.NORMAL, "Could not write license (%s)", err.Error())
		return fmt.Errorf("could not write license (%s)", err.Error())
	}

	_, err = WriteIgnoreFile()
	if err != nil {
		output.Print(output.NORMAL, "Could not write ignore file\n")
		return fmt.Errorf("could not write ignore file")
	}

	additionalPackages := cnf.AdditionalPackages
	for _, extraPackage := range extraPackages {
		p, err := project.GetLocationForPackage(extraPackage)
		if err != nil {
			output.Print(output.NORMAL, "Could not add addition packages\n")
			return fmt.Errorf("could not add addition packages")
		}
		additionalPackages = append(additionalPackages, p)
	}

	err = SyncPackages(i, *cnf, additionalPackages, forceDownload)
	if err != nil {
		output.Printf(output.NORMAL, "Error while syncing packages. (%s)", err.Error())
		return fmt.Errorf("error while syncing packages. (%s)", err.Error())
	}

	err = Start(i, foreground)
	if err != nil {
		output.Printf(output.NORMAL, "Could not unpack start AEM (%s)", err.Error())
		return fmt.Errorf("could not unpack start AEM (%s)", err.Error())
	}
	return nil
}

func Destroy(i objects.Instance, force bool, cnf objects.Config) error {
	p, err := project.GetInstanceDirLocation(i)
	if err != nil {
		return err
	}
	if project.Exists(p) {
		Stop(i)
		project.RemoveAll(p)
	}
	removeJar(i, cnf)
	return nil
}

func removeJar(currentInstance objects.Instance, cnf objects.Config){
	version, err := FindJarVersion(cnf.DefaultVersion, currentInstance.Version, &cnf)
	if err != nil {
		return
	}

	location, err := project.GetJarFileLocation(version)
	if err != nil {
		return
	}

	project.Remove(location)
}

// SyncPackages Sync the packages in the install dir
func SyncPackages(i objects.Instance, c objects.Config, additionalPackages []string, forceDownload bool) error {
	installDir, err := project.CreateInstallDir(i)
	if err != nil {
		return err
	}
	for _, cPkg := range additionalPackages {
		if cPkg[0:1] == "/" {
			project.Copy(cPkg, installDir+path.Base(cPkg))
		} else {
			_, err := http.DownloadFile(installDir+path.Base(cPkg), cPkg, ``, ``, forceDownload)
			if err != nil {
				return err
			}
		}
	}

	return cleanupPackages(installDir, c, additionalPackages)
}

func getAdditionalPackageNames(AdditionalPackages []string) []string {
	list := make([]string, 0)
	for _, value := range AdditionalPackages {
		list = append(list, path.Base(value))
	}

	return list
}

func getAdditionalPackageList(c objects.Config) []string {
	list := make([]string, 0)
	for _, value := range c.AdditionalPackages {
		list = append(list, path.Base(value))
	}

	return list
}

func cleanupPackages(p string, c objects.Config, AdditionalPackages []string) error {
	files, err := project.ReadDir(p)
	pkgList := getAdditionalPackageNames(AdditionalPackages)
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

// PidHandler Handle pid or stop
func PidHandler(ignorePid bool, i objects.Instance) bool {
	prompt := &survey.Confirm{
		Message: "Do you want to ignore the PID file.",
		Help: "A existing PID file is found. This could be caused by already running instance or " +
			"an improper shutdown.",
	}
	err := survey.AskOne(prompt, &ignorePid)
	if err != nil {
		if err.Error() == "interrupt" {
			output.Printf(output.NORMAL, "Program interrupted. (CTRL+C)")
			return false
		}
		output.Printf(output.NORMAL, "Unexpected error: %s", err.Error())
		return false
	}

	if !ignorePid {
		p, _ := project.GetPidFileLocation(i)
		output.Printf(output.NORMAL, "Pid already in place. AEM properly already running. (%s)", p)
		return false

	}
	return true
}

// PidExists Does the pid exists?
func PidExists(i objects.Instance) bool {
	path, _ := project.GetPidFileLocation(i)
	return project.Exists(path)
}
