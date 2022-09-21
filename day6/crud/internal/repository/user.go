package repository

import (
	"crud/internal/middleware"
	"crud/pkg/models"

	"gorm.io/gorm"
)

type UserInt interface {
	GetUsers() (interface{}, error)
	GetUserById(id string) (interface{}, error)
	CreateUser(userReq models.RequestUser) error
	UpdateUser(userReq models.RequestUser) error
	DeleteUser(id uint) error
	LoginUsers(user *models.User) (interface{}, error)
	GetDetailUsers(userId int) (interface{}, error)
}

type user struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{
		db,
	}
}
func (u *user) GetUsers() (interface{}, error) {
	var users []models.User
	if e := u.DB.Find(&users).Error; e != nil {
		return nil, e

	}
	return users, nil

}

func (u *user) GetUserById(id string) (interface{}, error) {
	var users models.User
	if e := u.DB.Where("id=?", id).Error; e != nil {
		return nil, e

	}
	return users, nil

}

func (u *user) CreateUser(userReq models.RequestUser) error {
	users := models.User{
		Email:    userReq.Email,
		Password: userReq.Password,
		Name:     userReq.Name,
	}
	if e := u.DB.Create(&users).Error; e != nil {
		return e

	}
	return nil

}

func (u *user) UpdateUser(userReq models.RequestUser) error {
	users := models.User{
		Email:    userReq.Email,
		Password: userReq.Password,
		Name:     userReq.Name,
	}
	if e := u.DB.Model(&users).Where("email=?", users.Email).Update("name", users.Name).Error; e != nil {
		return e
	}
	return nil

}
func (u *user) DeleteUser(id uint) error {
	if e := u.DB.Delete(&models.User{}, id).Error; e != nil {
		return e
	}
	return nil

}

func (u *user) LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = u.DB.Where("email=? AND password =?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := u.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if e := u.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil

}
