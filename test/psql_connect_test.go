package test

import (
	"service-account/config"
	"service-account/pkg/database/psql"
	"testing"
)

func TestPsqlConnect(t *testing.T) {
	config.LoadConfig("../.env")
	db := psql.PsqlConnect()
	defer db.Close()
}
