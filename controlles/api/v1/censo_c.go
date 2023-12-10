package v1

import (
	"fmt"
	"net/http"

	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

func CreateCenso(c *gin.Context) {
	var cp models.Censo
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Create()
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func UpdateCenso(c *gin.Context) {
	var cp models.Censo
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Update("id = ?", "id")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func SearchCenso(c *gin.Context) {
	var cp models.Censo
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Search("id = ?", "id")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func AllCenso(c *gin.Context) {
	var cp models.Censo

	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Get()
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func DeleteCenso(c *gin.Context) {
	var cp models.Censo

	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Delete("id = ?", "id")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func TodosLosCensos(c *gin.Context) {
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

	censoT, err := models.ObtenerCenso(int(usuario.EntidadID))

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr los censos:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
	} else {
		msg := Message(true, "Censos obtenidos con exito")
		msg["payload"] = censoT
		Respond(c.Writer, http.StatusOK, msg)
	}

}
