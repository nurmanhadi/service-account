package psql

import (
	"database/sql"
	"fmt"
	"log"
	"service-account/config"
	"time"

	_ "github.com/lib/pq"
)

func PsqlConnect() *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Set.Db.User, config.Set.Db.Password, config.Set.Db.Host, config.Set.Db.Port, config.Set.Db.Dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(config.Set.Db.ConnMaxIdletime)
	db.SetMaxOpenConns(config.Set.Db.MaxOpenConns)
	db.SetConnMaxIdleTime(time.Minute * time.Duration(config.Set.Db.ConnMaxIdletime))
	db.SetConnMaxLifetime(time.Hour * time.Duration(config.Set.Db.ConnMaxLifetime))

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("connect to postgres")
	return db
}
