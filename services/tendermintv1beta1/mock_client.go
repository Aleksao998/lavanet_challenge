//nolint:lll
package tendermintv1beta1

import (
	"context"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	types "cosmossdk.io/api/tendermint/types"
	"google.golang.org/grpc"
)

type MockClient struct {
	// Client represents network client
	Client tendermintv1beta1.ServiceClient
}

func (c MockClient) GetLatestBlock(ctx context.Context, in *tendermintv1beta1.GetLatestBlockRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetLatestBlockResponse, error) {
	block := &types.Block{
		Header: &types.Header{
			Height: 10,
		},
	}

	blockID := &types.BlockID{
		Hash: []byte{},
	}

	resp := &tendermintv1beta1.GetLatestBlockResponse{
		Block:   block,
		BlockId: blockID,
	}

	return resp, nil
}

func (c MockClient) GetNodeInfo(ctx context.Context, in *tendermintv1beta1.GetNodeInfoRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetNodeInfoResponse, error) {
	return &tendermintv1beta1.GetNodeInfoResponse{}, nil
}
func (c MockClient) GetSyncing(ctx context.Context, in *tendermintv1beta1.GetSyncingRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetSyncingResponse, error) {
	return &tendermintv1beta1.GetSyncingResponse{}, nil
}

func (c MockClient) GetBlockByHeight(ctx context.Context, in *tendermintv1beta1.GetBlockByHeightRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetBlockByHeightResponse, error) {
	block := &types.Block{
		Header: &types.Header{
			Height: 10,
		},
	}

	blockID := &types.BlockID{
		Hash: []byte{},
	}

	resp := &tendermintv1beta1.GetBlockByHeightResponse{
		Block:   block,
		BlockId: blockID,
	}

	return resp, nil
}
func (c MockClient) GetLatestValidatorSet(ctx context.Context, in *tendermintv1beta1.GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetLatestValidatorSetResponse, error) {
	return &tendermintv1beta1.GetLatestValidatorSetResponse{}, nil
}
func (c MockClient) GetValidatorSetByHeight(ctx context.Context, in *tendermintv1beta1.GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*tendermintv1beta1.GetValidatorSetByHeightResponse, error) {
	return &tendermintv1beta1.GetValidatorSetByHeightResponse{}, nil
}
func (c MockClient) ABCIQuery(ctx context.Context, in *tendermintv1beta1.ABCIQueryRequest, opts ...grpc.CallOption) (*tendermintv1beta1.ABCIQueryResponse, error) {
	return nil, nil
}
