package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/output"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func rebuildPackage(i *objects.Instance, p string) {
	output.Printf(output.NORMAL, "\U0001F525 %s building package: %s\n", i.Name, p)
	_, err := pkg.RebuildbyName(i, p)
	if err != nil {
		output.Printf(output.NORMAL, "Rebuild failed: %s", err.Error())
		os.Exit(ExitError)
	}
}

func installPackage(i *objects.Instance, p string) {
	output.Printf(output.NORMAL, "\U0001F48A %s installing package: %s\n", i.Name, p)
	_, err := pkg.InstallByName(i, p)
	if err != nil {
		output.Printf(output.NORMAL, "Install failed: %s", err.Error())
		os.Exit(ExitError)
	}
}

func rebuildPackageSearch(i *objects.Instance) {
	pkgs, err := pkg.PackageList(*i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not retrieve list from server %s", err.Error())
		os.Exit(ExitError)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ .Name | cyan }} ({{ .Version | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Version | red }})",
		Selected: "\U0001F6E0  Rebuilding... {{ .Name | red | cyan }}",
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

	output.Printf(output.NORMAL, "\U0001F525 %s", i.Name)
	_, err = pkg.Rebuild(i, &pkgs[in])
	if err != nil {
		output.Printf(output.NORMAL, "Rebuild failed: %s", err.Error())
		os.Exit(ExitError)
	}
}

func installPackageSearch(i *objects.Instance) {
	pkgs, err := pkg.PackageList(*i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not retrieve list from server %s", err.Error())
		os.Exit(ExitError)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ .Name | cyan }} ({{ .Version | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Version | red }})",
		Selected: "\U0001F48A  installing... {{ .Name | red | cyan }}",
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

	output.Printf(output.NORMAL, "\U0001F525 %s", i.Name)
	_, err = pkg.Install(i, &pkgs[in])
	if err != nil {
		output.Printf(output.NORMAL, "Install failed: %s", err.Error())
		os.Exit(ExitError)
	}
}
