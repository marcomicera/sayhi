# Including env file
cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFUZZBUILD=go-fuzz-build
GOFUZZ=go-fuzz
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLINTERS=golangci-lint

# Project information
PORT=8080
DEVELOPER=marcomicera
BINARY_NAME=sayhi
GIT_COMMIT := $(shell git rev-list -1 HEAD)
PROJECT_NAME := $(shell basename `git rev-parse --show-toplevel`)
FULLY_QUALIFIED_NAME=github.com/$(DEVELOPER)/$(BINARY_NAME)

# Build-time variables
define BUILD_TIME_VARS
-ldflags "-X $(FULLY_QUALIFIED_NAME)/go.GitCommit=$(GIT_COMMIT) -X $(FULLY_QUALIFIED_NAME)/go.ProjectName=$(PROJECT_NAME)"
endef

all: linters build
linters:
		$(GOLINTERS) run -v ./...
deps:
		$(GOGET) -d -v ./...
build: deps test
		$(GOBUILD) -v $(BUILD_TIME_VARS) -o $(BINARY_NAME)
fuzz:
		echo "Press CTRL+C to stop go-fuzz..."
		cd go/fuzz && GO111MODULE=off $(GOFUZZBUILD) && GO111MODULE=off $(GOFUZZ) && cd ../..
test:
		$(GOTEST) -v $(BUILD_TIME_VARS) ./...
run: build
		./$(BINARY_NAME) -port=$(PORT)
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
		--env-file=./config.env \
		-p $(PORT):8080 \
		$(DEVELOPER)/$(BINARY_NAME)
