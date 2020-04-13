package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shiki-tak/shiki-web/app/server/interfaces/database"
	"github.com/shiki-tak/shiki-web/app/server/usecase"
)

const (
	Path = "/climbed"
	Id   = "id"
)

type ClimbedController struct {
	Interactor usecase.ClimbedMountainInteractor
}

func NewClimbedController(mongoDBHandler database.MongoDBHandler) *ClimbedController {
	return &ClimbedController{
		Interactor: usecase.ClimbedMountainInteractor{
			ClimbedRepository: &database.ClimbedMountainRepository{
				MongoDBHandler: mongoDBHandler,
			},
		},
	}
}

func (p *ClimbedController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "create api")
}

func (p *ClimbedController) Edit(c echo.Context) error {
	return c.JSON(http.StatusOK, "edit api")
}

func (p *ClimbedController) Show(c echo.Context) error {
	return c.JSON(http.StatusOK, "sho api")
}

func (p *ClimbedController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "delete")
}
