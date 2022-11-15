package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type Photo struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   *string    `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      User       `gorm:"foreignKey:UserID;references:ID"`
}

func (p *Photo) Validate() error {
	if govalidator.IsNull(p.Title) {
		return fmt.Errorf("title is null")
	}

	if govalidator.IsNull(p.PhotoUrl) {
		return fmt.Errorf("photo url is null")
	}

	return nil
}

func (Photo) TableName() string {
	return "photo"
}
