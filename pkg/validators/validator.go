package validators

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func ExtractValidationError(req interface{}) []error {
	var errs []error
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := v.Struct(req); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, errors.New(e.Field()+": "+e.Tag()))
		}
	}
	return errs
}
