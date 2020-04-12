package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
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
			// Echo instance
			e := echo.New()

			// Middleware
			e.Use(middleware.Logger())
			e.Use(middleware.Recover())

			// request-id=c.Response().Header().Get(echo.HeaderXRequestID)
			e.Use(middleware.RequestID())
			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "Hello, Shiki-Web!!")
			})

			e.Logger.Fatal(e.Start(":" + port))
			return nil
		},
	}
	return cmd
}
