# Makefile

GO = go
GOTEST = $(GO) test
GOBUILD = $(GO) build

all: test build

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o main main.go

run:
	./main

clean:
	rm main
