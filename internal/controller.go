package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	// binding
	var request CreateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.service.Create(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) GetByID(c echo.Context) error {
	// binding
	id := c.Param("id")
	res, err := controller.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) Update(c echo.Context) error {
	// binding
	var request UpdateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.service.Update(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) DeleteById(c echo.Context) error {
	// binding
	id := c.Param("id")
	err := controller.service.DeleteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
