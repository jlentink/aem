.PHONY: all golint vet fmt test coverage scan build linux osx windows clean
BUILT_HASH=$(shell git rev-parse --short HEAD)
BUILT_VERSION=1.0.0rc10
LDFLAGS=-ldflags "-X github.com/jlentink/aem/internal/version.Build=${BUILT_HASH} -X github.com/jlentink/aem/internal/version.Main=${BUILT_VERSION} -w -s"
TRAVISBUILD?=off

all: clean get test code-test coverage build

clean:
	@-rm test-report.out
	@-rm coverage.out
	@-rm -rf build/*
	@-rm *.zip
	@-rm *.tbz2
	@-rm *.tgz
	@-rm aem

code-test: golintci

get:
	echo "get"

lint:
	golint -set_exit_status ./...

golintci:
	golangci-lint run

brew:
	GO111MODULE="on"
	GOPATH=$(shell pwd)/vendor
	go get
	go build ${LDFLAGS}

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

build: packr linux osx windows packr-clean

packr:
	packr2

packr-clean:
	packr2 clean

linux:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./build/linux/aem
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./build/linux/aem
endif
	@cp README.md ./build/linux/
	@cd build/linux/ && tar -jcf ../../linux-v${BUILT_VERSION}.tbz2 aem README.md
	@cd build/linux/ && tar -zcf ../../linux-v${BUILT_VERSION}.tgz aem README.md

osx:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ./build/osx/aem
ifeq ("$(TRAVISBUILD)","off")
	upx --brute ./build/osx/aem
endif
	@cp README.md ./build/osx/
	@cd build/osx/ && zip ../../osx-v${BUILT_VERSION}.zip aem README.md

windows:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ./build/windows.64/aem.exe
	@cp README.md ./build/windows.64/
	@cd build/windows.64/ && zip ../../windows-v${BUILT_VERSION}.amd64.zip aem.exe README.md
	env GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ./build/windows.32/aem.exe
	@cp README.md ./build/windows.32/
	@cd build/windows.32/ && zip ../../windows-v${BUILT_VERSION}.amd32.zip aem.exe README.md
