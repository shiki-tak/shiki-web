package server

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shiki-tak/shiki-web/server"
)

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = "1323"
)

func CMD() *cobra.Command {
	serverCMD := &cobra.Command{
		Use:              "server",
		TraverseChildren: true,
		Short:            "subcommand for generating server",
	}

	serverCMD.AddCommand(
		startCmd(),
	)
	return serverCMD
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		TraverseChildren: true,
		Use:              "start [port]",
		Short:            fmt.Sprintf("server start with specific port(default:%s)", DefaultPort),
		Args:             cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			port := DefaultPort
			if len(args) == 1 {
				port = args[0]
			}

			server.Init(port)
			return nil
		},
	}
	return cmd
}
