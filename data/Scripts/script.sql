CREATE TABLE 'Artikel' (
	'ArtikelID' INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	'Bezeichnung' varchar(50) NOT NULL,
	'Kategorie' varchar(20) NOT NULL,
	'Lagerort' varchar(30)NOT NULL,
	'Anzahl' INTEGER,
	'Hinweis' varchar(100),
	'BildURL' varchar(100)
)

CREATE TABLE 'Kunde' (
	'KundeID' INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	'Benutzername' varchar(20) NOT NULL,
	'Passwort' varchar(12) NOT NULL,
	'Email' varchar(30) NOT NULL
)

CREATE TABLE 'Verleiher' (
	'VerleiherID' INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	'Benutzername' varchar(20) NOT NULL,
	'Email' varchar(30) NOT NULL
);