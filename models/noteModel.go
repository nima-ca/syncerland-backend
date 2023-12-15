package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model

	Text string `gorm:"type:text;not null"`

	InterviewID uint `gorm:"index"`
	Interview   Interview

	ApplicationID uint `gorm:"index"`
	Application   Application
}
