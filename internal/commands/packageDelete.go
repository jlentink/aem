package commands

import (
	"bufio"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type commandPackageDelete struct {
	verbose      bool
	plain      bool
	instanceName string
	group string
	force bool
	packageName string
	packageFile string
}

func (c *commandPackageDelete) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete packages",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().StringVarP(&c.packageName, "package", "p", ``, "Package name. E.g: name,1.0.0,group")
	cmd.Flags().StringVarP(&c.packageFile, "package-file", "F", ``, "file to read packages from")
	cmd.Flags().BoolVarP(&c.force, "force", "f", false, "Force delete")
	cmd.MarkFlagRequired("name") // nolint: errcheck
	return cmd
}

func (c *commandPackageDelete) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageDelete) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if len(c.packageFile) > 0 {
		pkgs := c.readPackagesFromFile()
		c.delete(pkgs, i)
	} else {
		parts := strings.Split(c.packageName, ",")
		if len(parts) == 3 {
			cPkg, err := pkg.GetPackageByNameAndGroupAndVersion(*i, parts[0], parts[1], parts[2])
			if err != nil {
				output.Printf(output.NORMAL, "Could not find package on instance.")
				os.Exit(ExitError)
			}
			c.delete([]objects.Package{*cPkg}, i)
		}
	}
}

func (c *commandPackageDelete) delete(pkgs []objects.Package, i *objects.Instance){
	for _, cPkg := range pkgs {
		if !c.force {
			if c.confirm(cPkg) {
				fmt.Printf("🚮 Deleting package: %s - %s (%s)\n", cPkg.Name, cPkg.Group, cPkg.Version)
				pkg.Delete(i, &cPkg)

			}
		} else {
			fmt.Printf("🚮 Deleting package: %s - %s (%s)\n", cPkg.Name, cPkg.Group, cPkg.Version)
			pkg.Delete(i, &cPkg)
		}
	}
}

func (c *commandPackageDelete) readPackagesFromFile() []objects.Package {
	pkgs := make([]objects.Package, 0)
	file, err := os.Open(c.packageFile)
	if err != nil {
		output.Printf(output.NORMAL, "Could not read: %s (%s)\n", c.packageFile, err.Error())
		os.Exit(ExitError)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err = sc.Err(); err != nil {
		output.Printf(output.NORMAL, "Scan error %s \n", err.Error())
		os.Exit(ExitError)
	}

	for i, line := range lines {
		if len(line) != 0 {
			newPackage := objects.Package{}
			status := newPackage.FromString(line)
			if status != true {
				output.Printf(output.NORMAL, "Package description does not contain all data should be package,version,group,Download file. Line %d (%s)", i, line)
				os.Exit(ExitError)
			}
			pkgs = append(pkgs, newPackage)
		}
	}

	return pkgs
}

func (c *commandPackageDelete) confirm(cPkg objects.Package) bool {
	confirm := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete the package %s - %s (%s)\n", cPkg.Name, cPkg.Group, cPkg.Version),
		Help: "",
	}
	err := survey.AskOne(prompt, &confirm)
	if err != nil {
		if err.Error() == "interrupt" {
			output.Printf(output.NORMAL, "Program interrupted. (CTRL+C)")
			return false
		}
		output.Printf(output.NORMAL, "Unexpected error: %s", err.Error())
		return false
	}

	return confirm
}