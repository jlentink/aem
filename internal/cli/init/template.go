package init

const (
	configTemplate = `
#
# Example .aem config file.
# Update based on your need
#
verbose = false
#
# When set to true passwords will be stored in your local OS
# keyring. When disabled the passwords from this config file will be used.
#
use-keyring = %t
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
jarLocation = "%s"
jarUsername = "%s"
jarPassword = "%s"
#
# Project license information.
# licenses are sensitive material and should not be shared outside of the project
# or miss used. So be very careful when using this functonality!!!!
#
licenseCustomer = "%s"
licenseVersion = "%s"
licenseDownloadID = "%s"
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
	"content:1.0.0",
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
# Additional packages to auto install.
# Example given is ACS commons to auto install
# Urls can have username and password added for basic authentication
# Example: https://username@password:somedomain.tld/....
#
additionalPackages = [
%s
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
