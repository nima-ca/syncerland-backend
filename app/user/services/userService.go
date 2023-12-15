package services

import (
	"syncerland/app/user/dto"
	"syncerland/core/initializers"
	"syncerland/models"
	"time"

	"golang.org/x/crypto/bcrypt"
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
		return nil, err
	}

	return &user, nil
}

func FindUserByEmail(email string) *models.User {
	user := models.User{}
	initializers.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return nil
	}

	return &user
}

func FindUserById(id string) *models.User {
	user := models.User{}
	initializers.DB.Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return nil
	}

	return &user
}
