package server

import "github.com/hashicorp/go-hclog"

// Config is used to parametrize the lavanet_challenge client
type Config struct {
	// network gRPC endpoint
	Network string

	// gRPC server port
	Port uint64

	// logLevel represent a log type [ERROR, INFO, DEBUG]
	LogLevel hclog.Level

	// logFileLocation location of log file
	LogFilePath string
}
