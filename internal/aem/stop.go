package aem

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func stopWindows(pid, runDir string) error {
	cmd := exec.Command("taskkill", "/PID", string(pid), "/f")
	cmd.Dir = runDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func stopNix(pid, runDir string) error {
	cmd := exec.Command("kill", string(pid))
	cmd.Dir = runDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Signal sends signal to pid
func Signal(pid, runDir string) error {
	cmd := exec.Command("kill", "-0", string(pid))
	cmd.Dir = runDir
	return cmd.Run()
}

func readPid(pidFile string) (string, error) {
	pid, err := ioutil.ReadFile(pidFile)
	return string(pid), err
}

// Stop an instance
func Stop(i objects.Instance) error {
	pidPath, _ := project.GetPidFileLocation(i)
	runPath, _ := project.GetRunDirLocation(i)
	if !project.Exists(pidPath) {
		return fmt.Errorf("could not find pidfile for : %s", i.Name)
	}

	pid, err := readPid(pidPath)
	if err != nil {
		return fmt.Errorf("could not find pidfile for : %s", err.Error())
	}

	if strings.ToLower(runtime.GOOS) == "windows" {
		err := stopWindows(pid, runPath)
		project.Remove(pidPath)
		return err
	}

	err = stopNix(pid, runPath)
	if err != nil {
		return err
	}
	fmt.Printf("Waiting for aem to stop (pid: %s)", pid)
	retry := 0
	for retry <= 180 {
		err = Signal(pid, runPath)
		if err != nil {
			project.Remove(pidPath)
			fmt.Print("\n")
			return nil
		}
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		retry++
	}
	fmt.Print("\n")
	return err
}
