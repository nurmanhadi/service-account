package response

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, statusCode int, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"data": data,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
