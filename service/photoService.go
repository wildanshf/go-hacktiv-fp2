package service

import (
	"fmt"
	"project-2/config"
	"project-2/model"
)

func GetPhotoDetail(id int) (*model.Photo, error) {
	var (
		photo model.Photo
		db    = config.GetDB()
	)

	err := db.Model(&model.Photo{}).Where("id = ?", id).First(&photo).Error

	return &photo, err
}

func CreatePhoto(photo model.Photo) (*model.Photo, error) {
	var (
		db = config.GetDB()
	)

	err := db.Model(&model.Photo{}).Omit("User").Create(&photo).Error

	return &photo, err
}

func GetPhotos(userID int) ([]*model.Photo, error) {
	var (
		photos []*model.Photo
		db     = config.GetDB()
	)

	err := db.Model(&model.Photo{}).Preload("User").Where("user_id = ?", userID).Find(&photos).Error

	return photos, err
}

func UpdatePhoto(photo model.Photo, userID int) (*model.Photo, error) {
	var (
		db = config.GetDB()
	)

	res := db.Model(&model.Photo{}).Omit("User").Where("id = ? AND user_id = ?", photo.ID, userID).Updates(&photo)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("user not allowed to edit this photo")
	}

	return GetPhotoDetail(photo.ID)
}

func DeletePhoto(id, userID int) (int, error) {
	var (
		db = config.GetDB()
	)

	res := db.Model(&model.Photo{}).Where("id = ? AND user_id = ?", id, userID).Delete(&model.Photo{})
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, fmt.Errorf("user not allowed to delete this photo")
	}

	return id, nil
}
