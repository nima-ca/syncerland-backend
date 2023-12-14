package models

import (
	"syncerland/app/user/models"
	"time"

	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	Applied    ApplicationStatus = "Applied"
	InProgress ApplicationStatus = "InProgress"
	Rejected   ApplicationStatus = "Rejected"
	Offered    ApplicationStatus = "Offered"
	Accepted   ApplicationStatus = "Accepted"
	Withdrawn  ApplicationStatus = "Withdrawn"
)

type Application struct {
	gorm.Model

	DateApplied time.Time         `gorm:"not null"`
	Status      ApplicationStatus `gorm:"type:varchar(20);default:'Applied'"`
	Resume      string            `gorm:"type:text;not null"`
	CoverLetter string            `gorm:"type:text;not null"`

	UserID uint
	User   models.User
}
