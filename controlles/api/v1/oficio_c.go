package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

func CreateOficio(c *gin.Context) {
	var cp models.Oficio
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosOficio(), c)
	r := crud.Create()
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func UpdateOficio(c *gin.Context) {
	var cp models.Oficio
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosOficio(), c)
	r := crud.Update("numero = ?", "numero")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func SearchOficio(c *gin.Context) {
	var cp models.Oficio
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosOficio(), c)
	r := crud.Search("numero = ?", "numero")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func AllOficio(c *gin.Context) {
	var cp models.Oficio

	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosOficio(), c)
	r := crud.Get()
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func DeleteOficio(c *gin.Context) {
	var cp models.Oficio

	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosOficio(), c)
	r := crud.Delete("numero = ?", "numero")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

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

	// Obtiene el nombre del archivo
	archivo := c.Param("archivo")
	archivo = "/home/orlando/app/files/" + archivo
	// archivo = "/Users/orlando/Downloads/oficios/" + archivo
	f, err := os.Open(archivo)
	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el archivo:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadGateway, msg)
		return
	}
	defer f.Close()

	RespondPdf(c.Writer, http.StatusOK, f)

}
