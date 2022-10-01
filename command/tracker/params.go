package tracker

import (
	"net"

	"github.com/Aleksao998/lavanet_challenge/command"
	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/Aleksao998/lavanet_challenge/tracker"
	"github.com/hashicorp/go-hclog"
)

const (
	clientGrpcAddress   = "client-grpc-address"
	logFileLocationFlag = "log-to"
	logLevelFlag        = "log-level"
	pollingTime         = "polling-time"
)

var (
	params = &trackerParams{}
)

type trackerParams struct {
	// clientGrpcAddress is client gRPC endpoint
	clientGrpcAddress *net.TCPAddr

	// clientGrpcAddressRaw is a raw network gRPC endpoint
	clientGrpcAddressRaw string

	// logLevel is a log type [ERROR, INFO, DEBUG]
	logLevel string

	// logFileLocation location of log file
	logFileLocation string

	// pollingTime is a polling time in seconds
	pollingTime uint64
}

func (p *trackerParams) initRawParams() error {
	return p.initGRPCAddress()
}

func (p *trackerParams) initGRPCAddress() error {
	var parseErr error

	if p.clientGrpcAddress, parseErr = helper.ResolveAddr(
		p.clientGrpcAddressRaw,
		command.LocalHostBinding,
	); parseErr != nil {
		return parseErr
	}

	return nil
}

func (p *trackerParams) generateConfig() *tracker.Config {
	return &tracker.Config{
		ClientGrpcAddress: p.clientGrpcAddress,
		PollingTime:       p.pollingTime,
		LogLevel:          hclog.LevelFromString(p.logLevel),
		LogFilePath:       p.logFileLocation,
	}
}
