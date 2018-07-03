package controller

import (
	"../structs"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"../../config"
	"golang.org/x/crypto/bcrypt"
	"../model"
)

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

var funcMap = template.FuncMap{"split": func(s string, index int) string {
	arr := strings.Split(s, ",")
	if s == "" {
		return ""
	} else {
		return arr[index]
	}
},}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Index(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	
	Menu := structs.Menu{}
	
	if auth, ok := session.Values["logged"].(bool); !auth || !ok {
		
		fmt.Println("Index -> Gast")
		
		Menu = structs.Menu{
			Item1:     "Equipment,equipment", Item2: "Login,login", Item3: "",
			Name:      "", Type: "",
			Basket:    false,
			EmptySide: false,
			Profil:    false, ProfilBild: ""}
	} else {
		
		fmt.Println("Index -> User")
		
		kunden_id_int := session.Values["KundenID"].(int)
		client,_ := model.Get_Kunden_By_ID(kunden_id_int)
		
		Menu = structs.Menu{
			Item1:  "Equipment,equipment", Item2: "Meine Geräte,myEquipment", Item3: "Logout,logout",
			Name:   client.Benutzername, Type: client.Typ,
			Basket: false, EmptySide: false, Profil: true, ProfilBild: client.BildUrl}
	}
	var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/static_imports.html", "template/index.html"))
	
	tmpl.ExecuteTemplate(w, "main", Menu)
	tmpl.ExecuteTemplate(w, "static_imports", Menu)
	tmpl.ExecuteTemplate(w, "header", Menu)
	
	bilderUrlArray := []string{
		"static/Media/Equipment/1.jpg",
		"static/Media/Equipment/2.jpg",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png",
		"static/images/empty3.png"}
	
	foo := map[string]interface{}{"structs.Menu": Menu, "bilder": bilderUrlArray}
	
	tmpl.ExecuteTemplate(w, "index", foo)
	
}

func Login(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "POST" {
		
		session, _ := config.CookieStore.Get(r, "session")
		
		user := r.FormValue("user")
		password := r.FormValue("password")
		
		fmt.Println("Suche nach Nutzer : ", user)
		
		result, fehler := model.Get_Kunden_By_Name(user)
		
		if (fehler) {
			http.Redirect(w, r, "/register", 301)
		}
		
		fmt.Println("Kunden mit Passwort (", result.Passwort, ") in DB vorhanden.")
		
		passwordDB := []byte(result.Passwort)
		
		compr := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
		
		if compr == nil {
			
			session.Values["logged"] = true
			session.Values["KundenID"] = result.KundenID
			session.Values["accountTyp"] = result.Typ
			session.Save(r, w)
			
			http.Redirect(w, r, "/profil", 301)
		}
	}
	
	p := structs.Menu{
		Item1:     "Equipment,equipment", Item2: "Login,login", Item3: "",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profile:   false, ProfilBild: ""}
	
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/static_imports.html", "template/login.html"))
	
	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "login", p)
	
}

func Logout(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	
	fmt.Println(session.Values)
	
	session.Values["logged"] = false
	session.Values["KundenID"] = ""
	session.Values["accountTyp"] = ""
	session.Save(r, w)
	
	http.Redirect(w, r, "/index", 301)
}

func Register(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "POST" {
		
		user := r.FormValue("user")
		mail := r.FormValue("mail")
		password := r.FormValue("password")
		
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
		
		if model.Test_For_Kunden_By_Name_Mail(user, mail) {
			
			fmt.Println("Abbruch :: Kunde bereits in Datenbank")
			http.Redirect(w, r, "/register", 301)
			
		} else {
			
			model.Register_Kunden(user, mail, string(hash))
			
			fmt.Println("ERFOLG :: Kunde wurde angelegt !")
			
			http.Redirect(w, r, "/login", 301)
			
		}
	}
	
	// REGISTER
	p := structs.Menu{
		Title:     "borgdir.media, index",
		Item1:     "Equipment,equipment", Item2: "Login,login", Item3: "",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profile:   false, ProfilBild: ""}
		
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/register.html", "template/static_imports.html", "template/header.html"))
	
	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	
	tmpl.ExecuteTemplate(w, "register", p)
}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Test(w http.ResponseWriter, r *http.Request) {
	
	/*p := structs.Menu{
	Title:     "borgdir.media, index",
	Item1:     "Equipment,equipment",
	Item2:     "Login,login",
	Item3:     "",
	Basket:    false,
	Name:      "",
	Type:      "",
	SecondCart: false,
	Profile:   false}
	*/
	
	// fmt.Println(p)
	
	// tOk := template.New("first")
	
	// var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/static_imports.html", "template/index.html"))
	
	// tmpl.ExecuteTemplate(w, "main", p)
	// tmpl.ExecuteTemplate(w, "static_imports", p)
	// tmpl.ExecuteTemplate(w, "header", p)
	
	// tmpl.ExecuteTemplate(w, "index", p)
	
	// tmpl.Execute(os.Stdout, "HALLO")
	const html_code = `{{index . 3}}`
	
	/*type Text struct {
		text string
	}*/
	
	t := template.Must(template.New("html_code").Parse(html_code))
	
	// t.Execute(w,"test")
	
	EquipmentArr := model.GetEquipment()
	
	fmt.Println(EquipmentArr)
	
	t.Execute(w, EquipmentArr)
	
	// Info := make(map[string]string)
	// Info["test"] = "About Page"
	
	// tmpl.ExecuteTemplate(w, "equipment", EquipmentArr)
	
	// tmpl.ExecuteTemplate(w, "equipment", map[string]interface{}{"mymap": map[string]string{"key": "value"}})
	
	// t := template.Must(template.New("html_code").Parse(html_code))
	
}