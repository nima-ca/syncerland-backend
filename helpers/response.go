package helpers

import (
	"net/http"
)

type APIResponse[T any] struct {
	Status int      `json:"statusCode"`
	Data   *T       `json:"data"` // Use a pointer to T to allow nil values
	Errors []string `json:"errors"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func NewAPIResponse[T any](statusCode int, data *T, errors []string) APIResponse[T] {
	return APIResponse[T]{
		Status: statusCode,
		Data:   data,
		Errors: errors,
	}
}

func OkResponse[T any](data T) APIResponse[T] {
	return NewAPIResponse[T](http.StatusOK, &data, []string{})
}

func ErrorResponse[T any](statusCode int, errors []string) APIResponse[T] {
	return NewAPIResponse[T](statusCode, nil, errors)
}
