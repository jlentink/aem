.PHONY: all golint vet fmt test coverage scan build linux osx windows clean
BUILT_HASH=$(shell git rev-parse HEAD)
BUILT_VERSION=1.0.0rc3
LDFLAGS=-ldflags "-w -s -X internal.commands.versionBuild=${BUILT_HASH} -X internal.commands.versionMain=${BUILT_VERSION}"

all: clean get test code-test coverage build

clean:
	@-rm test-report.out
	@-rm coverage.out
	@-rm build/linux/aem
	@-rm build/windows/aem.exe
	@-rm build/osx/aem
	@-rm *.zip
	@-rm *.tbz2
	@-rm *.tgz

code-test: lint vet fmt cyclo ineffassign card

get:
	go get golang.org/x/tools/cmd/cover
	go get -u golang.org/x/lint/golint
	go get github.com/fzipp/gocyclo
	go get github.com/gordonklaus/ineffassign
	go get github.com/alecthomas/gometalinter
	go get github.com/gojp/goreportcard/cmd/goreportcard-cli
	go get github.com/client9/misspell/cmd/misspell
	go get github.com/spf13/pflag
	go get github.com/daviddengcn/go-colortext
	go get github.com/inconshreveable/mousetrap
	go get -t -v

lint:
	golint -set_exit_status ./...

golintci:
	golangci-lint run


card:
	goreportcard-cli -v -t 100

cyclo:
	@test -z $$(gocyclo -over 15 .)

vet:
	@go vet -all

fmt:
	gofmt -l .
	@test -z $$(go fmt)

ineffassign:
	@test -z $$(ineffassign .)

test:
	@export UNIT_TEST=1; go test -json > test-report.out

coverage:
	@export UNIT_TEST=1; go test -coverprofile=coverage.out

scan:
	/usr/local/sonar-scanner/bin/sonar-scanner

build: linux osx windows

linux:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./build/linux/aem
	upx --brute ./build/linux/aem
	@cp README.md ./build/linux/
	@cd build/linux/ && tar -jcf ../../linux-v${BUILT_VERSION}.tbz2 aem README.md
	@cd build/linux/ && tar -zcf ../../linux-v${BUILT_VERSION}.tgz aem README.md

osx:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ./build/osx/aem
	upx --brute ./build/osx/aem
	@cp README.md ./build/osx/
	@cd build/osx/ && zip ../../osx-v${BUILT_VERSION}.zip aem README.md

windows:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./build/windows/aem.exe
	@cp README.md ./build/windows/
	@cd build/windows/ && zip ../../windows-v${BUILT_VERSION}.zip aem.exe README.md

