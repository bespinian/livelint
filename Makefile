BUILDDIR = bin/livelint
BUILDVERSION := $(or $(shell git describe --abbrev=0 --tags),$(development))
BUILDDATE := $(shell date +%Y-%m-%d\ %H:%M)
GITHASH := $(shell git rev-list -1 HEAD)

LDFLAGS=-ldflags="-w -s -X 'main.buildversion=${BUILDVERSION}' -X 'main.builddate=${BUILDDATE}' -X 'main.githash=${GITHASH}'"

.PHONY: build
all:
	go build -o ${BUILDDIR}

.PHONY: build_ldflags
build:
	go build ${LDFLAGS} -o ${BUILDDIR}

.PHONY: install
install:
	go install

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: clean
clean:
	rm -rf bin

