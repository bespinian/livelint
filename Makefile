.PHONY: build
all:
	chmod +x ./scripts/run.sh
	./scripts/run.sh
	
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

