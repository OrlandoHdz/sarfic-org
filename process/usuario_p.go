package process

import (
	"fmt"
	"time"

	"github.com/OrlandoHdz/sarfic-org/encrypt"
	"github.com/OrlandoHdz/sarfic-org/models"
)

// CrearUsuarioAdmin crea el usuario administrador en la base de datos
func CrearUsuarioAdmin() {

	key := "rfq74564bdb6fc123706eea85ec98431"
	pas := "S4rf1cP"
	password, err := encrypt.Sencrypt(pas, key)
	if err != nil {
		fmt.Println("Ocurrio un error Sencrypt%v", err)
		return
	}

	u := models.Usuario{
		Usuario:         "admin",
		Email:           "admin@sarfic.com.mx",
		Nombre:          "Usuario Administrador",
		Password:        password,
		TipoAutt:        1,
		TipoAuttDesc:    "Sistema",
		TipoUsuario:     1,
		TipoUsuarioDesc: "Administrador",
		EntidadID:       1,
		Token:           "N/A",
		FechaToken:      time.Now(),
	}

	err = models.CrearUsuario(u)

	if err != nil {
		fmt.Println("Ocurrio un error al crear el usuario administrador:")
		fmt.Println(err.Error())
	}

}
