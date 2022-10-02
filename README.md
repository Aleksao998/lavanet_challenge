# Lavanet challenge

Lavanet challenge is an interview task for lavanet. The challenge is structured in 2 parts

* **Server** - basically it needs to works as forward proxy. It needs to be able to register grpc service, forward the request to specific network and returns the response
* **Listener** - service which needs to connect to the server, listen for new blocks and prints test results


## Architecture
![lavanetChallenge](https://user-images.githubusercontent.com/42786413/193462380-9d5f3da4-3fbf-4643-8508-09a890ed4245.png)


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

## Manual testing

To test manually first thing you need to do is to install https://github.com/fullstorydev/grpcurl.

**Steps to do**:

* Start server with default flags (default server grpc address: **localhost::9632**)
```bash
  go run main.go server
```
* Start tracker (enable test results after each new block so you dont need to wait for 5 blocks)
```bash
  go run main.go tracker --output-after 1
```

To query server you can use grpcurl commands:

Method | request | 
--- | --- | 
GetLatestBlock | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetLatestBlock |
GetLatestValidatorSet | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetLatestValidatorSet |
GetNodeInfo | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetNodeInfo |
GetSyncing | grpcurl --plaintext localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetSyncing |
GetBlockByHeight | grpcurl --plaintext -d '{"height":5350708}' localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetBlockByHeight  |
GetValidatorSetByHeight | grpcurl --plaintext -d '{"height":5350708}' localhost:9632 cosmos.base.tendermint.v1beta1.Service.GetValidatorSetByHeight 
