.PHONY: help clean build

all: help

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

include .env
PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
BASE=$(shell pwd)
BIN=$(BASE)/bin
FILES=$(wildcard **/*.go)


build: go-build

clean: go-clean


## go-build: build binary `go build` internally.
go-build: 
	@echo "  >  Building binary..."
	go build -o $(BIN)/$(PROJECTNAME) $(BASE)/cmd/main.go

## go-clean: Clean build files. Runs `go clean` internally.
go-clean:
	@echo "  >  Cleaning build cache"
	rm -rf ./bin 
	go clean