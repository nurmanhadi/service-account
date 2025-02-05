package response

import (
	"errors"
	"fmt"
	"service-account/pkg/exception"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, statusCode int, err string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": err,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func ResponseError(c *fiber.Ctx, err error) error {
	if errors.Is(err, exception.NasabahCountAlreadyExist) {
		return ErrorResponse(c, 400, err.Error())
	} else if validationErr, ok := err.(validator.ValidationErrors); ok {
		var values []string
		for _, fieldErr := range validationErr {
			value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
			values = append(values, value)
		}
		str := strings.Join(values, ", ")
		return ErrorResponse(c, 400, str)
	} else {
		return ErrorResponse(c, 500, "internal server error")
	}
}
