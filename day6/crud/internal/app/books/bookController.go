package books

import (
	"crud/internal/factory"
	"crud/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}
func (h *handler) GetBookControllers(c echo.Context) error {
	book := h.service.GetAllBook()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  book,
	})

}

func (h *handler) GetBookByIDControllers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	book := h.service.GetBookByID(idInt)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  book,
	})

}
func (h *handler) DeleteBookByIDControllers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	h.service.DeleteBookByID(idInt)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     id,
	})

}
func (h *handler) UpdateBookByIDControllers(c echo.Context) error {
	id := c.Param("id")
	name := c.Param("name")
	idInt, err := strconv.ParseInt(id, 2, 64)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	h.service.UpdateBook(idInt, name)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     id,
	})

}
func (h *handler) CreateBookControllers(c echo.Context) error {
	name := c.Param("name")
	book := models.Book{
		Name:      name,
		CreatedAt: time.Now(),
	}

	h.service.CreateBook(rand.Int63n(2000), book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     name,
	})

}
