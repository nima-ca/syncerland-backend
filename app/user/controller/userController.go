package controller

import (
	"net/http"
	emailService "syncerland/app/email/services"
	"syncerland/app/jwt"
	userDto "syncerland/app/user/dto"
	userService "syncerland/app/user/services"
	"syncerland/core/initializers"
	"syncerland/helpers"
	"syncerland/packages/errors"
	"syncerland/packages/validators"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserAlreadyExistErrorMsg       string = "User already Exist"
	FailedToCreateUserErrorMsg     string = "Failed to create user"
	InvalidEmailOrPasswordErrorMsg string = "Invalid email or password"
	IncorrectEmailErrorMsg         string = "Incorrect email address"
	IncorrectOTPErrorMsg           string = "Incorrect OTP"
	OTPExpiredErrorMsg             string = "OTP expired"
	UserAlreadyVerifiedErrorMsg    string = "User is already verified"
)

func RegisterHandler(ctx *fiber.Ctx) error {

	// Get Body of the Request
	var body userDto.RegisterHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(helpers.ErrorResponse[any](http.StatusBadRequest, []string{
				errors.FailedToParseBodyErrorMsg,
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
				[]string{errors.InternalServerErrorErrorMsg},
			))
	}

	// check if user exists
	if user != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{UserAlreadyExistErrorMsg},
			))
	}

	otp := userService.GenerateOTP()

	_, err := userService.CreateUser(userDto.CreateUserDto{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Otp:      otp,
	})

	// DOC: send OTP and not wait for the response
	go emailService.SendOTP(body.Email, otp)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{FailedToCreateUserErrorMsg},
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
				errors.FailedToParseBodyErrorMsg,
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
				[]string{errors.InternalServerErrorErrorMsg},
			))
	}

	// check if user exists
	if user == nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{InvalidEmailOrPasswordErrorMsg},
			))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{InvalidEmailOrPasswordErrorMsg},
			))
	}

	accessToken, refreshToken, generateTokenErr := jwt.GenerateAccessAndRefreshTokens(user.ID)
	if generateTokenErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				[]string{errors.InternalServerErrorErrorMsg},
			))
	}

	jwt.SetGeneratedTokensInCookie(ctx, accessToken, refreshToken)

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func VerifyUserHandler(ctx *fiber.Ctx) error {
	// Get Body of the Request
	var body userDto.VerifyHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			errors.FailedToParseBodyErrorMsg)
	}

	// Validate Body of request
	if errs := validators.Validate(body); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest,
				errors.GetValidationErrors(errs),
			))
	}

	// Find the user
	user, err := userService.FindUserByEmail(body.Email)
	if err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	// Verify user exists
	if user == nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			IncorrectEmailErrorMsg)
	}

	// Check if user is already verified
	if user.IsVerified {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			UserAlreadyVerifiedErrorMsg)
	}

	// Verify that an OTP is sent to user
	if user.Otp == "" {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			IncorrectEmailErrorMsg)
	}

	// Check if OTP is expired
	if userService.IsOTPExpired(user.OtpSendTime) {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			OTPExpiredErrorMsg)
	}

	// Verify OTP
	err = bcrypt.CompareHashAndPassword([]byte(user.Otp), []byte(body.Otp))
	if err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			IncorrectOTPErrorMsg)
	}

	result := initializers.DB.Model(&user).Updates(map[string]interface{}{"otp": "", "is_verified": true})

	if result.Error != nil || result.RowsAffected != 1 {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}
