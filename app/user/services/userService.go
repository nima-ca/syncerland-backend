package services

import (
	"errors"
	"fmt"
	"math/rand"
	"syncerland/app/user/dto"
	"syncerland/core/initializers"
	"syncerland/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(createUserDto dto.CreateUserDto) (*models.User, error) {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 10)
	if err != nil {
		return nil, err
	}

	// Hash the Otp before storing it
	hashedOtp, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Otp), 10)
	if err != nil {
		return nil, err
	}

	// create user
	user := models.User{
		Name:        createUserDto.Name,
		Email:       createUserDto.Email,
		Password:    string(hashedPassword),
		IsVerified:  false,
		Otp:         string(hashedOtp),
		OtpSendTime: time.Now(),
	}

	// save user in DB
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := initializers.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func FindUserById(id string) (*models.User, error) {
	user := models.User{}
	err := initializers.DB.Where("email = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GenerateOTP() string {
	otp := rand.Intn(900000) + 100000
	return fmt.Sprint(otp)
}
