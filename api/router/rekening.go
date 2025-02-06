package router

import (
	"service-account/api/handler"

	"github.com/gofiber/fiber/v2"
)

func RekeningRouter(app *fiber.App, u *handler.RekeningHandlerImpl) {
	api := app.Group("/api/v1")
	rekening := api.Group("/rekening")
	rekening.Post("/tabung", u.AddSaldo)
	rekening.Post("/tarik", u.WithdrawSaldo)
	rekening.Get("/saldo/:no_rekening", u.GetSaldoByNoRekening)
}
