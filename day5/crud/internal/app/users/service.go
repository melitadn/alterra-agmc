package users

import (
	"crud/internal/factory"
	"crud/internal/repository"
	"crud/pkg/models"
)

type service struct {
	UserRepository repository.UserInt
}

type Service interface {
	GetUsers() (interface{}, error)
	GetUserById(id string) (interface{}, error)
	CreateUser(userReq models.RequestUser) error
	UpdateUser(userReq models.RequestUser) error
	DeleteUser(id uint) error
	LoginUsers(user *models.User) (interface{}, error)
	GetDetailUsers(userId int) (interface{}, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (u *service) GetUsers() (interface{}, error) {
	return u.UserRepository.GetUsers()

}
func (u *service) GetUserById(id string) (interface{}, error) {
	return u.UserRepository.GetUserById(id)

}
func (u *service) CreateUser(userReq models.RequestUser) error {
	return u.UserRepository.CreateUser(userReq)

}
func (u *service) UpdateUser(userReq models.RequestUser) error {
	return u.UserRepository.UpdateUser(userReq)

}
func (u *service) DeleteUser(id uint) error {
	return u.UserRepository.DeleteUser(id)

}
func (u *service) LoginUsers(user *models.User) (interface{}, error) {
	return u.UserRepository.LoginUsers(user)

}
func (u *service) GetDetailUsers(userId int) (interface{}, error) {
	return u.UserRepository.GetDetailUsers(userId)

}
