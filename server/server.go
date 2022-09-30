package server

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/go-hclog"
)

// Server is the central manager of the blockchain client
type Server struct {
	logger hclog.Logger
	config *Config
}

// NewServer creates a new lavanet_challenge server, using the passed in configuration
func NewServer(config *Config) (*Server, error) {
	// initialize logger
	logger, err := newLoggerFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not setup new logger instance, %w", err)
	}

	// initialize server
	server := &Server{
		logger: logger,
		config: config,
	}

	// TODO start services and remove dummy code
	server.logger.Info("", "network", server.config.Network)
	server.logger.Info("", "port", server.config.Port)
	go func() {
		for i := 1; i <= 5; i++ {
			server.logger.Info("INFO LOG", "num", i)
			server.logger.Error("ERROR LOG", "num", i)
			server.logger.Debug("DEBBUG LOG", "num", i)
			time.Sleep(3 * time.Second)
		}
	}()

	return server, nil
}

// newLoggerFromConfig creates a new logger which logs to a specified file or standard output.
func newLoggerFromConfig(config *Config) (hclog.Logger, error) {
	if config.LogFilePath != "" {
		fileLoggerInstance, err := newFileLogger(config)
		if err != nil {
			return nil, err
		}

		return fileLoggerInstance, nil
	}

	return newCLILogger(config), nil
}

// newFileLogger returns logger instance that writes all logs to a specified file.
func newFileLogger(config *Config) (hclog.Logger, error) {
	logFileWriter, err := os.Create(config.LogFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create log file, %w", err)
	}

	return hclog.New(&hclog.LoggerOptions{
		Name:   "lavanet_challenge",
		Level:  config.LogLevel,
		Output: logFileWriter,
	}), nil
}

// newCLILogger returns logger instance that sends all logs to standard output
func newCLILogger(config *Config) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:  "lavanet_challenge",
		Level: config.LogLevel,
	})
}

// Close closes the lavanet_challenge server
func (s *Server) Close() {}
