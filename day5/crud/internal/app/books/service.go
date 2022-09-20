package books

import (
	"crud/internal/factory"
	"crud/internal/repository"
	"crud/pkg/models"
)

type service struct {
	BookRepository repository.Book
}

type Service interface {
	GetAllBook() repository.BookMap
	GetBookByID(id int64) models.Book
	CreateBook(id int64, book models.Book)
	UpdateBook(id int64, name string)
	DeleteBookByID(id int64)
}

func NewService(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (u *service) GetAllBook() repository.BookMap {
	return u.BookRepository.GetAllBook()

}
func (u *service) GetBookByID(id int64) models.Book {
	return u.BookRepository.GetBookByID(id)

}
func (u *service) CreateBook(id int64, book models.Book) {
	u.BookRepository.CreateBook(id, book)

}
func (u *service) UpdateBook(id int64, name string) {
	u.BookRepository.UpdateBook(id, name)

}
func (u *service) DeleteBookByID(id int64) {
	u.BookRepository.DeleteBookByID(id)

}
