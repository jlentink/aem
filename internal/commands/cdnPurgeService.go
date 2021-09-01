package commands

import (
	"github.com/fastly/go-fastly/v3/fastly"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandCdnPurgeService struct {
	verbose bool
	cmd     *cobra.Command
	name    string
	apiKey  string

}

func (c *commandCdnPurgeService) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "purge-service",
		Aliases: []string{},
		Short:   "Purge CDN service",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	c.cmd.Flags().StringVarP(&c.name, "name", "n", "", "CDN name")
	c.cmd.Flags().StringVarP(&c.apiKey, "credentials", "c", "", "API key / Credentials")
	c.cmd.MarkFlagRequired("name") // nolint: errcheck



	return c.cmd
}

func (c *commandCdnPurgeService) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCdnPurgeService) run(cmd *cobra.Command, args []string) {
	cdn, err := getCdnConfig(c.name)
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	if strings.EqualFold(cdn.CdnType, "fastly") {
		c.fastlyPurgeService(cdn)
	}
}

func (c *commandCdnPurgeService) fastlyPurgeService(cdn *objects.CDN) {

	if c.apiKey == "" {
		apiKey, err := cdn.GetAPIKey()
		if err != nil {
			output.Printf(output.NORMAL, "%s", err.Error())
			os.Exit(ExitError)
		}
		c.apiKey = apiKey
	}

	client, err := fastly.NewClient(c.apiKey)
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	purge, err := client.PurgeAll(&fastly.PurgeAllInput{
		ServiceID: cdn.ServiceID,
	})

	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}
	if strings.EqualFold(purge.Status, "ok") {
		output.Printf(output.NORMAL, "🚮 %s purged - %s", cdn.ServiceID, purge.ID)
		os.Exit(ExitNormal)
	}
	output.Printf(output.NORMAL, "🤬 Error purging - %s", purge.Status)
	os.Exit(ExitError)
}
