package controller

import (
	"../model"
	"../../config"
)

func GetAllVerleihe() (verleihe [] model.Verleih) {
	rows, err := config.Db.Query("select * from Verleih")

	if err != nil {
		return
	}

	for rows.Next() {
		verleih := model.Verleih{}

		err = rows.Scan(&verleih.VerleihID,&verleih.KundenID,&verleih.ArtikelID,&verleih.Beginn,&verleih.Rueckgabe)

		if err != nil {
			return
		}

		verleihe= append(verleihe, verleih)
	}
	rows.Close()
	return
}
