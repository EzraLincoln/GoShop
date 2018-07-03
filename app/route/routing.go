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
	"strconv"
	"log"
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
		client, fehler := Kunden.Get_Kunden_By_ID(kunden_id_int)
		
		/*<*/ check(w,r,fehler,"Fehler : Get Kunden By ID : login") /*>*/
		
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
		
		result, fehler := Kunden.Get_Kunden_By_Name(user)
		
		/*<*/ check(w,r,fehler,"Login : Get Kunden By Name : login") /*>*/
		
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
		
		if Kunden.Test_For_Kunden_By_Name_Mail(user, mail) {
			
			fehler := Kunden.Register_Kunden(user, mail, string(hash))
			
			/*<*/ check(w,r,fehler,"Register : Fehler bei Register_Kunden : register") /*>*/
			
			fmt.Println("ERFOLG :: Kunde wurde angelegt !")
			
			http.Redirect(w, r, "/login", 301)
			
		} else {
			
			fmt.Println("Abbruch :: Kunde bereits in Datenbank")
			http.Redirect(w, r, "/register", 301)
			
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
		client, fehler := Kunden.Get_Kunden_By_ID(kunden_id_int)
		
		/*<*/ check(w,r,fehler,"Index : Get Kunden By ID : login") /*>*/
		
		Menu = structs.Menu{
			Item1:  "Meine Geräte,myEquipment", Item2: "Logout,logout", Item3: "",
			Name:   client.Benutzername, Type: client.Typ,
			Basket: false, EmptySide: false, Profil: true, ProfilBild: client.BildUrl}
	}
	
	EquipmentArr := Equipments.GetEquipment()
	
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
		client, fehler := Kunden.Get_Kunden_By_ID(kunden_id_int)
		
		/*<*/ check(w,r,fehler,"Meine Geräte : Get Kunden By ID : login") /*>*/
		
		Menu := structs.Menu{
			Item1:     "Equipment,equipment", Item2: "Logout,logout", Item3: "",
			Basket:    false,
			Name:      client.Benutzername, Type: client.Typ,
			EmptySide: false,
			Profil:    true, ProfilBild: client.BildUrl,
		}
		// ArtikelArr := Equipments.GetUserEquipment(1)
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
		client, fehler := Kunden.Get_Kunden_By_ID(kunden_id_int)
		
		/*<*/ check(w,r,fehler,"Warenkorb : Get Kunden By ID : login") /*>*/
		
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
			
			fmt.Println(r.FormValue("Todo"))
			fmt.Println(r.FormValue("ID"))
			
			id, _ := strconv.Atoi(r.FormValue("ID"))
			
			if strings.Compare(r.FormValue("Todo"), "update") != 0 {
				
				fmt.Println("Lösche Kunden : ", id)
				
				Kunden.Delete_Kunden_By_Name("Alex")
				
				// http.Redirect(w, r, "/logout", 301)
			}
			
		}
	}
	//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm
	kunden_id_int := session.Values["KundenID"].(int)
	client, fehler := Kunden.Get_Kunden_By_ID(kunden_id_int)
	
	/*<*/ check(w,r,fehler,"Profil : Get Kunden By ID : logout") /*>*/
	
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

//mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm

func Admin(w http.ResponseWriter, r *http.Request) {
	
	session, _ := config.CookieStore.Get(r, "session")
	
	auth, ok := session.Values["logged"];
	if !(ok) || !(auth.(bool)) {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["user-type"].(string) == "Verleiher" {
			
			kunden_id_int := session.Values["KundenID"].(int)
			client, fehler := Kunden.Get_Kunden_By_ID(kunden_id_int)
			
			/*<*/ check(w,r,fehler,"Admin : Get Kunden By ID : logout") /*>*/
			
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
	
	kunde, fehler := Kunden.Get_Kunden_By_ID(1)
	
	/*<*/ check(w,r,fehler,"Admin Kunden Verwalten : Get Kunden By ID : logout") /*>*/
	
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
	
	kunde, fehler := Kunden.Get_Kunden_By_ID(1)
	
	/*<*/ check(w,r,fehler,"Admin Kunden Bearbeiten : Get Kunden By ID : logout") /*>*/
	
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

func check(w http.ResponseWriter, r *http.Request, f bool, msg string) {
	if (f) {
		s := strings.Split(msg, " : ")
		log.Fatalln(s[0] + " : " + s[1])
		http.Redirect(w, r, "/"+s[2], 301)
	}
}