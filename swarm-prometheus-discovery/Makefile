PROJECT_NAME ?= swarm-prometheus-discovery
SHELL 	 := /bin/bash
executable := swarm-prometheus-discovery

all: fmt test $(executable)

clean:
	@rm -rf bin/*

$(executable):
	@echo Building $(executable)
	GO111MODULE=on go build -o bin/$@ main.go

build:
	@echo Building $(executable)
	GOOS=linux GO111MODULE=on go build -o bin/$(executable)-linux-amd64 main.go
	docker build -t someprefix/$(executable) -f Dockerfile .

fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -w -s

test:
	go test ./...

.PHONY: clean $(executable) build fmt test
