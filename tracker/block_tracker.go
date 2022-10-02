package tracker

import (
	"context"
	"errors"
	"net"
	"time"

	tendermintv1beta1proto "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/Aleksao998/lavanet_challenge/services/tendermintv1beta1"
	"github.com/hashicorp/go-hclog"
)

var (
	errUnmarshalHash = errors.New("unable to unmarshal hash")
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

	// results parses tracker data and outputs in the specified file
	results blockTrackerResults

	// pendingBlocks are blocks which are waiting to be printed
	pendingBlocks []blockInfo

	// outputAfter is a number after which results will be generated
	outputAfter uint64
}

// blockInfo is necessary information from block
type blockInfo struct {
	// number is a block number
	Number uint64 `json:"number"`

	// hash is a block hash
	Hash *tmbytes.HexBytes `json:"hash"`
}

// NewBlockTracker creates a new block tracker service
// it tracks the latest block from specified client
// using cosmos.base.tendermint.v1beta1.Service.GetLatestBlock
func NewBlockTracker(
	logger hclog.Logger,
	clientGrpcAddress *net.TCPAddr,
	pollingTime uint64,
	outputFilePath string,
	outputAfter uint64,
) (*blockTracker, error) {
	blockTracker := &blockTracker{
		logger:                  logger.Named("block-tracker"),
		tendermintV1Beta1Client: tendermintv1beta1.NewClient(logger, clientGrpcAddress),
		pollingTime:             pollingTime,
		ticker:                  time.NewTicker(time.Second * time.Duration(pollingTime)),
		results:                 newBlockTrackerResults(logger, outputFilePath),
		pendingBlocks:           make([]blockInfo, 0, outputAfter),
		outputAfter:             outputAfter,
	}

	// fetch the latest block from client
	blockInfo, err := blockTracker.getLatestBlock()
	if err != nil {
		return nil, err
	}

	// set from block
	blockTracker.fromBlock = blockInfo.Number

	// append latest block
	if err := blockTracker.addBlockInPending(*blockInfo); err != nil {
		return nil, err
	}

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

	s.logger.Debug("Latest block received", "number", blockInfo.Number, "hash", blockInfo.Hash.String())

	// if next block arrived
	// set new block
	if s.isNextBlock(blockInfo.Number) {
		// add block in the pending queue
		if err := s.addBlockInPending(*blockInfo); err != nil {
			return err
		}

		// set new from
		s.fromBlock = blockInfo.Number

		s.logger.Debug("BlockInfo", "number", blockInfo.Number, "hash", blockInfo.Hash.String())
	}

	// if gap exists
	// fetch all blocks in between
	// and set new block
	if s.gapExists(blockInfo.Number) {
		s.logger.Debug("Gap exists", "from", s.fromBlock, "latest", blockInfo.Number)

		// get missing blocks
		blocks, err := s.getBlocks(s.fromBlock, blockInfo.Number)
		if err != nil {
			return err
		}

		// add blocks in the pending queue
		for _, block := range blocks {
			s.logger.Debug("BlockInfo", "number", blockInfo.Number, "hash", blockInfo.Hash.String())

			if err := s.addBlockInPending(block); err != nil {
				return err
			}
		}

		// add the latest block in the pending queue
		if err := s.addBlockInPending(*blockInfo); err != nil {
			return err
		}

		// set new from
		s.fromBlock = blockInfo.Number
	}

	return nil
}

// addBlockInPending adds block in pending queue
// if queue gets full print results and empty
func (s *blockTracker) addBlockInPending(block blockInfo) error {
	s.pendingBlocks = append(s.pendingBlocks, block)

	s.logger.Info("New block added in the pending queue", "number", block.Number, "hash", block.Hash.String())

	// if full print and empty
	if len(s.pendingBlocks) == int(s.outputAfter) {
		if err := s.results.writeResults(s.pendingBlocks); err != nil {
			return err
		}

		s.pendingBlocks = nil
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
		Number: uint64(res.GetBlock().Header.Height),
		Hash:   &tmbytes.HexBytes{},
	}

	// unmarshal hash
	err = blockInfo.Hash.Unmarshal(res.BlockId.Hash)
	if err != nil {
		return nil, errUnmarshalHash
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
		Number: uint64(res.GetBlock().Header.Height),
		Hash:   &tmbytes.HexBytes{},
	}

	// unmarshal hash
	err = blockInfo.Hash.Unmarshal(res.BlockId.Hash)
	if err != nil {
		return nil, errUnmarshalHash
	}

	return blockInfo, nil
}
