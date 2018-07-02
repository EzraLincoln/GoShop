package model

// Artikel data structure
type Equipment struct {
	EquipmentID    int
	Bezeichnung    string
	Kategorie      string
	InventarNummer int
	Lagerort       string
	Anzahl         int
	Hinweis        string
	BildURL        string
	VerleiherID    int
	Status         string
}

// /admin/equipment Seiten Struct
type Admin_Equipment struct {
	BildURL        string
	Bezeichnung    string
	InventarNummer int
	Lagerort       string
	Hinweis        string
	Benutzername   string
	Rueckgabe      string
}
