package tracker

import (
	"fmt"

	"github.com/Aleksao998/lavanet_challenge/command"
	"github.com/Aleksao998/lavanet_challenge/command/helper"
	"github.com/Aleksao998/lavanet_challenge/tracker"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	tracker := &cobra.Command{
		Use:     "tracker",
		Short:   "The default command that starts tracker client",
		PreRunE: runPreRun,
		Run:     runCommand,
	}

	setFlags(tracker)

	return tracker
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&params.clientGrpcAddressRaw,
		clientGrpcAddressFlag,
		fmt.Sprintf("%s:%d", command.LocalHostBinding, command.DefaultGRPCPort),
		"client gRPC endpoint",
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
		&params.outputFileLocation,
		outputFileLocationFlag,
		"test_results.txt",
		"write tracker data to the file at specified location",
	)
	cmd.Flags().Uint64Var(
		&params.pollingTime,
		pollingTimeFlag,
		2,
		"polling time in seconds",
	)
	cmd.Flags().Uint64Var(
		&params.outputAfter,
		outputAfterFlag,
		5,
		"number after which results will be generated",
	)
}

func runPreRun(cmd *cobra.Command, _ []string) error {
	// init raw params
	return params.initRawParams()
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter()
	if err := runServerLoop(params.generateConfig(), outputter); err != nil {
		outputter.SetError(err)
		outputter.WriteOutput()

		return
	}
}

func runServerLoop(
	config *tracker.Config,
	outputter command.OutputFormatter,
) error {
	trackerInstance, err := tracker.NewTracker(config)
	if err != nil {
		return err
	}

	return helper.HandleSignals(trackerInstance.Close, outputter)
}
