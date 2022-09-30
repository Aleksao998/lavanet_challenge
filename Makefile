.PHONY: lint
lint:
	golangci-lint run --config .golangci.yml

.PHONY: build
build:
	go build -o build/lavanet_challenge main.go