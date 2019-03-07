package main

import (
	"github.com/zalando/go-keyring"
)

type instance struct {
}

func (i *instance) getPasswordForInstance(instance aemInstanceConfig) string {
	if config.KeyRing {
		pw, err := keyring.Get(i.serviceName(instance), instance.Username)
		exitFatal(err, "Error retrieving password")
		return pw
	}
	return instance.Password
}

func (i *instance) getByName(instanceName string) aemInstanceConfig {
	for _, instance := range config.Instances {
		if instanceName == instance.Name {
			return instance
		}
	}
	exitProgram("Instance %s is not defined.\n", instanceName)
	return aemInstanceConfig{}
}

func (i *instance) serviceName(instance aemInstanceConfig) string {
	return ServiceName + "-" + instance.Name
}

func (i *instance) keyRingSetPassword(instance aemInstanceConfig, password string) error {
	return keyring.Set(i.serviceName(instance), instance.Username, password)
}

func (i *instance) keyRingGetPassword(instance aemInstanceConfig) (string, error) {
	return keyring.Get(i.serviceName(instance), instance.Username)
}
