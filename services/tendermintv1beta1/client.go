package tendermintv1beta1

import (
	"net"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	// Client represents network client
	Client tendermintv1beta1.ServiceClient

	// connection is grpc client connection
	connection *grpc.ClientConn

	// GrpcAddress is network gRPC endpoint
	GrpcAddress *net.TCPAddr

	logger hclog.Logger
}

func NewClient(
	logger hclog.Logger,
	networkGrpcAddress *net.TCPAddr,
) Client {
	logger = logger.Named("tendermint-v1beta1-client")

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
	return Client{
		Client:      tendermintv1beta1.NewServiceClient(conn),
		GrpcAddress: networkGrpcAddress,
		connection:  conn,
		logger:      logger,
	}
}

// Close closes network client
func (s *Client) Close() {
	s.logger.Debug("Closing gRPC client connection", "src", s.GrpcAddress.String())

	s.connection.Close()
}
