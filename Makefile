

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
	GOOS= GOARCH= GOARM= GOFLAGS= CGO_ENABLED= go build -o $@ $<
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
completions: bin/gh$(EXE)
	mkdir -p ./share/bash-completion/completions ./share/fish/vendor_completions.d ./share/zsh/site-functions
	bin/gh$(EXE) completion -s bash > ./share/bash-completion/completions/gh
	bin/gh$(EXE) completion -s fish > ./share/fish/vendor_completions.d/gh.fish
	bin/gh$(EXE) completion -s zsh > ./share/zsh/site-functions/_gh

# just a convenience task around `go test`
.PHONY: test
test:
	go test ./...

ray:
	@go get github.com/octoper/go-ray

update:
	go get -u ./... && go mod tidy

lint:
	@golangci-lint run
