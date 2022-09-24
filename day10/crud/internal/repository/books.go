package repository

import (
	"crud/pkg/models"
)

var BookInit BookMap

type BookMap struct {
	BookMap map[int64]models.Book
}
type Book interface {
	GetAllBook() BookMap
	GetBookByID(id int64) models.Book
	CreateBook(id int64, book models.Book)
	UpdateBook(id int64, name string)
	DeleteBookByID(id int64)
}
type book struct {
	BookInit BookMap
}

func NewBookRepository() *book {
	p := make(map[int64]models.Book)
	BookInit.BookMap = p
	return &book{
		BookInit: BookInit,
	}
}

func (b *book) GetAllBook() BookMap {
	return BookInit
}
func (b *book) GetBookByID(id int64) models.Book {
	return b.BookInit.BookMap[id]
}
func (b *book) CreateBook(id int64, book models.Book) {
	b.BookInit.BookMap[id] = book
}
func (b *book) UpdateBook(id int64, name string) {
	currentBook := b.BookInit.BookMap[id]
	currentBook.Name = name
	b.BookInit.BookMap[id] = currentBook
}
func (b *book) DeleteBookByID(id int64) {
	delete(b.BookInit.BookMap, id)
}
