package service

import (
	"project-2/config"
	"project-2/model"
	"project-2/tool"
	"time"
)

func IsLoginAllowed(param model.LoginParam) (bool, *model.User) {
	user, err := GetUserByEmail(param.Email)
	if err != nil {
		return false, nil
	}

	return tool.CheckPasswordHash(param.Password, user.Password), user
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	db := config.GetDB()

	err := db.Model(&model.User{}).Where("email = ?", email).First(&user).Error

	return &user, err
}

func CreateUser(user model.User) (*model.User, error) {
	db := config.GetDB()

	hashedPassword := tool.HashPassword(user.Password)
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	err := db.Model(&model.User{}).Create(&user).Error

	return &user, err
}

func GetUserDetail(id int) (*model.User, error) {
	var (
		user model.User
		db   = config.GetDB()
	)

	err := db.Model(&model.User{}).Where("id = ?", id).First(&user).Error

	return &user, err
}

func UpdateUser(user model.User) (*model.User, error) {
	var (
		currentTime = time.Now()
		db          = config.GetDB()
	)

	user.UpdatedAt = &currentTime

	err := db.Model(&model.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return GetUserDetail(user.ID)
}

func DeleteUser(id int) (int, error) {
	db := config.GetDB()

	err := db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error

	return id, err
}
