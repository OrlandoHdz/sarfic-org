package v1

import (
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
