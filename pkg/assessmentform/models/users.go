package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname        string `gorm:"type:varchar(200);not null" json:"fullname"`
	Email           string `gorm:"type:varchar(50);not null" json:"email"`
	Phone           string `gorm:"type:varchar(20);not null" json:"phone"`
	Password        string `gorm:"type:varchar(50);not null" json:"password"`
	ConfirmPassword string `gorm:"type:varchar(50);not null" json:"confirmPassword"`
}
