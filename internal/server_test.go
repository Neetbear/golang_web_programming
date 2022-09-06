package internal

import (
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
	"golang_web_programming/internal/controller"
	"golang_web_programming/internal/dto"
	"golang_web_programming/internal/repository"
	"golang_web_programming/internal/service"
	"net/http"
	"testing"
)

func New() *controller.MemberController {
	repo := repository.NewRepository(map[string]repository.Membership{})
	serv := service.NewApplication(repo)
	cont := controller.NewMemberController(serv)
	return cont
}
func Test_CreateMembership(t *testing.T) {
	echoServer := NewServer(New())
	echoServer.Routes()
	echoServer.ErrorHandler(NewHttpErrorHandler(map[error]int{
		controller.ErrPathValue:      http.StatusBadRequest,
		controller.ErrInvalidRequest: http.StatusBadRequest,
		service.ErrCantSignUp:        http.StatusBadRequest,
		repository.ErrDuplicateName:  http.StatusBadRequest,
		repository.ErrNotExist:       http.StatusBadRequest,
	}).Handler)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(echoServer.e),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Run("Create Membership", func(t *testing.T) {
		createdResult := e.POST("/v1/memberships").
			WithJSON(dto.CreateRequest{
				UserName:       "zico",
				MembershipType: "naver",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		e.DELETE("/v1/memberships/" + createdResult.Value("ID").Raw().(string)).
			Expect().
			Status(http.StatusOK)

		actual := e.POST("/v1/memberships").
			WithJSON(dto.CreateRequest{
				UserName:       "zico",
				MembershipType: "naver",
			}).
			Expect().
			Status(http.StatusBadRequest).
			JSON().
			Object().Raw()

		assert.Equal(t, service.ErrCantSignUp.Error(), actual["message"])
	})
}
