package route

import (
	"../controller"
	"../model"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

type menu struct {
	Title  string
	Item1  string
	Item2  string
	Item3  string
	Basket bool
	Name   string
	Type   string
	// Profil    bool
	EmptySide bool
	Profil    bool
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

var store *sessions.CookieStore

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

//var artikelList = make(model.Artikels)

var Equipments = controller.Equipments{}
var Kunden = controller.Kunden{}
var Verleihe = controller.Verleihe{}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("index(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media, index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profil:    false}

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

		user := r.FormValue("user")
		password := r.FormValue("password")

		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 4)

		fmt.Println("User : ", user, " Password : ", password)

		result := Kunden.Get_Kunden_By_Name(user)

		if result == "" {
			fmt.Println("Kunde in der DB nicht gefunden.")
			http.Redirect(w, r, "/login", 301)
		} else {

			fmt.Println(result)
			passwordDB := []byte(result)

			fmt.Println(passwordDB)
			fmt.Println(hash)

			bcrypt.CompareHashAndPassword(passwordDB, hash)

			// http.Redirect(w, r, "/register", 301)
		}
	}

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: true,
		Profil:    false}

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

		fmt.Println("User : ", user)
		fmt.Println("Mail : ", mail)
		fmt.Println("Password : ", hash)

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
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: true,
		Profil:    false}

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
		Title:     "borgdir.media,index", Item1: "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Name:      "", Type: "",
		Basket:    false,
		EmptySide: false,
		Profil:    false}
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	/// FÜR USER
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	p2 := menu{
		Title:     "borgdir.media,index", Item1: "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Name:      "", Type: "",
		Basket:    false,
		EmptySide: false,
		Profil:    false}
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

func equipmentAlternative(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")

	// p := PageData{}
	// c := model.Client{}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("Not authenticated")
		// p.Client = c
		// c.Gesperrt = false //Gespert wird hier genutzt, um festzustellen, ob ein Benutzer angemeldet ist (für Links in navbar)

	} else {
		fmt.Println("authenticated")
		// c, err := model.ReadClientByName(DB, session.Values["username"].(string))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// c.Gesperrt = true
		// p.Client = c
	}

	// var equipment []model.Equipment
	// var err error
	// equipment, err = model.ReadAllEquipment(DB)
	// if err != nil {
	//	log.Fatal(err)
	// }
	//
	// for k := 0; k < len(equipment); k++ {
	//	if equipment[k].Amount > 0 {
	//		equipment[k].Active = true
	//	} else {
	//		equipment[k].Active = false
	//	}
	// }
	//
	// var check bool = false
	//
	// for l := 0; l < len(equipment); l++ {
	//	for m := 0; m < len(p.Category); m++ {
	//		if strings.Compare(equipment[l].Category, p.Category[m]) == 0 {
	//			check = true
	//			break
	//		}
	//	}
	//	if !check {
	//		p.Category = append(p.Category, equipment[l].Category)
	//	}
	//	check = false
	// }
	//
	// for i := 0; i < len(equipment); i++ {
	//	p.StoreEquipment = append(p.StoreEquipment, equipment[i])
	// }

	// tmpl, _ := template.ParseFiles("template/equipment.html", "template/head.html", "template/foot.html")
	// tmpl.Execute(w, p)
}

func user_Meine_Geräte(w http.ResponseWriter, r *http.Request) {

	fmt.Println("myequipment(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profil:    true}

	// Alle Artikel von eingeloggtem Kunden -> var logged_id
	ArtikelArr := Equipments.GetUserEquipment(1)

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/myequipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "myequipment", MyEquipment{Items: ArtikelArr})

}

func cart(w http.ResponseWriter, r *http.Request) {

	fmt.Println("cart(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profil:    true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "cart", p)

}

func user_profil(w http.ResponseWriter, r *http.Request) {

	fmt.Println("profil(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Meine Geräte,myequipment",
		Item3:     "Logout,logout",
		Basket:    true,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profil:    true}

	request := Kunden.Get_Kunden_By_ID(1)

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/profile.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	//tmpl.ExecuteTemplate(w, "profile", Profiles{Items: ProfilesArr})

	tmpl.ExecuteTemplate(w, "profile", request)
}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func admin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "Peter",
		Type:      "Verleiher",
		EmptySide: false,
		Profil:    true}

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
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profil:    true}

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
			Item1:     "Equipment,equipment",
			Item2:     "Kunden,clients",
			Item3:     "Logout,logout",
			Basket:    false,
			Name:      "",
			Type:      "",
			EmptySide: false,
			Profil:    true}

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
			Item1:     "Equipment,equipment",
			Item2:     "Kunden,clients",
			Item3:     "Logout,logout",
			Basket:    false,
			Name:      "",
			Type:      "",
			EmptySide: false,
			Profil:    true}

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
			Item1:     "Equipment,equipment",
			Item2:     "Kunden,clients",
			Item3:     "Logout,logout",
			Basket:    false,
			Name:      "",
			Type:      "",
			EmptySide: false,
			Profil:    true}

		ClientsArr := []Client{}

		//Alle Kunden auslesen
		KundenArr := Kunden.Get_Alle_Kunden()

		for _, element := range KundenArr {

			artikelFromUser := Equipments.GetUserEquipment(element.KundeID)

			var EquipmentString = []Bez{}

			for _, element := range artikelFromUser {

				EquipmentString = append(EquipmentString, Bez{element.Bezeichnung})
			}

			ClientsArr = append(ClientsArr, Client{element.BildUrl, element.Benutzername, element.KundeID, element.Typ, EquipmentString, element.Status})
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
		Item1:     "Equipment,equipment",
		Item2:     "Kunden,clients",
		Item3:     "Logout,logout",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: false,
		Profil:    true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEditProfile.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	tmpl.ExecuteTemplate(w, "adminEditProfile", Kunden.Get_Kunden_By_ID(1))

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
	EmptySide: false,
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
	http.HandleFunc("/cart", cart)
	http.HandleFunc("/equipment", equipment)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/myequipment", user_Meine_Geräte)
	http.HandleFunc("/profil", user_profil)
	///////////////////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/test", test)

	return
}
