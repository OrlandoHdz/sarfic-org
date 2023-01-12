package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	// Db variable global de la base de datos
	Db *gorm.DB
)

// ConectarDb Se conecta a la base de datos obteniendo los parametros de .env, retorna un puntero a la base de datos
func ConectarDb() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Monterrey",
		dbHost,
		username,
		password,
		dbName,
		dbPort)

	// fmt.Println("dsn:", dsn)

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "org",
				SingularTable: false,
			},
		})

	if err != nil {
		panic(err.Error())
	}

	Db = db
}
