package model

import (
	"../../config"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)


type Verleih struct {
	VerleihID int
	KundenID  int
	ArtikelID int
	Beginn    string
	Rueckgabe string
}
type Verleihe map[int]*Verleih

func Add_Verleih(EquipmentID int, KundenID int, beginn string, rückgabe string, anzahl int) {
	//defer stmt.Close()
	_, err := config.Db.Exec("insert into Verleih (KundenID,ArtikelID,Beginn,Rueckgabe) values ($1,$2,$3,$4)", EquipmentID, KundenID, beginn, rückgabe)

if err != nil {
	fmt.Println("Fehler beim Verleih Insert")
}
return
}

func ReadVerleiher(id int) (verleiher Verleih, err error) {
	verleiher = Verleih{}

	// err = config.Db.QueryRow("select *  from Verleiher where VerleiherID = $1", id).Scan(&verleiher.VerleiherID, &verleiher.Benutzername, &verleiher.Email)

	return
}

func UpdateVerleiher(id int, bname string, mail string) (err error) {
	_, err = config.Db.Exec("update Verleiher set Benutzername = $1 where VerleiherID = $2", bname, id)
	_, err = config.Db.Exec("update Verleiher set Email = $1 where VerleiherID = $2", mail, id)
	return
}

func Delete_Verleih_By_Kunden_ID(KundenID int) (err error) {
	_, err = config.Db.Exec("delete from Verleih where KundenID = $1", KundenID)
	return
}

func Delete_Verleih_By_Artikel_ID(ArtikelID int) (err error) {
	_, err = config.Db.Exec("delete from Verleih where ArtikelID = $1", ArtikelID)
	return
}

func GetAllVerleihe() (verleihe []Verleih) {
	
	rows, err := config.Db.Query("select * from Verleih")
	
	if err != nil {
		return
	}
	for rows.Next() {
		
		verleih := Verleih{}
		
		err = rows.Scan(&verleih.VerleihID, &verleih.KundenID, &verleih.ArtikelID, &verleih.Beginn, &verleih.Rueckgabe)
		
		if err != nil {
			return
		}
		verleihe = append(verleihe, verleih)
	}
	rows.Close()
	
	return
}
