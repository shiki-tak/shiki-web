package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shiki-tak/shiki-web/app/server/domain"
	"github.com/shiki-tak/shiki-web/app/server/interfaces/database"
	"github.com/shiki-tak/shiki-web/app/server/usecase"
)

const (
	Path = "/climbed_mountains"
	Id   = "id"
)

type ClimbedController struct {
	Interactor usecase.ClimbedMountainInteractor
}

func NewClimbedController(sqlHandler database.SqlHandler) *ClimbedController {
	return &ClimbedController{
		Interactor: usecase.ClimbedMountainInteractor{
			ClimbedRepository: &database.ClimbedMountainRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (cc *ClimbedController) Create(c echo.Context) error {
	cm := domain.ClimbedMountain{}
	if err := c.Bind(&cm); err != nil {
		return err
	}

	err := cc.Interactor.Add(cm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// TODO: implement
func (cc *ClimbedController) Edit(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (cc *ClimbedController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	mountain, err := cc.Interactor.ClimbedMountainById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, mountain)
}

// TODO: implement
func (cc *ClimbedController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (cc *ClimbedController) Gets(c echo.Context) error {
	mountains, err := cc.Interactor.Gets()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, mountains)
}
