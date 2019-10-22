package commands

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
)

type commandGenerateDump struct {
	verbose bool
}

func (c *commandGenerateDump) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Write default templates to disk",
		Aliases: []string{},
		PreRun:  c.preRun,
		Run:     c.run,
	}

	return cmd
}

func (c *commandGenerateDump) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)

	ConfigCheckListProjects()
	RegisterProject()
}

func (c *commandGenerateDump) run(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("templates"); !os.IsNotExist(err) {
		output.Printf(output.NORMAL, "Found a template dir. No need to run. exiting.")
		os.Exit(ExitNormal)
	}

	box := packr.New("templates", "../../_templates/")
	files := box.List()

	for _, entry := range files {
		f := path.Base(entry)
		d := "templates/" + path.Dir(entry)
		c, err := box.Find(entry)
		if err != nil {
			output.Printf(output.NORMAL, "Found a template dir. No need to run. exiting.")
			os.Exit(ExitNormal)
		}
		err = os.MkdirAll(d, 0777)
		if err != nil {
			output.Printf(output.NORMAL, "Error while writing directory to disk. %s\n", err.Error())
			os.Exit(ExitNormal)
		}

		err = ioutil.WriteFile(d+"/"+f, c, 0666)
		if err != nil {
			output.Printf(output.NORMAL, "Error while writing file to disk. %s\n", err.Error())
			os.Exit(ExitNormal)
		}
		output.Printf(output.VERBOSE, "Creating new template file: %s\n", entry)
		fmt.Print(".")
	}
	fmt.Print("\n Files written to disk. You can now edit or start generating code.\n")
}
