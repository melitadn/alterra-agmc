package static

import "crud/models"

var BookInit BookMap

type BookMap struct {
	BookMap map[int64]models.Book
}

func Init() BookMap {
	p := make(map[int64]models.Book)

	return BookMap{BookMap: p}
}
func GetAllBook() BookMap {
	return BookInit
}
func GetBookByID(id int64) models.Book {
	return BookInit.BookMap[id]
}
func CreateBook(id int64, book models.Book) {
	BookInit.BookMap[id] = book
}
func UpdateBook(id int64, name string) {
	currentBook := BookInit.BookMap[id]
	currentBook.Name = name
	BookInit.BookMap[id] = currentBook
}
func DeleteBookByID(id int64) {
	delete(BookInit.BookMap, id)
}
