package v1

import (
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
	r := crud.Update("numero = ?", "numero")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func SearchCenso(c *gin.Context) {
	var cp models.Censo
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Search("numero = ?", "numero")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func AllCenso(c *gin.Context) {
	var cp models.Censo

	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Get()
	//r := crud.GetQry("tipo = ?", "tipo")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func DeleteCenso(c *gin.Context) {
	var cp models.Censo

	crud := models.NewCrud(&cp, &cp, cp.CamposObligatoriosCenso(), c)
	r := crud.Delete("numero = ?", "numero")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}
