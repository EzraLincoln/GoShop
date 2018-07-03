package route

import (
	"../controller"
	"../structs"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"../../config"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
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
var Equipments = controller.Equipments{}
var Kunden = controller.Kunden{}
var Verleihe = controller.Verleihe{}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Index(w http.ResponseWriter, r *http.Request) {

	p := structs.Menu{
		Item1:     "Equipment,equipment", Item2: "Login,login", Item3: "",
		Basket:    false,
		Name:      "", Type: "",
		EmptySide: false,
		Profile:   false, ProfilBild: ""}

	var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/static_imports.html", "template/index.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)

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

	foo := map[string]interface{}{"structs.Menu": p, "bilder": bilderUrlArray}

	tmpl.ExecuteTemplate(w, "index", foo)

}
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		session, _ := config.CookieStore.Get(r, "session")
		user := r.FormValue("user")
		password := r.FormValue("password")
		result := Kunden.Get_Kunden_By_Name(user)
		if result.Passwort == "" {
			fmt.Println("Kunde in der DB nicht gefunden.")

			http.Redirect(w, r, "/login", 301)

		} else {

			passwordDB := []byte(result.Passwort)

			compr := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))

			if compr == nil {

				session.Values["logged"] = true
				session.Values["KundenID"] = result.KundenID
				session.Save(r, w)

				http.Redirect(w, r, "/profil", 301)
			}

			// http.Redirect(w, r, "/register", 301)

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
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := r.FormValue("user")
		mail := r.FormValue("mail")
		password := r.FormValue("password")
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
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

func Equipment(w http.ResponseWriter, r *http.Request) {

	// session, _ := config.CookieStore.Get(r, "session")

	if r.Method == "POST" {

		sucheNach := r.FormValue("sucheNach")
		kategorie := r.FormValue("kategorie")
		sortierung := r.FormValue("sortierung")

		fmt.Println(sucheNach, kategorie, sortierung)

		http.Redirect(w, r, "/equipment/"+sucheNach+"-"+kategorie+"-"+sortierung, 301)
	}

	// options,_:= strconv.Atoi(r.URL.Path[len("/equipment/"):])
	// fmt.Println(options)

	vars := mux.Vars(r)
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Println(vars["options"])

	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	/// FÜR GÄSTE OHNE EINGELOGGT
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	guestMenu := structs.Menu{
		Item1:  "Login,login", Item2: "Registrieren,register", Item3: "",
		Name:   "", Type: "",
		Basket: false, EmptySide: false, Profil: false, ProfilBild: ""}
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	/// FÜR USER
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	userMenu := structs.Menu{
		Item1:     "", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Name:      "", Type: "",
		Basket:    false,
		EmptySide: false,
		Profil:    false, ProfilBild: ""}
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	fmt.Println(guestMenu)
	EquipmentArr := Equipments.GetEquipment()
	fmt.Println(EquipmentArr)
	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/static_imports.html"))
	tmpl.ExecuteTemplate(w, "main", userMenu)
	tmpl.ExecuteTemplate(w, "static_imports", userMenu)
	tmpl.ExecuteTemplate(w, "header", userMenu)
	tmpl.ExecuteTemplate(w, "equipment", structs.Equipment_Collection{Kategorien: []string{"Kameras", "Mikrofone", "Monitore", "Beleuchtung"}, Items: EquipmentArr})

}

func Meine_Geräte(w http.ResponseWriter, r *http.Request) {
	session, _ := config.CookieStore.Get(r, "session")
	auth, ok := session.Values["logged"];
	if !(ok) || !(auth.(bool)) {
		http.Redirect(w, r, "/login", 301)
	} else {
		kunden_id_int := session.Values["KundenID"].(int)
		client := Kunden.Get_Kunden_By_ID(kunden_id_int)
		p := structs.Menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Logout,logout", Item3: "",
			Basket:    true,
			Name:      client.Benutzername, Type: client.Typ,
			EmptySide: false,
			Profil:    true, ProfilBild: client.BildUrl,
		}
		// ArtikelArr := Equipments.GetUserEquipment(1)
		ArtikelArr := []string{}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/user_MeineGeräte.html", "template/header.html", "template/static_imports.html"))
		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "static_imports", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "myequipment", ArtikelArr)
	}
}
func Warenkorb(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Warenkorb(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := structs.Menu{
		Title:     "borgdir.media,index",
		Item1:     "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
		Basket:    true,
		Name:      "", Type: "",
		EmptySide: false,
		Profil:    true, ProfilBild: ""}

	tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/warenkorb.html", "template/header.html", "template/static_imports.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "static_imports", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "Warenkorb", p)

}
func Profil(w http.ResponseWriter, r *http.Request) {
	session, _ := config.CookieStore.Get(r, "session")
	auth, ok := session.Values["logged"];
	if !(ok) || !(auth.(bool)) {
		http.Redirect(w, r, "/login", 301)
	} else {
		kunden_id_int := session.Values["KundenID"].(int)
		client := Kunden.Get_Kunden_By_ID(kunden_id_int)
		p := structs.Menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment", Item2: "Meine Geräte,myequipment", Item3: "Logout,logout",
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

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Admin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	p := structs.Menu{
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
func Admin_Equipment_Hinzufügen(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin_Equipment_Hinzufügen(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		Equipments.CreateEquipment(w, r)
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

	// ArtikelArr := Equipments.Get_Alle_Equipment()

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
	kunde := Kunden.Get_Kunden_By_ID(1)
	tmpl.ExecuteTemplate(w, "adminEditClients", kunde)
}
func Admin_Equipment_Bearbeiten(w http.ResponseWriter, r *http.Request) {

	fmt.Println("admin_Equipment_Bearbeiten(w http.ResponseWriter, r *http.Request)")
	fmt.Println()

	if r.Method == "POST" {
		Equipments.CreateEquipment(w, r)
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
	kunde := Kunden.Get_Kunden_By_ID(1)
	tmpl.ExecuteTemplate(w, "adminEditClients", kunde)
}

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test(w,r)")
	fmt.Println()

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

	EquipmentArr := Equipments.GetEquipment()

	fmt.Println(EquipmentArr)

	t.Execute(w, EquipmentArr)

	// Info := make(map[string]string)
	// Info["test"] = "About Page"

	// tmpl.ExecuteTemplate(w, "equipment", EquipmentArr)

	// tmpl.ExecuteTemplate(w, "equipment", map[string]interface{}{"mymap": map[string]string{"key": "value"}})

	// t := template.Must(template.New("html_code").Parse(html_code))

}
