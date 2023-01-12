// Este paquete es para la creación y actualizacion de usuarios

package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/OrlandoHdz/sarfic-org/encrypt"
	"github.com/OrlandoHdz/sarfic-org/models"
	"github.com/gin-gonic/gin"
)

type JsonRequestUsuario struct {
	Usuario     string `json:"usuario"`
	Email       string `json:"email"`
	Nombre      string `json:"nombre"`
	Password    string `json:"password"`
	TipoAutt    string `json:"tipo_autt"`
	TipoUsuario string `json:"tipo_usuario"`
	EntidadID   string `json:"empresa_id"`
}

// CrearUsuario ...
func CrearUsuario(c *gin.Context) {
	// Valida el Token
	token := c.Request.Header["Authorization"][0]
	usuario, err := models.ValidaToken(token)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el token:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	// Obtiene los datso del Body y los valida
	mBody := JsonRequestUsuario{}
	err = c.ShouldBind(&mBody)
	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el body:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}
	if mBody.Usuario == "" ||
		mBody.Email == "" ||
		mBody.Nombre == "" ||
		mBody.Password == "" ||
		mBody.TipoAutt == "" ||
		mBody.TipoUsuario == "" ||
		mBody.EntidadID == "" {
		msg := Message(false, "Ocurrio un error al obtenr el body, falta información ")
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}

	// validar que tenga privilegios para poder crear el usuario
	if !models.TienePrivilejios(usuario) {
		msg := Message(false, "Ocurrio un error usted no tiene privilegios")
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}
	password, _ := encrypt.Sencrypt(mBody.Password, models.KeyEncrypt)
	mTipoAutt, _ := strconv.ParseUint(mBody.TipoAutt, 10, 32)
	mTipoUsuario, _ := strconv.ParseUint(mBody.TipoUsuario, 10, 32)
	mEntidadID, _ := strconv.ParseUint(mBody.EntidadID, 10, 32)

	usr := models.Usuario{
		Usuario:         mBody.Usuario,
		Email:           strings.ToLower(mBody.Email),
		Nombre:          mBody.Nombre,
		Password:        password,
		TipoAutt:        uint(mTipoAutt),
		TipoAuttDesc:    models.DescTipoAutt[mTipoAutt-1],
		TipoUsuario:     uint(mTipoUsuario),
		TipoUsuarioDesc: models.DescTipoUsuario[mTipoUsuario-1],
		EntidadID:       uint(mEntidadID),
		Token:           "na",
		FechaToken:      time.Now(),
	}

	err = models.CrearUsuario(usr)

	if err != nil {
		msg := Message(false, "Ocurrio un error al crear usuario:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}

	msg := Message(true, "Usuarios creado")
	Respond(c.Writer, http.StatusOK, msg)

}

// ActualizaUsuario ...
func ActualizaUsuario(c *gin.Context) {
	// Valida el Token
	token := c.Request.Header["Authorization"][0]
	usuario, err := models.ValidaToken(token)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el token:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	// Obtiene los datso del Body y los valida
	mBody := JsonRequestUsuario{}
	err = c.ShouldBind(&mBody)
	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el body:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}
	if mBody.Usuario == "" ||
		mBody.Email == "" ||
		mBody.Nombre == "" ||
		mBody.Password == "" ||
		mBody.TipoAutt == "" ||
		mBody.TipoUsuario == "" ||
		mBody.EntidadID == "" {
		msg := Message(false, "Ocurrio un error al obtenr el body, falta información ")
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}

	// validar que tenga privilegios para poder actualizaar el usuario
	if !models.TienePrivilejios(usuario) {
		msg := Message(false, "Ocurrio un error usted no tiene privilegios")
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}
	// busca usuario
	var usr models.Usuario

	r := models.Db.Where(models.Usuario{Email: strings.ToLower(mBody.Email)}).First(&usr)

	if r.Error != nil {
		msg := Message(false, "Ocurrio un error en base de datos:"+fmt.Sprint(r.Error))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}

	password, _ := encrypt.Sencrypt(mBody.Password, models.KeyEncrypt)
	mTipoAutt, _ := strconv.ParseUint(mBody.TipoAutt, 10, 32)
	mTipoUsuario, _ := strconv.ParseUint(mBody.TipoUsuario, 10, 32)
	mEntidadID, _ := strconv.ParseUint(mBody.EntidadID, 10, 32)

	usr.Usuario = mBody.Usuario
	usr.Email = strings.ToLower(mBody.Email)
	usr.Nombre = mBody.Nombre
	usr.Password = password
	usr.TipoAutt = uint(mTipoAutt)
	usr.TipoAuttDesc = models.DescTipoAutt[mTipoAutt-1]
	usr.TipoUsuario = uint(mTipoUsuario)
	usr.TipoUsuarioDesc = models.DescTipoUsuario[mTipoUsuario-1]
	usr.EntidadID = uint(mEntidadID)
	usr.Token = "na"
	usr.FechaToken = time.Now()

	err = models.ActualizaUsuario(usr)

	if err != nil {
		msg := Message(false, "Ocurrio un error al actualizar usuario:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	}

	msg := Message(true, "Usuarios actualizado")
	Respond(c.Writer, http.StatusOK, msg)

}

// BuscarUsuario ...
func BuscarUsuario(c *gin.Context) {
	// Valida el Token
	token := c.Request.Header["Authorization"][0]
	usuario, err := models.ValidaToken(token)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el token:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	// Obtiene email
	email := c.Param("email")

	// validar que tenga privilegios para poder actualizaar el usuario
	if !models.TienePrivilejios(usuario) {
		msg := Message(false, "Ocurrio un error usted no tiene privilegios")
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	usr, err := models.ObtenerUsuario(email)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el usuario:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	} else {
		msg := Message(true, "Usuario obtenido")
		msg["usuario"] = usr
		Respond(c.Writer, http.StatusOK, msg)
	}

}

// TodosUsuarios ...
func TodosUsuarios(c *gin.Context) {
	// Valida el Token
	token := c.Request.Header["Authorization"][0]
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
	usrs, err := models.ObtenerTodosUsuario(usuario.EntidadID)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr los usuarios:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	} else {
		msg := Message(true, "Usuarios obtenido")
		msg["usuarios"] = usrs
		Respond(c.Writer, http.StatusOK, msg)
	}
}

// EliminarUsuario ...
func EliminarUsuario(c *gin.Context) {
	// Valida el Token
	token := c.Request.Header["Authorization"][0]
	usuario, err := models.ValidaToken(token)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el token:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	// Obtiene email
	email := c.Param("email")

	// validar que tenga privilegios para poder actualizaar el usuario
	if !models.TienePrivilejios(usuario) {
		msg := Message(false, "Ocurrio un error usted no tiene privilegios")
		Respond(c.Writer, http.StatusUnauthorized, msg)
		return
	}

	elimino, err := models.EliminarUsuario(email)

	if err != nil {
		msg := Message(false, "Ocurrio un error al obtenr el usuario:"+fmt.Sprint(err))
		Respond(c.Writer, http.StatusBadRequest, msg)
		return
	} else {
		if elimino {
			msg := Message(true, "Usuario fue eliminado")
			Respond(c.Writer, http.StatusOK, msg)
		} else {
			msg := Message(false, "Ocurrio un error")
			Respond(c.Writer, http.StatusBadRequest, msg)
		}
	}

}
