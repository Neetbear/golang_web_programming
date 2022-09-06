package internal

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type httpErrorHandler struct {
	codes map[error]int
}

func NewHttpErrorHandler(codes map[error]int) *httpErrorHandler {
	return &httpErrorHandler{codes: codes}
}

func (h *httpErrorHandler) getStatusCode(err error) int {
	for key, value := range h.codes {
		if errors.Is(err, key) {
			return value
		}
	}
	return http.StatusInternalServerError
}

func (self *httpErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    self.getStatusCode(err),
			Message: err.Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = map[string]interface{}{"message": err.Error()}
	}

	if !c.Response().Committed {
		err = c.JSON(code, message)
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
