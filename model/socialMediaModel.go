package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	User           User       `gorm:"foreignKey:UserID;references:ID"`
}

func (s *SocialMedia) Validate() error {
	if govalidator.IsNull(s.Name) {
		return fmt.Errorf("name is null")
	}

	if govalidator.IsNull(s.SocialMediaUrl) {
		return fmt.Errorf("social media url is null")
	}

	return nil
}

func (SocialMedia) TableName() string {
	return "social_media"
}
