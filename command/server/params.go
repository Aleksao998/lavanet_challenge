package server

import (
	"net"

	"github.com/Aleksao998/lavanet_challenge/command"
	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/Aleksao998/lavanet_challenge/server"
	"github.com/hashicorp/go-hclog"
)

const (
	networkGrpcAddress  = "network-grpc-address"
	logFileLocationFlag = "log-to"
	logLevelFlag        = "log-level"
	grpcAddress         = "grpc-address"
)

var (
	params = &serverParams{}
)

type serverParams struct {
	// networkGrpcAddress is network gRPC endpoint
	networkGrpcAddress *net.TCPAddr

	// networkGrpcAddressRaw is a raw network gRPC endpoint
	networkGrpcAddressRaw string

	// logLevel represent a log type [ERROR, INFO, DEBUG]
	logLevel string

	// logFileLocation location of log file
	logFileLocation string

	// grpcAddress is gRPC address of lavanet_challenge client
	grpcAddress *net.TCPAddr

	// grpcAddress is raw gRPC address of lavanet_challenge client
	grpcAddressRaw string
}

func (p *serverParams) initRawParams() error {
	return p.initGRPCAddresses()
}

func (p *serverParams) generateConfig() *server.Config {
	return &server.Config{
		NetworkGrpcAddress: p.networkGrpcAddress,
		GrpcAddress:        p.grpcAddress,
		LogLevel:           hclog.LevelFromString(p.logLevel),
		LogFilePath:        p.logFileLocation,
	}
}

func (p *serverParams) initGRPCAddresses() error {
	var parseErr error

	if p.grpcAddress, parseErr = helper.ResolveAddr(
		p.grpcAddressRaw,
		command.LocalHostBinding,
	); parseErr != nil {
		return parseErr
	}

	if p.networkGrpcAddress, parseErr = helper.ResolveAddr(
		p.networkGrpcAddressRaw,
		command.OsmosisMainnetGrpcEndpoint,
	); parseErr != nil {
		return parseErr
	}

	return nil
}
