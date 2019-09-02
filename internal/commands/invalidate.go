package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/dispatcher"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandInvalidate struct {
	verbose       bool
	instanceName  string
	instanceGroup string
	group         string
	path          string
	paths         []string
}

func (c *commandInvalidate) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "invalidate",
		Short:   "Invalidate path's on dispatcher",
		Aliases: []string{"flush"},
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", ``, "Instance to sent invalidate to")
	cmd.Flags().StringVarP(&c.instanceGroup, "group", "g", ``, "Instance group to sent invalidate to")
	cmd.Flags().StringVarP(&c.path, "path", "p", ``, "Package to rebuild")
	return cmd
}

func (c *commandInvalidate) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)
}

func (c *commandInvalidate) run(cmd *cobra.Command, args []string) {
	var instances []objects.Instance
	if len(c.instanceName) == 0 && len(c.instanceGroup) == 0 {
		output.Print(output.NORMAL, "Please set group (-g|--group) or instance(-n|--name) to sent invalidate to.\n")
		os.Exit(ExitError)
	}
	if len(c.instanceName) > 0 {
		_, i, errorString, err := getConfigAndInstance(c.instanceName)
		instances = append(instances, *i)
		if err != nil {
			output.Printf(output.NORMAL, errorString, err.Error())
			os.Exit(ExitError)
		}
	} else {
		_, is, errorString, err := getConfigAndGroupWithRole(c.instanceGroup, aem.RoleDispatcher)
		instances = is
		if err != nil {
			output.Printf(output.NORMAL, errorString, err.Error())
			os.Exit(ExitError)
		}

	}

	if len(c.path) > 0 {
		aem.Cnf.InvalidatePaths = []string{c.path}
	}

	for _, instance := range instances {
		for _, p := range aem.Cnf.InvalidatePaths {
			output.Printf(output.NORMAL, "\U0001F5D1 Invalidating: %s (%s)\n", p, instance.Name)
			err := dispatcher.Invalidate(&instance, p)
			if err != nil {
				output.Printf(output.NORMAL, "Could not invalidate path: %s\n", err.Error())
				os.Exit(ExitError)
			}
		}
	}
}
