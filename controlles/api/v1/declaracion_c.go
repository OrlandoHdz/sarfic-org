package v1

import (
	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

func CreateDeclaracion(c *gin.Context) {
	var cd models.Declaracion
	crud := models.NewCrud(&cd, &cd, cd.CamposObligatorios(), c)
	r := crud.Create()
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func UpdateDeclaracion(c *gin.Context) {
	var cd models.Declaracion
	crud := models.NewCrud(&cd, &cd, cd.CamposObligatorios(), c)
	r := crud.Update("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func SearchDeclaracion(c *gin.Context) {
	var cd models.Declaracion
	crud := models.NewCrud(&cd, &cd, cd.CamposObligatorios(), c)
	r := crud.Search("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func AllDeclaracion(c *gin.Context) {
	var cd models.Declaracion
	crud := models.NewCrud(&cd, &cd, cd.CamposObligatorios(), c)
	r := crud.Get()
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func DeleteDeclaracion(c *gin.Context) {
	var cd models.Declaracion
	crud := models.NewCrud(&cd, &cd, cd.CamposObligatorios(), c)
	r := crud.Delete("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}
