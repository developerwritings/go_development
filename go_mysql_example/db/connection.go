package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/todo")

	if err != nil {
		panic(err.Error())
	}

	return db, err
}
