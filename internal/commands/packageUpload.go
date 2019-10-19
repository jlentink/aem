package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandPackageUpload struct {
	verbose       bool
	force         bool
	install       bool
	path          string
	instanceName  string
	instanceGroup string
}

func (c *commandPackageUpload) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "package-upload",
		Aliases: []string{"pupload", "pup"},
		Short:   "Upload package to aem",
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to install package on")
	cmd.Flags().StringVarP(&c.instanceGroup, "group", "g", ``, "Group to install package on")
	cmd.Flags().StringVarP(&c.path, "package", "p", ``, "Path to package")
	cmd.Flags().BoolVarP(&c.force, "force", "f", false, "Force upload")
	cmd.Flags().BoolVarP(&c.install, "install", "i", false, "Install package")
	cmd.MarkFlagRequired("package")          // nolint: errcheck
	cmd.MarkFlagFilename("package", "*.zip") // nolint: errcheck
	return cmd
}

func (c *commandPackageUpload) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandPackageUpload) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName, c.instanceGroup, []string{aem.RoleAuthor, aem.RolePublisher})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	for _, instance := range i {
		output.Printf(output.NORMAL, "\U0001F4A8 Uploading %s to: %s (install %t, force: %t)\n", c.path, instance.Name, c.install, c.force)
		crx, err := pkg.Upload(instance, c.path, c.install, c.force)
		if err != nil {
			output.Printf(output.NORMAL, "\U0000274C %s\n", err.Error())
			os.Exit(ExitError)
		}
		output.Printf(output.VERBOSE, "%s", crx.Response.Data.Log.Text)
	}
}
