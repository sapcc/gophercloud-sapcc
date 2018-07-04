DATE     = $(shell date +%Y%m%d%H%M)
VERSION  := $(DATE)
GOOS     ?= $(shell go env | grep GOOS | cut -d'"' -f2)
BINARIES := gophercloud-limes

LDFLAGS := -X github.com/sapcc/gophercloud-limes/main.VERSION=$(VERSION)
GOFLAGS := -ldflags "$(LDFLAGS)"

SRCDIRS  := .
PACKAGES := $(shell find $(SRCDIRS) -type d)
GOFILES  := $(addsuffix /*.go,$(PACKAGES))
GOFILES  := $(wildcard $(GOFILES))

.PHONY: all clean

all: $(BINARIES:%=bin/$(GOOS)/%)

bin/%: $(GOFILES) Makefile
	GOOS=$(*D) GOARCH=amd64 go build $(GOFLAGS) -v -i -o $(@D)/$(@F) .

test:
	go test -v github.com/sapcc/gophercloud-limes/... 
	
clean:
	rm -rf bin/*

