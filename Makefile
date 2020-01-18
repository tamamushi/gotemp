# vim: set ts=4 sts=2 sw=4 tw=0 fenc=utf-8: */

GOCMD	:= go
GOBUILD	:= $(GOCMD) build
GOCLEAN	:= $(GOCMD) clean

VERSION := $(shell git describe --tags --abbrev=0 >&/dev/null || echo "0.0.0.0")
REVISION := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

SRC_DIR	:= ./src
APP_DIR	:= $(SRC_DIR)/app/lambda
BIN_DIR := ./bin

.PHONY: build clean deploy

build: estimate invoice

%: $(APP_DIR)/%
	env GOOS=linux $(GOBUILD) $(LDFLAGS) -a -o $(BIN_DIR)/$@ $</main.go

deploy-%: %
	npx sls deploy function -f $<

clean:
	rm -rf ./bin

deploy: clean build
	npx sls deploy --verbose
