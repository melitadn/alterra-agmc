package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ShowStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
func main() {
	e := echo.New()
	e.GET("/status", ShowStatus)
}
