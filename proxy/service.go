package proxy

import (
	"context"
	"net"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type serviceServer struct {
	tendermintv1beta1.UnimplementedServiceServer

	// grpcAddress is gRPC address of lavanet_challenge client
	grpcAddress *net.TCPAddr

	// networkClient represents connection with network client
	networkClient networkClient

	logger hclog.Logger
}

func NewServiceServer(
	logger hclog.Logger,
	grpcAddress *net.TCPAddr,
	networkGrpcAddress *net.TCPAddr,
) serviceServer {
	return serviceServer{
		logger:        logger.Named("forward-proxy-service"),
		grpcAddress:   grpcAddress,
		networkClient: NewNetworkClient(logger, networkGrpcAddress),
	}
}

// Start starts all server services
func (s *serviceServer) Start() error {
	// create empty grpc server
	grpcServer := grpc.NewServer()

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	// register service
	tendermintv1beta1.RegisterServiceServer(grpcServer, s)

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

// Close closes all server service
func (s *serviceServer) Close() {
	s.logger.Debug("Closing service server")

	s.networkClient.Close()
}

// GetLatestBlock returns the latest block.
func (s *serviceServer) GetLatestBlock(
	ctx context.Context,
	in *tendermintv1beta1.GetLatestBlockRequest,
) (*tendermintv1beta1.GetLatestBlockResponse, error) {
	s.logger.Info("GetLatestBlock called", "dest", s.networkClient.grpcAddress.String())

	res, err := s.networkClient.client.GetLatestBlock(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}

// GetBlockByHeight queries block for given height.
func (s *serviceServer) GetBlockByHeight(
	ctx context.Context,
	in *tendermintv1beta1.GetBlockByHeightRequest,
) (*tendermintv1beta1.GetBlockByHeightResponse, error) {
	s.logger.Info("GetBlockByHeight called", "dest", s.networkClient.grpcAddress.String())

	res, err := s.networkClient.client.GetBlockByHeight(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}

// GetLatestValidatorSet queries latest validator-set.
func (s *serviceServer) GetLatestValidatorSet(
	ctx context.Context,
	in *tendermintv1beta1.GetLatestValidatorSetRequest,
) (*tendermintv1beta1.GetLatestValidatorSetResponse, error) {
	s.logger.Info("GetLatestValidatorSet called", "dest", s.networkClient.grpcAddress.String())

	res, err := s.networkClient.client.GetLatestValidatorSet(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}

// GetValidatorSetByHeight queries validator-set at a given height.
func (s *serviceServer) GetValidatorSetByHeight(
	ctx context.Context,
	in *tendermintv1beta1.GetValidatorSetByHeightRequest,
) (*tendermintv1beta1.GetValidatorSetByHeightResponse, error) {
	s.logger.Info("GetValidatorSetByHeight called", "dest", s.networkClient.grpcAddress.String())

	res, err := s.networkClient.client.GetValidatorSetByHeight(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}

// GetNodeInfo queries the current node info.
func (s *serviceServer) GetNodeInfo(
	ctx context.Context,
	in *tendermintv1beta1.GetNodeInfoRequest,
) (*tendermintv1beta1.GetNodeInfoResponse, error) {
	s.logger.Info("GetNodeInfo called", "dest", s.networkClient.grpcAddress.String())

	res, err := s.networkClient.client.GetNodeInfo(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}

// GetSyncing queries node syncing.
func (s *serviceServer) GetSyncing(
	ctx context.Context,
	in *tendermintv1beta1.GetSyncingRequest,
) (*tendermintv1beta1.GetSyncingResponse, error) {
	s.logger.Info("GetSyncing called", "dest", s.networkClient.grpcAddress.String())

	res, err := s.networkClient.client.GetSyncing(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}
