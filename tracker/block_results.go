package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/go-hclog"
)

type result struct {
	TestResult []blockInfo `json:"test_result"`
}

type blockTrackerResults struct {
	logger hclog.Logger

	// outputFilePath location of output file
	outputFilePath string
}

func newBlockTrackerResults(
	logger hclog.Logger,
	outputFilePath string,
) blockTrackerResults {
	return blockTrackerResults{
		logger:         logger.Named("tracker-results"),
		outputFilePath: outputFilePath,
	}
}

// writeResults writes results to the outputFilePath
func (b blockTrackerResults) writeResults(blocks []blockInfo) error {
	// marshall data
	data, err := json.MarshalIndent(
		result{
			TestResult: blocks,
		},
		"",
		"    ",
	)
	if err != nil {
		return fmt.Errorf("failed to generate data: %w", err)
	}

	// generate test name
	fileName := b.getTestFileName(
		blocks[0].Number,
		blocks[len(blocks)-1].Number,
	)

	// write to file
	if err := os.WriteFile(fileName, data, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write results: %w", err)
	}

	b.logger.Info("Results printed", "src", fileName)

	return nil
}

func (b blockTrackerResults) getTestFileName(from uint64, to uint64) string {
	fileName := strings.Split(b.outputFilePath, ".")

	// correct file name format should be filename.extension
	if len(fileName) != 2 {
		return b.outputFilePath
	}

	fromString := strconv.FormatUint(from, 10)
	toString := strconv.FormatUint(to, 10)

	return fileName[0] + "_" + fromString + "_" + toString + "." + fileName[1]
}
