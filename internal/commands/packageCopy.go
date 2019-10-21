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

type commandPackageCopy struct {
	verbose        bool
	instanceName   string
	toInstanceName string
	cPackage       string
}

func (c *commandPackageCopy) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "copy",
		Short:  "Copy packages from one instance to another",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "from", "f", ``, "Instance to copy from")
	cmd.Flags().StringVarP(&c.toInstanceName, "to", "t", ``, "Destination Instance")
	cmd.Flags().StringVarP(&c.cPackage, "package", "p", ``, "Package to copy")
	cmd.MarkFlagRequired("from") // nolint: errcheck
	cmd.MarkFlagRequired("to")   // nolint: errcheck
	return cmd
}

func (c *commandPackageCopy) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageCopy) run(cmd *cobra.Command, args []string) {
	_, f, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	_, t, errorString, err := getConfigAndInstance(c.toInstanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	var dp *objects.Package
	if len(c.cPackage) > 0 {
		output.Printf(output.NORMAL, "\U0001F69A %s => %s\n", f.Name, t.Name)
		dp, err = pkg.DownloadWithName(f, c.cPackage)
		if err != nil {
			output.Printf(output.NORMAL, "Could not download package from %s: %s", f.Name, err.Error())
			os.Exit(ExitError)
		}
	} else {
		dp = c.downloadSearch(f)
		output.Printf(output.NORMAL, "\U0001F69A %s => %s\n", f.Name, t.Name)
		_, err = pkg.Download(f, dp)
		if err != nil {
			output.Printf(output.NORMAL, "Could not download package from %s: %s", f.Name, err.Error())
			os.Exit(ExitError)
		}
	}

	p, err := project.GetLocationForPackage(dp)
	if err != nil {
		output.Printf(output.NORMAL, "Getting package location ended up in an error: %s", err.Error())
		os.Exit(ExitError)
	}

	crx, err := pkg.Upload(*t, p, true, true)
	if err != nil {
		output.Printf(output.NORMAL, "Issue while coping package: %s", err.Error())
		os.Exit(ExitError)
	}
	output.Printf(output.VERBOSE, "%s", crx.Response.Data.Log.Text)
}

func (c *commandPackageCopy) downloadSearch(i *objects.Instance) *objects.Package {
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
		return nil
	}

	return &pkgs[in]
}
