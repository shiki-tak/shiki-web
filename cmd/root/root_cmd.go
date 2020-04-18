package root

import (
	"github.com/shiki-tak/shiki-web/cmd/server"
	"github.com/spf13/cobra"
)

func CMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shiki-web",
		Short: "Command line interface for shiki-web",
	}
	// Construct Root Command
	cmd.AddCommand(
		server.CMD(),
	)
	return cmd
}
