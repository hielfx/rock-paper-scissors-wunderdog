BINDIR ?= $(CURDIR)/bin
BINNAME ?= rockpaperscissors-api
SERVERPATH ?= ./cmd/server

ARGS ?= -debug

GOBIN ?= $(shell which go)

#go options
CGO_ENABLED ?= 0
LDFLAGS := -s -w

.PHONY:all
all: build

.PHONY: run
run:
	$(GOBIN) run $(SERVERPATH) $(ARGS)

.PHONY:build
build: $(BINDIR)/$(BINNAME)

.PHONY: clean
clean:
	@rm -rf $(BINDIR)

$(BINDIR)/$(BINNAME): clean
	CGO_ENABLED=$(CGO_ENABLED) $(GOBIN) build -trimpath -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) $(SERVERPATH)

.PHONY: test
test:
	$(GOBIN) test -v -race -cover -failfast -coverprofile=./test/coverage.out ./...

.PHONY: coverage
coverage: test
	$(GOBIN) tool cover -html=./test/coverage.out