package structs

import "../model"

type Menu struct {
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
	Bezeichnungen []Bezeichnung
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
type Bezeichnung struct {
	Bezeichnung string
}
type Admin_Equipment_Collection struct {
	Items []model.Admin_Equipment
}
type Equipment_Collection struct {
	Kategorien []string
	Items      []model.Equipment
}
