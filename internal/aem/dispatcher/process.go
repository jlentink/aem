package dispatcher

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"os/exec"
	"strings"
)

func DaemonRunning() bool {
	cmd := exec.Command("docker", "version")
	if err := cmd.Run() ; err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return true
		} else {
			return false
		}
	}
	return true
}

func processName(cnf *objects.Config) string {
	return fmt.Sprintf("dispatcher-%s", strings.ToLower(cnf.ProjectName))
}