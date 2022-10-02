package tracker

import (
	"testing"

	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/Aleksao998/lavanet_challenge/services/tendermintv1beta1"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	osmosisMainnetGrpcRaw = "grpc.osmosis.zone:9090"
)

func setupBlockTrackerService() (*blockTracker, error) {
	// generate networkGrpcAddress from raw
	grpcAddress, err := helper.ResolveAddr(
		"",
		osmosisMainnetGrpcRaw,
	)
	if err != nil {
		return nil, err
	}

	client := tendermintv1beta1.Client{
		Client:      tendermintv1beta1.MockClient{},
		Connection:  nil,
		GrpcAddress: grpcAddress,
		Logger:      hclog.NewNullLogger(),
	}

	blockTracker := &blockTracker{
		hclog.NewNullLogger(),
		client,
		5,
		0,
		nil,
		blockTrackerResult{},
		nil,
		10,
	}

	return blockTracker, nil
}

// TestGetBlock tests if all get block methods works
// for grpc client mock is used
func TestGetBlock(t *testing.T) {
	t.Parallel()

	t.Run(
		"GetLatestBlock",
		func(t *testing.T) {
			t.Parallel()

			expectedOutput := &blockInfo{
				Number: 10,
				Hash:   &tmbytes.HexBytes{},
			}

			blockTracker, err := setupBlockTrackerService()
			if err != nil {
				t.Fatal(err)
			}

			actualOutput, err := blockTracker.getLatestBlock()
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"getBlock",
		func(t *testing.T) {
			t.Parallel()

			expectedOutput := &blockInfo{
				Number: 10,
				Hash:   &tmbytes.HexBytes{},
			}

			blockTracker, err := setupBlockTrackerService()
			if err != nil {
				t.Fatal(err)
			}

			actualOutput, err := blockTracker.getBlock(3)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedOutput, actualOutput)
		})
	t.Run(
		"getBlocks",
		func(t *testing.T) {
			t.Parallel()

			block := blockInfo{
				Number: 10,
				Hash:   &tmbytes.HexBytes{},
			}

			expectedOutput := []blockInfo{block, block}

			blockTracker, err := setupBlockTrackerService()
			if err != nil {
				t.Fatal(err)
			}

			actualOutput, err := blockTracker.getBlocks(1, 4)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedOutput, actualOutput)
		})
}
