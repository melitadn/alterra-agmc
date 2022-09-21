package middleware

import (
	"crud/internal/config"
	"crud/pkg/models"
	"github.com/labstack/echo"
)

var db = config.DB

func BasicAuthDB(username, password string, e echo.Context) (
	bool, error) {
	var db = config.DB
	var user models.User
	if err := db.Where("email=? AND password=?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil

}
