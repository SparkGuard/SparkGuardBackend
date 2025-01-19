package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

//go:embed init.sql
var initSQL string

var db *sql.DB

func initDB() {
	if _, err := db.Exec(initSQL); err != nil {
		panic(err)
	}
}

func init() {
	var err error

	var connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))

	if db, err = sql.Open("postgres", connectionString); err != nil {
		panic(err)
	}

	initDB()
}
