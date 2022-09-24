package users

import (
	"crud/internal/factory"
	"crud/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}
func (h *handler) CreateUserControllers(c echo.Context) error {
	reqUser := models.RequestUser{}
	c.Bind(reqUser)
	e := h.service.CreateUser(reqUser)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success creating new user",
	})

}
func (h *handler) GetUserControllers(c echo.Context) error {
	users, e := h.service.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})

}
func (h *handler) GetUserByIDControllers(c echo.Context) error {
	id := c.Param("id")
	user, e := h.service.GetUserById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"meesage": id,
		"users":   user,
	})

}
func (h *handler) UpdateUserControllers(c echo.Context) error {
	reqUser := models.RequestUser{}
	c.Bind(reqUser)
	e := h.service.UpdateUser(reqUser)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})

}
func (h *handler) DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	e := h.service.DeleteUser(uint(idInt))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"meesage": id,
	})

}

func (h *handler) LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := h.service.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})

}

func (h *handler) GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := h.service.GetDetailUsers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
