package route

import (
	"../controller"
	"../model"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"../../config"
	"golang.org/x/crypto/bcrypt"
)

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

type menu struct {
	Title      string
	Item1      string
	Item2      string
	Item3      string
	Basket     bool
	Name       string
	Type       string
	Profil     bool
	EmptySide  bool
	Profile    bool
	ProfilBild string
}

type Client_Collection struct {
	Items []Client
}
type Client struct {
	BildUrl       string
	Benutzername  string
	KundenID      int
	Typ           string
	Bezeichnungen []Bez
	Status        string
}

// /admin/edit-clients
type Profile struct {
	KundenID     int
	Benutzername string
	BildURL      string
	Mail         string
	Status       string
}

type Bez struct {
	Bezeichnung string
}

type MyEquipment struct {
	Items []model.MyEquipment
}
type admin_Equipment_Collection struct {
	Items []model.Admin_Equipment
}

type EquipmentCollection struct {
	Kategorien []string
	Items      []model.Equipment
}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

var funcMap = template.FuncMap{
	"split": func(s string, index int) string {
		arr := strings.Split(s, ",")

		if s == "" {
			return ""
		} else {
			return arr[index]
		}

	},
}

var Equipments = controller.Equipments{}
var Kunden = controller.Kunden{}
var Verleihe = controller.Verleihe{}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("index(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media, index",
		Item1:     "Equipment,equipment", Item2: "Login,login", Item3: "",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profile:   false, ProfilBild: ""}

	// fmt.Println(p)

	var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/static_imports.html", "template/index.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	// tmpl.ExecuteTemplate(w, "index", p)

	// map[string]interface{}{"mymap": map[string]string{"key": "value"}}

	// foo := map[string]interface{}{"menu" : p,"test" : map[string]string{"key": "value"}}

	fmt.Println(Verleihe.GetAllVerleihe())

	bilderUrlArray := []string{
		"static/images/empty3.png",
		"static/images/empty3.png",
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

	foo := map[string]interface{}{"menu": p, "bilder": bilderUrlArray}
	tmpl.ExecuteTemplate(w, "index", foo)

	// http://placehold.it/250x250"

}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("login(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {

		session, _ := config.CookieStore.Get(r, "session")

		user := r.FormValue("user")
		password := r.FormValue("password")

		fmt.Println("User : ", user, "\n")
		fmt.Println(" Clear Password : ", password, "\n")

		result := Kunden.Get_Kunden_By_Name(user)

		if result.Passwort == "" {
			fmt.Println("Kunde in der DB nicht gefunden.")
			http.Redirect(w, r, "/login", 301)
		} else {

			fmt.Println("KUNDEN GEFUNDEN")
			fmt.Println(result, "\n")
			passwordDB := []byte(result.Passwort)

			compr := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))

			if compr == nil {

				fmt.Println("compare -> nil")

				session.Values["logged"] = true
				session.Values["KundenID"] = result.KundenID
				session.Save(r, w)

				fmt.Println("\nWeiterleitung nach /Profil\n")

				http.Redirect(w, r, "/profil", 301)
			} else {
				fmt.Println("compare not nil")
			}

			// http.Redirect(w, r, "/register", 301)
		}
	}

	p := menu{
		Title:     "borgdir.media, index",
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

func register(w http.ResponseWriter, r *http.Request) {

	fmt.Println("register(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {

		user := r.FormValue("user")
		mail := r.FormValue("mail")
		password := r.FormValue("password")

		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 4)

		fmt.Println("User : ", user, "\n")
		fmt.Println("Mail : ", mail, "\n")
		fmt.Println("Password : ", hash, "\n")

		if Kunden.Test_For_Kunden_By_Name_Mail(user, mail) {

			if Kunden.Register_Kunden(user, mail, string(hash)) {
				fmt.Println("FEHLER :: Hinzufügen des Kunden")
			} else {
				fmt.Println("ERFOLG :: Kunde wurde angelegt !")
				http.Redirect(w, r, "/login", 301)
			}

		} else {

			http.Redirect(w, r, "/register", 301)
			fmt.Println("FEHLER :: Kunde bereits in Datenbank")
		}
	}

	// REGISTER
	p := menu{
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

func equipment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("equipment(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	/// FÜR GÄSTE OHNE EINGELOGGT
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	p := menu{
		Item1: "Login,login", Item2: "Registrieren,register", Item3: "",
		Name:      "", Type: "",
		Basket:    false,EmptySide: false,Profil:    false, ProfilBild: ""}
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	/// FÜR USER
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	p2 := menu{
		//Item1: "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Item1:     "", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Name:      "", Type: "",
		Basket:    false,
		EmptySide: false,
		Profil:    false, ProfilBild: ""}
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

	fmt.Println(p2)

	EquipmentArr := Equipments.GetEquipment()

	// KategorieArr := []string{"hallo","bubu","chingchong","donald"}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "equipment", EquipmentCollection{Kategorien: []string{"Kameras", "Mikrofone", "Monitore", "Beleuchtung"}, Items: EquipmentArr})

	// Info := make(map[string]string)
	// Info["test"] = "About Page"

	// tmpl.ExecuteTemplate(w, "equipment", EquipmentArr)
	// tmpl.ExecuteTemplate(w, "equipment", map[string]interface{}{"mymap": map[string]string{"key": "value"}})

}

func Meine_Geräte(w http.ResponseWriter, r *http.Request) {

	fmt.Println("myequipment(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	// mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

	session, _ := config.CookieStore.Get(r, "session")

	auth, ok := session.Values["logged"];
	fmt.Println("auth :", auth)
	fmt.Println("ok :", ok)

	if !(ok) || !(auth.(bool)) {

		http.Redirect(w, r, "/login", 301)

	} else {

		kunden_id_int := session.Values["KundenID"].(int)

		client := Kunden.Get_Kunden_By_ID(kunden_id_int)

		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Logout,logout", Item3: "",
			Basket:    true,
			Name:      client.Benutzername, Type: client.Typ,
			EmptySide: false,
			Profil:    true, ProfilBild: client.BildUrl}

		// mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

		// Alle Artikel von eingeloggtem Kunden -> var logged_id
		ArtikelArr := Equipments.GetUserEquipment(1)

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/myequipment.html", "template/header.html", "template/static_imports.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "myequipment", MyEquipment{Items: ArtikelArr})

	}
}

func Warenkorb(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Warenkorb(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Basket:    true,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "Warenkorb", p)

}

func Profil(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Profil(w http.ResponseWriter, r *http.Request)")
	fmt.Println("/n/n---------------------------------------------------/n/n")

	session, _ := config.CookieStore.Get(r, "session")

	auth, ok := session.Values["logged"];
	fmt.Println("auth :", auth)
	fmt.Println("ok :", ok)

	// auth, ok = session.Values["logged"].(bool);

	// fmt.Println("Bool Type Assertion !");

	if !(ok) || !(auth.(bool)) {
		http.Redirect(w, r, "/login", 301)
	} else {

		fmt.Println("/n/n---------------------------------------------------/n/n")

		kunden_id_int := session.Values["KundenID"].(int)

		client := Kunden.Get_Kunden_By_ID(kunden_id_int)

		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
			Basket:    false,
			Name:      client.Benutzername, Type: client.Typ,
			EmptySide: false,
			Profil:    true, ProfilBild: client.BildUrl}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/profile.html", "template/header.html", "template/static_imports.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)

		//tmpl.ExecuteTemplate(w, "profile", Profiles{Items: ProfilesArr})

		tmpl.ExecuteTemplate(w, "profile", client)
	}
}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func admin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
		Basket:    false,
		Name:      "Peter", Type: "Verleiher",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "admin", p)

}

