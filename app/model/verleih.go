package model

import (
	"../../config"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

// Verleiher data structure
type Verleih struct {
	VerleihID int
	KundenID  int
	ArtikelID int
	Beginn    string
	Rueckgabe string
}

type Verleihe map[int]*Verleih

// Create Verleiher
func Add_Verleih(EquipmentID int, KundenID int, beginn string, rückgabe string, anzahl int) {
	//defer stmt.Close()
	_, err := config.Db.Exec("insert into Verleih (KundenID,ArtikelID,Beginn,Rueckgabe) values ($1,$2,$3,$4)", EquipmentID, KundenID, beginn, rückgabe)

if err != nil {
	fmt.Println("Fehler beim Verleih Insert")
}
return
}

// Read Verleiher by VerleiherID
func ReadVerleiher(id int) (verleiher Verleih, err error) {
	verleiher = Verleih{}

	// err = config.Db.QueryRow("select *  from Verleiher where VerleiherID = $1", id).Scan(&verleiher.VerleiherID, &verleiher.Benutzername, &verleiher.Email)

	return
}

// Update the Verleiher by id
func UpdateVerleiher(id int, bname string, mail string) (err error) {
	_, err = config.Db.Exec("update Verleiher set Benutzername = $1 where VerleiherID = $2", bname, id)
	_, err = config.Db.Exec("update Verleiher set Email = $1 where VerleiherID = $2", mail, id)
	return
}

// Delete Verleiher by id
func Delete_Verleih_By_Kunden_ID(KundenID int) (err error) {
	_, err = config.Db.Exec("delete from Verleih where KundenID = $1", KundenID)
	return
}

func Delete_Verleih_By_Artikel_ID(ArtikelID int) (err error) {
	_, err = config.Db.Exec("delete from Verleih where ArtikelID = $1", ArtikelID)
	return
}
