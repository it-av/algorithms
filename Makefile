DESTDIR=/
SHELL=/bin/bash

.PHONY: algorithms

# App
all: algorithms

check-gopath:
ifneq ($(shell [ -n $(GOPATH) ] && echo set), set)
	@echo GOPATH is not set
	@exit 1
endif

deps: check-gopath

match-the-shoes:
	go build -o match-the-shoes hackerrank/zalando-codesprint/match-the-shoes.go

gofmt:
	for F in `find $(GOPATH)/src/ -name "*.go"` ; do \
	    go fmt $$F ; \
	done \