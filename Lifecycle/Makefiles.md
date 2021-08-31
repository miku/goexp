# Makefiles

Explicitly not required by Go, but found not completely absent in go projects either.

Example Makefile:

```
SHELL := /bin/bash
TARGETS := dcdump
VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOLDFLAGS += -w -s
GOFLAGS = -ldflags "$(GOLDFLAGS)"

.PHONY: all
all: $(TARGETS)

%: cmd/%/main.go
        go build -o $@ -ldflags "$(GOLDFLAGS)" $<

.PHONY: clean
clean:
        rm -f $(TARGETS)
```

Notes:

* "-X" populates a variable at link time