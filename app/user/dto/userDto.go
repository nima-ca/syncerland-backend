package dto

type CreateUserDto struct {
	Name     string
	Email    string
	Password string
	Otp      string
}

type RegisterHandlerBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type LoginHandlerBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type LoginHandlerResponse struct {
	Success bool `json:"success"`
}
