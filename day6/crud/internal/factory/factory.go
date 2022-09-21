package factory

import (
	"crud/internal/config"
	"crud/internal/repository"
)

type Factory struct {
	UserRepository repository.UserInt
	BookRepository repository.Book
}

func NewFactory() *Factory {
	db := config.DB
	return &Factory{
		repository.NewUserRepository(db),
		repository.NewBookRepository(),
	}
}
