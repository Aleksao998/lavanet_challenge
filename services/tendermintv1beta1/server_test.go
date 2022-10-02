package tendermintv1beta1

import (
	"context"
	"testing"

	tendermintv1beta1 "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

const (
	osmosisMainnetGrpcRaw = "grpc.osmosis.zone:9090"
)

func setupServerService() (*Server, error) {
	// generate networkGrpcAddress from raw
	grpcAddress, err := helper.ResolveAddr(
		"",
		osmosisMainnetGrpcRaw,
	)
	if err != nil {
		return nil, err
	}

	client := Client{
		Client:      mockClient{},
		connection:  nil,
		GrpcAddress: grpcAddress,
		logger:      hclog.NewNullLogger(),
	}

	return NewService(client, hclog.NewNullLogger()), nil
}

// TestServices tests if all grpc calls returns
// reflection from client
func TestServices(t *testing.T) {
	t.Parallel()

	t.Run(
		"GetLatestBlock",
		func(t *testing.T) {
			t.Parallel()

			service, err := setupServerService()
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput, _ := mockClient{}.GetLatestBlock(
				context.Background(),
				&tendermintv1beta1.GetLatestBlockRequest{},
			)

			actualOutput, _ := service.GetLatestBlock(
				context.Background(),
				&tendermintv1beta1.GetLatestBlockRequest{},
			)

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"GetNodeInfo",
		func(t *testing.T) {
			t.Parallel()

			service, err := setupServerService()
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput, _ := mockClient{}.GetNodeInfo(
				context.Background(),
				&tendermintv1beta1.GetNodeInfoRequest{},
			)

			actualOutput, _ := service.GetNodeInfo(
				context.Background(),
				&tendermintv1beta1.GetNodeInfoRequest{},
			)

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"GetSyncing",
		func(t *testing.T) {
			t.Parallel()

			service, err := setupServerService()
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput, _ := mockClient{}.GetSyncing(
				context.Background(),
				&tendermintv1beta1.GetSyncingRequest{},
			)

			actualOutput, _ := service.GetSyncing(
				context.Background(),
				&tendermintv1beta1.GetSyncingRequest{},
			)

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"GetBlockByHeight",
		func(t *testing.T) {
			t.Parallel()

			service, err := setupServerService()
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput, _ := mockClient{}.GetBlockByHeight(
				context.Background(),
				&tendermintv1beta1.GetBlockByHeightRequest{},
			)

			actualOutput, _ := service.GetBlockByHeight(
				context.Background(),
				&tendermintv1beta1.GetBlockByHeightRequest{},
			)

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"GetLatestValidatorSet",
		func(t *testing.T) {
			t.Parallel()

			service, err := setupServerService()
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput, _ := mockClient{}.GetLatestValidatorSet(
				context.Background(),
				&tendermintv1beta1.GetLatestValidatorSetRequest{},
			)

			actualOutput, _ := service.GetLatestValidatorSet(
				context.Background(),
				&tendermintv1beta1.GetLatestValidatorSetRequest{},
			)

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"GetValidatorSetByHeight",
		func(t *testing.T) {
			t.Parallel()

			service, err := setupServerService()
			if err != nil {
				t.Fatal(err)
			}

			expectedOutput, _ := mockClient{}.GetValidatorSetByHeight(
				context.Background(),
				&tendermintv1beta1.GetValidatorSetByHeightRequest{},
			)

			actualOutput, _ := service.GetValidatorSetByHeight(
				context.Background(),
				&tendermintv1beta1.GetValidatorSetByHeightRequest{},
			)

			assert.Equal(t, expectedOutput, actualOutput)
		})
}
