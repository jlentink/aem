package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/aem/pkg"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
	"time"
)

type commandDestroy struct {
	verbose       bool
	force         bool
	create        bool
	instanceName  string
	instanceGroup string
	backup        bool
	startLevel    string
}

func (c *commandDestroy) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "destroy",
		Short:  "Destroy instance",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to install bundle on")
	cmd.Flags().StringVarP(&c.instanceGroup, "group", "g", ``, "Instance group to install bundle on")
	cmd.Flags().BoolVarP(&c.backup, "no-backup", "b", true, "No Backup of content first.")
	cmd.Flags().BoolVarP(&c.force, "force", "f", false, "Force delete don't ask for confirmation.")
	cmd.Flags().BoolVarP(&c.create, "create", "c", false, "Spin up again after destroy")
	return cmd
}

func (c *commandDestroy) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandDestroy) run(cmd *cobra.Command, args []string) {
	config, is, errorString, err := getConfigAndInstanceOrGroupWithRoles(c.instanceName, c.instanceGroup, []string{aem.RoleAuthor, aem.RolePublisher})
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	for _, i := range is {
		if project.Confirm(fmt.Sprintf("\U0001F198 Are you sure you want to delete %s? (This cannot be undone!!)", i.Name), "Deletion of instance confirmation.", c.force) {
			cPkgs := make([]*objects.Package, 0)
			if !c.backup {
				output.Printf(output.NORMAL, "\U0001F477 Creating package \n")
				now := time.Now()
				versionStr := fmt.Sprintf("%s.%d", now.Format("20060102"), now.UnixNano())
				cPkg, err := pkg.Create(i, config.ContentBackupName, config.ContentBackupGroup, versionStr, config.ContentBackupPaths, false)
				cPkgs = append(cPkgs, cPkg)
				if err != nil {
					output.Printf(output.NORMAL, "Unable to build package. %s", err.Error())
					os.Exit(ExitError)
				}
				output.Printf(output.NORMAL, "\U0001F525 Building package %s\n", cPkg.Name)
				pkg.Rebuild(&i, cPkg)
				pkg.AwaitBuild(&i, cPkg)
				output.Printf(output.NORMAL, "\U0001F64C Downloading package %s\n", cPkg.Name)
				pkg.Download(&i, cPkg)
			}

			aem.Destroy(i, c.force, *config)
			if c.create {
				aem.FullStart(i, true, true, false, config, cPkgs)
			}
		}
	}
}
