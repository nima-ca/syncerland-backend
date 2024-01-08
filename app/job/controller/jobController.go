package controller

import (
	"net/http"
	"syncerland/app/job/dto"
	"syncerland/app/job/services"
	"syncerland/app/jwt"
	"syncerland/helpers"
	"syncerland/models"
	"syncerland/packages/errors"
	"syncerland/packages/validators"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateJobHandler(ctx *fiber.Ctx) error {

	// Get Body of the Request
	var body dto.CreateJobHandlerDto
	if err := ctx.BodyParser(&body); err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, errors.FailedToParseBodyErrorMsg)
	}

	// Validate Body of request
	if errs := validators.Validate(body); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				errors.GetValidationErrors(errs),
			))
	}

	user := ctx.Locals(jwt.UserLocalKey).(*models.User)

	const layout = "2006-01-02"
	parsedDeadline, _ := time.Parse(layout, body.Deadline)
	parsedPostedDate, _ := time.Parse(layout, body.PostedDate)

	services.CreateJob(dto.CreateJobDto{
		CompanyName:    body.CompanyName,
		Title:          body.Title,
		Description:    body.Description,
		EmploymentType: body.EmploymentType,
		UserID:         user.ID,
		CountryID:      body.CountryID,
		Deadline:       parsedDeadline,
		PostedDate:     parsedPostedDate,
	})

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func UpdateJobHandler(ctx *fiber.Ctx) error {
	return nil
}

func DeleteJobHandler(ctx *fiber.Ctx) error {
	return nil
}

func FindAllJobsHandler(ctx *fiber.Ctx) error {
	return nil
}

func FindOneHandler(ctx *fiber.Ctx) error {
	return nil
}
