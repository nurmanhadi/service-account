package test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"service-account/config"
	"service-account/internal/nasabah"
	"service-account/pkg/database/psql"
	"testing"
	"time"
)

func TestNasabahAdd(t *testing.T) {
	ctx := context.Background()
	config.LoadConfig()
	db := psql.PsqlConnect()
	defer db.Close()
	repo := nasabah.NewNasabahRepository(db, ctx)

	nasabah := &nasabah.Nasabah{
		Nama:      "nurman",
		Nik:       "92387483",
		NoHp:      "827378",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	nasabahId, err := repo.Add(nasabah)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(nasabahId)
}
func TestYali(t *testing.T) {
	fmt.Println(rand.Int64N(9999999999))
}
