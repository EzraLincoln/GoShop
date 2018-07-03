package main

import (
	"./config"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"./app/controller"
)

func main() {

	// makeHashString("abcd")

	config.InitSQLiteDB()

	// KLAPPT
	// model.Add_Verleih(10,20,"asdad","asdasd",3)

	// KLAPPT
	// model.Delete_Verleih_By_Artikel_ID(20)
	
	// KLAPPT
	// model.DeleteKunde(44)
	
	http.HandleFunc("/", controller.Index)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/admin", controller.Admin)
	http.HandleFunc("/admin/equipment{", controller.Admin_Equipment_Verwalten)
	http.HandleFunc("/admin/add", controller.Admin_Equipment_Hinzufügen)
	http.HandleFunc("/admin/edit-equipment", controller.Admin_Equipment_Bearbeiten)
	http.HandleFunc("/admin/clients", controller.Admin_Kunden_Verwalten)
	http.HandleFunc("/admin/edit-client", controller.Admin_Kunden_Bearbeiten)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/warenkorb", controller.Warenkorb)
	http.HandleFunc("/equipment", controller.Equipment)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/myEquipment", controller.MeineGeräte)
	http.HandleFunc("/profil", controller.Profil)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/test", controller.Test)
	
	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)
}

func makeHashString(cleartext  string) ([]byte) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(cleartext), 4)
	fmt.Println(string(hash))
	return hash

}



