package commands

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandLog struct {
	verbose      bool
	instanceName string
	listLogs     bool
	follow       bool
	log          string
}

func (c *commandLog) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "log",
		Aliases: []string{"logs"},
		Short:   "List error log or application log",
		PreRun:  c.preRun,
		Run:     c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to stop")
	cmd.Flags().BoolVarP(&c.listLogs, "list", "", false, "List available log files")
	cmd.Flags().BoolVarP(&c.follow, "follow", "f", false, "Actively follow lines when they come in")
	cmd.Flags().StringVarP(&c.log, "log", "l", "error.log", "Which file(s) to follow")
	return cmd
}

func (c *commandLog) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandLog) list(i *objects.Instance) {
	files, err := aem.ListLogFiles(i)
	if err != nil {
		output.Printf(output.NORMAL, "Could not list log files (%s)", err.Error())
		os.Exit(ExitError)
	}

	output.Printf(output.NORMAL, "Available log files.\n")
	for _, file := range files {
		if !file.IsDir() {
			output.Printf(output.NORMAL, " - %s\n", file.Name())
		}
	}
	os.Exit(ExitNormal)
}

func (c *commandLog) run(cmd *cobra.Command, args []string) {
	_, i, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if c.listLogs {
		c.list(i)
	}

	path, _ := project.GetLogDirLocation(*i)
	t, _ := tail.TailFile(path+c.log, tail.Config{Follow: c.follow})
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
