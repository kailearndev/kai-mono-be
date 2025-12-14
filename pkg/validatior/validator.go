package pkg

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

func FormatValidatorErrors(err error) map[string]string {
	errors := make(map[string]string)
	if err == nil {
		return errors
	}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		errors["error"] = err.Error()
		return errors
	}
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag() // ví dụ: "Name": "required"
	}

	return errors
}
