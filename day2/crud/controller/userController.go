package controller

import (
	"crud/lib/database"
	"crud/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func CreateUserControllers(c echo.Context) error {
	reqUser := models.RequestUser{}
	c.Bind(reqUser)
	e := database.CreateUser(reqUser)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success creating new user",
	})

}
func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})

}
func GetUserByIDControllers(c echo.Context) error {
	id := c.Param("id")
	user, e := database.GetUserById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"meesage": id,
		"users":   user,
	})

}
func UpdateUserControllers(c echo.Context) error {
	reqUser := models.RequestUser{}
	c.Bind(reqUser)
	e := database.UpdateUser(reqUser)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})

}
func DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	e := database.DeleteUser(uint(idInt))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"meesage": id,
	})

}
