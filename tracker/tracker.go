package tracker

import (
	"fmt"
	"time"

	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/hashicorp/go-hclog"
)

// Tracker is the tracker service for the lavanet_challenge
type Tracker struct {
	logger hclog.Logger

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

	// TODO start services and remove dummy code
	tracker.logger.Info("", "ClientGrpcAddress", tracker.config.ClientGrpcAddress)

	go func() {
		for i := 1; i <= 5; i++ {
			tracker.logger.Info("INFO LOG", "num", i)
			tracker.logger.Error("ERROR LOG", "num", i)
			tracker.logger.Debug("DEBBUG LOG", "num", i)
			time.Sleep(3 * time.Second)
		}
	}()

	return tracker, nil
}

// Close closes the tracker service
func (s *Tracker) Close() {
	s.logger.Debug("Closing tracker")
}
