package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) string {
	var values []string
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErr {
			value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
			values = append(values, value)
		}
	}
	msg := strings.Join(values, ", ")
	return msg
}
