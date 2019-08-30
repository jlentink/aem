package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandCopyPackage struct {
	verbose        bool
	dump           bool
	force          bool
	instanceName   string
	toInstanceName string
	cPackage       string
}

func (c *commandCopyPackage) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "package-copy",
		Short:  "Copy packages from one instance to another",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "from", "f", ``, "Instance to copy from")
	cmd.Flags().StringVarP(&c.toInstanceName, "to", "t", ``, "Destination Instance")
	cmd.Flags().StringVarP(&c.cPackage, "package", "p", ``, "Package to copy")
	return cmd
}

func (c *commandCopyPackage) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)
}

func (c *commandCopyPackage) run(cmd *cobra.Command, args []string) {
	_, f, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(EXIT_ERROR)
	}

	_, t, errorString, err := getConfigAndInstance(c.toInstanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(EXIT_ERROR)
	}

	var dp *objects.Package
	if len(c.cPackage) > 0 {
		output.Printf(output.NORMAL, "\U0001F69A %s => %s\n", f.Name, t.Name)
		dp, err = pkg.DownloadWithName(f, c.cPackage)
		if err != nil {
			output.Printf(output.NORMAL, "Could not download package from %s: %s", f.Name, err.Error())
			os.Exit(EXIT_ERROR)
		}
	} else {
		dp = c.downloadSearch(f)
		output.Printf(output.NORMAL, "\U0001F69A %s => %s\n", f.Name, t.Name)
		pkg.Download(f, dp)
	}

	p, err := project.GetLocationForPackage(dp)
	crx, err := pkg.Upload(*t, p, true, true)
	output.Printf(output.VERBOSE, "%s", crx.Response.Data.Log.Text)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(EXIT_ERROR)
	}
}

func (c *commandCopyPackage) downloadSearch(i *objects.Instance) *objects.Package {
	pkgs, err := pkg.PackageList(*i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not retrieve list from server %s", err.Error())
		os.Exit(EXIT_ERROR)
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
		return nil
	}

	return &pkgs[in]
}
