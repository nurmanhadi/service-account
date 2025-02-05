package router

import (
	"service-account/api/handler"

	"github.com/gofiber/fiber/v2"
)

func NasabahRouter(app *fiber.App, u *handler.NasabahHandlerImpl) {
	api := app.Group("/api/v1")
	nasabah := api.Group("/nasabah")
	nasabah.Post("/", u.Daftar)
}
