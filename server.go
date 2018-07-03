package main

import (
	"./config"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"./app"
	"./app/model"
)

func main() {

	// makeHashString("abcd")

	config.InitSQLiteDB()

	// KLAPPT
	// model.Add_Verleih(10,20,"asdad","asdasd",3)

	// KLAPPT
	// model.Delete_Verleih_By_Artikel_ID(20)
	
	model.DeleteKunde(44)
	
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



