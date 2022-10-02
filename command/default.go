package command

const (
	// OsmosisMainnetGrpcEndpoint default lavanet_challenge network gRPC endpoint
	OsmosisMainnetGrpcEndpoint = "grpc.osmosis.zone"

	// OsmosisMainnetGrpcPort default lavanet_challenge gRPC endpoint port
	OsmosisMainnetGrpcPort = 9090
)

type IPBinding string

const (
	// LocalHostBinding default local network gRPC endpoint
	LocalHostBinding IPBinding = "127.0.0.1"

	// DefaultGRPCPort default local gRPC endpoint port
	DefaultGRPCPort int = 9632
)
