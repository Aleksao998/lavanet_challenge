package tracker

import (
	"fmt"

	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/hashicorp/go-hclog"
)

// Tracker is the tracker service for the lavanet_challenge
type Tracker struct {
	logger hclog.Logger

	// blockTracker tracks latest block
	blockTracker *blockTracker

	// config server config
	config *Config
}

// NewTracker creates a new tracker service, using the passed in configuration
func NewTracker(config *Config) (*Tracker, error) {
	// initialize logger
	logger, err := helper.NewLoggerFromConfig(config.LogLevel, config.LogFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not setup new logger instance, %w", err)
	}

	// initialize server
	tracker := &Tracker{
		logger: logger,
		config: config,
	}

	// initialize block tracker
	blockTracker, err := NewBlockTracker(
		tracker.logger,
		tracker.config.ClientGrpcAddress,
		tracker.config.PollingTime,
		tracker.config.OutputFilePath,
		tracker.config.OutputAfter,
	)
	if err != nil {
		return nil, err
	}

	// assign block tracker to the tracker
	tracker.blockTracker = blockTracker

	// setup and start block tracker
	if err := blockTracker.Start(); err != nil {
		return nil, err
	}

	return tracker, nil
}

// Close closes the tracker service
func (s *Tracker) Close() {
	s.logger.Debug("Closing tracker")

	s.blockTracker.Close()
}
