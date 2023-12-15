package helpers

import (
	"net/http"
)

type APIResponse[T any] struct {
	Status int   `json:"statusCode"`
	Data   *T    `json:"data"` // Use a pointer to T to allow nil values
	Err    error `json:"error"`
}

func NewAPIResponse[T any](statusCode int, data *T, err error) APIResponse[T] {
	return APIResponse[T]{
		Status: statusCode,
		Data:   data,
		Err:    err,
	}
}

func OkResponse[T any](data T) APIResponse[T] {
	return NewAPIResponse[T](http.StatusOK, &data, nil)
}

func ErrorResponse[T any](statusCode int, err error) APIResponse[T] {
	return NewAPIResponse[T](statusCode, nil, err)
}
