################################################################################
# This file is AUTOGENERATED with <https://github.com/sapcc/go-makefile-maker> #
# Edit Makefile.maker.yaml instead.                                            #
################################################################################

# Copyright 2024 SAP SE
# SPDX-License-Identifier: Apache-2.0

MAKEFLAGS=--warn-undefined-variables
# /bin/sh is dash on Debian which does not support all features of ash/bash
# to fix that we use /bin/bash only on Debian to not break Alpine
ifneq (,$(wildcard /etc/os-release)) # check file existence
	ifneq ($(shell grep -c debian /etc/os-release),0)
		SHELL := /bin/bash
	endif
endif

default: FORCE
	@echo 'There is nothing to build, use `make check` for running the test suite or `make help` for a list of available targets.'

install-golangci-lint: FORCE
	@if ! hash golangci-lint 2>/dev/null; then printf "\e[1;36m>> Installing golangci-lint (this may take a while)...\e[0m\n"; go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; fi

install-go-licence-detector: FORCE
	@if ! hash go-licence-detector 2>/dev/null; then printf "\e[1;36m>> Installing go-licence-detector...\e[0m\n"; go install go.elastic.co/go-licence-detector@latest; fi

install-addlicense: FORCE
	@if ! hash addlicense 2>/dev/null; then  printf "\e[1;36m>> Installing addlicense...\e[0m\n";  go install github.com/google/addlicense@latest; fi

prepare-static-check: FORCE install-golangci-lint install-go-licence-detector install-addlicense

GO_BUILDFLAGS =
GO_LDFLAGS =
GO_TESTENV =
GO_BUILDENV =

# These definitions are overridable, e.g. to provide fixed version/commit values when
# no .git directory is present or to provide a fixed build date for reproducibility.
BININFO_VERSION     ?= $(shell git describe --tags --always --abbrev=7)
BININFO_COMMIT_HASH ?= $(shell git rev-parse --verify HEAD)
BININFO_BUILD_DATE  ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# which packages to test with test runner
GO_TESTPKGS := $(shell go list -f '{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./... | grep -Ev '/acceptance')
ifeq ($(GO_TESTPKGS),)
GO_TESTPKGS := ./...
endif
# which packages to measure coverage for
GO_COVERPKGS := $(shell go list ./... | grep -Ev '/acceptance')
# to get around weird Makefile syntax restrictions, we need variables containing nothing, a space and comma
null :=
space := $(null) $(null)
comma := ,

check: FORCE static-check build/cover.html
	@printf "\e[1;32m>> All checks successful.\e[0m\n"

run-golangci-lint: FORCE install-golangci-lint
	@printf "\e[1;36m>> golangci-lint\e[0m\n"
	@golangci-lint run

build/cover.out: FORCE | build
	@printf "\e[1;36m>> Running tests\e[0m\n"
	@env $(GO_TESTENV) go test -shuffle=on -p 1 -coverprofile=$@ $(GO_BUILDFLAGS) -ldflags '-s -w -X github.com/sapcc/go-api-declarations/bininfo.binName=v2 -X github.com/sapcc/go-api-declarations/bininfo.version=$(BININFO_VERSION) -X github.com/sapcc/go-api-declarations/bininfo.commit=$(BININFO_COMMIT_HASH) -X github.com/sapcc/go-api-declarations/bininfo.buildDate=$(BININFO_BUILD_DATE) $(GO_LDFLAGS)' -covermode=count -coverpkg=$(subst $(space),$(comma),$(GO_COVERPKGS)) $(GO_TESTPKGS)

build/cover.html: build/cover.out
	@printf "\e[1;36m>> go tool cover > build/cover.html\e[0m\n"
	@go tool cover -html $< -o $@

static-check: FORCE run-golangci-lint check-dependency-licenses check-license-headers

build:
	@mkdir $@

tidy-deps: FORCE
	go mod tidy
	go mod verify

force-license-headers: FORCE install-addlicense
	@printf "\e[1;36m>> addlicense\e[0m\n"
	echo -n $(patsubst $(shell awk '$$1 == "module" {print $$2}' go.mod)%,.%/*.go,$(shell go list ./...)) | xargs -d" " -I{} bash -c 'year="$$(rg -P "Copyright (....) SAP SE" -Nor "\$$1" {})"; awk -i inplace '"'"'{if (display) {print} else {!/^\/\*/ && !/^\*/ && !/^\$$/}}; /^package /{print;display=1}'"'"' {}; addlicense -c "SAP SE" -s=only -y "$$year" -ignore "./internal/acceptance/clients/http.go" -ignore "./internal/acceptance/tools/tools.go" -- {}'

