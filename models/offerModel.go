package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency string

const (
	AUDCurrency Currency = "AUD"
	GBPCurrency Currency = "GBP"
	EURCurrency Currency = "EUR"
	JPYCurrency Currency = "JPY"
	CHFCurrency Currency = "CHF"
	USDCurrency Currency = "USD"
	CADCurrency Currency = "CAD"
	IRRCurrency Currency = "IRR"
	NZDCurrency Currency = "NZD"
	SEKCurrency Currency = "SEK"
	AEDCurrency Currency = "AED"
)

type Offer struct {
	gorm.Model

	BaseSalary string    `gorm:"not null"`
	Benefits   string    `gorm:"type:text"`
	Currency   Currency  `gorm:"type:varchar(5)"`
	Deadline   time.Time `gorm:"type:timestamptz"`

	ApplicationID uint `gorm:"index"`
	Application   Application
}
