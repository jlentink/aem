package main

// General
const CONFIG_FILENAME = ".aem"
const CONFIG_AEM_JAR_NAME = "AEM.jar"
const CONFIG_AEM_RUN_DIR = "crx-quickstart"
const CONFIG_AEM_INSTALL_DIR = "install"
const CONFIG_INSTANCE_DIR = "instance"
const CONFIG_PACKAGES_DIR = "packages"
const CONFIG_APP_DIR = "app"
const CONFIG_INSTANCE_GIT_IGNORE = ".gitignore"
const CONFIG_INSTANCE_GIT_IGNORE_CONTENT = "# Ignore everything in this directory\n*\n# Except this file\n!.gitignore"
const CONFIG_AEM_LOG = "logs/error.log"
const CONFIG_AEM_PID = "conf/cq.pid"
const CONFIG_DEFAULT_INSTANCE = "local-author"

const REGEX_ZIP = "\\.zip$"
const REGEX_PACKAGE_VERSIONED = "^(.*)-([0-9\\.]*)\\.zip$"
const REGEX_PACKAGE_VERSIONED_SNAPSHOT = "^(.*)-([0-9\\.]*)-SNAPSHOT(.*)\\.zip$"
const REGEX_PACKAGE_UNVERSIONED = "^(.*)\\.zip$"

// Template config file
const CONFIG_TEMPLATE_URL = "http://127."+"0.0.1:4502"
const CONFIG_TEMPLATE_USERNAME = "admin"
const CONFIG_TEMPLATE_PASSWORD = "admin"
const CONFIG_HTTP_NO_CACHE = "no-cache"

const CONFIG_TEMPLATE_LOCAL_PORT = 4502
const CONFIG_TEMPLATE_LOCAL_ROLE = "author"
const CONFIG_TEMPLATE_LOCAL_JAR = "http://doain.tld/some.jar"
const CONFIG_TEMPLATE_LOCAL_NAME = "local"
const CONFIG_TEMPLATE_LOCAL_GROUP = "local"

const CONFIG_AEM_CLI_SERVICE = "aem-cli"
