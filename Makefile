.PHONY: all clean release testrelease snapshot packr lint

all: snapshot

clean:
	@-rm test-report.out
	@-rm coverage.out
	@-rm -rf build
	@-rm -rf dist
	@-rm -rf completions
	@-rm *.zip
	@-rm *.tbz2
	@-rm *.tgz
	@-rm aem
	@-rm ./dist

release: clean lint
	./set-gh-key
	goreleaser release

testrelease:
	goreleaser --skip-publish --skip-validate

snapshot: clean lint
	goreleaser --snapshot

packr:
	packr2

lint:
	golint -set_exit_status ./...
