GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLINTERS=golangci-lint
DEVELOPER=marcomicera
BINARY_NAME=sayhi

all: linters test build
linters:
		$(GOLINTERS) run -v ./...
deps:
		$(GOGET) -d -v ./...
build: deps
		$(GOBUILD) -v -o $(BINARY_NAME)
test:
		$(GOTEST) -v ./...
run: build
		./$(BINARY_NAME)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
image:
		docker build -t $(DEVELOPER)/$(BINARY_NAME) .
