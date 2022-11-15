package service

import (
	"fmt"
	"project-2/config"
	"project-2/model"
)

func GetSocialMediaDetail(id int) (*model.SocialMedia, error) {
	var (
		socialMedia model.SocialMedia
		db          = config.GetDB()
	)

	err := db.Model(&model.SocialMedia{}).Where("id = ?", id).First(&socialMedia).Error

	return &socialMedia, err
}

func AddPSocialMedia(socialMedia model.SocialMedia) (*model.SocialMedia, error) {
	var (
		db = config.GetDB()
	)

	err := db.Model(&model.SocialMedia{}).Omit("User").Create(&socialMedia).Error

	return &socialMedia, err
}

func GetSocialMedias(userID int) ([]*model.SocialMedia, error) {
	var (
		socialMedias []*model.SocialMedia
		db           = config.GetDB()
	)

	err := db.Model(&model.SocialMedia{}).Preload("User").Where("user_id = ?", userID).Find(&socialMedias).Error

	return socialMedias, err
}

func UpdateSocialMedia(socialMedia model.SocialMedia, userID int) (*model.SocialMedia, error) {
	var (
		db = config.GetDB()
	)

	res := db.Model(&model.SocialMedia{}).Omit("User").Where("id = ? AND user_id = ?", socialMedia.ID, userID).Updates(&socialMedia)
	if res.Error != nil {
		return nil, res.Error
	}

	return GetSocialMediaDetail(socialMedia.ID)
}

func DeleteSocialMedia(id, userID int) (int, error) {
	var (
		db = config.GetDB()
	)

	res := db.Model(&model.SocialMedia{}).Where("id = ? AND user_id = ?", id, userID).Delete(&model.SocialMedia{})
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, fmt.Errorf("user not allowed to delete this social media")
	}

	return id, nil
}
