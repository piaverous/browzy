
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=browzy

all: build
build: 
		mkdir -p bin
		$(GOBUILD) ./cmd
		$(GOBUILD) -o bin/$(BINARY_NAME)
clean: 
		rm -rf bin/*
