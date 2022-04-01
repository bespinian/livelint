.PHONY: build
all:
	go build -o bin/livelint ./cmd/livelint

.PHONY: install
install:
	go install ./cmd/livelint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: clean
clean:
	rm -rf bin

