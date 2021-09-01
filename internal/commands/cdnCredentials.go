package commands

import (
	"bufio"
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandCdnCredentials struct {
	verbose bool
	cmd     *cobra.Command
	name    string
	url     string
	apiKey  string
	soft  bool
}

func (c *commandCdnCredentials) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "credentials",
		Aliases: []string{"passwords"},
		Short:   "Store credentials for CDN's",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	c.cmd.Flags().StringVarP(&c.name, "name", "n", "", "CDN name")
	return c.cmd
}

func (c *commandCdnCredentials) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCdnCredentials) run(cmd *cobra.Command, args []string) {
	if c.name != "" {
		cdn, err := getCdnConfig(c.name)
		if err != nil {
			output.Printf(output.NORMAL, "%s", err.Error())
			os.Exit(ExitError)
		}
		c.setCredentials(cdn)
	} else {
		cnf, err := getConfig()
		if err != nil {
			output.Printf(output.NORMAL, "%s", err.Error())
			os.Exit(ExitError)
		}
		for _, cdn := range cnf.CDNs {
			c.setCredentials(&cdn)
		}
	}
}

func (c *commandCdnCredentials) setCredentials(cdn *objects.CDN){
	if strings.EqualFold(cdn.CdnType, "fastly") {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Fastly API key: ")
		pw, _ := reader.ReadString('\n')
		err := cdn.SetAPIKey(pw[:len(pw)-1])
		if err != nil {
			output.Printf(output.NORMAL, "Could not set/update password: %s", err.Error())
			os.Exit(ExitError)
		}
	}
}