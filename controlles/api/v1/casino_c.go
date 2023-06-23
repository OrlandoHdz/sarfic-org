package v1

import (
	"fmt"
	"net/http"

	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

func CreateCasino(c *gin.Context) {
	var cp models.Casino
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Create()
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func UpdateCasino(c *gin.Context) {
	var cp models.Casino
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Update("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func SearchCasino(c *gin.Context) {
	var cp models.Casino
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Search("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func AllCasino(c *gin.Context) {
	var cp models.Casino
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Get()
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func DeleteCasino(c *gin.Context) {
	var cp models.Casino
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Delete("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func TodosLosCasinos(c *gin.Context) {
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

	casinos, err := models.ObtenerCasinos(int(usuario.EntidadID))

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr los casinos:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
	} else {
		msg := Message(true, "Casinos obtenidos con exito")
		msg["payload"] = casinos
		Respond(c.Writer, http.StatusOK, msg)
	}

}
