package routes

import (
	"crud/constants"
	"crud/controller"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Group("v1")
	eAuth := e.Group("/jwt")
	eAuth.Use(mw.JWT([]byte(constants.SECRET_JWT)))
	eAuth.GET("/users", controller.GetUserControllers)
	eAuth.GET("/users/:id", controller.GetUserByIDControllers)
	e.POST("/users/", controller.CreateUserControllers)
	eAuth.PUT("/users/:id", controller.UpdateUserControllers)
	eAuth.DELETE("/users/:id", controller.DeleteUserControllers)

	e.GET("/books", controller.GetBookControllers)
	e.GET("/books/:id", controller.GetBookByIDControllers)
	eAuth.POST("/books/", controller.CreateBookControllers)
	eAuth.PUT("/books/:id", controller.UpdateBookByIDControllers)
	eAuth.DELETE("/books/:id", controller.DeleteBookByIDControllers)

	e.POST("/login", controller.LoginUserController)
	e.GET("/users/login", controller.GetUserDetailControllers)
	return e

}
