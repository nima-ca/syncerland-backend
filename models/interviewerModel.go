package models

import "gorm.io/gorm"

type Interviewer struct {
	gorm.Model

	Name        string `gorm:"type:varchar(250); not null"`
	Email       string `gorm:"type:varchar(255);"`
	PhoneNumber string `gorm:"type:varchar(20)"`

	InterviewID uint `gorm:"index"`
	Interview   Interview
}
