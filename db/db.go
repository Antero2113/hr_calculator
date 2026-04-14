package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	dsn := "antero:MSOj9x2vnZjGrgZZ@tcp(localhost:3307)/antero?charset=utf8mb4&parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}