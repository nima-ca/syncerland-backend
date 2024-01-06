package controller

import (
	"fmt"
	"net/http"
	emailService "syncerland/app/email/services"
	"syncerland/app/jwt"
	userDto "syncerland/app/user/dto"
	userService "syncerland/app/user/services"
	"syncerland/core/initializers"
	"syncerland/helpers"
	"syncerland/packages/errors"
	"syncerland/packages/validators"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserAlreadyExistErrorMsg       string = "Email is already in use"
	FailedToCreateUserErrorMsg     string = "Failed to create user"
	InvalidEmailOrPasswordErrorMsg string = "Invalid email or password"
	IncorrectEmailErrorMsg         string = "Please send a correct email address"
	IncorrectOTPErrorMsg           string = "OTP is incorrect"
	OTPExpiredErrorMsg             string = "OTP is expired"
	UserAlreadyVerifiedErrorMsg    string = "You are already verified"
)

func RegisterHandler(ctx *fiber.Ctx) error {

	// Get Body of the Request
	var body userDto.RegisterHandlerBody
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

	user, userErr := userService.FindUserByEmail(body.Email)
	if userErr != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	// check if user exists
	if user != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, UserAlreadyExistErrorMsg)
	}

	otp := userService.GenerateOTP()

	_, err := userService.CreateUser(userDto.CreateUserDto{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Otp:      otp,
	})

	if err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, FailedToCreateUserErrorMsg)
	}

	// DOC: send OTP and not wait for the response
	go emailService.SendOTP(body.Email, otp)

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func LoginHandler(ctx *fiber.Ctx) error {
	// Get Body of the Request
	var body userDto.LoginHandlerBody
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

	user, userErr := userService.FindUserByEmail(body.Email)
	if userErr != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	// check if user exists
	if user == nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, InvalidEmailOrPasswordErrorMsg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, InvalidEmailOrPasswordErrorMsg)
	}

	accessToken, refreshToken, generateTokenErr := jwt.GenerateAccessAndRefreshTokens(user.ID)
	if generateTokenErr != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	jwt.SetGeneratedTokensInCookie(ctx, accessToken, refreshToken)

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func VerifyUserHandler(ctx *fiber.Ctx) error {
	// Get Body of the Request
	var body userDto.VerifyHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, errors.FailedToParseBodyErrorMsg)
	}

	// Validate Body of request
	if errs := validators.Validate(body); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest, errors.GetValidationErrors(errs)))
	}

	// Find the user
	user, err := userService.FindUserByEmail(body.Email)
	if err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	// Verify user exists
	if user == nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, IncorrectEmailErrorMsg)
	}

	// Check if user is already verified
	if user.IsVerified {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, UserAlreadyVerifiedErrorMsg)
	}

	// Verify that an OTP is sent to user
	if user.Otp == "" {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, IncorrectEmailErrorMsg)
	}

	// Check if OTP is expired
	if time.Now().After(user.OtpExpireTime) {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, OTPExpiredErrorMsg)
	}

	// Verify OTP
	if err := bcrypt.CompareHashAndPassword([]byte(user.Otp), []byte(body.Otp)); err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, IncorrectOTPErrorMsg)
	}

	result := initializers.DB.Model(&user).Updates(map[string]interface{}{"otp": "", "is_verified": true})

	if result.Error != nil || result.RowsAffected != 1 {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}

func ResendOTPHandler(ctx *fiber.Ctx) error {
	// Get Body of the Request
	var body userDto.ResendOTPHandlerBody
	if err := ctx.BodyParser(&body); err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, errors.FailedToParseBodyErrorMsg)
	}

	// Validate Body of request
	if errs := validators.Validate(body); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			helpers.ErrorResponse[any](http.StatusBadRequest, errors.GetValidationErrors(errs)))
	}

	// Find the user
	user, err := userService.FindUserByEmail(body.Email)
	if err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	// Verify user exists
	if user == nil {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, IncorrectEmailErrorMsg)
	}

	// Check if user is already verified
	if user.IsVerified {
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest, UserAlreadyVerifiedErrorMsg)
	}

	if time.Now().Before(user.OtpExpireTime) {
		diff := time.Until(user.OtpExpireTime)
		return helpers.SendErrorResponse(ctx, http.StatusBadRequest,
			fmt.Sprintf("You should wait %d seconds for next otp", int(diff.Seconds())),
		)
	}

	otp := userService.GenerateOTP()
	// Hash the Otp before storing it
	hashedOtp, err := bcrypt.GenerateFromPassword([]byte(otp), 10)
	if err != nil {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	result := initializers.DB.Model(&user).Updates(map[string]interface{}{"otp": hashedOtp,
		"otp_expire_time": userService.GetOTPExpireTime()})

	if result.Error != nil || result.RowsAffected != 1 {
		return helpers.SendErrorResponse(ctx, http.StatusInternalServerError,
			errors.InternalServerErrorErrorMsg)
	}

	// DOC: send OTP and not wait for the response
	go emailService.SendOTP(body.Email, otp)

	return ctx.Status(http.StatusOK).
		JSON(helpers.OkResponse[helpers.SuccessResponse](helpers.SuccessResponse{Success: true}))
}
