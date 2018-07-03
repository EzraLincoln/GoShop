package model

import (
	"../../config"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type Kunden map[int]*Kunde

type Kunde struct {
	KundenID     int
	Benutzername string
	BildUrl      string
	Typ          string
	Status       string
	Passwort     string
	Email        string
}

func ReadKunde(id int) (kunde Kunde, err error) {
	kunde = Kunde{}
	err = config.Db.QueryRow("select *  from Kunde where KundenID = $1", id).Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.Passwort, &kunde.Email)
	return
}

func UpdateKunde(id int, bname string, psw string, mail string) (err error) {
	_, err = config.Db.Exec("update Kunde set Benutzername = $1 where KundenID = $2", bname, id)
	_, err = config.Db.Exec("update Kunde set Passwort = $1 where KundenID = $2", psw, id)
	_, err = config.Db.Exec("update Kunde set Email = $1 where KundenID = $2", mail, id)
	return
}

func DeleteKunde(id int) (err error) {
	_, err = config.Db.Exec("delete from Kunde where KundenID = $1", id)
	return
}

func Register_Kunden(user string, mail string, password string) (bool) {
	
	v1, err := config.Db.Prepare("Insert into Kunde(Benutzername,BildUrl,Typ,Status,Passwort,Email) values (?,?,?,?,?,?)");
	defer v1.Close()
	
	_, err = v1.Exec(user, "empty.jpg", "Benutzer", "aktiv", password, mail);
	if err != nil {
		return true
	} else {
		return false
	}
}

func Get_Kunden_By_ID(kunde_id int) (*Kunde, bool) {
	
	result := config.Db.QueryRow("Select * From Kunde Where KundenID = $1", kunde_id)
	
	kunde := Kunde{}
	
	if result.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email) != nil {
		return &kunde, true
	} else {
		return &kunde, false
	}
	
}

func Get_Kunden_By_Name(name string) (*Kunde, bool) {
	
	kunde := Kunde{}
	
	result := config.Db.QueryRow("Select * from Kunde WHERE Kunde.Benutzername= $1", name)
	
	if result.Scan(&kunde.KundenID, &kunde.Benutzername, &kunde.BildUrl, &kunde.Typ, &kunde.Status, &kunde.Passwort, &kunde.Email) != nil {
		return &kunde, true
	} else {
		return &kunde, false
	}
	
}

func Test_For_Kunden_By_Name_Mail(user string, mail string) (bool) {
	
	var id int
	
	if config.Db.QueryRow("Select KundenID from Kunde WHERE Kunde.Benutzername= $1 AND Kunde.Email=$2", user, mail).Scan(&id) != nil {
		return false // KORREKT ... Kunde noch nicht vorhanden !
	} else {
		return true // FEHER ... Kunde in DB !
	}
}

func Delete_Kunden_By_Name(name string) (bool) {
	
	fmt.Println("Delete From Kunde Where Benutzername = ", name)
	
	_, err := config.Db.Query("Delete From Kunde Where Benutzername = $1", name)
	if err != nil {
		return true
	} else {
		return false
	}
}

func Get_Alle_Kunden() ([]Kunde, bool) {
	
	rows, err := config.Db.Query("Select * From Kunde Where Typ = 'Benutzer'")
	
	kunde := Kunde{}    // Benennung des Rückgabe Werts innerhalb der Funktion // Nicht über ParameterListe
	kunden := []Kunde{} // Benennung des Rückgabe Werts innerhalb der Funktion // Nicht über ParameterListe
	
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
