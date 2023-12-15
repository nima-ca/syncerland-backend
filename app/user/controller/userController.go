package controller

import (
	"net/http"
	"syncerland/app/user/dto"
	"syncerland/helpers"
	"syncerland/packages/errors"
	"syncerland/packages/validators"

	"github.com/gofiber/fiber/v2"
)

const (
	FailedToParseBodyError string = "Failed to Parse Request Body"
)

func RegisterHandler(ctx *fiber.Ctx) error {

	// Get Body of the Request
	var body dto.RegisterHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(helpers.ErrorResponse[any](http.StatusBadRequest, []string{
				FailedToParseBodyError,
			}))
	}

	errs := validators.Validate(body)
	if len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				errors.GetValidationErrors(errs),
			))
	}

	return ctx.Status(http.StatusOK).JSON("")

}
