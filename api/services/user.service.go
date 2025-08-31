package services

import (
	"fiber/api/models"
	"fiber/common/types"
	"fiber/configs"
)

func GetUserService() []models.User {
	var users []models.User
	configs.Db.Find(&users)
	return users
}

func GetUserByIdService(id int) models.User {
	var user models.User
	configs.Db.Find(&user, id)
	return user
}

func CreateUserService(data types.CreateUser) (*models.User, error) {
	user := &models.User{
		UserName: data.UserName,
		Email:    data.Email,
		Password: data.Password,
	}

	result := configs.Db.Create(&user)
	if result.Error != nil {
		return   nil ,result.Error
	}
	
	return user, nil
}
