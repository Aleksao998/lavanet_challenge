package server

import (
	"fmt"

	"github.com/Aleksao998/lavanet_challenge/command"
	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/Aleksao998/lavanet_challenge/server"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	showCmd := &cobra.Command{
		Use:     "server",
		Short:   "The default command that starts lavanet_challenge client",
		PreRunE: runPreRun,
		Run:     runCommand,
	}

	setFlags(showCmd)

	return showCmd
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&params.networkGrpcAddressRaw,
		networkGrpcAddressFlag,
		fmt.Sprintf("%s:%d", command.OsmosisMainnetGrpcEndpoint, command.OsmosisMainnetGrpcPort),
		"network gRPC endpoint",
	)
	cmd.Flags().StringVar(
		&params.logLevel,
		logLevelFlag,
		"INFO",
		"the log level for console output",
	)
	cmd.Flags().StringVar(
		&params.logFileLocation,
		logFileLocationFlag,
		"",
		"write all logs to the file at specified location instead of writing them to console",
	)
	cmd.Flags().StringVar(
		&params.grpcAddressRaw,
		grpcAddressFlag,
		fmt.Sprintf("%s:%d", command.LocalHostBinding, command.DefaultGRPCPort),
		"the GRPC interface",
	)
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter()
	if err := runServerLoop(params.generateConfig(), outputter); err != nil {
		outputter.SetError(err)
		outputter.WriteOutput()

		return
	}
}

func runPreRun(cmd *cobra.Command, _ []string) error {
	return params.initRawParams()
}

func runServerLoop(
	config *server.Config,
	outputter command.OutputFormatter,
) error {
	serverInstance, err := server.NewServer(config)
	if err != nil {
		return err
	}

	return helper.HandleSignals(serverInstance.Close, outputter)
}
