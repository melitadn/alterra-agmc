package controller

import (
	"crud/lib/static"
	"crud/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func GetBookControllers(c echo.Context) error {
	book := static.GetAllBook()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  book,
	})

}

func GetBookByIDControllers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	book := static.GetBookByID(idInt)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  book,
	})

}
func DeleteBookByIDControllers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	static.DeleteBookByID(idInt)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     id,
	})

}
func UpdateBookByIDControllers(c echo.Context) error {
	id := c.Param("id")
	name := c.Param("name")
	idInt, err := strconv.ParseInt(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	static.UpdateBook(idInt, name)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     id,
	})

}
func CreateBookControllers(c echo.Context) error {
	name := c.Param("name")
	book := models.Book{
		Name:      name,
		CreatedAt: time.Now(),
	}

	static.CreateBook(rand.Int63n(2000), book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     name,
	})

}
