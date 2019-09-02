package commands

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type commandWatch struct {
	verbose      bool
	instanceName string
	watchers     []string
}

func (c *commandWatch) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "watch",
		Short:  "Watch file changes and deploy (Work in progress)",
		PreRun: c.preRun,
		Run:    c.run,
	}
	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Send changes to this instance")
	return cmd
}

func (c *commandWatch) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(c.verbose)
}

func (c *commandWatch) run(cmd *cobra.Command, args []string) {
	_, _, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		output.Printf(output.NORMAL, "Could not start watcher: %s", err.Error())
		os.Exit(ExitError)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}

				fmt.Printf("%s\n", event.Name)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/tmp/")
	err = watcher.Add("/tmp/bla1")
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
