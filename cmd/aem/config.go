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
	AemJar             string              `toml:"jar"`
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

const (
	configTemplate = `
#
# Example .aem config file.
# Update based on your need
#
verbose = true

#
# When set to true passwords will be stored in your local OS
# keyring. When disabled the passwords from this config file will be used.
#
use-keyring = true

#
# AEM jar. The AEM jar to start AEM.
#
# Can be an user with username and password or without.
# you can also define a path to a local jar on disk.
#
# Urls can have username and password added for basic authentication
# Example: https://username@password:somedomain.tld/aem-author-4502.jar
#
# Beware to not store the AEM jar in a public path!
#
jar = "http://username:password@localhost/aem-author-4502.jar"

#
# paths to watch and sync during development.
# aemsync needs to be installed to use this function
# npm install aemsync -g
#
watchPath = [
    "ui.apps/src/main/content/jcr_root",
    "ui.content/src/main/content/jcr_root"
]

#
# Content packages to use for this project
#
contentPackages = [
	"content:1.0.0",
	"assets:1.0.0",
]


#
# Additional packages to auto install.
# Example given is ACS commons to auto install
# Urls can have username and password added for basic authentication
# Example: https://username@password:somedomain.tld/....
#
additionalPackages = [
    "https://github.com/Adobe-Consulting-Services/acs-aem-commons/releases/download/acs-aem-commons-3.19.0/acs-aem-commons-content-3.19.0.zip"
]

#
# General JVM options to use when starting AEM
#
jvm-options = [
    "-server",
    "-Xmx4048m",
    "-XX:MaxPermSize=512M",
    "-Djava.awt.headless=true",
    "-Dcom.sun.management.jmxremote.port=3233",
    "-Dcom.sun.management.jmxremote.ssl=false",
    "-Dcom.sun.management.jmxremote.authenticate=false"
]

#
# General JVM debug options to use when starting AEM
#
jvm-debug-options = [
    "-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005",
]

#
# Instances for your project.
# Make sure you have local-author and local-publisher available
# and you can define as many as you want name needs to be unique.
#
[[instance]]
name = "local-author"
group = "local"
debug = true
proto = "http"
hostname = "127.0.0.1"
port = 4502
username = "admin"
password = ""
type = "author"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "local-publish"
group = "local"
debug = false
proto = "http"
hostname = "127.0.0.1"
port = 4503
username = "admin"
password = ""
type = "publish"
jvm-options = []
jvm-debug-options = []


#
# Not used yet
#
packages = [
]
`
)
