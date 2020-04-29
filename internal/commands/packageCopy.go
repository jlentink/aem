package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
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
	toGroup        string
	install        bool
	force          bool
	cPackage       []string
}

func (c *commandPackageCopy) setup() *cobra.Command {
	c.install = false
	c.force = false
	cmd := &cobra.Command{
		Use:    "copy",
		Short:  "Copy packages from one instance to another",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "from", "f", ``, "Instance to copy from")
	cmd.Flags().StringVarP(&c.toInstanceName, "to", "t", aem.GetDefaultInstanceName(), "Destination Instance")
	cmd.Flags().StringVarP(&c.toGroup, "group", "g", ``, "Destination Instance")
	cmd.Flags().StringArrayVarP(&c.cPackage, "package", "p", []string{}, "Package to copy")
	cmd.Flags().BoolVarP(&c.install, "install", "i", false, "Install package after upload")
	cmd.Flags().BoolVarP(&c.force, "force", "F", false, "Force package")
	cmd.MarkFlagRequired("from") // nolint: errcheck
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

	_, t, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.toInstanceName, c.toGroup, []string{aem.RoleAuthor, aem.RolePublisher})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if len(t) == 0 {
		output.Printf(output.NORMAL, "No hosts in group %s", c.toGroup)
		os.Exit(ExitError)
	}

	dp := make([]*objects.Package, 0)
	if len(c.cPackage) > 0 {
		output.Printf(output.NORMAL, "\U0001F69A %s => %s\n", f.Name, t[0].Name)
		for _, cPackage := range c.cPackage {
			po, ierr := pkg.DownloadWithName(f, cPackage)
			dp = append(dp, po)
			if ierr != nil {
				output.Printf(output.NORMAL, "Could not download package from %s: %s", f.Name, ierr.Error())
				os.Exit(ExitError)
			}

		}
	} else {
		po := c.downloadSearch(f)
		dp = append(dp, po)
		output.Printf(output.NORMAL, "\U0001F69A Downloading from %s\n", f.Name)
		_, err = pkg.Download(f, po)
		if err != nil {
			output.Printf(output.NORMAL, "Could not download package from %s: %s", f.Name, err.Error())
			os.Exit(ExitError)
		}
	}

	for _, po := range dp {
		p, err := project.GetLocationForPackage(po)
		if err != nil {
			output.Printf(output.NORMAL, "Getting package location ended up in an error: %s", err.Error())
			os.Exit(ExitError)
		}
		for _, toInstance := range t {
			output.Printf(output.NORMAL, "\U0001F69A %s => %s (install: %t, force: %t)\n",
				f.Name, toInstance.Name, c.install, c.force)
			crx, err := pkg.Upload(toInstance, p, c.install, c.force)
			if err != nil {
				output.Printf(output.NORMAL, "Issue while coping package: %s", err.Error())
				os.Exit(ExitError)
			}
			output.Printf(output.VERBOSE, "%s", crx.Response.Data.Log.Text)
		}

	}
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
