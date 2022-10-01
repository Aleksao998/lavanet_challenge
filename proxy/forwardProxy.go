package proxy

import (
	"net"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/Aleksao998/lavanet_challenge/proxy/client"
	"github.com/Aleksao998/lavanet_challenge/proxy/services"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ForwardProxy struct {
	// grpcAddress is gRPC address of lavanet_challenge client
	grpcAddress *net.TCPAddr

	// networkClient represents connection with network client
	networkClient client.NetworkClient

	logger hclog.Logger
}

func NewForwardProxy(
	logger hclog.Logger,
	grpcAddress *net.TCPAddr,
	networkGrpcAddress *net.TCPAddr,
) *ForwardProxy {
	return &ForwardProxy{
		logger:        logger.Named("forward-proxy"),
		grpcAddress:   grpcAddress,
		networkClient: client.NewNetworkClient(logger, networkGrpcAddress),
	}
}

// Start starts service
func (s *ForwardProxy) Start() error {
	// create empty grpc server
	grpcServer := grpc.NewServer()

	// register all supported services
	s.registerServices(grpcServer)

	// create listener for grpcAddress
	lis, err := net.Listen("tcp", s.grpcAddress.String())
	if err != nil {
		return err
	}

	// Start listening on grpcAddress
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			s.logger.Error(err.Error())
		}
	}()

	s.logger.Info("GRPC server running", "addr", s.grpcAddress.String())

	return nil
}

// registerServices registers all supported services
func (s *ForwardProxy) registerServices(grpcServer *grpc.Server) {
	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	// register tendermintv1beta1 service
	tendermintv1beta1.RegisterServiceServer(
		grpcServer,
		services.NewService(
			s.networkClient,
			s.logger,
		),
	)
}

// Close closes service
func (s *ForwardProxy) Close() {
	s.logger.Debug("Closing service server")

	s.networkClient.Close()
}
