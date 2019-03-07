.PHONY: all test coverage scan build linux osx windows
BUILT_HASH=$(shell git rev-parse HEAD)
BUILT_VERSION=1.0

all: test coverage scan build

test:
	export UNIT_TEST=1; go test -json > test-report.out

coverage:
	export UNIT_TEST=1; go test -coverprofile=coverage.out

scan:
	/usr/local/sonar-scanner/bin/sonar-scanner

build: linux osx windows

LDFLAGS=-ldflags "-w -s -X main.BuiltHash=${BUILT_HASH} -X main.BuiltVersion=${BUILT_VERSION}"
linux:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux/aem

osx:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/osx/aem

windows:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/windows/aem.exe

