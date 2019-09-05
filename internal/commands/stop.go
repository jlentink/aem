package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandStop struct {
	verbose       bool
	instanceName  string
}

func (c *commandStop) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "stop",
		Short:  "stop Adobe Experience Manager instance",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	return cmd
}

func (c *commandStop) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandStop) run(cmd *cobra.Command, args []string) {
	_, currentInstance, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	err = aem.Stop(*currentInstance)
	if err != nil {
		output.Printf(output.NORMAL, "Could not stop instance. (%s)", err.Error())
		os.Exit(ExitError)
	}
}
