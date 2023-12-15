package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	FullTimeEmploymentType   uint8 = 0
	PartTimeEmploymentType   uint8 = 1
	ContractEmploymentType   uint8 = 2
	InternshipEmploymentType uint8 = 3
	FreelancerEmploymentType uint8 = 4
	ConsultantEmploymentType uint8 = 5
)

type Job struct {
	gorm.Model

	CompanyName    string    `gorm:"type:varchar(250);not null"`
	Title          string    `gorm:"type:varchar(250);not null"`
	Description    string    `gorm:"type:text"`
	EmploymentType uint8     `gorm:"not null"`
	PostedDate     time.Time `gorm:"type:timestamptz"`
	Deadline       time.Time `gorm:"type:timestamptz"`

	UserID uint `gorm:"index"`
	User   User

	Applications []Application

	CountryID int `gorm:"index"`
	Country   Country
}
