package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	PhoneInterviewType  uint8 = 0
	OnlineInterviewType uint8 = 1
	OnSiteInterviewType uint8 = 2
)

const (
	PlannedInterviewStatus  uint8 = 0
	FinishedInterviewStatus uint8 = 1
	CanceledInterviewStatus uint8 = 2
)

type Interview struct {
	gorm.Model

	Date          time.Time `gorm:"type:timestamptz; not null"`
	Status        uint8
	InterviewType uint8

	Interviewers []Interviewer
	Notes        []Note

	ApplicationID uint `gorm:"index"`
	Application   Application

	UserID uint `gorm:"index"`
	User   User
}
