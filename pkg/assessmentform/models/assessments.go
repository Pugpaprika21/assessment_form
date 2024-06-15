package models

import "gorm.io/gorm"

type Assessments struct {
	gorm.Model
	Question1  int `gorm:"type:int;not null"`
	Question2  int `gorm:"type:int;not null"`
	Question3  int `gorm:"type:int;not null"`
	Question4  int `gorm:"type:int;not null"`
	Question5  int `gorm:"type:int;not null"`
	Question6  int `gorm:"type:int;not null"`
	Question7  int `gorm:"type:int;not null"`
	Question8  int `gorm:"type:int;not null"`
	Question9  int `gorm:"type:int;not null"`
	Question10 int `gorm:"type:int;not null"`
	UserID     uint
}
