package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Otp         string
	IsVerified  bool      `gorm:"default:false"`
	OtpSendTime time.Time `gorm:"type:timestamptz"`
	LastLogin   time.Time `gorm:"type:timestamptz"`

	Jobs         []Job
	Interviews   []Interview
	Applications []Application
}
