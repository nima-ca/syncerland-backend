package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	FullTimeEmploymentType   uint8 = 1
	PartTimeEmploymentType   uint8 = 2
	ContractEmploymentType   uint8 = 3
	InternshipEmploymentType uint8 = 4
	FreelancerEmploymentType uint8 = 5
	ConsultantEmploymentType uint8 = 6
)

type Job struct {
	gorm.Model

	CompanyName    string    `gorm:"type:varchar(250);not null"`
	Title          string    `gorm:"type:varchar(250);not null"`
	Description    string    `gorm:"type:text"`
	EmploymentType uint8     `gorm:"not null"`
	PostedDate     time.Time `gorm:"type:timestamptz;default:null"`
	Deadline       time.Time `gorm:"type:timestamptz;default:null"`

	UserID uint `gorm:"index"`
	User   User

	Applications []Application

	CountryID int `gorm:"index"`
	Country   Country
}
