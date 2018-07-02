package controller

import (
	"../../config"
	"../model"
	"log"
)

// type Verleih model.Kunde

type Kunden struct{}

func (v Kunden) Register_Kunden(user string, mail string, password string) (bool) {

	if stmt, err := config.Db.Prepare("Insert into Kunde values (?,?,?,?,?,?,?)"); err != nil {
		return true
	} else {
		defer stmt.Close()
		if _, err = stmt.Exec(nil, user, "Benutzer", "aktiv", password, mail); err != nil {
			return true
		} else {
			return false
		}
	}
}

func (v Kunden) Get_Kunden_By_ID(kunde_id int) (profiles []model.Profile) {
	rows, err := config.Db.Query("select Kunde.KundeID,Kunde.Benutzername,Kunde.BildUrl,Kunde.Email,Kunde.Status from Kunde WHERE Kunde.KundeID = $1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		profile := model.Profile{}

		err = rows.Scan(&profile.KundenID, &profile.Benutzername, &profile.BildURL, &profile.Mail, &profile.Status)

		if err != nil {
			return
		}

		profiles = append(profiles, profile)
	}
	rows.Close()
	return
}
func (v Kunden) Get_Kunden_By_Name(name string) (profiles []model.Profile) {
	rows, err := config.Db.Query("select Kunde.KundeID,Kunde.Benutzername,Kunde.BildUrl,Kunde.Email,Kunde.Status from Kunde WHERE Kunde.Benutzername= $1", name)

	if err != nil {
		return
	}
	for rows.Next() {
		profile := model.Profile{}

		err = rows.Scan(&profile.KundenID, &profile.Benutzername, &profile.BildURL, &profile.Mail, &profile.Status)

		if err != nil {
			return
		}

		profiles = append(profiles, profile)
	}
	rows.Close()
	return
}

func (v Kunden) Get_Kunden_By_Name_Mail(user string, mail string) (bool) {

	var id int

	if config.Db.QueryRow("Select KundeID from Kunde WHERE Kunde.Benutzername= $1 AND Kunde.Email=$2", user, mail).Scan(&id) != nil {
		return true
	} else {
		log.Fatalln("FEHLER", id)
		return false
	}
}

func (v Kunden) Delete_Kunden_By_ID(kunde_id int) (bool) {
	_, err := config.Db.Query("Delete From Kunde Where Kunde.KundeID = $1", kunde_id)
	if err != nil {
		return false
	} else {
		return true
	}
}
func (v Kunden) Delete_Kunden_By_Name(name string) (bool) {
	_, err := config.Db.Query("Delete From Kunde Where Kunde.Benutzername= $1", name)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (v Kunden) Get_Alle_Kunden() (kunden []model.Kunde) {
	rows, err := config.Db.Query("select * from Kunde where Typ = 'Benutzer'")

	if err != nil {
		return
	}

	for rows.Next() {
		kunde := model.Kunde{}
		err = rows.Scan(&kunde.KundeID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Email, &kunde.Passwort)

		if err != nil {
			return
		}

		kunden = append(kunden, kunde)
	}
	rows.Close()
	return
}
