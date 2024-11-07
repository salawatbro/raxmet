package validators

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var customMessages = map[string]string{
	"required": "This field is required",
	"email":    "Invalid email format",
	"min":      "Value is too short, minimum required is %s characters",
	"max":      "Value is too long, maximum allowed is %s characters",
}

func ExtractValidationError(req interface{}) []string {
	var messages []string
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := v.Struct(req); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := customMessages[err.Tag()]
			if message == "" {
				message = "Invalid value"
			}
			if err.Tag() == "min" || err.Tag() == "max" {
				message = strings.Replace(message, "%s", err.Param(), 1)
			}
			messages = append(messages, err.Field()+": "+message)
		}
	}
	return messages
}
