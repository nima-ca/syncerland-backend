package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model

	Name   string `gorm:"type:varchar(100);not null"`
	Alpha2 string `gorm:"type:varchar(2);not null"`
	Alpha3 string `gorm:"type:varchar(3);not null"`
}
