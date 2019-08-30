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
project-name = "my-project"


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
# - Any other string - will used exactly
#
version-suffix="-GIT_SHORT"

#
# What parameters to send to maven for this build
#
buildCommand = "clean install -P adobe-public"

#
# Default AEM version to use for this project.
# This version will be used when not defined with the instance
#
default-version = "6.4.0"


#
# Validate SSL on server when https is used.
#
ssl-validate = true

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
    "ui.apps/src/main/content/jcr_root",
    "ui.content/src/main/content/jcr_root"
]

#
# Content packages to use for this project
#
contentPackages = [
	"content",
	"assets:1.0.0",
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
    "/content",
    "/etc.clientlibs",
]

#
# Additional packages to auto install.
# Example given is ACS commons to auto install
# Urls can have username and password added for basic authentication
# Example: https://username@password:somedomain.tld/....
#
additionalPackages = [
    "http://aem:aem@aem.test/aem/acs-aem-commons-content-3.19.0.zip",
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
version = "6.4.0"
location = "http://url.tld/AEM-6.4.jar"
username = "aem"
password = "aem"

[[aemJar]]
version = "6.5.0"
location = "http://url.tld/aem/AEM-6.5.jar"
username = "aem"
password = "aem"

#
# Instances for your project.
# Make sure you have local-author and local-publisher available
# and you can define as many as you want name needs to be unique.
#
#
# Definitions:
# type: author, publish, dispatch
# proto: http, https
#
[[instance]]
name = "local-author"
group = "local"
debug = true
proto = "http"
hostname = "127.0.0.1"
port = 4502
username = "admin"
password = "admin"
type = "author"
runmode = "crx3,crx3tar"
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
password = "admin"
type = "publish"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "dev-author"
group = "dev"
debug = true
proto = "https"
hostname = "author.dev"
port = 4502
username = "admin"
password = "admin"
type = "author"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "dev-publish"
group = "dev"
debug = false
proto = "http"
hostname = "publish.dev"
port = 4503
username = "admin"
password = "admin"
type = "publish"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "dev-dispatcher"
group = "dev"
debug = false
proto = "http"
hostname = "dispatcher.dev"
port = 80
username = ""
password = ""
type = "dispatch"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "stage-author"
group = "stage"
debug = true
proto = "https"
hostname = "author.stage"
port = 4502
username = "admin"
password = "admin"
type = "author"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "dev-stage"
group = "stage"
debug = false
proto = "http"
hostname = "publish.stage"
port = 4503
username = "admin"
password = "admin"
type = "publish"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "stage-dispatcher"
group = "dev"
debug = false
proto = "http"
hostname = "dispatcher.stage"
port = 80
username = ""
password = ""
type = "dispatch"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "prod-author"
group = "dev"
debug = true
proto = "https"
hostname = "author.prod"
port = 4502
username = "admin"
password = "admin"
type = "author"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "prod-publish"
group = "dev"
debug = false
proto = "http"
hostname = "publish.prod"
port = 4503
username = "admin"
password = "admin"
type = "publish"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []

[[instance]]
name = "prod-dispatcher"
group = "prod"
debug = false
proto = "http"
hostname = "dispatcher.prod"
port = 80
username = ""
password = ""
type = "dispatch"
runmode = "crx3,crx3tar"
jvm-options = []
jvm-debug-options = []
`
)
