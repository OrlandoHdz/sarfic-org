package main

import (
	"fmt"

	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/joho/godotenv"
)

func main() {
	// carga las variables de ambient del archivo .env
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// Conecta a la base de datos
	models.ConectarDb()

	// Migrar tablas
	models.MigrarUsuarios()

}
