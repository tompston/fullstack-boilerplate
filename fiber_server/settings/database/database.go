package database

import (
	"database/sql"
	"fmt"
	"os"

	"fiber_server/settings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// a function that returns the dsn string
func DsnString() string {
	// load all the env variables
	settings.LoadEnvFile()

	HOST := os.Getenv("HOST")
	USER := os.Getenv("POSTGRES_USER")
	PASSW := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DB")
	PORT := os.Getenv("HOST_PORT")
	SSLMODE := os.Getenv("SSLMODE")
	TZ := os.Getenv("TZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASSW, DBNAME, PORT, SSLMODE, TZ)

	return dsn
}

// get a db connection with the sqlx package
func GetDbConnSqlx() (*sqlx.DB, error) {

	dsn := DsnString()

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, err
}

// get a connection with the inbuilt package
func GetDbConnSql() (*sql.DB, error) {

	dsn := DsnString()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, err
}
