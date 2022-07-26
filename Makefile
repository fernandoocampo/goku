GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_UNIX=$(BINARY_NAME)-amd64-linux
BINARY_NAME=bin/goku
SRC_FOLDER=cmd/goku

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./$(SRC_FOLDER)