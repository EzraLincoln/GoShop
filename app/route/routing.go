package route

import (
	"../controller"
	"../model"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	//	"log"
)

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

type menu struct {
	Title     string
	Item1     string
	Item2     string
	Item3     string
	Basket    bool
	Name      string
	Type      string
	Profil    bool
	EmptySide bool
	Profile   bool
}

type Clients struct {
	Items []client
}

type client struct {
	BildUrl       string
	Benutzername  string
	KundenID      int
	Typ           string
	Bezeichnungen []Bez
	Status        string
}

type Bez struct {
	Bezeichnung string
}

type MyEquipment struct {
	Items []model.MyEquipment
}

type AdminEquipments struct {
	Items []model.AdminEquipments
}

type Equipment struct {
	Kategorien []string
	Items      []model.Equipment
}

type Profiles struct {
	Items []model.Profile
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

//var artikelList = make(model.Artikels)

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm


var Equipments = controller.Equipments{}
var Kunden= controller.Kunden{}
var Verleihe = controller.Verleihe{}


//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

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
		Profile:   false}

	// fmt.Println(p)

	var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/static_imports.html", "template/index.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	// tmpl.ExecuteTemplate(w, "index", p)

	// map[string]interface{}{"mymap": map[string]string{"key": "value"}}

	// foo := map[string]interface{}{"menu" : p,"test" : map[string]string{"key": "value"}}

	fmt.Println(Verleih.GetAllVerleihe())

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
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "admin", p)

}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("login(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		EmptySide: true,
		Profile:   false}

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

		controller.RegisterKunden(w, r)
		//userName := r.FormValue("KundenID")
		//fmt.Print(userName)
		index(w, r)
	} else {

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
			Profile:   false}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/register.html", "template/static_imports.html", "template/header.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "register", p)
	}
}

func equipment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("equipment(w http.ResponseWriter, r *http.Request)")
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
		Profile:   true}

	EquipmentArr := controller.GetEquipment()

	// KategorieArr := []string{"hallo","bubu","chingchong","donald"}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "equipment", Equipment{Kategorien: []string{"Kameras", "Mikrofone", "Monitore", "Beleuchtung"}, Items: EquipmentArr})

	// Info := make(map[string]string)
	// Info["test"] = "About Page"

	// tmpl.ExecuteTemplate(w, "equipment", EquipmentArr)
	// tmpl.ExecuteTemplate(w, "equipment", map[string]interface{}{"mymap": map[string]string{"key": "value"}})

}
func myequipment(w http.ResponseWriter, r *http.Request) {

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
		Profile:   true}

	// Alle Artikel von eingeloggtem Kunden -> var logged_id
	ArtikelArr := controller.GetUserEquipment(1)

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
		Profile:   true}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "cart", p)

}

func profile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("profile(w http.ResponseWriter, r *http.Request)")
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
		Profile:   true}

	ProfilesArr := controller.GetProfile(1)

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/profile.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	tmpl.ExecuteTemplate(w, "profile", Profiles{Items: ProfilesArr})

}

func adminEquipment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adminEquipment(w http.ResponseWriter, r *http.Request)")
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
		Profile:   true}

	ArtikelArr := controller.GetAdminEquipment(1)

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEquipment.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "adminEquipment", AdminEquipments{Items: ArtikelArr})

}

func adminAddEquipment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adminAddEquipment(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		controller.CreateArtikel(w, r)
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
			Profile:   true}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminAddEquipment.html", "template/header.html", "template/static_imports.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "adminAddEquipment", p)
	}

}
func adminEditEquipment(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adminEditEquipment(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		controller.CreateArtikel(w, r)
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
			Profile:   true}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEditEquipment.html", "template/header.html", "template/static_imports.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "adminEditEquipment", p)
	}

}

func adminProfiles(w http.ResponseWriter, r *http.Request) {

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
			Profile:   true}

		//Alle Kunden auslesen
		KundenArr := controller.GetAllUser()

		var ClientsArr = []client{}

		// for index := range ClientsArr {
		for _, element := range KundenArr {
			// ClientsArr = append(ClientsArr,client{controller.getKundenById(controller.getVerleihById(index).kundeID)).bildUrl,"asdasd","asdasd","asdasd","asdasd","asdasdad",},)

			artikelFromUser := controller.GetAllBezeichnungenFromKundenArtikel(element.KundeID)

			var EquipmentString = []Bez{}

			for _, element := range artikelFromUser {

				EquipmentString = append(EquipmentString, Bez{element})
			}

			ClientsArr = append(ClientsArr, client{element.BildUrl, element.Benutzername, element.KundeID, element.Typ, EquipmentString, element.Status})
		}

		data := Clients{
			Items: ClientsArr,
		}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/clients.html", "template/header.html", "template/static_imports.html"))

		tmpl.ExecuteTemplate(w, "main", nil)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "clients", data)

	}
}
func adminEditProfile(w http.ResponseWriter, r *http.Request) {

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
		Profile:   true}

	ClientArr := controller.GetProfile(1)

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEditProfile.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

	tmpl.ExecuteTemplate(w, "adminEditProfile", Profiles{Items: ClientArr})

}

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

	EquipmentArr := controller.GetEquipment()

	fmt.Println(EquipmentArr)

	t.Execute(w, EquipmentArr)

	// Info := make(map[string]string)
	// Info["test"] = "About Page"

	// tmpl.ExecuteTemplate(w, "equipment", EquipmentArr)
	// tmpl.ExecuteTemplate(w, "equipment", map[string]interface{}{"mymap": map[string]string{"key": "value"}})

	// t := template.Must(template.New("html_code").Parse(html_code))

}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Handler() {

	fmt.Println("Aufruf Handler()")

	fmt.Println()

	http.HandleFunc("/", index)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/equipment", adminEquipment)
	http.HandleFunc("/admin/add", adminAddEquipment)
	http.HandleFunc("/admin/clients", adminProfiles)
	http.HandleFunc("/admin/edit-clients", adminEditProfile)
	http.HandleFunc("/admin/edit-equipment", adminEditEquipment)
	http.HandleFunc("/login", login)
	http.HandleFunc("/equipment", equipment)
	http.HandleFunc("/myequipment", myequipment)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/register", register)
	http.HandleFunc("/cart", cart)

	http.HandleFunc("/test", test)

	return
}
