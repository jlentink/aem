package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

type commandOpen struct {
	verbose      bool
	useIP        bool
	useSSH       bool
	instanceName string
}

func (c *commandOpen) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "open",
		Short:  "Open URL for Adobe Experience Manager instance in browser",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().BoolVarP(&c.useIP, "ip", "i", false, "Show Ip instead of hostname")
	cmd.Flags().BoolVarP(&c.useSSH, "ssh", "s", false, "Show SSH url")
	return cmd
}

func (c *commandOpen) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")

	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandOpen) run(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		c.instanceName = args[0]
	}
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if i.Hostname == "" {
		c.useIP = false
	}

	if i.Type == "dispatch" || c.useSSH {
		if i.SSHUsername != "" {
			i.Username = i.SSHUsername
		}
		if c.useIP {
			fmt.Printf("use:\n ssh %s@%s\n", i.SSHUsername, i.IP)
		} else {
			fmt.Printf("use:\n ssh %s@%s\n", i.SSHUsername, i.Hostname)
		}

		return
	}

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", aem.URLString(i, c.useIP))
		cmd.Start() // nolint: errcheck
	case "darwin":
		cmd := exec.Command("open", aem.URLString(i, c.useIP))
		cmd.Start() // nolint: errcheck
	case "linux":
		cmd := exec.Command("xdg-open", aem.URLString(i, c.useIP))
		cmd.Start() // nolint: errcheck
	default:
		fmt.Printf("unsuported operating systen %s", runtime.GOOS)
	}
}
