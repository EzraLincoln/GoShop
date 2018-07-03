package controller

import (
	"../../config"
	"../model"
	"fmt"
)

type Kunden struct{}

func (v Kunden) Register_Kunden(user string, mail string, password string) (bool) {

	v1, err := config.Db.Prepare("Insert into Kunde(Benutzername,BildUrl,Typ,Status,Passwort,Email) values (?,?,?,?,?,?)");
	defer v1.Close()

	_, err = v1.Exec(user, "empty.jpg", "Benutzer", "aktiv", password, mail);
	if err != nil {
		return true
	} else {
		return false
	}
}
func (v Kunden) Get_Kunden_By_ID(kunde_id int) (*model.Kunde, bool) {

	result := config.Db.QueryRow("Select * From Kunde Where KundenID = $1", kunde_id)

	kunde := model.Kunde{}

	if result.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email) != nil {
		return nil, true
	} else {
		return &kunde, false
	}

}
func (v Kunden) Get_Kunden_By_Name(name string) (*model.Kunde, bool) {

	kunde := model.Kunde{}

	result := config.Db.QueryRow("Select * from Kunde WHERE Kunde.Benutzername= $1", name)

	if result.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email) != nil {
		return nil, true
	} else {
		return &kunde, false
	}

}
func (v Kunden) Test_For_Kunden_By_Name_Mail(user string, mail string) (bool) {

	var id int

	if config.Db.QueryRow("Select KundenID from Kunde WHERE Kunde.Benutzername= $1 AND Kunde.Email=$2", user, mail).Scan(&id) != nil {
		return true // KORREKT ... Kunde noch nicht vorhanden !
	} else {
		return false // FEHER ... Kunde in DB !
	}
}

/*func (v Kunden) Delete_Kunden_By_ID(kunde_id int) (bool) {
	_, err := config.Db.Query("Delete From Kunde Where Kunde.KundenID = $1", kunde_id)
	if err != nil {
		fmt.Println("Fehler bei Delete Kunden By ID")
		return true
	} else {
		return false
	}
}*/

func (v Kunden) Delete_Kunden_By_Name(name string) (bool) {
	
	fmt.Println("Delete From Kunde Where Benutzername = ",name)
	
	_, err := config.Db.Query("Delete From Kunde Where Benutzername = $1", name)
	if err != nil {
		return true
	} else {
		return false
	}
}

func (v Kunden) Get_Alle_Kunden() ([]model.Kunde, bool) {

	rows, err := config.Db.Query("Select * From Kunde Where Typ = 'Benutzer'")

	kunde := model.Kunde{}    // Benennung des Rückgabe Werts innerhalb der Funktion // Nicht über ParameterListe
	kunden := []model.Kunde{} // Benennung des Rückgabe Werts innerhalb der Funktion // Nicht über ParameterListe

	if err != nil {
		return nil, true // Db Query hat einen Fehler returned !
	}
	for rows.Next() {

		err = rows.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email)

		if err != nil {
			return nil, true // Db Row Scan hat einen Fehler returned !
		}

		kunden = append(kunden, kunde) // Hänge dem leeren Kunden Array die neu gefundenen Kunden Einträge an.
	}

	rows.Close() // Schließen der Row Instanz um weiterer Enumeration zu vermeiden

	return kunden, false // Gebe das Kunden Array zurück und keine Fehler !

}
