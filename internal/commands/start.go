package commands

import (
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"github.com/spf13/cobra"
	"os"
)

type commandStart struct {
	verbose       bool
	instanceName  string
	allowRoot     bool
	foreground    bool
	forceDownload bool
	ignorePid     bool
}

func (c *commandStart) setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "start",
		Short:  "Start Adobe Experience Manager instance",
		PreRun: c.preRun,
		Run:    c.run,
	}

	cmd.Flags().StringVarP(&c.instanceName, "name", "n", aem.GetDefaultInstanceName(), "Instance to start")
	cmd.Flags().BoolVarP(&c.forceDownload, "download", "d", false, "Force re-download")
	cmd.Flags().BoolVarP(&c.foreground, "foreground", "f", false, "on't detach aem from current tty")
	cmd.Flags().BoolVarP(&c.allowRoot, "allow-root", "r", false, "Allow to start as root user (UID: 0)")
	cmd.Flags().BoolVarP(&c.ignorePid, "ignore-pid", "p", false, "Ignore existing PID file and start AEM")

	return cmd
}

func (c *commandStart) preRun(cmd *cobra.Command, args []string) {
	c.verbose, _ = cmd.Flags().GetBool("verbose")
	output.SetVerbose(verbose)
}

func (c *commandStart) run(cmd *cobra.Command, args []string) {
	if !aem.AllowUserStart(c.allowRoot) {
		output.Print(output.NORMAL, "You are starting aem as a root. This is not allowed. override with: --allow-root\n")
		os.Exit(ExitError)
	}

	cnf, currentInstance, errorString, err := getConfigAndInstance(c.instanceName)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	if aem.PidExists(*currentInstance) && !c.ignorePid {
		p, _ := project.GetPidFileLocation(*currentInstance)
		output.Printf(output.NORMAL, "Pid already in place. AEM properly already running. (%s)", p)
		os.Exit(ExitError)

	}

	if aem.TCPPortOpen(currentInstance.Port) {
		output.Printf(output.NORMAL, "Port already taken by other application (%d)", currentInstance.Port)
		os.Exit(ExitError)
	}

	version, err := aem.FindJarVersion(cnf.DefaultVersion, currentInstance.Version, cnf)
	if err != nil {
		output.Printf(output.NORMAL, errorString, err.Error())
		os.Exit(ExitError)
	}

	_, err = aem.GetJar(c.forceDownload, version)
	if err != nil {
		output.Printf(output.NORMAL, err.Error())
		os.Exit(ExitError)
	}

	err = aem.Unpack(*currentInstance, version)
	if err != nil {
		output.Printf(output.NORMAL, "Could not unpack AEM jar (%s)", err.Error())
		os.Exit(ExitError)
	}

	_, err = aem.WriteLicense(currentInstance, cnf)
	if err != nil {
		output.Printf(output.NORMAL, "Could not unpack AEM jar (%s)", err.Error())
		os.Exit(ExitError)
	}

	_, err = aem.WriteIgnoreFile()
	if err != nil {
		output.Print(output.NORMAL, "Could not write ignore file\n")
		os.Exit(ExitError)
	}

	err = aem.SyncPackages(*currentInstance, *cnf, c.forceDownload)
	if err != nil {
		output.Printf(output.NORMAL, "Error while syncing packages. (%s)", err.Error())
		os.Exit(ExitError)
	}

	err = aem.Start(*currentInstance, c.foreground)
	if err != nil {
		output.Printf(output.NORMAL, "Could not unpack start AEM (%s)", err.Error())
		os.Exit(ExitError)
	}
}
