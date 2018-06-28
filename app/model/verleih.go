package model

import (
	_ "github.com/mattn/go-sqlite3"
	"../../config"
)

// Verleiher data structure
type Verleih struct {
	VerleihID int
	KundenID int
	ArtikelID int
	Beginn string
	Rueckgabe string
}

type Verleihe map[int]* Verleih

// Create Verleiher
func  CreateVerleiher(bname string, mail string)  (err error) {
	//defer stmt.Close()
	_, err = config.Db.Exec("insert into Verleiher (Benutzername, Email) values (bname, mail)")
	return
}

// Read Verleiher by VerleiherID
func ReadVerleiher(id int) (verleiher Verleih, err error) {
	verleiher = Verleih{}

	// err = config.Db.QueryRow("select *  from Verleiher where VerleiherID = $1", id).Scan(&verleiher.VerleiherID, &verleiher.Benutzername, &verleiher.Email)

	return
}

// Update the Verleiher by id
func  UpdateVerleiher(id int, bname string, mail string) (err error) {
	_, err = config.Db.Exec("update Verleiher set Benutzername = $1 where VerleiherID = $2",bname, id)
	_, err = config.Db.Exec("update Verleiher set Email = $1 where VerleiherID = $2",mail, id)
	return
}

// Delete Verleiher by id
func DeleteVerleiher(id int) (err error) {
	_, err = config.Db.Exec("delete from Verleiher where VerleiherID = $1", id)
	return
}