package controller

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang_web_programming/internal/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMember_Create_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected error
	}{
		{
			name:     "Casting fail",
			input:    "{\"id\": \"uuid\"}",
			expected: ErrInvalidRequest,
		},
	}
	controller := NewMemberController(nil)
	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshal, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/v1/memberships", bytes.NewBuffer(marshal))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			actual := controller.Create(c)

			assert.ErrorAs(t, tt.expected, &actual)
		})
	}
}

func TestMemberController_Update_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       string
		input    interface{}
		expected error
	}{
		{
			name:     "Casting fail",
			id:       "uuid",
			input:    "{\"id\": \"uuid\"}",
			expected: ErrInvalidRequest,
		},
		{
			name: "Not exist id",
			id:   "",
			input: dto.UpdateRequestBody{
				UserName:       "zico",
				MembershipType: "naver",
			},
			expected: ErrPathValue,
		},
	}
	controller := NewMemberController(nil)
	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshal, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPut, "/memberships/:id", bytes.NewBuffer(marshal))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.id)

			actual := controller.Update(c)

			assert.ErrorAs(t, tt.expected, &actual)
		})
	}
}

func TestMemberController_Get_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name:     "Not exist id",
			id:       "",
			expected: ErrPathValue,
		},
	}
	controller := NewMemberController(nil)
	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/memberships", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.id)

			actual := controller.Get(c)

			assert.ErrorAs(t, tt.expected, &actual)
		})
	}
}

func TestMemberController_Delete_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name:     "Not exist id",
			id:       "",
			expected: ErrPathValue,
		},
	}
	controller := NewMemberController(nil)
	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/v1/memberships", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.id)

			actual := controller.Get(c)

			assert.ErrorAs(t, tt.expected, &actual)
		})
	}
}

func TestValidCreateRequest_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    dto.CreateRequest
		expected error
	}{
		{
			name:     "Not enter username",
			input:    dto.CreateRequest{"", "naver"},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Not enter type",
			input:    dto.CreateRequest{"zico", ""},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Not enter all",
			input:    dto.CreateRequest{"", ""},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Bad type : kakao",
			input:    dto.CreateRequest{"zico", "kakao"},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Bad type : nh",
			input:    dto.CreateRequest{"zico", "nh"},
			expected: ErrInvalidRequest,
		},
	}
	controller := NewMemberController(nil)
	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshal, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/v1/memberships", bytes.NewBuffer(marshal))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			actual := controller.Create(c)

			assert.ErrorAs(t, tt.expected, &actual)
		})
	}
}

func Test_ValidUpdateRequest_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    dto.UpdateRequestBody
		expected error
	}{
		{
			name:     "Not enter username",
			input:    dto.UpdateRequestBody{"", "naver"},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Not enter type",
			input:    dto.UpdateRequestBody{"zico", ""},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Bad type : kakao",
			input:    dto.UpdateRequestBody{"zico", "kakao"},
			expected: ErrInvalidRequest,
		},
		{
			name:     "Bad type : nh",
			input:    dto.UpdateRequestBody{"zico", "nh"},
			expected: ErrInvalidRequest,
		},
	}
	controller := NewMemberController(nil)
	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshal, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPut, "/v1/memberships", bytes.NewBuffer(marshal))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			actual := controller.Update(c)

			assert.ErrorAs(t, tt.expected, &actual)
		})
	}
}
