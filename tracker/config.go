package tracker

import (
	"net"

	"github.com/hashicorp/go-hclog"
)

// Config is used to parametrize the tracker client
type Config struct {
	// NetworkGrpcAddress is network gRPC endpoint
	ClientGrpcAddress *net.TCPAddr

	// logLevel represent a log type [ERROR, INFO, DEBUG]
	LogLevel hclog.Level

	// logFileLocation location of log file
	LogFilePath string

	// pollingTime is a polling time in seconds
	PollingTime uint64
}
