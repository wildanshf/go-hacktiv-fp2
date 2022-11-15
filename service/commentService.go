package service

import (
	"fmt"
	"project-2/config"
	"project-2/model"
)

func GetCommentDetail(id int) (*model.Comment, error) {
	var (
		comment model.Comment
		db      = config.GetDB()
	)

	err := db.Model(&model.Comment{}).Where("id = ?", id).First(&comment).Error

	return &comment, err
}

func CreateComment(comment model.Comment) (*model.Comment, error) {
	var (
		db = config.GetDB()
	)

	err := db.Model(&model.Comment{}).Omit("User", "Photo").Create(&comment).Error

	return &comment, err
}

func GetComments(userID int) ([]*model.Comment, error) {
	var (
		comments []*model.Comment
		db       = config.GetDB()
	)

	err := db.Model(&model.Photo{}).Preload("User").Preload("Photo").Where("user_id = ?", userID).Find(&comments).Error

	return comments, err
}

func UpdateComment(comment model.Comment, userID int) (*model.Comment, error) {
	var (
		db = config.GetDB()
	)

	res := db.Model(&model.Comment{}).Omit("User", "Photo").Where("id = ? AND user_id = ?", comment.ID, userID).Updates(&comment)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("user not allowed to edit this comment")
	}

	return GetCommentDetail(comment.ID)
}

func DeleteComment(id, userID int) (int, error) {
	var (
		db = config.GetDB()
	)

	res := db.Model(&model.Comment{}).Where("id = ? AND user_id = ?", id, userID).Delete(&model.Comment{})
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return 0, fmt.Errorf("user not allowed to delete this photo")
	}

	return id, nil
}
