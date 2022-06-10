.PHONY: build
all:
	GitTag=$(git describe --tags)
	Date=$(date +'%Y-%m-%d_%T')
	GitHash=$(git rev-parse --short HEAD)
	go build -o bin/livelint -ldflags "-X github.com/ropes/go-linker-vars-example/src/version.GitTag=${GitTag} -X github.com/ropes/go-linker-vars-example/src/version.Date=${Date} -X github.com/ropes/go-linker-vars-example/src/version.Version=${GitHash}"

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

