package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(port string) {
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
}
