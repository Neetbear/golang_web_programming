package internal

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const _defaultPort = 8080

type Server struct {
	controller Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	service := NewService(*NewRepository(data))
	controller := NewController(*service)
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(BasicLogger())
	e.Use(BodyLogger())
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.controller)
}

func RouteMemberships(e *echo.Group, c Controller) {
	e.GET("/memberships/:id", c.GetByID)       // GET    /v1/memverships/:id
	e.POST("/memberships", c.Create)           // POST   /v1/memverships
	e.PUT("/merberships", c.Update)            // Update /v1/memverships
	e.DELETE("/memberships/:id", c.DeleteById) // Delete /v1/memverships/:id
}

// 1. e.POST("/memberships", c.Create, middleware.RequestID)
//// 에코 내장 미들웨어 사용

// 2. “X-My-Request-Header” 헤더
//// middleware.RequestIDWithConfig && middleware.RequestIDConfig

// 3. 미들웨어 함수 생성하여 사용
// func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Request().Header.Set(echo.HeaderXRequestID, uuid.New().String())
// 		return next(c)
// 	}
// }
