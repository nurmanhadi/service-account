package test

import (
	"context"
	"service-account/config"
	"service-account/internal/rekening"
	"service-account/pkg/database/psql"
	"testing"
	"time"
)

func TestRekeningRepository(t *testing.T) {
	ctx := context.Background()
	config.LoadConfig()
	db := psql.PsqlConnect()
	defer db.Close()
	repo := rekening.NewRekeningRepository(db, ctx)
	nasabah := &rekening.Rekening{
		NasabahId:  1,
		NoRekening: "123456",
		Saldo:      0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	err := repo.Add(nasabah)
	if err != nil {
		t.Error(err)
	}
}
