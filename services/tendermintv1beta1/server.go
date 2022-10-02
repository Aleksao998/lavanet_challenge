package tendermintv1beta1

import (
	"context"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/hashicorp/go-hclog"
)

type Server struct {
	tendermintv1beta1.UnimplementedServiceServer

	// client represents connection with client
	client Client

	logger hclog.Logger
}

func NewService(
	client Client,
	logger hclog.Logger,
) *Server {
	return &Server{
		logger: logger.Named("tendermint-v1beta1-server"),
		client: client,
	}
}

// GetLatestBlock returns the latest block.
func (s *Server) GetLatestBlock(
	ctx context.Context,
	in *tendermintv1beta1.GetLatestBlockRequest,
) (*tendermintv1beta1.GetLatestBlockResponse, error) {
	s.logger.Info("GetLatestBlock called", "dest", s.client.GrpcAddress.String())

	return s.client.Client.GetLatestBlock(ctx, in)
}

// GetBlockByHeight queries block for given height.
func (s *Server) GetBlockByHeight(
	ctx context.Context,
	in *tendermintv1beta1.GetBlockByHeightRequest,
) (*tendermintv1beta1.GetBlockByHeightResponse, error) {
	s.logger.Info("GetBlockByHeight called", "dest", s.client.GrpcAddress.String())

	return s.client.Client.GetBlockByHeight(ctx, in)
}

// GetLatestValidatorSet queries latest validator-set.
func (s *Server) GetLatestValidatorSet(
	ctx context.Context,
	in *tendermintv1beta1.GetLatestValidatorSetRequest,
) (*tendermintv1beta1.GetLatestValidatorSetResponse, error) {
	s.logger.Info("GetLatestValidatorSet called", "dest", s.client.GrpcAddress.String())

	return s.client.Client.GetLatestValidatorSet(ctx, in)
}

// GetValidatorSetByHeight queries validator-set at a given height.
func (s *Server) GetValidatorSetByHeight(
	ctx context.Context,
	in *tendermintv1beta1.GetValidatorSetByHeightRequest,
) (*tendermintv1beta1.GetValidatorSetByHeightResponse, error) {
	s.logger.Info("GetValidatorSetByHeight called", "dest", s.client.GrpcAddress.String())

	return s.client.Client.GetValidatorSetByHeight(ctx, in)
}

// GetNodeInfo queries the current node info.
func (s *Server) GetNodeInfo(
	ctx context.Context,
	in *tendermintv1beta1.GetNodeInfoRequest,
) (*tendermintv1beta1.GetNodeInfoResponse, error) {
	s.logger.Info("GetNodeInfo called", "dest", s.client.GrpcAddress.String())

	return s.client.Client.GetNodeInfo(ctx, in)
}

// GetSyncing queries node syncing.
func (s *Server) GetSyncing(
	ctx context.Context,
	in *tendermintv1beta1.GetSyncingRequest,
) (*tendermintv1beta1.GetSyncingResponse, error) {
	s.logger.Info("GetSyncing called", "dest", s.client.GrpcAddress.String())

	return s.client.Client.GetSyncing(ctx, in)
}
