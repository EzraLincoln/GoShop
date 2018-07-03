package controller

import (
	"../../config"
	"../model"
)

type Kunden struct{}

func (v Kunden) Register_Kunden(user string, mail string, password string) (bool) {

	v1, err := config.Db.Prepare("Insert into Kunde(Benutzername,BildUrl,Typ,Status,Passwort,Email) values (?,?,?,?,?,?)");
	defer v1.Close()

	_,err = v1.Exec(user, "empty.jpg", "Benutzer", "aktiv", password, mail);

	if err != nil {
		return true
	} else {
		return false
	}
}

func (v Kunden) Get_Kunden_By_ID(kunde_id int) (*model.Kunde) {

	result := config.Db.QueryRow("Select * From Kunde Where KundenID = $1", kunde_id)

	kunde := model.Kunde{}

	if result.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email) != nil {
		return nil
	} else {
		return &kunde
	}

}

func (v Kunden) Get_Kunden_By_Name(name string) (*model.Kunde) {

	kunde := model.Kunde{}

	result := config.Db.QueryRow("Select * from Kunde WHERE Kunde.Benutzername= $1", name)

	if result.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email) != nil {
		return nil
	} else {
		return &kunde
	}

}

func (v Kunden) Test_For_Kunden_By_Name_Mail(user string, mail string) (bool) {

	var id int

	if config.Db.QueryRow("Select KundenID from Kunde WHERE Kunde.Benutzername= $1 AND Kunde.Email=$2", user, mail).Scan(&id) != nil {
		return true
	} else {
		return false
	}
}

func (v Kunden) Delete_Kunden_By_ID(kunde_id int) (bool) {
	_, err := config.Db.Query("Delete From Kunde Where Kunde.KundenID = $1", kunde_id)
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

	rows, err := config.Db.Query("Select * From Kunde Where Typ = 'Benutzer'")

	if err != nil {
		return
	}
	for rows.Next() {

		kunde := model.Kunde{}

		err = rows.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email)

		if err != nil {
			return
		}
		kunden = append(kunden, kunde)
	}
	rows.Close()

	return

}
