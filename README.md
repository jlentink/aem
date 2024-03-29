# aemCLI (command line interface)

[![](https://travis-ci.org/jlentink/aem.svg?branch=master)](https://travis-ci.org/jlentink/aem)
[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=jlentink_aem&metric=alert_status)](https://sonarcloud.io/dashboard?id=jlentink_aem)
[![Go Report Card](https://goreportcard.com/badge/github.com/jlentink/aem)](https://goreportcard.com/report/github.com/jlentink/aem)
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)


**This tool is work in progress**<br />
*If you find any bugs or miss any feature feel free to pitch in or create a ticket so the issue can be resolved quickly or the new feature can be added.*


When using AEM in projects there are a couple of things that happen quite often. This tool is like a swiss army knife that tries to help you on everyday tasks for developer and dev-ops.

Use cases:

* Stop searching for the cURL commands use them through this tool. It will help you to do the tasks quicker.
* Need to work on multiple projects create one configuration file for that project and pull in the requirements by typing aem start.
* Update dependencies every time you start aem based on the configuration file.
* Copy packages from one instance to the other with absolute ease.



Let me know what you think. happy AEM-ing.

## Getting Started
Build the project from scratch (needs go >= v1.14.1) or download the prebuild binary available for your operating system.

Prebuild versions are:

* OSX (Intel (64 bit), Apple Silicone) 
* Linux (64bit)
* Windows (64bit)

### Installing
Download or build and copy the binary to a location in your path.
Latest prebuild version can be found [here](https://github.com/jlentink/aem/releases/latest).

Or use the following OSX/Linux oneliner to install:

#### using home brew (OSX only)

Check out the amazing [Homebrew](https://brew.sh/) and when you have homebrew installed use the following commands to install


	brew tap jlentink/aem
	brew install aem

### Other methods

##### install to /usr/local/bin
`curl -sfL https://raw.githubusercontent.com/jlentink/aem/master/install.sh | bash -s -- -b /usr/local/bin`

##### install to ~/bin
`curl -sfL https://raw.githubusercontent.com/jlentink/aem/master/install.sh | bash`

##### install to other location
`curl -sfL https://raw.githubusercontent.com/jlentink/aem/master/install.sh | bash -s -- -b YOUR_LOCATION`


Example install locations.

### OSX & Linux
execute `echo ${PATH}` and validate that */usr/local/bin* is in your current path. If in path use `cp aem /usr/local/bin` to place the executable ont he correct spot. If not in path add the following line `export PATH="${PATH}:/usr/local/bin"` to your ~/.bash_profile, ~/.profile or ~/.zprofile with your favorite editor.

Don't forget to set the executable permission. `chmod a+x aem`



### Windows
Place the executable in for example `"C:\Program files\aem"` and follow the [tutorials](https://www.google.com/search?q=windows+change+path) on the internet how to add them to add this directory to your path.


## Usage

The command line tool is broken up in different sub-commands. The commands can be used by typing `aem <command>` eg. `aem start` All the possible commands are listed below. Every command has the option to request help on the specifications of that commands. eg. `aem start -h`


	Usage:
	  aem [command]
	
	Available Commands:
	  activation         Activation commands
	  build              Build application
	  bundle             Bundle commands
	  deploy             Deploy to server(s)
	  generate           Generate code block
	  help               Help about any command
	  indexes            index commands
	  init               Init new project
	  invalidate         Invalidate path's on dispatcher
	  log                List error log or application log
	  oak                Oak commands
	  open               Open URL for Adobe Experience Manager instance in browser
	  package            Package commands
	  passwords          Set passwords into your keychain
	  projects           List know projects
	  pull-content       Pull content in from instance via packages
	  setup-check        Check if all needed binaries are available for all functionality
	  shell              Shell completion commands
	  start              Start Adobe Experience Manager instance
	  stop               stop Adobe Experience Manager instance
	  system-information Get system information from Adobe Experience Manager instance
	  version            Show version of aemcli
	
	Flags:
	  -h, --help             help for aem
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### activation
Activation commands for (de)activation a page or a tree.

	Activation commands
	
	Usage:
	  aem activation [flags]
	  aem activation [command]
	
	Available Commands:
	  page        Activate / Deactivate page
	  tree        Activate Tree
	
	Flags:
	  -h, --help   help for activation
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output
	  
### activation invalidate 
Invalidate path's on dispatcher

Usage:
  aem activation invalidate [flags]

Aliases:
  invalidate, flush

Flags:
  -g, --group string   Instance group to sent invalidate to
  -h, --help           help for invalidate
  -n, --name string    Instance to sent invalidate to
  -p, --path string    Package to rebuild

Global Flags:
  -P, --project string   Run command for project. (if not current working directory)
  -v, --verbose          verbose output

### activation page 
Activate / Deactivate page on instance

	Usage:
	  aem activation page [flags]
	
	Flags:
	  -a, --activate       Activate page
	  -d, --deactivate     Deactivate
	  -g, --group string   Instance group to (de)activate page on
	  -h, --help           help for page
	  -n, --name string    Instance to (de)activate page on (default "local-author")
	  -p, --path string    Path to (de)activate
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### activation tree 
Activate tree

	Usage:
	  aem activation tree [flags]
	
	Flags:
	  -g, --group string         Instance group to (de)activate page on
	  -h, --help                 help for tree
	  -d, --ignore-deactivated   Ignore Deactivated
	  -n, --name string          Instance to (de)activate page on (default "local-author")
	  -o, --only-modified        Only Modified
	  -p, --path string          Path to (de)activate
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output


### build
Build application using the maven pom file

	Usage:
	  aem build [flags]
	
	Flags:
	  -h, --help   help for build
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output
 
### bundle
All commands available for bundles

	Usage:
	  aem bundle [flags]
	  aem bundle [command]
	
	Available Commands:
	  install     Install bundle
	  list        List bundles
	  start       Start bundle
	  stop        Stop bundle
	
	Flags:
	  -h, --help   help for bundle
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### bundle install
Install bundle on instance

	Usage:
	  aem bundle install [flags]
	
	Flags:
	  -b, --bundle string   Instance group to install bundle on
	  -g, --group string    Instance group to install bundle on
	  -h, --help            help for install
	  -l, --level string    Bundle start level (default "20")
	  -n, --name string     Instance to install bundle on (default "local-author")
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output
	
### bundle list
List bundles on instance

	Usage:
	  aem bundle list [flags]
	
	Flags:
	  -h, --help          help for list
	  -n, --name string   Instance to list bundles off (default "local-author")
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output


### bundle start

Start bundle on an aem instance

	Usage:
	  aem bundle start [flags]
	
	Flags:
	  -b, --bundle string   Instance group to install bundle on
	  -g, --group string    Instance group to install bundle on
	  -h, --help            help for start
	  -n, --name string     Instance to install bundle on (default "local-author")
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### bundle stop
Stop bundle on an aem instance

	Usage:
	  aem bundle stop [flags]
	
	Flags:
	  -b, --bundle string   Instance group to install bundle on
	  -g, --group string    Instance group to install bundle on
	  -h, --help            help for stop
	  -n, --name string     Instance to install bundle on (default "local-author")
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### deploy
Deploy (or deploy and build) to server(s)

	Usage:
	  aem deploy [flags]
	
	Flags:
	  -a, --artifact string   Deploy one a single artifact
	  -b, --build             Build before deploy
	  -g, --group string      Group to deploy to
	  -h, --help              help for deploy
	  -n, --name string       Instance to deploy to (default "local-author")
	  -p, --password string   Overwrite password to use if not using the one from config file
	  -u, --username string   Overwrite username to use if not using the one from config file
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### generate
Generate code block based on templates. Every template or filename can have a variable. Variables are defined as {{.variablename}}. Altering the case typing can be done via following functions:

* CC - camelCase
* UC - UPPER CASE
* LC - lower case
* PC - PascalCase
* SC - snake_case
* KC - kabab-case
* TC - Title case

More documentation on templates can be found [here](https://golang.org/pkg/text/template/).

		
		Usage:
		  aem generate [flags]
		  aem generate [command]
			
		Available Commands:
		  init        Write default templates to disk
			
		Flags:
		  -h, --help   help for generate
			
		Global Flags:
		  -P, --project string   Run command for project. (if not current working directory)
		  -v, --verbose          verbose output

### generate init
Write default templates to disk for usage and customization

	Usage:
	  aem generate init [flags]
	
	Flags:
	  -h, --help   help for init
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### indexes
Index commands

	Usage:
	  aem indexes [flags]
	  aem indexes [command]
	
	Available Commands:
	  indexes     Show indexes on instance
	  list        Reindex index on instance
	
	Flags:
	  -h, --help   help for indexes
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### indexes indexes
Show indexes on instance

Usage:
  aem indexes indexes [flags]

Flags:
  -h, --help          help for indexes
  -n, --name string   Instance to stop (default "local-author")

Global Flags:
  -P, --project string   Run command for project. (if not current working directory)
  -v, --verbose          verbose output

### indexes reindex
Reindex index on instance

	Usage:
	  aem indexes list [flags]
	
	Aliases:
	  list, reindex
	
	Flags:
	  -h, --help           help for list
	  -i, --index string   Index to reindex
	  -n, --name string    Instance to stop (default "local-author")
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### init
Creates a config file. The config file allows you to define the instances used during the project. (E.g. local author, local dev etc...)
	
	Usage:
	  aem init [flags]
	
	Flags:
	  -f, --force   Force override of current configuration
	  -h, --help    help for init
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

In the config, you can also define the packages you want to install at boot as for the location of the AEM jar to download for the project.

### log
See the log file for an instance running locally. Use -f to follow the log file for more log information coming in. use CTRL+c to stop following the log file.

	Usage:
	  aem log [flags]
	
	Aliases:
	  log, logs
	
	Flags:
	  -f, --follow        Actively follow lines when they come in
	  -h, --help          help for log
	      --list          List available log files
	  -l, --log string    Which file(s) to follow (default "error.log")
	  -n, --name string   Instance to stop (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
     -P, --project string   Run command for project. (if not current working directory)

### oak
Oak commands

	Usage:
	  aem oak [flags]
	  aem oak [command]
	
	Available Commands:
	  check       Run oak check
	  checkpoints Run oak checkpoints
	  compact     Run oak compact
	  console     Run oak console
	  explorer    Run oak explorer
	
	Flags:
	  -h, --help   help for oak
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output


### oak check 
Run oak check

	Usage:
	  aem oak check [flags]
	
	Flags:
	  -a, --aem string    Version of AEM to use oak-run on. (use matching AEM version of oak-run)
	  -h, --help          help for check
	  -n, --name string   Instance to stop (default "local-author")
	  -o, --oak string    Define version of oak-run to use
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### oak checkpoints
Run oak checkpoints

	Usage:
	  aem oak checkpoints [flags]
	
	Flags:
	  -a, --aem string    Version of AEM to use oak-run on. (use matching AEM version of oak-run)
	  -h, --help          help for checkpoints
	  -n, --name string   Instance to stop (default "local-author")
	  -o, --oak string    Define version of oak-run to use
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### oak compact
Run oak compact
Run oak-run check on instance. Check the FileStore for inconsistencies. More information see [Oak-run](https://github.com/apache/jackrabbit-oak/tree/trunk/oak-run).

Use ```--aem``` to define which AEM version you are running this against and aem cli will set the corresponding oak-run version. 
When you wan't to define a specific version use the ```--oak```. 
The oak jar will be placed in the bin folder under instance and downloaded if it not exists yet.


	Usage:
	  aem oak compact [flags]
	
	Flags:
	  -a, --aem string    Version of AEM to use oak-run on. (use matching AEM version of oak-run)
	  -h, --help          help for compact
	  -n, --name string   Instance to stop (default "local-author")
	  -o, --oak string    Define version of oak-run to use
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### oak console
Run oak console
Run oak-run check on instance. Check the FileStore for inconsistencies. More information see [Oak-run](https://github.com/apache/jackrabbit-oak/tree/trunk/oak-run).

Use ```--aem``` to define which AEM version you are running this against and aem cli will set the corresponding oak-run version. 
When you wan't to define a specific version use the ```--oak```. 
The oak jar will be placed in the bin folder under instance and downloaded if it not exists yet.

	Usage:
	  aem oak console [flags]
	
	Flags:
	  -a, --aem string    Version of AEM to use oak-run on. (use matching AEM version of oak-run)
	  -h, --help          help for console
	  -m, --metrics       Enables metrics based statistics collection
	  -n, --name string   Instance to stop (default "local-author")
	  -o, --oak string    Define version of oak-run to use
	  -w, --read-write    Connect to repository in read-write mode
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### oak explorer
Run oak explorer
Run oak-run check on instance. Check the FileStore for inconsistencies. More information see [Oak-run](https://github.com/apache/jackrabbit-oak/tree/trunk/oak-run).

Use ```--aem``` to define which AEM version you are running this against and aem cli will set the corresponding oak-run version. 
When you wan't to define a specific version use the ```--oak```. 
The oak jar will be placed in the bin folder under instance and downloaded if it not exists yet.


	Usage:
	  aem oak explorer [flags]
	
	Flags:
	  -a, --aem string    Version of AEM to use oak-run on. (use matching AEM version of oak-run)
	  -h, --help          help for explorer
	  -n, --name string   Instance to stop (default "local-author")
	  -o, --oak string    Define version of oak-run to use
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output


### open
Open a browser to the instance of your choosing.

	Usage:
	  aem open [instance name] [flags]
	
	Flags:
	  -h, --help          help for open
	  -n, --name string   Instance to stop (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

### package
Package commands

	Usage:
	  aem package [flags]
	  aem package [command]
	
	Available Commands:
	  copy        Copy packages from one instance to another
	  download    Download packages
	  install     Install uploaded package
	  list        List packages
	  rebuild     package rebuild
	  upload      Upload package to aem
	
	Flags:
	  -h, --help   help for package
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output
### package list
List the packages installed on an instance of your choosing.

	Usage:
	  aem package list [flags]
	
	Aliases:
	  package-list, plist
	
	Flags:
	  -h, --help          help for package-list
	  -n, --name string   Instance to stop (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

### package rebuild
Rebuild a package on an instance of your choosing.

	Usage:
	  aem package rebuild [flags]
	
	Flags:
	  -h, --help             help for package-rebuild
	  -n, --name string      Instance to rebuild package on (default "local-author")
	  -p, --package string   Package to rebuild
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

### package download
Download a package from any instance defined in the configuration file

	Usage:
	  aem package download [flags]
	
	Flags:
	  -h, --help             help for package-download
	  -n, --name string      Instance to stop (default "local-author")
	  -p, --package string   Package name. E.g: name, name:1.0.0
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

### package copy
Copy a package from one instance to another. The destination can be a group to easily install to all members of a group or a single target.

	Usage:
	  aem package copy [flags]
	
	Flags:
	  -f, --from string      Instance to copy from
	  -h, --help             help for package-copy
	  -p, --package string   Package to copy
	  -t, --to string        Destination Instance
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)      
	  
### package install
Install a package you have locally to one instance or to a complete group.
The name of the package will be extracted from the manifest in the package

	Usage:
	  aem package install [flags]
	
	Flags:
	  -h, --help             help for package-install
	  -n, --name string      Instance to rebuild package on (default "local-author")
	  -p, --package string   Package to rebuild
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

    
### passwords
You don't want to store passwords in a git repository for secure development. Although the tool allows you to define passwords in the configuration file there is an option to safely store the passwords in the key-ring (password managers eg. OSX key-chain) of the operating system. Use the passwords command to populate or update the stored passwords.

	Usage:
	  aem passwords [flags]
	
	Aliases:
	  passwords, password, passwd
	
	Flags:
	  -h, --help          help for passwords
	  -n, --name string   Update specific instance
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)
      
### Projects
shown the known projects from previous use. These projects can be triggered from everywhere using `-P` `--project` with every command

    Usage:
      aem projects [flags]
    
    Flags:
      -h, --help   help for projects
    
    Global Flags:
      -P, --project string   Run command for project. (if not current working directory)
      -v, --verbose          verbose output

### pull-content
Download the content packages defined in the configuration file and upload them to an instance of your choosing. Handy to sync content to developer instances during the project.

	Usage:
	  aem pull-content [flags]
	
	Aliases:
	  pull-content, cpull
	
	Flags:
	  -b, --build         Build before download
	  -f, --from string   Instance to copy from
	  -h, --help          help for pull-content
	  -t, --to string     Destination Instance (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
     -P, --project string   Run command for project. (if not current working directory)
     
### setup-check
Check if all needed binaries are available for all functionality

Usage:
  aem setup-check [flags]

Aliases:
  setup-check, setup

Flags:
  -h, --help   help for setup-check

Global Flags:
  -P, --project string   Run command for project. (if not current working directory)
  -v, --verbose          verbose output
      
### shell
Shell completion commands

	Usage:
	  aem shell [flags]
	  aem shell [command]
	
	Available Commands:
	  bash        Generate bash completion for aemCLI
	  zsh         Generate zsh completion for aemCLI
	
	Flags:
	  -h, --help   help for shell
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output
      
### shell bash
Create a bash completion script for to use on your system.<br />
**E.g: aem bash > aem-completion.bash**


	Usage:
	  aem shell bash[flags]
		
	Flags:
	  -h, --help          help for bash-completion
	  -n, --name string   Instance to stop (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
     -P, --project string   Run command for project. (if not current working directory)

### shell zsh
Create a zsh completion script for to use on your system.<br />
**E.g: aem zsh > aem-completion.bash**


	Usage:
	  aem shell zsh [flags]
	
	Aliases:
	  zsh, zsh
	
	Flags:
	  -h, --help   help for zsh
	
	Global Flags:
	  -P, --project string   Run command for project. (if not current working directory)
	  -v, --verbose          verbose output

### start
when using the start command is used it will download all files needed that are defined in the config file and automatically installed. Files removed from the configuration will be removed from disk so that AEM will automatically deinstall them on the next start.

By default, the start commands also checks that you are not starting the aem server as the root user.

start is compatible with the start and stop scripts provided by Adobe.


	Usage:
	  aem start [flags]
	
	Flags:
	  -r, --allow-root    Allow to start as root user (UID: 0)
	  -d, --download      Force re-download
	  -f, --foreground    on't detach aem from current tty
	  -h, --help          help for start
	  -p, --ignore-pid    Ignore existing PID file and start AEM
	  -n, --name string   Instance to start (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)


### stop
Stop AEM instances running. 

stop is compatible with the start and stop scripts provided by Adobe.

	Usage:
	  aem stop [flags]
	
	Flags:
	  -h, --help          help for stop
	  -n, --name string   Instance to stop (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)

### system-information or sysinfo
Display information about an instance. This feature is only available from AEM 6.4 or newer.

	Usage:
	  aem system-information [flags]
	
	Aliases:
	  system-information, sys, sysinfo
	
	Flags:
	  -r, --allow-root    Allow to start as root user (UID: 0)
	  -d, --download      Force re-download
	  -f, --foreground    on't detach aem from current tty
	  -h, --help          help for system-information
	      --ignore-pid    Ignore existing PID file and start AEM
	  -n, --name string   Instance to start (default "local-author")
	
	Global Flags:
	  -v, --verbose   verbose output
      -P, --project string   Run command for project. (if not current working directory)


### version
Output the current version of the aem command line interface you are using.

	Usage:
	  aem version [flags]
	
	Flags:
	  -h, --help      help for version
	  -m, --minimal   Show the minimal version information
	
	Global Flags:
	  -v, --verbose   verbose output


## Environment variables

### AEM_ME
Set the AEM_ME variable to change the default instance choosen by start, stop and log

```
export AEM_ME=<instance name>
```  

#### Instance name resolution
There are 2 ways to influence the default instance that is selected when not specified as parameter command. The resolution order can be seen below:

1. AEM_ME Variable
2. config file "defaultInstance" variable
3. default name set in application "local-author"

## Shell scripts
### Bash Completion

Terminals are fun. Completion in the terminal is even more fun. Add `aem-completion.bash` to you completion folder.
or execute the following commands
 

    mkdir ~/.bash-completion
    aem shell bash > ~/.bash-completion/aem-completion.bash
    
    echo "source ~/.bash-completion/aem-completion.bash" >> ~/.bashrc

### Zsh
The generated completion script should be put somewhere in your $fpath named _aem.
    Usage:
      aem shell zsh [flags]
    
    Aliases:
      zsh-completion, zsh
    
    Flags:
      -h, --help   help for zsh-completion
    
    Global Flags:
      -P, --project string   Run command for project. (if not current working directory)
      -v, --verbose          verbose output



Or replace the last line with the following if you use [zshell](https://sourceforge.net/p/zsh/code/ci/master/tree/)

    echo "source ~/.bash-completion/aem-completion.bash" >> ~/.zshrc
    
### Init script (System V)
This script helps with automaticly booting and stopping AEM when the server starts or stops.

Place the init script in the right location. To set it ready for usage.

	cp aem.init /etc/init.d/aem
	chown root: /etc/init.d/aem
	chmod u+x /etc/init.d/aem

Enable to start automaticly at boot with:

	update-rc.d aem defaults
	update-rc.d aem enable	

## Built With

* [Getopt](https://github.com/pborman/getopt/tree/master/v2) - For command line parsing
* [Go-pretty](github.com/jedib0t/go-pretty/table) - For table printing
* [Afero](https://github.com/spf13/afero) - For FileSystem Abstraction
* [Progressbar](https://github.com/schollz/progressbar) - For progress bar printing
* [Go-humanize](https://github.com/dustin/go-humanize) - Formatters for units to human friendly sizes
* [Go-keyring](https://github.com/zalando/go-keyring) - Store password in operating systems own keyring
* [Cobra](https://github.com/spf13/cobra) - Commandline library
* [TOML](https://github.com/BurntSushi/toml) - TOML parser for Golang with reflection.
* [Tail](https://github.com/hpcloud/tail) - For tailing files
* [Survey](https://github.com/AlecAivazis/survey) - For console survey
* [Go-colortext](https://github.com/daviddengcn/go-colortext) - Color text printing

Thank all authors and contributors of these libraries. For publishing such great software.

## Todo

- Code cleanup
- more code testing and coverage
- combine with lazybones
- Features
  - thread dumps
  - sling tracer

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

