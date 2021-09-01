package commands

import (
	"github.com/fastly/go-fastly/v3/fastly"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandCdnPurgeURL struct {
	verbose bool
	cmd     *cobra.Command
	name    string
	url     string
	apiKey  string
	soft  bool
}

func (c *commandCdnPurgeURL) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "purge-url",
		Aliases: []string{},
		Short:   "Purge CDN based on URL",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	c.cmd.Flags().StringVarP(&c.name, "name", "n", "", "CDN name")
	c.cmd.Flags().StringVarP(&c.apiKey, "credentials", "c", "", "API key / Credentials")
	c.cmd.Flags().StringVarP(&c.url, "url", "u", "", "url to purge")
	c.cmd.Flags().BoolVarP(&c.soft, "soft", "s", false, "Soft purge")

	c.cmd.MarkFlagRequired("name") // nolint: errcheck
	c.cmd.MarkFlagRequired("url") // nolint: errcheck

	return c.cmd
}

func (c *commandCdnPurgeURL) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCdnPurgeURL) run(cmd *cobra.Command, args []string) {
	cdn, err := getCdnConfig(c.name)
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	if strings.EqualFold(cdn.CdnType, "fastly") {
		c.fastlyPurge(cdn)
	}
}

func (c *commandCdnPurgeURL) fastlyPurge(cdn *objects.CDN){
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

	purge, err := client.Purge(&fastly.PurgeInput{
		URL:  c.url,
		Soft: c.soft,
	})
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	if strings.EqualFold(purge.Status, "ok") {
		output.Printf(output.NORMAL, "🚮 %s purged - %s", c.url, purge.ID)
		os.Exit(ExitNormal)
	}
	output.Printf(output.NORMAL, "🤬 Error purging - %s", purge.Status)
	os.Exit(ExitError)
}
