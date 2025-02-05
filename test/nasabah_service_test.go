package test

import (
	"context"
	"fmt"
	"service-account/config"
	"service-account/internal/nasabah"
	"service-account/internal/rekening"
	"service-account/pkg/database/psql"
	"service-account/pkg/dto"
	"service-account/pkg/logger"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestNasabahDaftar(t *testing.T) {
	loggger := logger.InitLogrus()
	ctx := context.Background()
	validation := validator.New()
	config.LoadConfig("../.env")
	db := psql.PsqlConnect()
	defer db.Close()
	nRepo := nasabah.NewNasabahRepository(db, ctx)
	rRepo := rekening.NewRekeningRepository(db, ctx)
	serv := nasabah.NewNasabahService(&nRepo, &rRepo, validation, loggger)
	req := &dto.NasabahRequestDto{
		Nama: "nurman hadi",
		Nik:  "1",
		NoHp: "1",
	}

	noRekening, err := serv.Daftar(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(noRekening)

}
