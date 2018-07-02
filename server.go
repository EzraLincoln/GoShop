package main

import (
	"./app/route"
	"./config"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {

	makeHashString("abcd")

	config.InitSQLiteDB()

	route.Handler()

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}

func makeHashString(cleartext  string) ([]byte) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(cleartext), 4)
	fmt.Println(string(hash))
	return hash
}



