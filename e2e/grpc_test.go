package e2e

import (
	"context"
	"testing"

	tendermintv1beta1proto "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	"github.com/Aleksao998/lavanet_challenge/e2e/framework"
	"github.com/stretchr/testify/assert"
)

func TestGetBlockByHeight(t *testing.T) {
	// connect with V1Beta1Client
	client, err := framework.NewV1Beta1Client(framework.OsmosisMainnetGrpcRaw)
	if err != nil {
		t.Fatal(err)
	}

	// get the latest block from OsmosisMainnet
	res, err := client.Client.GetLatestBlock(
		context.Background(),
		&tendermintv1beta1proto.GetLatestBlockRequest{},
	)
	if err != nil {
		t.Fatal(err)
	}

	// save block height
	blockHeight := res.GetBlock().Header.Height

	// get the block from OsmosisMainnet on blockHeight
	expectedOutput, err := client.Client.GetBlockByHeight(
		context.Background(),
		&tendermintv1beta1proto.GetBlockByHeightRequest{
			Height: blockHeight,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// start local server
	srvs := framework.NewTestServers(t, 1)

	// assigned started server
	srv := srvs[0]

	localClient, err := framework.NewV1Beta1Client(srv.Config.GrpcAddress.String())
	if err != nil {
		t.Fatal(err)
	}

	// get the block from lavanet_challenge server on blockHeight
	serverOutput, err := localClient.Client.GetBlockByHeight(
		context.Background(),
		&tendermintv1beta1proto.GetBlockByHeightRequest{
			Height: blockHeight,
		},
	)

	assert.Equal(t, serverOutput, expectedOutput)
}

func TestGetValidatorSetByHeight(t *testing.T) {
	// connect with V1Beta1Client
	client, err := framework.NewV1Beta1Client(framework.OsmosisMainnetGrpcRaw)
	if err != nil {
		t.Fatal(err)
	}

	// get the latest block from OsmosisMainnet
	res, err := client.Client.GetLatestBlock(
		context.Background(),
		&tendermintv1beta1proto.GetLatestBlockRequest{},
	)
	if err != nil {
		t.Fatal(err)
	}

	// save block height
	blockHeight := res.GetBlock().Header.Height

	// get the validator set from OsmosisMainnet on blockHeight
	expectedOutput, err := client.Client.GetValidatorSetByHeight(
		context.Background(),
		&tendermintv1beta1proto.GetValidatorSetByHeightRequest{
			Height: blockHeight,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// start local server
	srvs := framework.NewTestServers(t, 1)

	// assigned started server
	srv := srvs[0]

	localClient, err := framework.NewV1Beta1Client(srv.Config.GrpcAddress.String())
	if err != nil {
		t.Fatal(err)
	}

	// get the validator set from lavanet_challenge server on blockHeight
	serverOutput, err := localClient.Client.GetValidatorSetByHeight(
		context.Background(),
		&tendermintv1beta1proto.GetValidatorSetByHeightRequest{
			Height: blockHeight,
		},
	)

	assert.Equal(t, serverOutput, expectedOutput)
}
