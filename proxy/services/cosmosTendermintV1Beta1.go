package services

import (
	"context"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/Aleksao998/lavanet_challenge/proxy/client"
	"github.com/hashicorp/go-hclog"
)

type serviceServer struct {
	tendermintv1beta1.UnimplementedServiceServer

	// networkClient represents connection with network client
	networkClient client.NetworkClient

	logger hclog.Logger
}

func NewService(
	networkClient client.NetworkClient,
	logger hclog.Logger,
) *serviceServer {
	return &serviceServer{
		logger:        logger.Named("cosmosTendermintV1Beta1"),
		networkClient: networkClient,
	}
}

// GetLatestBlock returns the latest block.
func (s *serviceServer) GetLatestBlock(
	ctx context.Context,
	in *tendermintv1beta1.GetLatestBlockRequest,
) (*tendermintv1beta1.GetLatestBlockResponse, error) {
	s.logger.Info("GetLatestBlock called", "dest", s.networkClient.GrpcAddress.String())

	res, err := s.networkClient.Client.GetLatestBlock(ctx, in)
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
	s.logger.Info("GetBlockByHeight called", "dest", s.networkClient.GrpcAddress.String())

	res, err := s.networkClient.Client.GetBlockByHeight(ctx, in)
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
	s.logger.Info("GetLatestValidatorSet called", "dest", s.networkClient.GrpcAddress.String())

	res, err := s.networkClient.Client.GetLatestValidatorSet(ctx, in)
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
	s.logger.Info("GetValidatorSetByHeight called", "dest", s.networkClient.GrpcAddress.String())

	res, err := s.networkClient.Client.GetValidatorSetByHeight(ctx, in)
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
	s.logger.Info("GetNodeInfo called", "dest", s.networkClient.GrpcAddress.String())

	res, err := s.networkClient.Client.GetNodeInfo(ctx, in)
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
	s.logger.Info("GetSyncing called", "dest", s.networkClient.GrpcAddress.String())

	res, err := s.networkClient.Client.GetSyncing(ctx, in)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	return res, nil
}
