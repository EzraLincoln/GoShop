package controller

import (
	"../../config"
	"../model"
	"fmt"
	"net/http"
)

type Equipments struct{}

func (v Equipments) CreateEquipment(w http.ResponseWriter, r *http.Request) {

	bezeichnung := r.FormValue("bz")
	kategorie := r.FormValue("kat")
	inventarNummer := r.FormValue("invNum")
	lagerort := r.FormValue("lgo")
	inhalt := r.FormValue("inhalt")
	hinweis := r.FormValue("hinweis")
	anzahl := r.FormValue("anz")

	statement := "insert into Equipment (Bezeichnung, Kategorie,InventarNummer,Lagerort, Anzahl,Hinweis,BildURL) values (?,?,?,?,?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(bezeichnung, kategorie, inventarNummer, lagerort, inhalt, hinweis, anzahl)

	return
}

// Update the Equipment-Bezeichnung by id
func (v Equipments) UpdateEquipment(id int, bez string, kat string, lago string, anz int, hin string, url string) (err error) {
	_, err = config.Db.Exec("update Equipment set Bezeichnung = $1 where EquipmentID = $2", bez, id)
	_, err = config.Db.Exec("update Equipment set Kategorie = $1 where EquipmentID = $2", kat, id)
	_, err = config.Db.Exec("update Equipment set Lagerort = $1 where EquipmentID = $2", lago, id)
	_, err = config.Db.Exec("update Equipment set Anzahl = $1 where EquipmentID = $2", anz, id)
	_, err = config.Db.Exec("update Equipment set Hinweis = $1 where EquipmentID = $2", hin, id)
	_, err = config.Db.Exec("update Equipment set BildURL = $1 where EquipmentID = $2", url, id)
	return
}

// Delete Equipment by id
func (v Equipments) DeleteEquipment(id int) (err error) {
	_, err = config.Db.Exec("delete from Equipment where EquipmentID = $1", id)
	return
}

func (v Equipments) GetEquipment() (Equipments []model.Equipment) {
	// rows, err := config.Db.Query("select BildURL, Bezeichnung, Anzahl, Hinweis FROM Equipment")

	// fmt.Println("DB : ",config.ReturnDB())
	// fmt.Println("DB : ",config.Db)

	/* rows, err := config.ReturnDB().Query("SELECT bezeichnung FROM Equipment")
	fmt.Println(err)

	for rows.Next() {
		var bezeichnung string
		err = rows.Scan(&bezeichnung)
		fmt.Println(err)
		fmt.Println("bezeichnung")
		fmt.Printf("%8v\n", bezeichnung)
	}
	*/

	rows, err := config.Db.Query("SELECT * FROM Equipment")

	if err != nil {
		fmt.Println("Error (1) in Controller - GetEquipment()")
	}

	Equipment := model.Equipment{}

	for rows.Next() {

		err = rows.Scan(
			&Equipment.EquipmentID,
			&Equipment.Bezeichnung,
			&Equipment.Kategorie,
			&Equipment.InventarNummer,
			&Equipment.Lagerort,
			&Equipment.Anzahl,
			&Equipment.Hinweis,
			&Equipment.BildURL,
			&Equipment.VerleiherID,
			&Equipment.Status,
		)

		// err = rows.Scan(&Equipment.Bezeichnung)

		fmt.Println()

		fmt.Println(Equipment.EquipmentID)
		fmt.Println(Equipment.Bezeichnung)
		fmt.Println(Equipment.Kategorie)
		fmt.Println(Equipment.InventarNummer)
		fmt.Println(Equipment.Lagerort)
		fmt.Println(Equipment.Anzahl)
		fmt.Println(Equipment.Hinweis)
		fmt.Println(Equipment.BildURL)
		fmt.Println(Equipment.VerleiherID)
		fmt.Println(Equipment.Status)

		fmt.Println()

		Equipments = append(Equipments, Equipment)

		if err != nil {
			fmt.Println("Error (2) in Controller - GetEquipment()")
		}
	}
	rows.Close()

	return
}

func (v Equipments) GetAllBezeichnungenFromKundenEquipment(kunde_id int) (Bezeichnungen []string) {

	rows, err := config.Db.Query("select Equipment.Bezeichnung from Equipment,Verleih where Verleih.KundenID=$1 and Equipment.EquipmentID = Verleih.EquipmentID", kunde_id)

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

func (v Equipments) GetUserEquipment(kunde_id int) (equipments []model.MyEquipment) {
	rows, err := config.Db.Query("select Equipment.BildURL, Equipment.Bezeichnung, Equipment.InventarNummer, Equipment.Hinweis, Verleih.Beginn, Verleih.Rueckgabe from Equipment,Verleih WHERE Equipment.EquipmentID = Verleih.EquipmentID AND Verleih.KundenID=$1", kunde_id)

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

func (v Equipments) Get_Admin_Equipment_By_Kunden_ID(kunde_id int) (adminEquipments []model.Admin_Equipment) {
	rows, err := config.Db.Query("select Equipment.Bezeichnung, Equipment.InventarNummer, Equipment.Lagerort Equipment.Hinweis, Kunde.Benutzername, Verleih.Rueckgabe from Equipment,Verleih,Kunde WHERE Equipment.EquipmentID = Verleih.EquipmentID AND Verleih.KundenID=$1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		adminEquipment := model.Admin_Equipment{}

		err = rows.Scan(&adminEquipment.Bezeichnung, &adminEquipment.InventarNummer, &adminEquipment.Lagerort, &adminEquipment.Hinweis, &adminEquipment.Benutzername, &adminEquipment.Rueckgabe)

		if err != nil {
			return
		}

		adminEquipments = append(adminEquipments, adminEquipment)
	}
	rows.Close()
	return
}

func (v Equipments) Get_Alle_Equipment() (equipments []model.Equipment) {
	rows, err := config.Db.Query("select * from Equipment where Typ = 'Benutzer'")

	if err != nil {
		return
	}

	for rows.Next() {
		equipment := model.Equipment{}
		err = rows.Scan(&equipment.EquipmentID, &equipment.Bezeichnung, &equipment.Kategorie, &equipment.InventarNummer, &equipment.Anzahl, &equipment.Hinweis, &equipment.BildURL,&equipment.VerleiherID,&equipment.Status)

		if err != nil {
			return
		}

		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}
