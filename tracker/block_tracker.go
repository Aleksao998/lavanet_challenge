package tracker

import (
	"context"
	"net"
	"time"

	tendermintv1beta1proto "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"

	"github.com/Aleksao998/lavanet_challenge/services/tendermintv1beta1"
	"github.com/hashicorp/go-hclog"
)

type blockTracker struct {
	logger hclog.Logger

	// tendermintV1Beta1Client represents connection with client
	tendermintV1Beta1Client tendermintv1beta1.Client

	// pollingTime is a polling time in seconds
	pollingTime uint64

	// fromBlock from which block to fetch
	fromBlock uint64

	// ticker is used to trigger polling at specific interval
	ticker *time.Ticker
}

// blockInfo is necessary information from block
type blockInfo struct {
	// number is a block number
	number uint64

	// hash is a block hash
	hash string
}

// NewBlockTracker creates a new block tracker service
// it tracks the latest block from specified client
// using cosmos.base.tendermint.v1beta1.Service.GetLatestBlock
func NewBlockTracker(
	logger hclog.Logger,
	clientGrpcAddress *net.TCPAddr,
	pollingTime uint64,
) (*blockTracker, error) {
	blockTracker := &blockTracker{
		logger:                  logger.Named("block-tracker"),
		tendermintV1Beta1Client: tendermintv1beta1.NewClient(logger, clientGrpcAddress),
		pollingTime:             pollingTime,
		ticker:                  time.NewTicker(time.Second * time.Duration(pollingTime)),
	}

	// fetch the latest block from client
	blockInfo, err := blockTracker.getLatestBlock()
	if err != nil {
		return nil, err
	}

	// set from block
	blockTracker.fromBlock = blockInfo.number

	return blockTracker, nil
}

// Start starts block tracker service
func (s *blockTracker) Start() error {
	defer func() {
		s.ticker.Stop()
	}()

	for {
		select {
		case <-s.ticker.C:
			if err := s.fetchBlocks(); err != nil {
				return err
			}
		}
	}

	return nil
}

// Close closes service
func (s *blockTracker) Close() {
	s.logger.Debug("Closing block tracker server")

	s.tendermintV1Beta1Client.Close()
}

// fetchBlocks fetches all blocks in between tick intervals
func (s *blockTracker) fetchBlocks() error {
	// fetch the latest block from client
	blockInfo, err := s.getLatestBlock()
	if err != nil {
		return err
	}

	s.logger.Debug("BlockInfo", "number", blockInfo.number, "hash", blockInfo.hash)

	// if next block arrived
	// set new block
	if s.isNextBlock(blockInfo.number) {
		s.fromBlock = blockInfo.number
		s.logger.Info("BlockInfo", "number", blockInfo.number, "hash", blockInfo.hash)
	}

	// if gap exists
	// fetch all blocks in between
	// and set new block
	if s.gapExists(blockInfo.number) {
		s.logger.Debug("Gap exists", "from", s.fromBlock, "latest", blockInfo.number)

		blocks, err := s.getBlocks(s.fromBlock, blockInfo.number)
		if err != nil {
			return err
		}

		// TODO remove
		for _, block := range blocks {
			s.logger.Info("BlockInfo", "number", block.number, "hash", block.hash)
		}

		s.logger.Info("BlockInfo", "number", blockInfo.number, "hash", blockInfo.hash)
	}

	return nil
}

// isNextBlock checks if latest block is next block
func (s *blockTracker) isNextBlock(latestBlockNumber uint64) bool {
	return latestBlockNumber == s.fromBlock+1
}

// gapExists checks if there is a gap
// between last fetched and latest featched block
func (s *blockTracker) gapExists(latestBlockNumber uint64) bool {
	return latestBlockNumber-s.fromBlock > 1
}

// getBlocks fetches all block
// in between from and to
func (s *blockTracker) getBlocks(from uint64, to uint64) ([]blockInfo, error) {
	// calculate how many blocks are missing
	// exclude both
	length := (to - from) - 1

	// next block to fetch
	next := from + 1

	blocks := make([]blockInfo, length)

	for index := range blocks {
		// Get block from specific height
		blockInfo, err := s.getBlock(next)
		if err != nil {
			return nil, err
		}

		blocks[index] = *blockInfo
		next++
	}

	return blocks, nil
}

// getBlock fetches block from specific height
func (s *blockTracker) getBlock(number uint64) (*blockInfo, error) {
	// fetch the latest block
	res, err := s.tendermintV1Beta1Client.Client.GetBlockByHeight(
		context.Background(),
		&tendermintv1beta1proto.GetBlockByHeightRequest{
			Height: int64(number),
		},
	)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	// extract data from response
	blockInfo := &blockInfo{
		number: uint64(res.GetBlock().Header.Height),
		hash:   string(res.BlockId.Hash),
	}

	return blockInfo, nil
}

// getLatestBlock fetches the latest block
func (s *blockTracker) getLatestBlock() (*blockInfo, error) {
	// fetch the latest block
	res, err := s.tendermintV1Beta1Client.Client.GetLatestBlock(
		context.Background(),
		&tendermintv1beta1proto.GetLatestBlockRequest{},
	)
	if err != nil {
		s.logger.Error("connection failed:", "err", err)

		return nil, err
	}

	// extract data from response
	blockInfo := &blockInfo{
		number: uint64(res.GetBlock().Header.Height),
		hash:   string(res.BlockId.Hash),
	}

	return blockInfo, nil
}
