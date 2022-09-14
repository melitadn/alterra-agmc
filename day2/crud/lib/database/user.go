package database

import (
	"crud/config"
	"crud/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e

	}
	return users, nil

}

func GetUserById(id string) (interface{}, error) {
	var users models.User
	if e := config.DB.Where("id=?", id).Error; e != nil {
		return nil, e

	}
	return users, nil

}

func CreateUser(userReq models.RequestUser) error {
	users := models.User{
		Email:    userReq.Email,
		Password: userReq.Password,
		Name:     userReq.Name,
	}
	if e := config.DB.Create(&users).Error; e != nil {
		return e

	}
	return nil

}

func UpdateUser(userReq models.RequestUser) error {
	users := models.User{
		Email:    userReq.Email,
		Password: userReq.Password,
		Name:     userReq.Name,
	}
	if e := config.DB.Model(&users).Where("email=?", users.Email).Update("name", users.Name).Error; e != nil {
		return e
	}
	return nil

}
func DeleteUser(id uint) error {
	if e := config.DB.Delete(&models.User{}, id).Error; e != nil {
		return e
	}
	return nil

}
