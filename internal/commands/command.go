package commands

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/sliceutil"
	"github.com/spf13/cobra"
)

// Exit codes for AEM
const (
	ExitNormal = 0
	ExitError  = 1
)

// Command internal interface for commands
type Command interface {
	setup() *cobra.Command
	preRun(cmd *cobra.Command, args []string)
	run(cmd *cobra.Command, args []string)
}

func getConfig() (*objects.Config, error) {
	return aem.GetConfig()
}
func getConfigAndInstance(i string) (*objects.Config, *objects.Instance, string, error) {
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, nil, "Could not load config file. (%s)", err
	}

	currentInstance, err := aem.GetByName(i, cnf.Instances)
	if err != nil {
		return cnf, nil, "Could not find instance. (%s)", err
	}

	aem.Cnf = cnf
	return cnf, currentInstance, ``, nil
}

//nolint
func getConfigAndInstanceOrGroup(i, g string) (*objects.Config, []objects.Instance, string, error) {
	if len(g) > 0 {
		return getConfigAndGroup(g)
	}
	c, in, s, e := getConfigAndInstance(i)
	return c, []objects.Instance{*in}, s, e
}

func getConfigAndInstanceOrGroupWithRoles(i, g string, r []string) (*objects.Config, []objects.Instance, string, error) {
	if len(g) > 0 {
		return getConfigAndGroupWithRoles(g, r)
	}

	c, in, s, e := getConfigAndInstance(i)
	if e != nil {
		return c, nil, s, e
	}

	if !sliceutil.InSliceString(r, in.Type) {
		return c, []objects.Instance{*in}, s, fmt.Errorf("instance is not of type %s", in.Type)
	}

	return c, []objects.Instance{*in}, s, e
}

func getConfigAndGroup(i string) (*objects.Config, []objects.Instance, string, error) {
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, nil, "Could not load config file. (%s)", err
	}

	currentInstance, err := aem.GetByGroup(i, cnf.Instances)
	if err != nil {
		return cnf, nil, "Could not find instance. (%s)", err
	}
	aem.Cnf = cnf
	return cnf, currentInstance, ``, nil
}

func getConfigAndGroupWithRole(i, r string) (*objects.Config, []objects.Instance, string, error) {
	return getConfigAndGroupWithRoles(i, []string{r})
}

func getConfigAndGroupWithRoles(i string, r []string) (*objects.Config, []objects.Instance, string, error) {
	cnf, err := aem.GetConfig()
	if err != nil {
		return nil, nil, "Could not load config file. (%s)", err
	}

	currentInstance, err := aem.GetByGroupAndRoles(i, cnf.Instances, r)
	if err != nil {
		return cnf, nil, "Could not find instance. (%s)", err
	}
	aem.Cnf = cnf
	return cnf, currentInstance, ``, nil
}

// GetInstancesAndConfig gets config and configuration for instance or group
func GetInstancesAndConfig(i, g string) (*objects.Config, []objects.Instance, string, error) {
	if len(g) > 0 {
		return getConfigAndGroup(g)
	}
	cnf, instance, errString, err := getConfigAndInstance(i)
	aem.Cnf = cnf
	return cnf, []objects.Instance{*instance}, errString, err

}
