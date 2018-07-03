package app

import "net/http"
import "./route"

func DefineHandlers() {

	http.HandleFunc("/", route.Index)
	http.HandleFunc("/admin", route.Admin)
	http.HandleFunc("/admin/equipment{", route.Admin_Equipment_Verwalten)
	http.HandleFunc("/admin/add", route.Admin_Equipment_Hinzufügen)
	http.HandleFunc("/admin/edit-equipment", route.Admin_Equipment_Bearbeiten)
	http.HandleFunc("/admin/clients", route.Admin_Kunden_Verwalten)
	http.HandleFunc("/admin/edit-client", route.Admin_Kunden_Bearbeiten)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/login", route.Login)
	http.HandleFunc("/logout", route.Logout)
	http.HandleFunc("/register", route.Register)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/warenkorb", route.Warenkorb)
	http.HandleFunc("/equipment", route.Equipment)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/myEquipment", route.MeineGeräte)
	http.HandleFunc("/profil", route.Profil)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/test", route.Test)

	return

}
