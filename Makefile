VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo 0.0.0-dev)
PREFIX ?= $(HOME)/.local
BINDIR ?= $(PREFIX)/bin
PRIMARY_BIN ?= bin/namba-intent
SHIM_BIN ?= bin/ni
LDFLAGS := -X ni/internal/version.Version=$(VERSION)

.PHONY: test quality smoke build install-check release-check install-local

test:
	go test ./...

quality:
	bash scripts/quality.sh

smoke:
	bash scripts/smoke.sh

build:
	mkdir -p $(dir $(PRIMARY_BIN))
	go build -ldflags "$(LDFLAGS)" -o $(PRIMARY_BIN) ./cmd/namba-intent
	go build -ldflags "$(LDFLAGS)" -o $(SHIM_BIN) ./cmd/ni

install-check:
	bash scripts/install-check.sh

release-check:
	bash scripts/release-check.sh

install-local:
	mkdir -p $(BINDIR)
	go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/namba-intent ./cmd/namba-intent
	go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/ni ./cmd/ni
