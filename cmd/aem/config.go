package main

import "fmt"

type configStruct struct {
	ConfigFile         string
	Verbose            bool                `toml:"verbose"`
	Packages           []string            `toml:"commandPullContent"`
	Command            string              `toml:"command,omitempty"`
	CommandArgs        []string            `toml:"command,omitempty"`
	Instances          []aemInstanceConfig `toml:"instance"`
	JVMOptions         []string            `toml:"jvm-options"`
	JVMDebugOptions    []string            `toml:"jvm-debug-options"`
	AemJar             string              `toml:"jarLocation"`
	AemJarUsername     string              `toml:"jarUsername"`
	AemJarPassword     string              `toml:"jarPassword"`
	WatchPath          []string            `toml:"watchPath"`
	Port               int                 `toml:"port"`
	Role               string              `toml:"role"`
	KeyRing            bool                `toml:"use-keyring"`
	JcrRoot            string              `toml:"jcrRoot"`
	JVMOpts            []string            `toml:"jvmOptions"`
	AdditionalPackages []string            `toml:"additionalPackages"`
	ContentPackages    []string            `toml:"contentPackages"`
}

type aemInstanceConfig struct {
	Name            string   `toml:"name"`
	Group           string   `toml:"group"`
	Debug           bool     `toml:"debug"`
	Protocol        string   `toml:"proto"`
	Hostname        string   `toml:"hostname"`
	Port            int      `toml:"port"`
	Type            string   `toml:"type"`
	Username        string   `toml:"username"`
	Password        string   `toml:"password"`
	JVMOptions      []string `toml:"jvm-options"`
	JVMDebugOptions []string `toml:"jvm-debug-options"`
}

func (i *aemInstanceConfig) URL() string {
	return fmt.Sprintf("%s://%s:%d", i.Protocol, i.Hostname, i.Port)
}

func (i *aemInstanceConfig) PasswordURL() string {
	instance := new(instance)
	return fmt.Sprintf("%s://%s:%s@%s:%d", i.Protocol, i.Username, instance.getPasswordForInstance(*i), i.Hostname, i.Port)
}
