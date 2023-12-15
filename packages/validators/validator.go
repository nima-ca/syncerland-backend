package validators

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(data interface{}) []ValidationErrorResponse {
	validationErrors := []ValidationErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValidationErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Param = err.Param()       // Export field parameter

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
