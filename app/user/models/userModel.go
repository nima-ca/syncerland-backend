package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name        string `gorm:"not null;"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Otp         string
	IsVerified  bool `gorm:"default:false"`
	OtpSendTime time.Time
	LastLogin   time.Time
}
