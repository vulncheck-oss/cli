

EXE =
ifeq ($(shell go env GOOS),windows)
EXE = .exe
endif

## The following tasks delegate to `script/build.go` so they can be run cross-platform.

.PHONY: bin/vulncheck$(EXE)
bin/vulncheck$(EXE): script/build$(EXE)
	@script/build$(EXE) $@

script/build$(EXE): script/build.go
ifeq ($(EXE),)
	GOOS= GOARCH= GOARM= GOFLAGS= go build -o $@ $<
else
	go build -o $@ $<
endif

.PHONY: clean
clean: script/build$(EXE)
	@$< $@

.PHONY: manpages
manpages: script/build$(EXE)
	@$< $@

.PHONY: completions
completions: bin/vulncheck$(EXE)
	mkdir -p ./share/bash-completion/completions ./share/fish/vendor_completions.d ./share/zsh/site-functions
	bin/vulncheck$(EXE) completion bash > ./share/bash-completion/completions/vulncheck
	bin/vulncheck$(EXE) completion fish > ./share/fish/vendor_completions.d/vulncheck.fish
	bin/vulncheck$(EXE) completion zsh  > ./share/zsh/site-functions/_vulncheck

# just a convenience task around `go test`
.PHONY: test
test:
	go test ./...

# Opt-in offline↔online parity suite. Needs VC_TOKEN and synced offline
# indices; the tests skip themselves if either is missing. See pkg/parity.
.PHONY: test-parity
test-parity:
	go test -tags=parity ./pkg/parity/...

dbug:
	@go get github.com/dbugapp/dbug-go

nodbug:
	@go mod edit -droprequire github.com/dbugapp/dbug-go && go mod tidy

update:
	go get -u ./... && go mod tidy

format:
	go fmt  ./...

lint:
	@golangci-lint run

