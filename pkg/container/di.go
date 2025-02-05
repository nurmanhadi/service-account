package container

import (
	"context"
	"service-account/api/handler"
	"service-account/api/router"
	"service-account/internal/nasabah"
	"service-account/internal/rekening"
	"service-account/pkg/database/psql"
	"service-account/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func DependencyInjection(app *fiber.App) {
	ctx := context.Background()
	logger := logger.InitLogrus()
	validation := validator.New()
	db := psql.PsqlConnect()

	rekeningRepository := rekening.NewRekeningRepository(db, ctx)

	nasabahRepository := nasabah.NewNasabahRepository(db, ctx)
	nasabahService := nasabah.NewNasabahService(&nasabahRepository, &rekeningRepository, validation, logger)
	nasabahHandler := handler.NewNasabahHandler(&nasabahService)
	router.NasabahRouter(app, nasabahHandler)

}
