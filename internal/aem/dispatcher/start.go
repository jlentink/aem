package dispatcher

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"os"
	"os/exec"
)

func Start(i objects.Instance, cnf *objects.Config, forGround bool) error {
	if !DaemonRunning() {
		return fmt.Errorf("docker daemon is not running")
	}

	pwd, err := project.GetWorkDir()
	if err != nil {
		return fmt.Errorf("could not get working directory")
	}

	author, err := aem.GetByName(i.Author, cnf.Instances)
	if err != nil {
		return err
	}

	publisher, err := aem.GetByName(i.Publisher, cnf.Instances)
	if err != nil {
		return err
	}

	dispatcherEndpoint := i.DispatcherEndpoint
	if dispatcherEndpoint == "" {
		dispatcherEndpoint = "host.docker.internal"
	}

	fp, err := project.Create(fmt.Sprintf("%s/dispatcher/src/empty.txt", pwd))
	if err != nil {
		return fmt.Errorf("could not create empty file")
	}
	fp.Close()

	options := []string {
		"run",
		"--rm",
		"-p",
		fmt.Sprintf("%d:80", i.Port),
		"-p",
		fmt.Sprintf("%d:443", i.SPort),
		"--detach",
		fmt.Sprintf("--name=%s", processName(cnf)),
		"-v",
		fmt.Sprintf("%s/dispatcher/src/conf:/usr/local/apache2/conf", pwd),
		"-v",
		fmt.Sprintf("%s/dispatcher/src/conf.d:/usr/local/apache2/conf.d", pwd),
		"-v",
		fmt.Sprintf("%s/dispatcher/src/conf.dispatcher.d:/usr/local/apache2/conf.dispatcher.d", pwd),
		"-v",
		fmt.Sprintf("%s/dispatcher/src/conf.modules.d:/usr/local/apache2/conf.modules.d", pwd),
		"-v",
		fmt.Sprintf("%s/dispatcher/src/empty.txt:/usr/local/apache2/conf.modules.d/00-systemd.conf", pwd),
		"-e",
		fmt.Sprintf("AUTHOR_IP=%s", dispatcherEndpoint),
		"-e",
		fmt.Sprintf("AUTHOR_PORT=%d", author.Port),
		"-e",
		fmt.Sprintf("PUBLISH_IP=%s", dispatcherEndpoint),
		"-e",
		"CRX_FILTER=allow",
		"-e",
		fmt.Sprintf("PUBLISH_PORT=%d", publisher.Port),
		"-e",
		"AUTHOR_DOCROOT=/var/www/author",
		"-e",
		"PUBLISH_DOCROOT=/var/www/html",
		"-e",
		"PUBLISH_DEFAULT_HOSTNAME=publish",
		"-e",
		"AUTHOR_DEFAULT_HOSTNAME=author",
		"-e",
		fmt.Sprintf("DISP_ID=dispatcher-%s", processName(cnf)),
		fmt.Sprintf("jlentink/aem-dispatcher:%s", i.DispatcherVersion),
	}
	cmd := exec.Command("docker", options...)

	if !forGround {
		return cmd.Start()
		if err != nil {
			return err
		}
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return  cmd.Run()
}