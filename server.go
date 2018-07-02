package main

import (
	"./app/route"
	"./config"
	"net/http"
)

func main() {

	config.InitSQLiteDB()

	route.Handler()

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}



