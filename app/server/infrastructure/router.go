package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	climbed "github.com/shiki-tak/shiki-web/app/server/interfaces/controllers"
)

const (
	v1             = "/v1"
	climbedIdParam = "/:" + climbed.Id
	createPath     = v1 + climbed.Path
	editPath       = v1 + climbed.Path + climbedIdParam
	showPath       = v1 + climbed.Path + climbedIdParam
	deletePath     = v1 + climbed.Path + climbedIdParam

	mongoURL = "mongodb://localhost:27017"
)

func Init(port string) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// request-id=c.Response().Header().Get(echo.HeaderXRequestID)
	e.Use(middleware.RequestID())

	mongoDBhandler, err := NewMongoDBHandler(mongoURL)
	if err != nil {
		panic(err)
	}
	climbedController := climbed.NewClimbedController(mongoDBhandler)

	api := e.Group("/api")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Shiki-Web!!")
	})

	api.POST(createPath, climbedController.Create)
	api.PUT(editPath, climbedController.Edit)
	api.GET(showPath, climbedController.Show)
	api.DELETE(deletePath, climbedController.Delete)

	e.Logger.Fatal(e.Start(":" + port))
}
