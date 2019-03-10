#!/usr/bin/env bash

#
# Bash completion for AEMcli
# Add script into your bash/zsh completion path to auto complete on the aem command.
#
# aem<TAB><TAB>
#

__aem_completions()
{
    case ${COMP_WORDS[1]} in
        init)
            COMPREPLY+=("--verbose")
            ;;
        package-list)
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        package-copy)
            COMPREPLY+=("--force-download")
            COMPREPLY+=("--from-name=")
            COMPREPLY+=("--to-group=")
            COMPREPLY+=("--log")
            COMPREPLY+=("--package=")
            COMPREPLY+=("--to-name=")
            COMPREPLY+=("--verbose")
            ;;
        package-install)
            COMPREPLY+=("--to-group=")
            COMPREPLY+=("--no-install")
            COMPREPLY+=("--package=")
            COMPREPLY+=("--to-name=")
            COMPREPLY+=("--verbose")
            COMPREPLY+=("--yes")
            ;;
        package-download)
            COMPREPLY+=("--force-download")
            COMPREPLY+=("--from=")
            COMPREPLY+=("--package=")
            COMPREPLY+=("--verbose")
            ;;
        package-rebuild)
            COMPREPLY+=("--from-name=")
            COMPREPLY+=("--package=")
            COMPREPLY+=("--verbose")
            ;;
        pull-content)
            COMPREPLY+=("--force-download")
            COMPREPLY+=("--from-name=")
            COMPREPLY+=("--to-name=")
            COMPREPLY+=("--verbose")
            ;;
        bundle-list)
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        bundle-start)
            COMPREPLY+=("--bundle=")
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        bundle-stop)
            COMPREPLY+=("--bundle=")
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        bundle-install)
            COMPREPLY+=("--bundle=")
            COMPREPLY+=("--name=")
            COMPREPLY+=("--startlevel=")
            COMPREPLY+=("--verbose")
            ;;
        start)
            COMPREPLY+=("--download")
            COMPREPLY+=("--foreground")
            COMPREPLY+=("--name=")
            COMPREPLY+=("--root")
            COMPREPLY+=("--verbose")
            ;;
        stop)
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        log)
            COMPREPLY+=("--follow")
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        open)
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        system-information)
            COMPREPLY+=("--name=")
            COMPREPLY+=("--verbose")
            ;;
        replicate-page)
            COMPREPLY+=("--activate")
            COMPREPLY+=("--deactivate")
            COMPREPLY+=("--group=")
            COMPREPLY+=("--name=")
            COMPREPLY+=("--page=")
            COMPREPLY+=("--verbose")
            ;;
        activate-tree)
            COMPREPLY+=("--instance=")
            COMPREPLY+=("--path=")
            COMPREPLY+=("--verbose")
            ;;
        password)
            COMPREPLY+=("--all")
            COMPREPLY+=("--group=value")
            COMPREPLY+=("--name=value")
            COMPREPLY+=("--verbose")
            COMPREPLY+=("--yes")
            ;;
        version)
            ;;
        sync)
            COMPREPLY+=("--instance-group=")
            COMPREPLY+=("--instance-name=")
            COMPREPLY+=("--disable-log")
            COMPREPLY+=("--aemsync=")
            COMPREPLY+=("--verbose")
            ;;
        help)
            ;;
        *)
            COMPREPLY+=("init")
            COMPREPLY+=("package-list")
            COMPREPLY+=("package-copy")
            COMPREPLY+=("package-install")
            COMPREPLY+=("package-rebuild")
            COMPREPLY+=("package-download")
            COMPREPLY+=("bundle-list")
            COMPREPLY+=("bundle-start")
            COMPREPLY+=("bundle-stop")
            COMPREPLY+=("bundle-install")
            COMPREPLY+=("pull-content")
            COMPREPLY+=("start")
            COMPREPLY+=("stop")
            COMPREPLY+=("log")
            COMPREPLY+=("open")
            COMPREPLY+=("system-information")
            COMPREPLY+=("replicate-page")
            COMPREPLY+=("activate-tree")
            COMPREPLY+=("password")
            COMPREPLY+=("version")
            COMPREPLY+=("sync")
            COMPREPLY+=("help")
            ;;
    esac
}

if [[ -n ${ZSH_VERSION-} ]]; then
  autoload -U +X bashcompinit && bashcompinit
  autoload -U +X compinit && compinit
fi

complete -F __aem_completions aem
