package errors

import (
	"fmt"
	"syncerland/packages/validators"
)

func GetValidationErrors(errors []validators.ValidationErrorResponse) []string {

	var errorsSlice []string

	for _, err := range errors {
		// Customize error messages based on the field and tag
		switch err.Tag {
		case "required":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s is required.",
				err.FailedField))
		case "email":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be a valid email address.",
				err.FailedField))
		case "gte":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be greater or equal to %s.",
				err.FailedField, err.Param))
		case "lte":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be less than or equal to %s.",
				err.FailedField, err.Param))
		case "min":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be at least %s.",
				err.FailedField, err.Param))
		case "max":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must be at most %s.",
				err.FailedField, err.Param))
		case "len":
			errorsSlice = append(errorsSlice, fmt.Sprintf("%s must have exactly %s characters.",
				err.FailedField, err.Param))
		default:
			// Use the default error message
			errorsSlice = append(errorsSlice, fmt.Sprintf("Validation failed for field %s with tag %s.",
				err.FailedField, err.Tag))
		}
	}

	return errorsSlice
}
