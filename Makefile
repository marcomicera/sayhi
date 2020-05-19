GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLINTERS=golangci-lint
DEVELOPER=marcomicera
BINARY_NAME=sayhi
PORT=8080

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
		./$(BINARY_NAME) -port=$(PORT)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)

# Docker targets
build-image: build
		docker build -t $(DEVELOPER)/$(BINARY_NAME) .
run-image: build-image
		docker run \
		--name=$(BINARY_NAME) \
		--rm \
		-p $(PORT):8080 \
		$(DEVELOPER)/$(BINARY_NAME)
