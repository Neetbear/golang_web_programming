package internal

import (
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal/controller"
	"log"
)

type Server struct {
	controller *controller.MemberController
	e          *echo.Echo
}

func NewServer(controller *controller.MemberController) *Server {
	return &Server{controller: controller, e: echo.New()}
}

func (s *Server) Run(port string) {
	s.Routes()
	s.Logger(BodyLogger(), BasicLogger())
	log.Fatalln(s.e.Start(":" + port))
}

func (s *Server) Routes() {
	g := s.e.Group("/v1")
	s.routeMemberships(g)
}
func (s *Server) routeMemberships(e *echo.Group) {
	e.POST("/memberships", s.controller.Create)
	e.PUT("/memberships/:id", s.controller.Update)
	e.GET("/memberships/:id", s.controller.Get)
	e.DELETE("/memberships/:id", s.controller.Delete)
}

func (s *Server) Logger(middlewares ...echo.MiddlewareFunc) {
	for _, middleware := range middlewares {
		s.e.Use(middleware)
	}
}

func (s *Server) ErrorHandler(handler echo.HTTPErrorHandler) {
	s.e.HTTPErrorHandler = handler
}
