package command

const (
	// OsmosisMainnetGrpcEndpoint default lavanet_challenge network gRPC endpoint
	OsmosisMainnetGrpcEndpoint = "grpc.osmosis.zone"

	// OsmosisMainnetGrpcPort default lavanet_challenge gRPC endpoint port
	OsmosisMainnetGrpcPort = 9090
)

type IPBinding string

const (
	LocalHostBinding IPBinding = "127.0.0.1"
	DefaultGRPCPort  int       = 9632
)
