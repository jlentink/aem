package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/indexes"
	"github.com/jlentink/aem/internal/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandReindex struct {
	verbose      bool
	instanceName string
	index string
}

func (c *commandReindex) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "indexes-reindex",
		Short:   "Show indexes on instance",
		Aliases: []string{"reindex"},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().StringVarP(&c.index, "index", "i", "", "Index to reindex")
	return cmd
}

func (c *commandReindex) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandReindex) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if c.index == "" {
		searchIndexes, err := indexes.GetIndexes(i)
		if err != nil {
			output.Printf(output.NORMAL, "Error retrieving indexes. (%s)", err.Error())
			os.Exit(ExitError)
		}
		index := c.searchIndex(searchIndexes)
		if index == nil {
			output.Printf(output.NORMAL, "Process killed stopping...")
			os.Exit(ExitError)

		}
		c.index = index.Name
	}

	indexes.Reindex(i, c.index)

}

func (c *commandReindex) searchIndex(indexesList []*indexes.Index) *indexes.Index {

	localIndexes := make([]indexes.Index, 0)
	for _, cIndex := range indexesList {
		localIndexes = append(localIndexes, *cIndex)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ .Name | cyan }} ({{ .Info | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Info | red }})",
		Selected: "\U0001F522  ReIndexing... {{ .Name | red | cyan }}",
		Details: `--------- Package ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Reindex count:" | faint }}	{{ .ReindexCount }}
{{ "Info:" | faint }}	{{ .Info }}
`,
	}


	searcher := func(input string, index int) bool {
		cIndex := indexesList[index]
		name := strings.Replace(strings.ToLower(cIndex.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select Index",
		Items:     localIndexes,
		Templates: templates,
		Size:      20,
		Searcher:  searcher,
	}

	in, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return &localIndexes[in]

}
