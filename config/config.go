package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Db struct {
		User            string
		Password        string
		Host            string
		Port            int
		Dbname          string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxIdletime int
		ConnMaxLifetime int
	}
}

var Set *config

func LoadConfig(path string) {
	err := godotenv.Load(path)
	if err != nil {
		panic(err)
	}
	Set = &config{}

	Set.Db.User = os.Getenv("DB_USER")
	Set.Db.Password = os.Getenv("DB_PASSWORD")
	Set.Db.Host = os.Getenv("DB_HOST")
	Set.Db.Dbname = os.Getenv("DB_NAME")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Invalid DB_PORT value")
	}
	Set.Db.Port = port

	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		panic("Invalid DB_MAX_IDLE_CONNS value")
	}
	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		panic("Invalid DB_MAX_OPEN_CONNS value")
	}
	connMaxIdletime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_IDLETIME"))
	if err != nil {
		panic("Invalid DB_CONN_MAX_IDLETIME value")
	}
	connMaxLifetime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		panic("Invalid DB_CONN_MAX_LIFETIME value")
	}
	Set.Db.MaxIdleConns = maxIdleConns
	Set.Db.MaxOpenConns = maxOpenConns
	Set.Db.ConnMaxIdletime = connMaxIdletime
	Set.Db.ConnMaxLifetime = connMaxLifetime

}
