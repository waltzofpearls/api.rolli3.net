TIMESTAMP = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
VERSION = $(shell git rev-parse HEAD)

all: build

build: *.go
	go build \
		-ldflags '-X main.apiBuildTimestamp=$(TIMESTAMP) -X main.apiVersion=$(VERSION)' \
		-o api \
		.
