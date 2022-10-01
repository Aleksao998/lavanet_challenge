package proxy

import (
	"net"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type networkClient struct {
	// client represents network client
	client tendermintv1beta1.ServiceClient

	// connection is grpc client connection
	connection *grpc.ClientConn

	// grpcAddress is network gRPC endpoint
	grpcAddress *net.TCPAddr

	logger hclog.Logger
}

func NewNetworkClient(
	logger hclog.Logger,
	networkGrpcAddress *net.TCPAddr,
) networkClient {
	logger = logger.Named("forward-proxy-network-client")

	// Dial network grpc client
	conn, err := grpc.Dial(
		networkGrpcAddress.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("connection failed: ", "err", err)
	}

	logger.Info("GRPC network client running", "addr", networkGrpcAddress.String())

	// initialize network client
	return networkClient{
		client:      tendermintv1beta1.NewServiceClient(conn),
		connection:  conn,
		grpcAddress: networkGrpcAddress,
		logger:      logger,
	}
}

// Close closes network client
func (s *networkClient) Close() {
	s.logger.Debug("Closing gRPC client connection", "src", s.grpcAddress)

	s.connection.Close()
}
