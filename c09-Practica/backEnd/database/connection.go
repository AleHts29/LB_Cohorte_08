package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Connect() {
	dsn := "root:root@tcp(localhost:3307)/go_auth"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("No se pudo conectar a la DB: ", err)
	}

	//	verificamos conxion contra la DB
	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la DB: ", err)
	}

	DB = db

	// Crear la tabla de usuarios si no existe
	createTable := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password BLOB NOT NULL,
		PRIMARY KEY (id)
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
