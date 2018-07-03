package controller

import (
	"net/http"
	"fmt"
	"strconv"
	"../structs"
	"html/template"
	"../../config"
	"../model"
	"strings"
)

type Kunden struct{}

func Equipment(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "POST" {
	
	}
	
	session, _ := config.CookieStore.Get(r, "session")
	
	Menu := structs.Menu{}
	
	if auth, ok := session.Values["logged"].(bool); !auth || !ok {
		
		fmt.Println("Index -> Gast")
		
		Menu = structs.Menu{
			Item1:     "", Item2: "Meine Geräte,myEquipment", Item3: "Login,login",
			Name:      "", Type: "",
			Basket:    false,
			EmptySide: false,
			Profil:    false, ProfilBild: ""}
	} else {
		
		fmt.Println("Index -> User")
		
		kunden_id_int := session.Values["KundenID"].(int)
		client, _ := model.Get_Kunden_By_ID(kunden_id_int)
		
		Menu = structs.Menu{
			Item1:  "Meine Geräte,myEquipment", Item2: "Logout,logout", Item3: "",
			Name:   client.Benutzername, Type: client.Typ,
			Basket: false, EmptySide: false, Profil: true, ProfilBild: client.BildUrl}
	}
	
	EquipmentArr := model.GetEquipment()
	
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))
	
	tmpl.ExecuteTemplate(w, "main", Menu)
	tmpl.ExecuteTemplate(w, "static_imports", Menu)
	tmpl.ExecuteTemplate(w, "header", Menu)
	
	tmpl.ExecuteTemplate(w, "equipment", structs.Equipment_Collection{Kategorien: []string{"Kameras", "Mikrofone", "Monitore", "Beleuchtung"}, Items: EquipmentArr})
	
}

func MeineGeräte(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	
	if auth, ok := session.Values["logged"].(bool); !auth || !ok {
		
		fmt.Println("MeineGeräte : Gast : FEHLER")
		http.Redirect(w, r, "/login", 301)
		
	} else {
		
		fmt.Println("Meine Geräte -> User ")
		
		kunden_id_int := session.Values["KundenID"].(int)
		client, _ := model.Get_Kunden_By_ID(kunden_id_int)
		
		Menu := structs.Menu{
			Item1:     "Equipment,equipment", Item2: "Logout,logout", Item3: "",
			Basket:    false,
			Name:      client.Benutzername, Type: client.Typ,
			EmptySide: false,
			Profil:    true, ProfilBild: client.BildUrl,
		}
		// ArtikelArr := model.GetUserEquipment(1)
		ArtikelArr := []string{}
		
		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/user_MeineGeräte.html", "template/header.html", "template/static_imports.html"))
		tmpl.ExecuteTemplate(w, "main", Menu)
		tmpl.ExecuteTemplate(w, "static_imports", Menu)
		tmpl.ExecuteTemplate(w, "header", Menu)
		tmpl.ExecuteTemplate(w, "myequipment", ArtikelArr)
	}
}

func Warenkorb(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	
	Menu := structs.Menu{}
	
	if auth, ok := session.Values["logged"].(bool); !auth || !ok {
		
		fmt.Println("Warenkorb -> Gast")
		
		Menu = structs.Menu{
			Item1:     "Equipment,equipment", Item2: "Login,login", Item3: "",
			Name:      "", Type: "",
			Basket:    false,
			EmptySide: true,
			Profil:    false, ProfilBild: ""}
		
		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/warenkorb.html", "template/header.html", "template/static_imports.html"))
		
		tmpl.ExecuteTemplate(w, "main", Menu)
		tmpl.ExecuteTemplate(w, "static_imports", Menu)
		tmpl.ExecuteTemplate(w, "header", Menu)
		tmpl.ExecuteTemplate(w, "warenkorb", session.Values["cached_equipment"])
		
	} else {
		
		fmt.Println("Warenkorb -> User")
		
		if r.Method == "POST" {
			// MUSS NOCH
		}
		
		kunden_id_int := session.Values["KundenID"].(int)
		client, _ := model.Get_Kunden_By_ID(kunden_id_int)
		
		Menu = structs.Menu{
			Item1:  "Equipment,equipment", Item2: "Meine Geräte,myEquipment", Item3: "Logout,logout",
			Name:   client.Benutzername, Type: client.Typ,
			Basket: false, EmptySide: false, Profil: true, ProfilBild: client.BildUrl}
		
		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/warenkorb.html", "template/header.html", "template/static_imports.html"))
		
		tmpl.ExecuteTemplate(w, "main", Menu)
		tmpl.ExecuteTemplate(w, "static_imports", Menu)
		tmpl.ExecuteTemplate(w, "header", Menu)
		tmpl.ExecuteTemplate(w, "Warenkorb", Menu)
		
	}
}

func Profil(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	auth, ok := session.Values["logged"];
	
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	
	if !(ok) || !(auth.(bool)) {
		
		fmt.Println("Nicht eingeloggt !")
		http.Redirect(w, r, "/login", 301)
		
	} else {
		//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
		
		if r.Method == "POST" {
			
			fmt.Println("Todo : ", r.FormValue("Todo"))
			
			if strings.Compare(r.FormValue("Todo"), "delete") == 0 {
				
				id, _ := strconv.Atoi(r.FormValue("ID"))
				
				fmt.Println("Lösche Kunden : ", id)
				
				model.DeleteKunde(id)
				
				http.Redirect(w, r, "/logout", 301)
				
			} else {
				
				data_string := r.FormValue("data")
				
				data := strings.Split(data_string, ";")
				
				id, _ := strconv.Atoi(data[0])
				
				fmt.Println("Updaten Kunden : ", id)
				
				model.UpdateKunde(id, data[1], data[2], data[3])
				
				http.Redirect(w, r, "/index", 301)
				
			}
			
		}
		
		//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
		
		kunden_id_int := session.Values["KundenID"].(int)
		client, _ := model.Get_Kunden_By_ID(kunden_id_int)
		
		p := structs.Menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Meine Geräte,myEquipment", Item3: "Logout,logout",
			Basket:    false,
			Name:      client.Benutzername, Type: client.Typ,
			EmptySide: false,
			Profil:    true, ProfilBild: client.BildUrl}
		
		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/user_Profil.html", "template/header.html", "template/static_imports.html"))
		
		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		
		tmpl.ExecuteTemplate(w, "profile", client)
	}
}
