#!/usr/bin/env bash

BINDIR=${BINDIR:-~/bin}
VERSION=1.0.0rc5

while getopts "b:" opt; do
  case ${opt} in
    b )
      BINDIR="${OPTARG}"
      ;;
    \? ) echo "Usage: cmd [-h] [-t]"
      ;;
  esac
done

if [ -d "${BINDIR}" ]; then
  mkdir -p ${BINDIR}
fi

if [ "$(uname -s)" == "Darwin" ]; then
  echo "Downloading OSX version..."
  cd /tmp
  curl -sL https://github.com/jlentink/aem/releases/download/${VERSION}/osx-v${version}.zip --output /tmp/osx-${version}.zip
  unzip -o /tmp/osx-${VERSION}.zip
  echo "Placing aemCLI bineary in ${BINDIR}"
  mv -f /tmp/aem ${BINDIR}
else
  echo "Downloading Linux version..."
  cd /tmp
  mkdir -p ${BINDIR}
  curl -sL https://github.com/jlentink/aem/releases/download/${VERSION}/linux-v${version}.tgz --output /tmp/linux-v${VERSION}.tgz
  tar -zxf /tmp/linux-v${version}.tgz
  echo "Placing aemCLI bineary in ${BINDIR}"
  mv -f /tmp/aem ${BINDIR}
fi

echo "execute \"aem\" from a project folder"
echo "if aem could not be found add ${BINDIR} to your path"
