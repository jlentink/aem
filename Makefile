.PHONY: all golint vet fmt test coverage scan build linux osx windows clean
BUILT_HASH=$(shell git rev-parse HEAD)
BUILT_VERSION=1.2.2

all: clean get test code-test coverage build

clean:
	@-cd cmd/aem && rm test-report.out
	@-cd cmd/aem && rm coverage.out
	@-rm build/linux/aem
	@-rm build/windows/aem.exe
	@-rm build/osx/aem
	@-rm *.zip
	@-rm *.tbz2
	@-rm *.tgz

code-test: golint vet fmt gocyclo ineffassign goreportcard

get:
	go get golang.org/x/tools/cmd/cover
	go get -u golang.org/x/lint/golint
	go get github.com/fzipp/gocyclo
	go get github.com/gordonklaus/ineffassign
	go get github.com/alecthomas/gometalinter
	go get github.com/gojp/goreportcard/cmd/goreportcard-cli
	@cd cmd/aem && go get -t -v

golint:
	@cd cmd/aem && golint -set_exit_status

golintci:
	golangci-lint run

goreportcard:
	goreportcard-cli -t 100

gocyclo:
	@cd cmd/aem && test -z $$(gocyclo -over 15 .)

vet:
	@cd cmd/aem && go vet -all

fmt:
	gofmt -l .
	@cd cmd/aem && test -z $$(go fmt)

ineffassign:
	@cd cmd/aem && test -z $$(ineffassign .)

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
	@cd build/linux/ && tar -jcf ../../linux-v${BUILT_VERSION}.tbz2 aem README.md
	@cd build/linux/ && tar -zcf ../../linux-v${BUILT_VERSION}.tgz aem README.md

osx:
	@cd cmd/aem && env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ../../build/osx/aem
	@cp README.md ./build/osx/
	@cd build/osx/ && zip ../../osx-v${BUILT_VERSION}.zip aem README.md

windows:
	@cd cmd/aem && env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ../../build/windows/aem.exe
	@cp README.md ./build/windows/
	@cd build/windows/ && zip ../../windows-v${BUILT_VERSION}.zip aem.exe README.md