func admin_Equipment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin_Equipment(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	// ADMIN
	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}

	// ArtikelArr := Equipments.Get_Alle_Equipment()

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin_Equipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	// tmpl.ExecuteTemplate(w, "admin_Equipment", admin_Equipment_Collection{Items: ArtikelArr})

}

func admin_Equipment_Hinzufügen(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin_Equipment_Hinzufügen(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		Equipments.CreateEquipment(w, r)
		// equipment(w,r)
	} else {

		p := menu{
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

func admin_Equipment_Bearbeiten(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin_Equipment_Bearbeiten(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		Equipments.CreateEquipment(w, r)
		// equipment(w,r)
	} else {

		p := menu{
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

func admin_Kunden_Verwalten(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adminClients(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		// adminEditProfile(w,r)
		// userName := r.
		// KundenID = r.FormValue("KundenID")

		// adminEditProfile(r.PostFormValue())

	} else {

		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
			Basket:    false,
			Name:      "", Type: "",
			EmptySide: false,
			Profil:    true, ProfilBild: ""}

		ClientsArr := []Client{}

		//Alle Kunden auslesen
		KundenArr := Kunden.Get_Alle_Kunden()

		for _, element := range KundenArr {

			artikelFromUser := Equipments.GetUserEquipment(element.KundenID)

			var EquipmentString = []Bez{}

			for _, element := range artikelFromUser {

				EquipmentString = append(EquipmentString, Bez{element.Bezeichnung})
			}

			ClientsArr = append(ClientsArr, Client{element.BildUrl, element.Benutzername, element.KundenID, element.Typ, EquipmentString, element.Status})
		}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/clients.html", "template/header.html", "template/static_imports.html"))

		tmpl.ExecuteTemplate(w, "main", nil)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)

		tmpl.ExecuteTemplate(w, "clients", Client_Collection{Items: ClientsArr,})

	}
}

func admin_Kunden_Bearbeiten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adminEditProfile(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Kunden,clients", Item3: "Logout,logout",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEditClients.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	kunde := Kunden.Get_Kunden_By_ID(1)

	tmpl.ExecuteTemplate(w, "adminEditClients", kunde)

}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test(w,r)")
	fmt.Println()

	/*p := menu{
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

	EquipmentArr := Equipments.GetEquipment()

	fmt.Println(EquipmentArr)

	t.Execute(w, EquipmentArr)

	// Info := make(map[string]string)
	// Info["test"] = "About Page"

	// tmpl.ExecuteTemplate(w, "equipment", EquipmentArr)

	// tmpl.ExecuteTemplate(w, "equipment", map[string]interface{}{"mymap": map[string]string{"key": "value"}})

	// t := template.Must(template.New("html_code").Parse(html_code))

}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Handler() {

	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/", index)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/equipment", admin_Equipment)
	http.HandleFunc("/admin/add", admin_Equipment_Hinzufügen)
	http.HandleFunc("/admin/clients", admin_Kunden_Verwalten)
	http.HandleFunc("/admin/edit-client", admin_Kunden_Bearbeiten)
	http.HandleFunc("/admin/edit-equipment", admin_Equipment_Bearbeiten)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/Warenkorb", Warenkorb)
	http.HandleFunc("/equipment", equipment)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/myequipment", Meine_Geräte)
	http.HandleFunc("/profil", Profil)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/test", test)

	return
}
