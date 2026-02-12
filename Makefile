BINARY_NAME ?= spectra
GO ?= go

.PHONY: all build run test fmt tidy clean

all: build

build:
	$(GO) build -o bin/$(BINARY_NAME) ./cmd/spectra

run:
	$(GO) run ./cmd/spectra

test:
	$(GO) test ./...

fmt:
	$(GO) fmt ./...

tidy:
	$(GO) mod tidy

clean:
	rm -rf bin

