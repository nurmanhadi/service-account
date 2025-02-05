package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log := InitLogrus()
		start := time.Now()
		requestId := uuid.NewString()
		c.Locals("request_id", requestId)

		err := c.Next()
		latency := time.Since(start)
		log.WithFields(logrus.Fields{
			"request_id": requestId,
			"method":     c.Method(),
			"path":       c.OriginalURL(),
			"status":     c.Response().StatusCode(),
			"latency":    latency,
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}).Info("Request processed")
		return err
	}
}
