package controller

import (
	"../../config"
	"../model"
	"fmt"
	"net/http"
)

func CreateArtikel(w http.ResponseWriter, r *http.Request) {

	bezeichnung := r.FormValue("bz")
	kategorie := r.FormValue("kat")
	inventarNummer := r.FormValue("invNum")
	lagerort := r.FormValue("lgo")
	inhalt := r.FormValue("inhalt")
	hinweis := r.FormValue("hinweis")
	anzahl := r.FormValue("anz")

	statement := "insert into Artikel (Bezeichnung, Kategorie,InventarNummer,Lagerort, Anzahl,Hinweis,BildURL) values (?,?,?,?,?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(bezeichnung, kategorie, inventarNummer, lagerort, inhalt, hinweis, anzahl)

	return
}

// Update the Artikel-Bezeichnung by id
func UpdateArtikel(id int, bez string, kat string, lago string, anz int, hin string, url string) (err error) {
	_, err = config.Db.Exec("update Artikel set Bezeichnung = $1 where ArtikelID = $2", bez, id)
	_, err = config.Db.Exec("update Artikel set Kategorie = $1 where ArtikelID = $2", kat, id)
	_, err = config.Db.Exec("update Artikel set Lagerort = $1 where ArtikelID = $2", lago, id)
	_, err = config.Db.Exec("update Artikel set Anzahl = $1 where ArtikelID = $2", anz, id)
	_, err = config.Db.Exec("update Artikel set Hinweis = $1 where ArtikelID = $2", hin, id)
	_, err = config.Db.Exec("update Artikel set BildURL = $1 where ArtikelID = $2", url, id)
	return
}

// Delete Artikel by id
func DeleteArtikel(id int) (err error) {
	_, err = config.Db.Exec("delete from Artikel where ArtikelID = $1", id)
	return
}

func GetEquipment() (Artikels []model.Artikel) {
	// rows, err := config.Db.Query("select BildURL, Bezeichnung, Anzahl, Hinweis FROM Artikel")

	// fmt.Println("DB : ",config.ReturnDB())
	// fmt.Println("DB : ",config.Db)

	/* rows, err := config.ReturnDB().Query("SELECT bezeichnung FROM artikel")
	fmt.Println(err)

	for rows.Next() {
		var bezeichnung string
		err = rows.Scan(&bezeichnung)
		fmt.Println(err)
		fmt.Println("bezeichnung")
		fmt.Printf("%8v\n", bezeichnung)
	}
	*/

	rows, err := config.Db.Query("SELECT Bezeichnung FROM artikel")

	if err != nil {
		fmt.Println("Error (1) in Controller - GetEquipment()")
	}

	Artikel := model.Artikel{}

	for rows.Next() {

		// err = rows.Scan(&Artikel.ArtikelID,&Artikel.Bezeichnung,&Artikel.Kategorie,&Artikel.InventarNummer,&Artikel.Lagerort,&Artikel.Anzahl,&Artikel.Hinweis,&Artikel.BildURL)
		err = rows.Scan(&Artikel.Bezeichnung)

		fmt.Println(Artikel)

		Artikels = append(Artikels, Artikel)

		if err != nil {
			fmt.Println("Error (2) in Controller - GetEquipment()")
		}
	}
	rows.Close()

	return
}

func GetAllBezeichnungenFromKundenArtikel(kunde_id int) (Bezeichnungen []string) {

	rows, err := config.Db.Query("select Artikel.Bezeichnung from Artikel,Verleih where Verleih.KundenID=$1 and Artikel.ArtikelID = Verleih.ArtikelID", kunde_id)

	if err != nil {
		return
	}

	var temp = ""

	for rows.Next() {
		err = rows.Scan(&temp)

		Bezeichnungen = append(Bezeichnungen, temp)

		if err != nil {
			return
		}
	}
	rows.Close()

	return
}

func GetUserEquipment(kunde_id int) (equipments []model.MyEquipment) {
	rows, err := config.Db.Query("select Artikel.BildURL, Artikel.Bezeichnung, Artikel.InventarNummer, Artikel.Hinweis, Verleih.Beginn, Verleih.Rueckgabe from Artikel,Verleih WHERE Artikel.ArtikelID = Verleih.ArtikelID AND Verleih.KundenID=$1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		equipment := model.MyEquipment{}

		err = rows.Scan(&equipment.BildURL, &equipment.Bezeichnung, &equipment.InventarNummer, &equipment.Hinweis, &equipment.Beginn, &equipment.Rueckgabe)

		if err != nil {
			return
		}

		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

func GetAdminEquipment(kunde_id int) (adminEquipments []model.AdminEquipments) {
	rows, err := config.Db.Query("select Artikel.Bezeichnung, Artikel.InventarNummer, Artikel.Lagerort Artikel.Hinweis, Kunde.Benutzername, Verleih.Rueckgabe from Artikel,Verleih,Kunde WHERE Artikel.ArtikelID = Verleih.ArtikelID AND Verleih.KundenID=$1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		adminEquipment := model.AdminEquipments{}

		err = rows.Scan(&adminEquipment.Bezeichnung, &adminEquipment.InventarNummer, &adminEquipment.Lagerort, &adminEquipment.Hinweis, &adminEquipment.Benutzername, &adminEquipment.Rueckgabe)

		if err != nil {
			return
		}

		adminEquipments = append(adminEquipments, adminEquipment)
	}
	rows.Close()
	return
}
