
.PHONY: build clean test help default



BIN_NAME=k3ai

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "alefesta/k3ai"
GOOS:=$(shell go env GOOS)
GOARCH:=$(shell go env GOARCH)

LDFLAGS:= -w -s

default: test

help:
	@echo 'Management commands for k3ai:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-w -s -X github.com/k3ai/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/k3ai/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	dep ensure

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

