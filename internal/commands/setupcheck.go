package commands

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/jlentink/aem/internal/cli/setup"
	"github.com/spf13/cobra"
	"os"
)

type commandSetupCheck struct {
	verbose bool
	minimal bool
}

func (c *commandSetupCheck) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "setup-check",
		Aliases: []string{"setup"},
		Short:   "Check if all needed binaries are available for all functionality",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	return cmd
}

func (c *commandSetupCheck) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
}

func (c *commandSetupCheck) printStatus(bin setup.Description) {
	color := ct.Green
	statusMsg := " FOUND "
	if !bin.Found {
		statusMsg = "MISSING"
		switch bin.Required {
		case setup.Required:
			color = ct.Red
		case setup.Optional:
			color = ct.Yellow
		}
	}
	fmt.Print("[ ")
	ct.ChangeColor(color, false, ct.None, false)
	fmt.Print(statusMsg)
	ct.ResetColor()
	fmt.Printf(" ] %s -  %s\n", bin.Bin, bin.Description)
}
func (c *commandSetupCheck) run(cmd *cobra.Command, args []string) {
	for _, bin := range setup.Check() {
		c.printStatus(bin)
	}
	os.Exit(ExitNormal)
}
