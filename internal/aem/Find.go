package aem

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/sliceutil"
)

// Roles available for selection
const (
	RoleAuthor     = "author"
	RoleDispatcher = "dispatch"
	RolePublisher  = "publish"
)

// GetByName find instance by name
func GetByName(n string, i []objects.Instance) (*objects.Instance, error) {
	for _, instance := range i {
		instance := instance
		if n == instance.Name {
			return &instance, nil
		}
		if len(instance.Aliases) > 0 {
			for _, alias := range instance.Aliases {
				if n == alias {
					return &instance, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("instance %s is not defined", n)
}

// GetByGroup find instances by group
func GetByGroup(g string, i []objects.Instance) ([]objects.Instance, error) {
	instances := make([]objects.Instance, 0)
	for _, instance := range i {
		if g == instance.Group {
			instances = append(instances, instance)
		}
	}

	err := fmt.Errorf("could not find instance in group. (%s)", g)
	if len(instances) > 0 {
		err = nil
	}

	return instances, err
}

// GetByGroupAndRole find instances by group and role
func GetByGroupAndRole(g string, i []objects.Instance, r string) ([]objects.Instance, error) {
	return GetByGroupAndRoles(g, i, []string{r})
}

// GetByGroupAndRoles find instances by group and roles
func GetByGroupAndRoles(g string, i []objects.Instance, r []string) ([]objects.Instance, error) {
	instances := make([]objects.Instance, 0)
	for _, instance := range i {
		if g == instance.Group && sliceutil.InSliceString(r, instance.Type) {
			instances = append(instances, instance)
		}
	}

	err := fmt.Errorf("could not find instance in group. (%s)", g)
	if len(instances) > 0 {
		err = nil
	}

	return instances, err
}
