package controller

import (
	"net/http"
	emailService "syncerland/app/email/services"
	"syncerland/app/jwt"
	userDto "syncerland/app/user/dto"
	userService "syncerland/app/user/services"
	"syncerland/helpers"
	"syncerland/packages/errors"
	"syncerland/packages/validators"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserAlreadyExistError   string = "User already Exist"
	FailedToCreateUserError string = "Failed to create user"
	InvalidEmailOrPassword  string = "Invalid Email or password"
)

func RegisterHandler(ctx *fiber.Ctx) error {

	// Get Body of the Request
	var body userDto.RegisterHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(helpers.ErrorResponse[any](http.StatusBadRequest, []string{
				errors.FailedToParseBodyError,
			}))
	}

	// Validate Body of request
	if errs := validators.Validate(body); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				errors.GetValidationErrors(errs),
			))
	}

	user, userErr := userService.FindUserByEmail(body.Email)
	if userErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{errors.InternalServerError},
			))
	}

	// check if user exists
	if user != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{UserAlreadyExistError},
			))
	}

	otp := userService.GenerateOTP()

	_, err := userService.CreateUser(userDto.CreateUserDto{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Otp:      otp,
	})

	go emailService.SendOTP(body.Email, otp)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{FailedToCreateUserError},
			))
	}

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func LoginHandler(ctx *fiber.Ctx) error {
	// Get Body of the Request
	var body userDto.LoginHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(helpers.ErrorResponse[any](http.StatusBadRequest, []string{
				errors.FailedToParseBodyError,
			}))
	}

	// Validate Body of request
	if errs := validators.Validate(body); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				errors.GetValidationErrors(errs),
			))
	}

	user, userErr := userService.FindUserByEmail(body.Email)
	if userErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{errors.InternalServerError},
			))
	}

	// check if user exists
	if user == nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{InvalidEmailOrPassword},
			))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{InvalidEmailOrPassword},
			))
	}

	accessToken, refreshToken, generateTokenErr := jwt.GenerateAccessAndRefreshTokens(user.ID)
	if generateTokenErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{errors.InternalServerError},
			))
	}

	jwt.SetGeneratedTokensInCookie(ctx, accessToken, refreshToken)

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}
