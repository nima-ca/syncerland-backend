package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	AppliedApplicationStatus    uint8 = 0
	InProgressApplicationStatus uint8 = 1
	RejectedApplicationStatus   uint8 = 2
	OfferedApplicationStatus    uint8 = 3
	AcceptedApplicationStatus   uint8 = 4
	WithdrawnApplicationStatus  uint8 = 5
)

type Application struct {
	gorm.Model

	DateApplied time.Time `gorm:"type:timestamptz;not null"`
	Resume      string    `gorm:"type:text;not null"`
	CoverLetter string    `gorm:"type:text"`
	Status      uint8

	JobID uint `gorm:"index"`
	Job   Job

	UserID uint `gorm:"index"`
	User   User

	Interviews []Interview
	Notes      []Note
	Offers     []Offer
}
