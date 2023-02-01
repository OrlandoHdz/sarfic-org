package v1

import (
	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

func CreatePermisionaria(c *gin.Context) {
	var cp models.Permisionaria
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Create()
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func UpdatePermisionaria(c *gin.Context) {
	var cp models.Permisionaria
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Update("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}

func SearchPermisionaria(c *gin.Context) {
	var cp models.Permisionaria
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Search("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func AllPermisionaria(c *gin.Context) {
	var cp models.Permisionaria
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Get()
	msg := Message(r.Success, r.Message)
	msg["payload"] = r.Payload
	Respond(c.Writer, r.HttpStatus, msg)
}

func DeletePermisionaria(c *gin.Context) {
	var cp models.Permisionaria
	crud := models.NewCrud(&cp, &cp, cp.CamposObligatorios(), c)
	r := crud.Delete("rfc = ?", "rfc")
	msg := Message(r.Success, r.Message)
	Respond(c.Writer, r.HttpStatus, msg)
}
