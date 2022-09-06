package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal/dto"
	"golang_web_programming/internal/service"
	"net/http"
)

type MemberController struct {
	service service.MemberService
}

func NewMemberController(service service.MemberService) *MemberController {
	return &MemberController{service: service}
}
func (m *MemberController) Create(c echo.Context) error {
	var req dto.CreateRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("%v :%v", ErrInvalidRequest, err)
	}
	if !req.IsValid() {
		return ErrInvalidRequest
	}

	createRes, err := m.service.Create(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, &createRes)
}

func (m *MemberController) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrPathValue
	}
	var req dto.UpdateRequestBody
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("%v :%v", ErrInvalidRequest, err)
	}
	if !req.IsValid() {
		return ErrInvalidRequest
	}

	updateRes, err := m.service.Update(dto.UpdateRequest{
		ID:                id,
		UpdateRequestBody: req,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &updateRes)
}

func (m *MemberController) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrPathValue
	}

	getRes, err := m.service.Get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &getRes)
}

func (m *MemberController) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return ErrPathValue
	}

	err := m.service.Delete(id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}
