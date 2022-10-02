package tracker

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

// TestGenerateFileName verifies logic for
// generating test file name
func TestGenerateFileName(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		name           string
		outputFilePath string
		from           uint64
		to             uint64
		result         string
	}{
		{
			"correct file format",
			"results.txt",
			10,
			15,
			"results_10_15.txt",
		},
		{
			"incorrect file format, print to outputFilePath",
			"resultstest",
			10,
			15,
			"resultstest",
		},
	}
	for _, testCase := range testTable {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			blockTrackerResult := newBlockTrackerResults(
				hclog.NewNullLogger(),
				testCase.outputFilePath,
			)

			fileName := blockTrackerResult.getTestFileName(
				testCase.from,
				testCase.to,
			)

			assert.Equal(t, fileName, testCase.result)
		})
	}
}
