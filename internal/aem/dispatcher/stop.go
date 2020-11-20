package dispatcher

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"os/exec"
)

func Stop(i objects.Instance, cnf *objects.Config) error {
	if !DaemonRunning() {
		return fmt.Errorf("docker daemon is not running")
	}

	options := []string {
		"stop",
		processName(cnf),
	}
	cmd := exec.Command("docker", options...)
	cmd.Start()
	fmt.Printf("Waiting for dispatcher to stop...")
	err := cmd.Wait()
	fmt.Printf("Dispatcher stopped...")
	return err
}