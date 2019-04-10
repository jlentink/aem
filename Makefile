.PHONY: all golint vet fmt test coverage scan build linux osx windows clean
BUILT_HASH=$(shell git rev-parse HEAD)
BUILT_VERSION=1.2.1

all: clean get test golint vet fmt coverage build

clean:
	@-cd cmd/aem && rm test-report.out
	@-cd cmd/aem && rm coverage.out
	@-rm build/linux/aem
	@-rm build/windows/aem.exe
	@-rm build/osx/aem
	@-rm *.zip

get:
	@cd cmd/aem && go get -t -v

golint:
	@cd cmd/aem && golint -set_exit_status

vet:
	@cd cmd/aem && go vet -all

fmt:
	@cd cmd/aem && test -z $$(go fmt ./...)

test:
	@cd cmd/aem && export UNIT_TEST=1; go test -json > test-report.out

coverage:
	@cd cmd/aem && export UNIT_TEST=1; go test -coverprofile=coverage.out

scan:
	/usr/local/sonar-scanner/bin/sonar-scanner

build: linux osx windows

LDFLAGS=-ldflags "-w -s -X main.BuiltHash=${BUILT_HASH} -X main.BuiltVersion=${BUILT_VERSION}"
linux:
	@cd cmd/aem && env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ../../build/linux/aem
	@cp README.md ./build/linux/
	@cd build/linux/ && zip ../../linux-v${BUILT_VERSION}.zip aem README.md

osx:
	@cd cmd/aem && env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ../../build/osx/aem
	@cp README.md ./build/osx/
	@cd build/osx/ && zip ../../osx-v${BUILT_VERSION}.zip aem README.md

windows:
	@cd cmd/aem && env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ../../build/windows/aem.exe
	@cp README.md ./build/windows/
	@cd build/windows/ && zip ../../windows-v${BUILT_VERSION}.zip aem.exe README.md

