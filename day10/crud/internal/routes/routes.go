package routes

import (
	"crud/internal/app/books"
	controller2 "crud/internal/app/users"
	"crud/internal/factory"
	"crud/pkg/constants"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func New(f *factory.Factory) *echo.Echo {
	e := echo.New()
	e.Group("v1")
	eAuth := e.Group("/jwt")
	eAuth.Use(mw.JWT([]byte(constants.SECRET_JWT)))
	handlerUser := controller2.NewHandler(f)
	eAuth.GET("/users", handlerUser.GetUserControllers)
	eAuth.GET("/users/:id", handlerUser.GetUserByIDControllers)
	e.POST("/users/", handlerUser.CreateUserControllers)
	eAuth.PUT("/users/:id", handlerUser.UpdateUserControllers)
	eAuth.DELETE("/users/:id", handlerUser.DeleteUserControllers)

	handlerBook := books.NewHandler(f)
	e.GET("/books", handlerBook.GetBookControllers)
	e.GET("/books/:id", handlerBook.GetBookByIDControllers)
	eAuth.POST("/books/", handlerBook.CreateBookControllers)
	eAuth.PUT("/books/:id", handlerBook.UpdateBookByIDControllers)
	eAuth.DELETE("/books/:id", handlerBook.DeleteBookByIDControllers)

	e.POST("/login", handlerUser.LoginUserController)
	e.GET("/users/login", handlerUser.GetUserDetailControllers)
	return e

}
