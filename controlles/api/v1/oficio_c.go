package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

func RetornaOficio(c *gin.Context) {
	// Valida el Token
	aut := c.Request.Header["Authorization"]
	token, err := validaAuthorization(aut)
	if err != nil {
		msg := Message(false, err.Error())
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	usuario, err := models.ValidaToken(token)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el token:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	// validar que tenga privilegios para poder actualizaar el usuario
	if !models.TienePrivilejios(usuario) {
		msg := Message(false, "Ocurrio un error usted no tiene privilegios")
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	//archivo := "/home/orlando/app/files/oficio90.pdf"
	archivo := "/Users/orlando/Downloads/oficio90.pdf"
	f, err := os.Open(archivo)
	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el archivo:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadGateway, msg)
		return
	}
	defer f.Close()

	RespondPdf(c.Writer, http.StatusOK, f)

}
