REVISION = $(shell git rev-parse --short HEAD)

all: pre test build

pre: fmt vet lint

fmt:
	go fmt ./...

vet:
	go vet $$(go list ./... | grep -v /vendor/)

lint:
	golint $$(go list ./... | grep -v /vendor/)

test:
	go test -cover ./...

build:
	go build -i -v -o dist/dump ./cmd/dump

.PHONY: all pre fmt test vet lint build
