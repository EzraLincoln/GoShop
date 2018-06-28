package main

import (
	"net/http"
	"./app/route"
	"fmt"
	"./config"
)

const (
	DB_USER     = "borgdirmedia"
	DB_PASSWORD = "borgdirmedia"
	DB_NAME     = "borgdirmedia"
)

func main() {

	fmt.Println("Start")

	/*dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println(err)
	defer db.Close()

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT bezeichnung FROM artikel")
	fmt.Println(err)

	for rows.Next() {
		var bezeichnung string
		err = rows.Scan(&bezeichnung)
		fmt.Println(err)
		fmt.Println("bezeichnung")
		fmt.Printf("%8v\n", bezeichnung)
	}
	*/


	// config.InitSQLiteDB()

	fmt.Println(config.Db)

	// config.Db = config.InitPostgresDB()

	config.Db = config.InitSQLiteDB()

	route.Handler()

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)

	fmt.Println("Exit")
}
