go_bin := .go/bin
export GOBIN = $(PWD)/$(go_bin)

PATH := $(PWD)/$(go_bin):$(PATH)
SHELL := /usr/bin/env bash -eu -o pipefail
CPUS ?= $(shell (nproc --all || sysctl -n hw.ncpu) 2>/dev/null || echo 1)
MAKEFLAGS += --warn-undefined-variables --output-sync=line --jobs $(CPUS)
.DEFAULT_GOAL := help
.SECONDEXPANSION:
.DELETE_ON_ERROR:

.githooks.log:
	@git config core.hooksPath .githooks
	@git config --get core.hooksPath > $@
pre-reqs += .githooks.log

# ---

$(go_bin)/go:
	@git clone --depth=1 https://go.googlesource.com/go .go
	@cd .go/src; ./make.bash
pre-reqs += $(go_bin)/go

go.tools.log: go.tools.go go.tools.mod | $(go_bin)/go
	@date | sed -e :a -e 's/^.\{1,79\}$$/-&/;ta' >> $@
	@go mod tidy -modfile=go.tools.mod 2>&1 | tee -a $@
	@go get -modfile=go.tools.mod $$(go list -tags=tools -f '{{ join .Imports "\n" }}' .) 2>&1 | tee -a $@
	@go install -modfile=go.tools.mod $$(go list -tags=tools -f '{{ join .Imports "\n" }}' .) 2>&1 | tee -a $@
pre-reqs += go.tools.log

# ---

install: $(pre-reqs)
.PHONY: install

test: | $(pre-reqs)
	@gotestsum -- -race -cover -test.shuffle=on ./...
.PHONY: test

test.watch: | $(pre-reqs)
	@gotestsum --format dots --watch -- -cover ./...
.PHONY: test.watch

bench: | $(pre-reqs)
	@go test -bench=. -benchmem ./...
.PHONY: bench

lint: | $(pre-reqs)
	@gofmt -w $$({ git ls-files -- '*.go'; git ls-files --others --exclude-standard -- '*.go'; })
	@goimports -w $$({ git ls-files -- '*.go'; git ls-files --others --exclude-standard -- '*.go'; })
.PHONY: lint

clobber:
	@rm -rf .go .tools *.log
.PHONY: clobber

pre-commit: | $(pre-reqs)
	@MAKEFLAGS= $(MAKE) lint
	@gotestsum -- -cover ./...
.PHONY: pre-commit

tag-release:
	@[ -z $$(git tag -l $$(ccv)) ] && (git tag $$(ccv) && git push --tags) || echo "Latest version already tagged!"
.PHONY: tag-release

# --

help:
ifndef help
	@echo -e "\nAvailable phony targets:\n"
	@help=true MAKEFLAGS= $(MAKE) -rpn \
	| sed -rn "s/^\.PHONY: (.*)/\1/p" \
	| tr " " "\n" \
	| sort -u \
	| sed -re "s/^($(.DEFAULT_GOAL))$$/\1 $$(tput setaf 2)*$$(tput sgr0)/" \
	| sed -e "s/^/  $$(tput setaf 8)make$$(tput sgr0) /"
endif
.PHONY: help
