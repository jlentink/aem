.PHONY: all golint vet fmt test coverage scan build linux osx windows
BUILT_HASH=$(shell git rev-parse HEAD)
BUILT_VERSION=1.2.0

all: get test golint coverage build

get:
	cd cmd/aem && go get -t -v

golint:
	cd cmd/aem && golint -set_exit_status

vet:
	cd cmd/aem && go vet

fmt:
	cd cmd/aem && gofmt -s -l .

test:
	cd cmd/aem && export UNIT_TEST=1; go test -json > test-report.out

coverage:
	cd cmd/aem && export UNIT_TEST=1; go test -coverprofile=coverage.out

scan:
	/usr/local/sonar-scanner/bin/sonar-scanner

build: linux osx windows

LDFLAGS=-ldflags "-w -s -X main.BuiltHash=${BUILT_HASH} -X main.BuiltVersion=${BUILT_VERSION}"
linux:
	cd cmd/aem && env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ../../build/linux/aem
	zip linux-${BUILT_VERSION}.zip build/linux/aem

osx:
	cd cmd/aem && env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ../../build/osx/aem
	zip osx-${BUILT_VERSION}.zip build/osx/aem

windows:
	cd cmd/aem && env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ../../build/windows/aem.exe
	zip windows-${BUILT_VERSION}.zip build/windows/aem.exe

