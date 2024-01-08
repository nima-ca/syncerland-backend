package dto

import "time"

type CreateJobDto struct {
	CompanyName    string
	Title          string
	Description    string
	EmploymentType uint8
	UserID         uint
	CountryID      int
	PostedDate     time.Time
	Deadline       time.Time
}

type CreateJobHandlerDto struct {
	CompanyName    string `json:"companyName" validate:"required,min=1,max=250"`
	Title          string `json:"title" validate:"required,min=1,max=250"`
	Description    string `json:"description" validate:"gte=0,lte=3000"`
	EmploymentType uint8  `json:"employmentType" validate:"required,min=1,max=6"`
	CountryID      int    `json:"countryId" validate:"required,min=1"`
	PostedDate     string `json:"postedDate" validate:"omitempty,dateformat=2006-01-02"`
	Deadline       string `json:"deadline" validate:"omitempty,dateformat=2006-01-02"`
}
