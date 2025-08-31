package services

import (
	"errors"
	"fiber/api/models"
	"fiber/common/types"
	"fiber/configs"
	"strconv"
	"gorm.io/gorm"
)

func GetUserService() []models.User {
	var users []models.User
	configs.Db.Find(&users)
	return users
}

func GetUserByIdService(id int) (*models.User, error) {
	user := &models.User{}
	if err := configs.Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUserService(data types.CreateUser) (*models.User, error) {
	user := &models.User{
		UserName: data.UserName,
		Email:    data.Email,
		Password: data.Password,
	}

	var count int64
	if err := configs.Db.Model(&models.User{}).
		Where("Email = ?", data.Email).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("user already exist")
	}

	if err := configs.Db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUserById(id string) (*models.User, error) {
	user := &models.User{}
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	if err := configs.Db.
		First(&user, "user_id = ?", userId).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := configs.Db.
		Delete(&user).
		Error; err != nil {
		return nil, err
	}

	return user, nil

}

func DeleteAllUsers() (*[]models.User, error) {
	users := &[]models.User{}

	if err := configs.Db.
		Find(&users).Error; err != nil {
		return nil, err
	}

	if len(*users) == 0 {
		return nil, errors.New("user not found")
	}

	if err := configs.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{}).Error; err != nil {
		return nil, err
	}

	return users, nil
}
