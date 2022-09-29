lint:
	golangci-lint run --config .golangci.yml

build:
	go build -o lavanet_challenge main.go