VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo 0.0.0-dev)
PREFIX ?= $(HOME)/.local
BINDIR ?= $(PREFIX)/bin
BIN ?= bin/ni
LDFLAGS := -X ni/internal/version.Version=$(VERSION)

.PHONY: test quality smoke build install-local

test:
	go test ./...

quality:
	bash scripts/quality.sh

smoke:
	bash scripts/smoke.sh

build:
	mkdir -p $(dir $(BIN))
	go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/ni

install-local:
	mkdir -p $(BINDIR)
	go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/ni ./cmd/ni
