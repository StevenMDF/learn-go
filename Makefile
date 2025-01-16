.DEFAULT_GOAL := build

.PHONY: fmt vet build clean
.SILENT: clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean