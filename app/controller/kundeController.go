package controller

import (
	"../../config"
	"net/http"
	"../model"
)


func RegisterKunden(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("user")
	email := r.FormValue("mail")
	password := r.FormValue("psw")

	statement := "insert into Kunde (Benutzername, Passwort, Email) values (?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(userName,email,password)

	return
}

// GetAll Kunden
func GetAllUser() (kunden [] model.Kunde){
	rows, err := config.Db.Query("select * from Kunde where Typ = 'Benutzer'")

	if err != nil {
		return
	}

	for rows.Next() {
		kunde := model.Kunde{}
		err = rows.Scan(&kunde.KundeID,&kunde.Benutzername,&kunde.BildUrl,&kunde.Typ,&kunde.Status,&kunde.Email, &kunde.Passwort)

		if err != nil {
			return
		}

		kunden = append(kunden, kunde)
	}
	rows.Close()
	return
}

func GetProfile(kunde_id int) (profiles []model.Profile) {
	rows, err := config.Db.Query("select Kunde.KundeID,Kunde.Benutzername,Kunde.BildUrl,Kunde.Email,Kunde.Status from Kunde WHERE Kunde.KundeID = $1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		profile := model.Profile{}

		err = rows.Scan(&profile.KundenID,&profile.Benutzername, &profile.BildURL, &profile.Mail, &profile.Status)

		if err != nil {
			return
		}

		profiles = append(profiles, profile)
	}
	rows.Close()
	return
}