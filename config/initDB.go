package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

// Db handle
var Db *sql.DB
var err error

func InitSQLiteDB() *sql.DB {

	Db, err := sql.Open("sqlite3", "./data/borgdirmedia.db")
	test(err)
	return Db
}

func test(e error) {
	if (err != nil) {
		fmt.Println("FEHLER : ", e)
	}
}
