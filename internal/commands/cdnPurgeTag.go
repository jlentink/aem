package commands

import (
	"github.com/fastly/go-fastly/v3/fastly"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandCdnPurgeTag struct {
	verbose bool
	cmd     *cobra.Command
	name    string
	apiKey  string
	tags  string
	soft  bool
}

func (c *commandCdnPurgeTag) setup() *cobra.Command {
	c.cmd = &cobra.Command{
		Use:     "purge-tag",
		Aliases: []string{"purge-tags"},
		Short:   "Purge cdn based on tag.",
		PreRun:  c.preRun,
		Run:     c.run,
	}

	c.cmd.Flags().StringVarP(&c.name, "name", "n", "", "CDN name")
	c.cmd.Flags().StringVarP(&c.apiKey, "credentials", "c", "", "API key / Credentials")
	c.cmd.Flags().StringVarP(&c.tags, "tag", "t", "", "Tag(s) comma seperated")
	c.cmd.Flags().BoolVarP(&c.soft, "soft", "s", false, "Soft purge")
	c.cmd.MarkFlagRequired("name") // nolint: errcheck
	return c.cmd
}

func (c *commandCdnPurgeTag) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandCdnPurgeTag) run(cmd *cobra.Command, args []string) {
	cdn, err := getCdnConfig(c.name)
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}

	if strings.EqualFold(cdn.CdnType, "fastly") {
		c.fastlyPurgeTag(cdn)
	}
}

func (c *commandCdnPurgeTag) fastlyPurgeTag(cdn *objects.CDN){
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

	tags := strings.Split(c.tags, ",")

	purge, err := client.PurgeKeys(&fastly.PurgeKeysInput{
		ServiceID: cdn.ServiceID,
		Keys:      tags,
		Soft:      c.soft,
	})
	if err != nil {
		output.Printf(output.NORMAL, "%s", err.Error())
		os.Exit(ExitError)
	}
	if err == nil {
		output.Printf(output.NORMAL, "🚮 purged  %s - %s", cdn.ServiceID, c.tags)
		os.Exit(ExitNormal)
	}
	output.Printf(output.NORMAL, "🤬 Error purging - %s", purge)
	os.Exit(ExitError)
}