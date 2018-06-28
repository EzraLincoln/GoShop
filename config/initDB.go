package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"fmt"
)

// Db handle
var Db *sql.DB
var err error

const (
	DB_USER     = "borgdirmedia"
	DB_PASSWORD = "borgdirmedia"
	DB_NAME     = "borgdirmedia"
)

func InitSQLiteDB() *sql.DB {

	fmt.Println("Initialize SQLite Database")

	Db, err = sql.Open("sqlite3", "./data/borgdirmedia")

	if err != nil {
		fmt.Println("FEHLER")
		panic(err)
	} else {
		fmt.Println("Erfolgreich Verbindung mit Datenbank aufgebaut !")
	}

	return Db
}

func InitPostgresDB() *sql.DB {

	/* fmt.Println("Initialize Postgres Database")

	connStr := "user=borgdirmedia dbname=borgdirmedia password=borgdirmedia host=localhost port=5431 sslmode=disable"
	Db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("FEHLER")
		log.Fatal(err)
	} else {
		fmt.Println("Erfolgreich Verbindung mit Datenbank aufgebaut !")
	}

		fmt.Println("Start")
	*/

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println("Db = ",Db)
	Db, err := sql.Open("postgres", dbinfo)

	// fmt.Println("Db = ",Db)

	fmt.Println(err)

	// defer Db.Close()

	// fmt.Println("Db = ",Db)

	return Db
}

func ReturnDB() *sql.DB {
	return Db
}
