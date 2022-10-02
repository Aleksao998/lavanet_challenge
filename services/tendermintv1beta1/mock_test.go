package tendermintv1beta1

import (
	"context"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	types "cosmossdk.io/api/tendermint/types"
	"google.golang.org/grpc"
)

type mockClient struct {
	// Client represents network client
	Client tendermintv1beta1.ServiceClient
}

func (c mockClient) GetLatestBlock(ctx context.Context, in *tendermintv1beta1.GetLatestBlockRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetLatestBlockResponse, error) {
	block := &types.Block{
		Header: &types.Header{
			Height: 10,
		},
	}

	resp := &tendermintv1beta1.GetLatestBlockResponse{
		Block: block,
	}

	return resp, nil
}

func (c mockClient) GetNodeInfo(ctx context.Context, in *tendermintv1beta1.GetNodeInfoRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetNodeInfoResponse, error) {
	return &tendermintv1beta1.GetNodeInfoResponse{}, nil
}
func (c mockClient) GetSyncing(ctx context.Context, in *tendermintv1beta1.GetSyncingRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetSyncingResponse, error) {
	return &tendermintv1beta1.GetSyncingResponse{}, nil
}

func (c mockClient) GetBlockByHeight(ctx context.Context, in *tendermintv1beta1.GetBlockByHeightRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetBlockByHeightResponse, error) {
	block := &types.Block{
		Header: &types.Header{
			Height: 10,
		},
	}

	resp := &tendermintv1beta1.GetBlockByHeightResponse{
		Block: block,
	}

	return resp, nil
}
func (c mockClient) GetLatestValidatorSet(ctx context.Context, in *tendermintv1beta1.GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetLatestValidatorSetResponse, error) {
	return &tendermintv1beta1.GetLatestValidatorSetResponse{}, nil
}
func (c mockClient) GetValidatorSetByHeight(ctx context.Context, in *tendermintv1beta1.GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetValidatorSetByHeightResponse, error) {
	return &tendermintv1beta1.GetValidatorSetByHeightResponse{}, nil
}
func (c mockClient) ABCIQuery(ctx context.Context, in *tendermintv1beta1.ABCIQueryRequest, opts ...grpc.CallOption) (*tendermintv1beta1.ABCIQueryResponse, error) {
	return nil, nil
}
