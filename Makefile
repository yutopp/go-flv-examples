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
	go test -v -race -cover ./...

bench:
	go test -bench . -benchmem -gcflags="-m -l" ./...

dep-init:
	dep ensure

dep-update:
	dep ensure -update

build:
	go build -i -v -o dist/dump ./cmd/dump

.PHONY: all pre fmt vet lint test bench dep-init dep-update build
