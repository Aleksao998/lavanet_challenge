# Lavanet challenge

Lavanet challenge is an interview task for lavanet. The challenge is structured in 2 parts

* **Server** - basically it needs to work as a forward proxy. It needs to be able to register gRPC service, forward the request to a specific network, and return the response
* **Listener** - service which needs to connect to the server, listen for new events and prints test results


## Architecture
![lavanetChallenge](https://user-images.githubusercontent.com/42786413/193462380-9d5f3da4-3fbf-4643-8508-09a890ed4245.png)

### Server

The server is the main service of the lavanet_challenge. It can be extended to support many services. This version supports only *forward_proxy* service

- **Forward proxy** needs to listen for gRPC requests, forwards them to a specific client, and return a response. This version of code support only *cosmos.base.tendermint.v1beta1.Service* but it is envisioned to be easily extended.
  If you wish to contribute and extend services here are the changes needed to be implemented:
  - Create a server and client implementation from desired service [example](https://github.com/Aleksao998/lavanet_challenge/tree/develop/services/tendermintv1beta1)
  - Extended forward proxy object to keep track of new client [here](https://github.com/Aleksao998/lavanet_challenge/blob/develop/proxy/forwardProxy.go#L31)
  - Register server [here](https://github.com/Aleksao998/lavanet_challenge/blob/develop/proxy/forwardProxy.go#L62)

```bash
go run main.go server --help
```
```bash
The default command that starts lavanet_challenge client

Usage:
   server [flags]

Flags:
      --grpc-address string           the GRPC interface (default "127.0.0.1:9632")
  -h, --help                          help for server
      --log-level string              the log level for console output (default "INFO")
      --log-to string                 write all logs to the file at specified location instead of writing them to console
      --network-grpc-address string   network gRPC endpoint (default "grpc.osmosis.zone:9090")

```

### Tracker

Tracker is designed to support multiple specific trackers. This version supports only *block tracker* service. 

- **Block tracker** is a scheduled polling machine whose goal is to print test results for a specified block range. It can be connected to any network which supports *cosmos.base.tendermint.v1beta1.Service.GetLatestBlock* api. Depending on the polling time 3 situations can occur.
  - New block arrived, it just needs to save it in the pending queue
  - The same block arrived, it will be skipped
  - Gap occurred so we need to fetch all blocks in between and save them i the pending queue
  Once a pending queue gets full, the tests result will be exported in the format *fileName_from_to.extension*

To see all supported tracker commands:

```bash
go run main.go tracker --help
```
```bash
The default command that starts the tracker client

Usage:
   tracker [flags]

Flags:
      --client-grpc-address string   client gRPC endpoint (default "127.0.0.1:9632")
  -h, --help                         help for tracker
      --log-level string             the log level for console output (default "INFO")
      --log-to string                write all logs to the file at specified location instead of writing them to console
      --output-after uint            number after which results will be generated (default 5)
      --output-to string             write tracker data to the file at specified location (default "test_results.txt")
      --polling-time uint            polling time in seconds (default 2)
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/Aleksao998/lavanet_challenge.git
```

Go to the project directory

```bash
  cd lavanet_challenge
```

Start the server

```bash
  go run main.go server
```

Start the tracker

```bash
  go run main.go tracker
```

To run linter

```bash
  make lint
```

To build binary

```bash
  make build
```

## Manual testing

To test manually first thing you need to do is to install https://github.com/fullstorydev/grpcurl.

**Steps to do**:

* Start server with default flags (default server grpc address: **localhost::9632**)
```bash
  go run main.go server
```
* Start tracker (enable test results after each new block so you don't need to wait for 5 blocks)
```bash
  go run main.go tracker --output-after 1
```

To query the server you can use grpcurl commands:

Method | request | 
--- | --- | 
GetLatestBlock | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetLatestBlock |
GetLatestValidatorSet | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetLatestValidatorSet |
GetNodeInfo | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetNodeInfo |
GetSyncing | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetSyncing |
GetBlockByHeight | grpcurl --plaintext -d '{"height":5350708}' localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetBlockByHeight  |
GetValidatorSetByHeight | grpcurl --plaintext -d '{"height":5350708}' localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetValidatorSetByHeight 

## Running tests

Lavanet_challenge has both e2e and unit tests. 

- For the e2e test there is a [freamwork](https://github.com/Aleksao998/lavanet_challenge/tree/develop/e2e/framework) which starts how many server instances we want. All e2e tests should be placed in the e2e [folder](https://github.com/Aleksao998/lavanet_challenge/tree/develop/e2e)
- For unit tests, there is a [mock client](https://github.com/Aleksao998/lavanet_challenge/blob/develop/services/tendermintv1beta1/mock_client.go) for existing gRPC services

How to run:

- Build binary
```bash
  make buid
```
- Run all tests
```bash
  go test ./...
```

## Want to contribute

- The lavanet_challenge uses a relatively basic feature proposition [template](https://github.com/Aleksao998/lavanet_challenge/blob/develop/.github/pull_request_template.md), that is concise and to the point.

- For branching this project uses *branch_type/branch_name* format

- Before merging the PR, make sure that both *lint* and *test* git action passed

- Existing branching strategy and examples used so far [merged PR-s](https://github.com/Aleksao998/lavanet_challenge/pulls?q=is%3Apr+is%3Aclosed)

