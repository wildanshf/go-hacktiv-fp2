package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Age       int        `json:"age"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Validate() error {
	if !govalidator.IsEmail(u.Email) {
		return fmt.Errorf("email not valid")
	}

	if govalidator.IsNull(u.Email) {
		return fmt.Errorf("email is null")
	}

	if govalidator.IsNull(u.Username) {
		return fmt.Errorf("username is null")
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password too short")
	}

	if u.Age < 8 {
		return fmt.Errorf("under age")
	}

	return nil
}

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserParam struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
