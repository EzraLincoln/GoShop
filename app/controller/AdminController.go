package controller

import (
	"../structs"
	"fmt"
	"html/template"
	"net/http"
	"../../config"
	"../model"
)

type Equipments struct{}

func Admin(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	
	auth, ok := session.Values["logged"];
	if !(ok) || !(auth.(bool)) {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["user-type"].(string) == "Verleiher" {
			
			kunden_id_int := session.Values["KundenID"].(int)
			client,_ := model.Get_Kunden_By_ID(kunden_id_int)
			
			p := structs.Menu{
				Item1:     "Kunden,clients", Item2: "Equipment,equipment", Item3: "Logout,logout",
				Basket:    false,
				Name:      client.Benutzername, Type: client.Typ,
				EmptySide: false,
				Profil:    true, ProfilBild: client.BildUrl,
			}
			
			tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin.html", "template/header.html", "template/static_imports.html"))
			
			tmpl.ExecuteTemplate(w, "main", p)
			tmpl.ExecuteTemplate(w, "static_imports", p)
			tmpl.ExecuteTemplate(w, "header", p)
			tmpl.ExecuteTemplate(w, "admin", p)
			
		} else {
			http.Redirect(w, r, "/index", 301)
		}
		
	}
}

func Admin_Equipment_Hinzufügen(w http.ResponseWriter, r *http.Request) {
	fmt.Println("admin_Equipment_Hinzufügen(w http.ResponseWriter, r *http.Request)")
	fmt.Println()
	
	if r.Method == "POST" {
		model.CreateEquipment(w, r)
		// equipment(w,r)
	} else {
		
		p := structs.Menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
			Basket:    false,
			Name:      "", Type: "",
			EmptySide: false,
			Profil:    true, ProfilBild: ""}
		
		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin_Equipment_Hinzufügen.html", "template/header.html", "template/static_imports.html"))
		
		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "admin_Equipment_Hinzufügen", p)
	}
	
}

func Admin_Equipment_Verwalten(w http.ResponseWriter, r *http.Request) {
	// ADMIN
	p := structs.Menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}
	
	// ArtikelArr := model.Get_Alle_Equipment()
	
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin_Equipment_Verwalten.html", "template/header.html", "template/static_imports.html"))
	
	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	// tmpl.ExecuteTemplate(w, "admin_Equipment", admin_Equipment_Collection{Items: ArtikelArr})
	
}

func Admin_Kunden_Verwalten(w http.ResponseWriter, r *http.Request) {
	
	p := structs.Menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}
	
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin_Kunden_Verwalten.html", "template/header.html", "template/static_imports.html"))
	
	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	
	kunde,_:= model.Get_Kunden_By_ID(1)
	
	tmpl.ExecuteTemplate(w, "adminEditClients", kunde)
}

func Admin_Equipment_Bearbeiten(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("admin_Equipment_Bearbeiten(w http.ResponseWriter, r *http.Request)")
	fmt.Println()
	
	if r.Method == "POST" {
		model.CreateEquipment(w, r)
		// equipment(w,r)
	} else {
		
		p := structs.Menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
			Basket:    false,
			Name:      "", Type: "",
			EmptySide: false,
			Profil:    true, ProfilBild: ""}
		
		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin_Equipment_Bearbeiten.html", "template/header.html", "template/static_imports.html"))
		
		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "admin_Equipment_Bearbeiten", p)
	}
	
}

func Admin_Kunden_Bearbeiten(w http.ResponseWriter, r *http.Request) {
	
	p := structs.Menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}
	
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin_Kunden_Bearbeiten.html", "template/header.html", "template/static_imports.html"))
	
	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	
	kunde,_ := model.Get_Kunden_By_ID(1)
	
	tmpl.ExecuteTemplate(w, "adminEditClients", kunde)
}

