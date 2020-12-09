package aem

// Template variables
var (
	KeyChain           = true
	JarLocation        = ``
	JarUsername        = ``
	JarPassword        = ``
	LicenseCustomer    = ``
	LicenseVersion     = ``
	LicenseDownloadID  = ``
	DefaultInstanceStr = "local-author"
	AdditionalPackages []string
)

const (
	configTemplate = `
#
# Example .aem config file.
# Update based on your need
#
verbose = false

#
# Project name
#
project-name = ""

#
# When set to true passwords will be stored in your local OS
# keyring. When disabled the passwords from this config file will be used.
#
use-keyring = true

#
# What is the version used for this artifact build
#
version = "1.0.0"

#
# What to append as a suffix to the build version.
#
# Options:
#
# - GIT_LONG - will be replaced git commit hash long
# - GIT_SHORT - will be replaced with git commit hash short
# - DATE - ill be replaced with a time stamp
# - Any other string - will used exactly
#
version-suffix=".DATE"

#
# What parameters to send to maven for this build
#
buildCommand = "clean install -P adobe-public"

#
# Default AEM version to use for this project.
# This version will be used when not defined with the instance
#
default-version = "6.5.0"

#
# Validate SSL on server when https is used.
#
ssl-validate = false

#
# Project license information.
# licenses are sensitive material and should not be shared outside of the project
# or miss used. So be very careful when using this functonality!!!!
#
licenseCustomer = ""
licenseVersion = ""
licenseDownloadID = ""

#
# Default instance to use if not providing the detail in the command
#
defaultInstance = "local-author"
#
# paths to watch and sync during development.
# aemsync needs to be installed to use this function
# npm install aemsync -g
#
watchPath = [
]

#
# Content packages to use for this project
#
contentPackages = [
]

#
# Content page paths
#
contentBackupPaths = [
]

#
# Content page package name
#
contentBackupName = "content-download"

#
# Content page package group
#
contentBackupGroup = ""

#
# Exclude packages from deployment
#
packageExclude = [
  "\\.apps$",
  "\\.content$",
]

#
# Paths to copy when using vlt-sync
# Prepend path with ! to prevent recursive copy.
#
# e.g:
#   None recursive: "!/content/some-path/"
#   Recursive:      "/content/some-other-path/"
#
vltSyncPaths = [
]

#
# Invalidate path's
# which paths should be send to a dispatcher to invalidate.
#
invalidatePaths = [
]

#
# Additional packages to auto install.
# Example given is ACS commons to auto install
# Urls can have username and password added for basic authentication
# Example: https://username@password:somedomain.tld/....
#
additionalPackages = [
]

#
# Which extra arguments to use for oak-run
#
oakOptions = [
    "-mx8G",
    "-Dtar.memoryMapped=true",
]
oakDefaultVersion = "1.8.12"

#
# General JVM options to use when starting AEM
#
jvm-options = [
    "-server",
    "-Xmx4048m",
    "-XX:MaxPermSize=512M",
]

#
# General JVM debug options to use when starting AEM
#
jvm-debug-options = [
    "-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005",
]

#
# AEM jar. The AEM jar to start AEM.
#
# Can be an user with username and password or without.
# you can also define a path to a local jar on disk.
#
# Urls where to find the AEM jar
# Example: https://somedomain.tld/aem-author-4502.jar
#
# The jarUsername and jarPassword are the
# HTTP basic credentials that can be used.
#
# Beware to not store the AEM jar in a public unprotected path!
#
[[aemJar]]
version = "6.5.0"
location = "https://someurl/AEM-6.5.jar"
username = ""
password = ""

#
# Instances for your project.
# Make sure you have local-author and local-publisher available
# and you can define as many as you want name needs to be unique.
#
#
# Definitions:
#   name: instance name
#   group: group name eg: local, dev, test, stage, prod
#   aliases: Array of aliasses you want to use (not mandatory)
#   debug:  true, false to enable the debug parameters
#   proto: http, https
#   hostname: hostname to use
#   ip: ip address to use
#   port: port to open
#   username: for login
#   password: for login if not using keychain
#   type: author, publish, dispatch
#   runmode: Runmodes to be appended when starting
#   jvm-options: jvm extra options
#   jvm-debug-options: jvm debug extra options
#   secure-port
#   author
#   publisher
#   dispatcher-version
#   dispatcher-version
#
[[instance]]
name = "local-author"
group = "local"
aliases = ["example-1", "example-2"]
debug = true
proto = "http"
hostname = "127.0.0.1"
port = 4502
username = "admin"
password = ""
type = "author"
runmode = "crx3,crx3tar,dev"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "local-publish"
group = "local"
aliases = []
debug = false
proto = "http"
hostname = "127.0.0.1"
port = 4503
username = "admin"
password = ""
type = "publish"
runmode = "crx3,crx3tar,dev"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "local-dispatcher"
group = "local"
proto = "http"
hostname = "127.0.0.1"
port = 8888
secure-port = 8443
author = "local-author"
publisher = "local-publish"
dispatcher-endpoint=""
dispatcher-version = "4.3.2"
username = ""
password = ""
type = "dispatch"
`
)
