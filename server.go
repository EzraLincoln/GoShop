package main

import (
	"./config"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"./app"
)

func main() {

	makeHashString("abcd")

	config.InitSQLiteDB()

	app.DefineHandlers()

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}

func makeHashString(cleartext  string) ([]byte) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(cleartext), 4)
	fmt.Println(string(hash))
	return hash

}



