#!/usr/bin/env bash

BINDIR=${BINDIR:-~/bin}
VERSION=${VERSION:-1.0.1}

while getopts "b:v:" opt; do
  case ${opt} in
    b )
      BINDIR="${OPTARG}"
      ;;
    v )
      VERSION="${OPTARG}"
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
  curl -sL https://github.com/jlentink/aem/releases/download/${VERSION}/aem_Darwin_x86_64.tar.gz --output /tmp/osx-${VERSION}.zip
  unzip -o /tmp/osx-${VERSION}.zip
  echo "Placing aemCLI bineary in ${BINDIR}"
  mv -f /tmp/aem ${BINDIR}
else
  echo "Downloading Linux version..."
  cd /tmp
  mkdir -p ${BINDIR}
  curl -sL https://github.com/jlentink/aem/releases/download/${VERSION}/aem_Linux_x86_64.tar.gz --output /tmp/linux-v${VERSION}.tgz
  tar -zxf /tmp/linux-v${VERSION}.tgz
  echo "Placing aemCLI bineary in ${BINDIR}"
  mv -f /tmp/aem ${BINDIR}
fi

echo "execute \"aem\" from a project folder"
echo "if aem could not be found add ${BINDIR} to your path"
