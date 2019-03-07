package main

import (
	"github.com/zalando/go-keyring"
)

type Instance struct {
}

func (i *Instance) getPasswordForInstance(instance AEMInstanceConfig) string {
	if config.KeyRing {
		pw, err := keyring.Get(i.serviceName(instance), instance.Username)
		exitFatal(err, "Error retrieving password")
		return pw
	} else {
		return instance.Password
	}
}

func (i *Instance) getByName(instanceName string) AEMInstanceConfig {
	for _, instance := range config.Instances {
		if instanceName == instance.Name {
			return instance
		}
	}
	exitProgram("Instance %s is not defined.\n", instanceName)
	return AEMInstanceConfig{}
}


func (i *Instance) serviceName(instance AEMInstanceConfig) string {
	return ServiceName + "-" + instance.Name
}

func (i *Instance) keyRingSetPassword(instance AEMInstanceConfig, password string) error {
	return keyring.Set(i.serviceName(instance), instance.Username, password)
}

func (i *Instance) keyRingGetPassword(instance AEMInstanceConfig) (string, error) {
	return keyring.Get(i.serviceName(instance), instance.Username)
}
