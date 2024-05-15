# Variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=simpledns

# Targets
all: test build

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
	./bin/$(BINARY_NAME) serve

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)

build-docker:
	docker build -t simpledns .

run-docker:
	docker run -p 5353:53/udp -p 5353:53/tcp simpledns