license-headers: FORCE install-addlicense
	@printf "\e[1;36m>> addlicense\e[0m\n"
	@addlicense -c "SAP SE" -s=only -ignore "./internal/acceptance/clients/http.go" -ignore "./internal/acceptance/tools/tools.go" -- $(patsubst $(shell awk '$$1 == "module" {print $$2}' go.mod)%,.%/*.go,$(shell go list ./...))

check-license-headers: FORCE install-addlicense
	@printf "\e[1;36m>> addlicense --check\e[0m\n"
	@addlicense --check -ignore "./internal/acceptance/clients/http.go" -ignore "./internal/acceptance/tools/tools.go" -- $(patsubst $(shell awk '$$1 == "module" {print $$2}' go.mod)%,.%/*.go,$(shell go list ./...))

check-dependency-licenses: FORCE install-go-licence-detector
	@printf "\e[1;36m>> go-licence-detector\e[0m\n"
	@go list -m -mod=readonly -json all | go-licence-detector -includeIndirect -rules .license-scan-rules.json -overrides .license-scan-overrides.jsonl

goimports: FORCE
	@printf "\e[1;36m>> goimports -w -local https://github.com/sapcc/gophercloud-sapcc\e[0m\n"
	@goimports -w -local https://github.com/sapcc/gophercloud-sapcc internal/ $(patsubst $(shell awk '$$1 == "module" {print $$2}' go.mod)%,.%/*.go,$(shell go list ./...))

clean: FORCE
	git clean -dxf build

vars: FORCE
	@printf "BININFO_BUILD_DATE=$(BININFO_BUILD_DATE)\n"
	@printf "BININFO_COMMIT_HASH=$(BININFO_COMMIT_HASH)\n"
	@printf "BININFO_VERSION=$(BININFO_VERSION)\n"
	@printf "GO_BUILDFLAGS=$(GO_BUILDFLAGS)\n"
	@printf "GO_COVERPKGS=$(GO_COVERPKGS)\n"
	@printf "GO_LDFLAGS=$(GO_LDFLAGS)\n"
	@printf "GO_TESTENV=$(GO_TESTENV)\n"
	@printf "GO_TESTPKGS=$(GO_TESTPKGS)\n"
help: FORCE
	@printf "\n"
	@printf "\e[1mUsage:\e[0m\n"
	@printf "  make \e[36m<target>\e[0m\n"
	@printf "\n"
	@printf "\e[1mGeneral\e[0m\n"
	@printf "  \e[36mvars\e[0m                         Display values of relevant Makefile variables.\n"
	@printf "  \e[36mhelp\e[0m                         Display this help.\n"
	@printf "\n"
	@printf "\e[1mPrepare\e[0m\n"
	@printf "  \e[36minstall-golangci-lint\e[0m        Install golangci-lint required by run-golangci-lint/static-check\n"
	@printf "  \e[36minstall-go-licence-detector\e[0m  Install-go-licence-detector required by check-dependency-licenses/static-check\n"
	@printf "  \e[36minstall-addlicense\e[0m           Install addlicense required by check-license-headers/license-headers/static-check\n"
	@printf "  \e[36mprepare-static-check\e[0m         Install any tools required by static-check. This is used in CI before dropping privileges, you should probably install all the tools using your package manager\n"
	@printf "\n"
	@printf "\e[1mTest\e[0m\n"
	@printf "  \e[36mcheck\e[0m                        Run the test suite (unit tests and golangci-lint).\n"
	@printf "  \e[36mrun-golangci-lint\e[0m            Install and run golangci-lint. Installing is used in CI, but you should probably install golangci-lint using your package manager.\n"
	@printf "  \e[36mbuild/cover.out\e[0m              Run tests and generate coverage report.\n"
	@printf "  \e[36mbuild/cover.html\e[0m             Generate an HTML file with source code annotations from the coverage report.\n"
	@printf "  \e[36mstatic-check\e[0m                 Run static code checks\n"
	@printf "\n"
	@printf "\e[1mDevelopment\e[0m\n"
	@printf "  \e[36mtidy-deps\e[0m                    Run go mod tidy and go mod verify.\n"
	@printf "  \e[36mforce-license-headers\e[0m        Remove and re-add all license headers to all non-vendored source code files.\n"
	@printf "  \e[36mlicense-headers\e[0m              Add license headers to all non-vendored source code files.\n"
	@printf "  \e[36mcheck-license-headers\e[0m        Check license headers in all non-vendored .go files.\n"
	@printf "  \e[36mcheck-dependency-licenses\e[0m    Check all dependency licenses using go-licence-detector.\n"
	@printf "  \e[36mgoimports\e[0m                    Run goimports on all non-vendored .go files\n"
	@printf "  \e[36mclean\e[0m                        Run git clean.\n"

.PHONY: FORCE
