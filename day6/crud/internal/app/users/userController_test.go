package users_test

import (
	"crud/internal/config"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func InitEcho() *echo.Echo {
	//Setup
	config.InitDB()
	e := echo.New()
	return e

}

func TestGetUserControllers(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "berhasil",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		if assert.NoError(t, GetUserControllers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))

		}
	}
}

func GetUserByIDControllers(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "failed",
			path:                 "/users/:id",
			expectBodyStartsWith: "{\"status\":\"failed\",\"users\":[",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		if assert.Error(t, GetUserControllers(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))

		}
	}
}
