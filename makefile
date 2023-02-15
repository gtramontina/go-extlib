PATH := $(PWD)/.bin:$(PATH)
SHELL := /usr/bin/env bash -eu -o pipefail
CPUS ?= $(shell (nproc --all || sysctl -n hw.ncpu) 2>/dev/null || echo 1)
MAKEFLAGS += --warn-undefined-variables --output-sync=line --jobs $(CPUS)

include makefile.golangci.mk
include makefile.gotestsum.mk

.git/.hooks.log:
	@git config core.hooksPath .githooks
	@git config --get core.hooksPath > $@
pre-reqs += .git/.hooks.log

test: $(pre-reqs)
	@gotestsum --format-hide-empty-pkg -- -race -cover -timeout=10m -shuffle=on ./...
.PHONY: test

test.failfast: $(pre-reqs)
	@gotestsum --format-hide-empty-pkg --format=dots --max-fails=1 -- -timeout=10m -failfast ./...
.PHONY: test.failfast

test.failfast.mutation: $(pre-reqs)
	@GOCACHE=$(PWD)/.gocache MAKEFLAGS= $(MAKE) test.failfast
.PHONY: test.failfast.mutation

test.mutation: $(pre-reqs)
	@go test -timeout=30m -count=1 -ooze.v -tags=mutation
.PHONY: test.mutation

lint: $(pre-reqs)
	@golangci-lint -v run
.PHONY: lint
