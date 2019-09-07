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
	return cmd
}

func (c *commandOpen) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
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

	if i.Type == "dispatch" {
		fmt.Printf("use: ssh %s@%s\n", i.Username, i.Hostname)
		return
	}

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", aem.URLString(i))
		cmd.Start()
	case "darwin":
		cmd := exec.Command("open", aem.URLString(i))
		cmd.Start()
	case "linux":
		cmd := exec.Command("xdg-open", aem.URLString(i))
		cmd.Start()

	default:
		fmt.Printf("unsuported operating systen %s", runtime.GOOS)
	}
}
