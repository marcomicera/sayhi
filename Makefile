# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLINTERS=golangci-lint

# Project information
DEFAULT_PORT=8080
DEVELOPER=marcomicera
BINARY_NAME=sayhi
GIT_COMMIT := $(shell git rev-list -1 HEAD)
PROJECT_NAME := $(shell basename `git rev-parse --show-toplevel`)

all: linters test build
linters:
		$(GOLINTERS) run -v ./...
deps:
		$(GOGET) -d -v ./...
build: deps
		$(GOBUILD) -v -ldflags "-X github.com/marcomicera/sayhi/go.GitCommit=$(GIT_COMMIT) \
		-X github.com/marcomicera/sayhi/go.ProjectName=$(PROJECT_NAME)" -o $(BINARY_NAME)
test:
		$(GOTEST) -v ./...
run: build
		./$(BINARY_NAME) -port=$(DEFAULT_PORT)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
build-image:
		docker build \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg PROJECT_NAME=$(PROJECT_NAME) \
		-t $(DEVELOPER)/$(BINARY_NAME) .
run-image: build-image
		docker run \
		--name=$(BINARY_NAME) \
		--rm \
		-p $(DEFAULT_PORT):8080 \
		$(DEVELOPER)/$(BINARY_NAME)
