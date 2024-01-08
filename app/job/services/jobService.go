package services

import (
	"syncerland/app/job/dto"
	"syncerland/core/initializers"
	"syncerland/models"
)

func CreateJob(dto dto.CreateJobDto) (*models.Job, error) {

	newJob := models.Job{
		CompanyName:    dto.CompanyName,
		Title:          dto.CompanyName,
		EmploymentType: dto.EmploymentType,
		UserID:         dto.UserID,
		CountryID:      dto.CountryID,
		Description:    dto.Description,
		PostedDate:     dto.PostedDate,
		Deadline:       dto.Deadline,
	}

	// save job in DB
	result := initializers.DB.Create(&newJob)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newJob, nil
}
