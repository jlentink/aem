package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandPackageDownload struct {
	verbose      bool
	instanceName string
	packageName  string
}

func (c *commandPackageDownload) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "package-download",
		Short:   "List packages",
		Aliases: []string{"pdownload", "pdown"},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().StringVarP(&c.packageName, "package", "p", ``, "Package name. E.g: name, name:1.0.0")
	return cmd
}

func (c *commandPackageDownload) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageDownload) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if len(c.packageName) > 0 {
		c.downloadByName(i)
	} else {
		c.downloadSearch(i)
	}
}

func (c *commandPackageDownload) downloadByName(i *objects.Instance) {
	_, err := pkg.DownloadWithName(i, c.packageName)

	if err != nil {
		output.Printf(output.NORMAL, "Could not download package. %s", err.Error())
		os.Exit(ExitError)
	}
}

func (c *commandPackageDownload) downloadSearch(i *objects.Instance) {
	pkgs, err := pkg.PackageList(*i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not retrieve list from server %s", err.Error())
		os.Exit(ExitError)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ .Name | cyan }} ({{ .Version | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Version | red }})",
		Selected: "\U00002705  Downloading... {{ .Name | red | cyan }}",
		Details: `--------- Package ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Version:" | faint }}	{{ .Version }}
{{ "Size:" | faint }}	{{ .SizeHuman }}
{{ "Created:" | faint }}	{{ .CreatedStr }}
{{ "Modified:" | faint }}	{{ .LastModifiedByStr }}
`,
	}

	searcher := func(input string, index int) bool {
		pkg := pkgs[index]
		name := strings.Replace(strings.ToLower(pkg.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select package:",
		Items:     pkgs,
		Templates: templates,
		Size:      20,
		Searcher:  searcher,
	}

	in, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	pkg.Download(i, &pkgs[in])
}
