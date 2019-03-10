# AEMcli (command line interface)

[![](https://travis-ci.org/jlentink/aem.svg?branch=master)](https://travis-ci.org/jlentink/aem)
[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=jlentink_aem&metric=alert_status)](https://sonarcloud.io/dashboard?id=jlentink_aem)
[![Go Report Card](https://goreportcard.com/badge/github.com/jlentink/aem)](https://goreportcard.com/report/github.com/jlentink/aem)
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)

**This tool is work in progress**<br />
*If you find any bugs or miss any feature feel free to pitch in or create a ticket so the issue can be resolved quickly or the new feature can be added.*


When using AEM in projects there are a couple of things that happen quite often. This tool is like a swiss army knife that tries to help you on everyday tasks for developer and dev-ops.
So you don't need to search for all the common cURL commands they are included.

Let me know what you think. happy AEM-ing.

## Getting Started
Build the project from scratch (needs go >= v1.11.4) or download the prebuild binary available for your operating system.

Prebuild versions are:

* osx (64bit) 
* linux (64bit)
* windows (64bit)

### Installing
Download or build and copy the binary to a location in your path.
Latest prebuild version can be found [here](https://github.com/jlentink/aem/releases/latest).

Example install locations.

### OSX & Linux
execute `echo ${PATH}` and validate that */usr/local/bin* is in your current path. If in path use `cp aem /usr/local/bin` to place the executable ont he correct spot. If not in path add the following line `export PATH="${PATH}:/usr/local/bin"` to your ~/.bash_profile, ~/.profile or ~/.zprofile with your favorite editor.

Dont forget to set the executable permision. `chmod a+x aem`



### Windows
Place the executable in for example `"C:\Program files\aem"` and follow the [tutorials](https://www.google.com/search?q=windows+change+path) on the internet how to add them to add this directory to your path.


## Usage

The command line tool is broken up in different sub-commands. The commands can be used by typing `aem <command>` eg. `aem start` All the possible commands are listed below. Every command has the option to request help on the specifications of that commands. eg. `aem start -h`

### help
Use help to get a list + description of possible sub-commands.

Available options:
   
    No options available

### init
Creates a config file. The config file allows you to define the instances used during the project. (E.g. local author, local dev etc...)

Available options:

     -d, --dump              Write default config file without setup questions
     -f, --force-overwrite   Overwrite current configuration
     -v, --verbose           Enable verbose



In the config, you can also define the packages you want to install at boot as for the location of the AEM jar to download for the project.

### start
when using the start command is used it will download all files needed that are defined in the config file and automatically installed. Files removed from the configuration will be removed from disk so that AEM will automatically deinstall them on the next start.

By default, the start commands also checks that you are not starting the aem server as the root user.

start is compatible with the start and stop scripts provided by Adobe.

Available options:

     -d, --download    Force new download
     -f, --for-ground  Don't detach aem from current tty.
     -n, --name=value  Instance to start. (default: local-author)
     -r, --root        Allow root
     -v, --verbose     Enable verbose


### stop
Stop AEM instances running. 

stop is compatible with the start and stop scripts provided by Adobe.

Available options:

    -n, --name=value  Instance to stop. (default: local-author)
    -v, --verbose     Enable verbose
    
### sync
sync triggers [aemsync](https://github.com/gavoja/aemsync). [aemsync](https://github.com/gavoja/aemsync) needs to be installed on your system to make this work. In the .aem file, you can define which directories to watch. and sync. When using `instance-name` one instance will be used to sync the content to. When using `instance-group` all instances that are defined in the group will be synced Eg. `aem sync -g local` can sync to the local author and publish.

Available options:

    -g, --instance-group=value  Instance group to sync to.
    -i, --instance-name=value   Instance to sync to. (default: local-author)
    -l, --disable-log           Disable AEM log output
    -s, --aemsync=value         Path to AEM sync executable
    -v, --verbose               Enable verbose

### pull-content
Download the content packages defined in the configuration file and upload them to an instance of your choosing. Handy to sync content to developer instances during the project.

Available options:

    -d, --force-download  Force new download
    -f, --from-name=value Pull content from
    -t, --to-name=value   Push content to (default: local-author)
    -v, --verbose         Enable verbose
    
### passwords
You don't want to store passwords in a git repository for secure development. Although the tool allows you to define passwords in the configuration file there is an option to safely store the passwords in the key-ring (password managers eg. OSX key-chain) of the operating system. Use the passwords command to populate or update the stored passwords.

Available options:

    -a, --all          Update all
    -g, --group=value  Instance group to update.
    -n, --name=value   Instance to update.
    -v, --verbose      Enable verbose
    -y, --yes          Confirm all questions with yes

### system-information or sysinfo
Display information about an instance. This feature is only available from AEM 6.4 or newer.

Available options:

    -n, --name=value  Instance to start. (default: local-author)
    -v, --verbose     Enable verbose

### package-list
List the packages installed on an instance of your choosing.

Available options:

    -n, --name=value  Instance to start. (default: local-author)
    -v, --verbose     Enable verbose

### package-rebuild
Rebuild a package on an instance of your choosing.

Available options:

    -f, --from-name=value  Rebuild package on instance (Default: local-author)
    -p, --package=value    Define package package:version (no interactive mode)
    -v, --verbose          Enable verbose

### package-download
Download a package from any instance defined in the configuration file

Available options:

     -d, --force-download  Force new download
     -f, --from=value      Pull content from (default: local-author)
     -p, --package=value   Define package package:version (no interactive mode)
     -v, --verbose         Enable verbose

### package-copy
Copy a package from one instance to another. The destination can be a group to easily install to all members of a group or a single target.

Available options:

      -d, --force-download  Force new download
      -f, --from-name=value Pull package from (default: local-author)
      -g, --to-group=value  Push package to group
      -l, --log             Show AEM log output
      -p, --package=value   Packages (multiple use comma separated list.)
      -t, --to-name=value   Push package to instance
      -v, --verbose         Enable verbose
      
### package-install
Install a package you have locally to one instance or to a complete group.
The name of the package will be extracted from the manifest in the package

Available options:

    -g, --to-group=value  Push package to group
    -n, --no-install      Do not install package
    -p, --package=value   Package to install (path to file)
    -t, --to-name=value   Push package to instance
    -v, --verbose         Enable verbose
    -y, --yes             Skip confirmation
    
### bundle-list
List all bundles on an instance.

Available options:

    -n, --name=value  Name of instance to list bundles from from (default: local-author)
    -v, --verbose     Enable verbose

### bundle-start
Start a bundle based by its symbolic name

Available options:

    -b, --bundle=value  Bundle to start (Symbolic name)
    -n, --name=value    Name of instance to list bundles from from (default: local-author)
    -v, --verbose       Enable verbose


### bundle-stop
Stop a bundle based on it's symbolic name

Available options:

    -b, --bundle=value  Bundle to stop (Symbolic name)
    -n, --name=value    Name of instance to list bundles from from (default: local-author)
    -v, --verbose       Enable verbose


### bundle-install
Install a bundle based on it's symbolic name

Available options:

    -b, --bundle=value  Path to bundle (.jar)
    -n, --name=value    Name of instance to list bundles from from (default: local-author)
    -s --startlevel     Bundle start level (default: 20)
    -v, --verbose       Enable verbose


### log
See the log file for an instance running locally. Use -f to follow the log file for more log information coming in. use CTRL+c to stop following the log file.

Available options:

    -f, --follow      Follow log file. Show new lines if they come in.
    -n, --name=value  Instance to start. (default: local-author)
    -v, --verbose     Enable verbose

### page-replicate
Activate or deactivate a page. use the page path to define which page to activate.

Available options:

    -a, --activate     Activate
    -d, --deactivate   Deactivate
    -g, --group=value  Instances to target based on group
    -n, --name=value   Instance to target based on name
    -p, --page=value   Page to activate
    -v, --verbose      Enable verbose

### activate-tree 
Activate a piece of the tree. use the path to define which part.

Available options:

    -i, --instance=value  Activate Tree on instance (Default: local-author)
    -p, --path=value      Path to activate
    -v, --verbose         Enable verbose

### open
Open a browser to the instance of your choosing.

Available options:

    -n, --name=value  Instance to open. (default: local-author)
    -v, --verbose     Enable verbose


### version
Output the current version of the aem command line interface you are using.

Available options:
   
    No options available

## Todo

- Code cleanup
- more code testing and coverage
- Features
  - pull/copy content over vlt
  - thread dumps
  - sling tracer


## Built With

* [Getopt](https://github.com/pborman/getopt/tree/master/v2) - For command line parsing
* [Go-pretty](github.com/jedib0t/go-pretty/table) - For table printing
* [Logrus](https://github.com/sirupsen/logrus) - For logging
* [Afero](https://github.com/spf13/afero) - For FileSystem Abstraction
* [Progressbar](https://github.com/schollz/progressbar) - For progress bar printing
* [Go-humanize](https://github.com/dustin/go-humanize) - Formatters for units to human friendly sizes
* [Go-keyring](https://github.com/zalando/go-keyring) - Store password in operating systems own keyring
* [TOML](https://github.com/BurntSushi/toml) - TOML parser for Golang with reflection.
* [Tail](https://github.com/hpcloud/tail) - For tailing files

Thank all authors and contributors of these libraries. For publishing such great software.


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

