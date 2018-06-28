package model

// /myequipment Seiten Struct
type MyEquipment struct {
	BildURL        string
	Bezeichnung    string
	InventarNummer int
	Hinweis        string
	Beginn         string
	Rueckgabe      string
}

// /admin/equipment Seiten Struct
type AdminEquipments struct {
	BildURL        string
	Bezeichnung    string
	InventarNummer int
	Lagerort       string
	Hinweis        string
	Benutzername   string
	Rueckgabe      string
}

// /equipment Seiten Struct
type Equipment struct {
	BildURL     string
	Bezeichnung string
	Anzahl      int
	Hinweis     string
}

// /admin/edit-clients
type Profile struct {
	KundenID     int
	Benutzername string
	BildURL      string
	Mail         string
	Status       string
}
