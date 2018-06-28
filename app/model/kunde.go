package model

import (
	_ "github.com/mattn/go-sqlite3"
	"../../config"
)

type Kunden map[int]* Kunde

// Kunde data structure
type Kunde struct {
	KundeID int
	Benutzername string
	BildUrl string
	Typ string
	Status string
	Passwort string
	Email string
}

// Read Kunde by KundeID
func ReadKunde(id int) (kunde Kunde, err error) {
	kunde = Kunde{}
	err = config.Db.QueryRow("select *  from Kunde where KundeID = $1", id).Scan(&kunde.KundeID, &kunde.Benutzername, &kunde.Passwort, &kunde.Email)
	return
}

// Update the Kunde by id
func  UpdateKunde(id int, bname string, psw string, mail string) (err error) {
	_, err = config.Db.Exec("update Kunde set Benutzername = $1 where KundeID = $2",bname, id)
	_, err = config.Db.Exec("update Kunde set Passwort = $1 where KundeID = $2",psw, id)
	_, err = config.Db.Exec("update Kunde set Email = $1 where KundeID = $2",mail, id)
	return
}

// Delete Kunde by id
func DeleteKunde(id int) (err error) {
	_, err = config.Db.Exec("delete from Kunde where KundeID = $1", id)
	return
}