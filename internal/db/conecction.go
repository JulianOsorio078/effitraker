package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB inicializa la conexión a la base de datos
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}

// CloseDB cierra la conexión a la base de datos
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Println("Error cerrando la base de datos:", err)
	}
}

func ConfigureConection() {
	err := InitDB("root:root@tcp(localhost:3309)/EffiManager")
	if err != nil {
		log.Fatal(err)
	}
}
