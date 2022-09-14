package routes

import (
	"crud/controller"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.Group("v1")
	e.GET("/users", controller.GetUserControllers)
	e.GET("/users/:id", controller.GetUserByIDControllers)
	e.POST("/users/", controller.CreateUserControllers)
	e.PUT("/users/:id", controller.UpdateUserControllers)
	e.DELETE("/users/:id", controller.DeleteUserControllers)

	e.GET("/books", controller.GetBookControllers)
	e.GET("/books/:id", controller.GetBookByIDControllers)
	e.POST("/books/", controller.CreateBookControllers)
	e.PUT("/books/:id", controller.UpdateBookByIDControllers)
	e.DELETE("/books/:id", controller.DeleteBookByIDControllers)
	return e

}
