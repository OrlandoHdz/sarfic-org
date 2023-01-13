package models

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success    bool
	Error      error
	Message    string
	Payload    interface{}
	HttpStatus int
}

type Crud struct {
	Model          interface{}
	Values         interface{}
	ValidateFields []string
	Context        *gin.Context
}

func NewCrud(model interface{}, values interface{}, validateFields []string, contexto *gin.Context) Crud {
	return Crud{
		Model:          model,
		Values:         values,
		ValidateFields: validateFields,
		Context:        contexto,
	}
}

func (c *Crud) Create() *Response {
	rs, usuario := validaToken(c.Context)
	if rs.Error != nil {
		return rs
	}
	// validar que tenga privilegios
	if !TienePrivilejios(usuario) {
		return &Response{
			Success:    false,
			Error:      errors.New("no tiene privilejios"),
			Message:    "Ocurrio un error usted no tiene privilegios",
			HttpStatus: http.StatusUnauthorized,
		}
	}
	// obtiene los datos del Body
	err := c.Context.ShouldBind(&c.Values)
	if err != nil {
		return &Response{
			Success:    false,
			Error:      err,
			Message:    fmt.Sprintf("Ocurrio un error al obtener el Body:%v", err),
			HttpStatus: http.StatusBadRequest,
		}
	}
	// valida que contengan datos
	val := reflect.ValueOf(c.Values).Elem()
	for _, field := range c.ValidateFields {
		n := fmt.Sprintf("%v", val.FieldByName(field).Interface())
		if n == "" || n == "0" {
			return &Response{
				Success:    false,
				Error:      errors.New("faltan datos en el body"),
				Message:    fmt.Sprintf("Ocurrio un error en el Body, faltan datos[%v]", field),
				HttpStatus: http.StatusBadRequest,
			}
		}
	}

	// Crea el registro en la base de datos
	r := Db.Create(c.Values)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al crear el registro:%v", r.Error),
			HttpStatus: http.StatusBadRequest,
		}
	}
	// Si fue exitos
	return &Response{
		Success:    true,
		Error:      nil,
		Message:    "Registro creado",
		HttpStatus: http.StatusOK,
	}
}

func (c *Crud) Update(query string, param string) *Response {
	rs, usuario := validaToken(c.Context)
	if rs.Error != nil {
		return rs
	}
	// validar que tenga privilegios
	if !TienePrivilejios(usuario) {
		return &Response{
			Success:    false,
			Error:      errors.New("no tiene privilejios"),
			Message:    "Ocurrio un error usted no tiene privilegios",
			HttpStatus: http.StatusUnauthorized,
			Payload:    nil,
		}
	}
	// obtiene el parametro
	cparam := c.Context.Param(param)
	// obtiene los datos del Body
	err := c.Context.ShouldBind(&c.Values)
	if err != nil {
		return &Response{
			Success:    false,
			Error:      err,
			Message:    fmt.Sprintf("Ocurrio un error al obtener el Body[2]:%v", err),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}
	}
	// valida que contengan datos
	val := reflect.ValueOf(c.Values).Elem()
	for _, field := range c.ValidateFields {
		n := fmt.Sprintf("%v", val.FieldByName(field).Interface())
		if n == "" || n == "0" {
			return &Response{
				Success:    false,
				Error:      errors.New("faltan datos en el body"),
				Message:    "Ocurrio un error en el Body, faltan datos",
				HttpStatus: http.StatusBadRequest,
				Payload:    nil,
			}
		}
	}
	// Busca el registro en la base de datos
	reg := map[string]interface{}{}
	r := Db.Model(c.Model).Where(query, cparam).First(&reg)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al buscar en base de datos: %v", r.Error),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}

	}
	// Actualiza el registro
	r = Db.Model(c.Model).Where(query, cparam).Updates(c.Values)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al actualizar en base de datos: %v", r.Error),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}

	}
	// Si fue exitos
	return &Response{
		Success:    true,
		Error:      nil,
		Message:    "Registro actualizado",
		HttpStatus: http.StatusOK,
		Payload:    nil,
	}
}

func (c *Crud) Search(query string, param string) *Response {
	rs, usuario := validaToken(c.Context)
	if rs.Error != nil {
		return rs
	}
	// validar que tenga privilegios
	if !TienePrivilejios(usuario) {
		return &Response{
			Success:    false,
			Error:      errors.New("no tiene privilejios"),
			Message:    "Ocurrio un error usted no tiene privilegios",
			HttpStatus: http.StatusUnauthorized,
			Payload:    nil,
		}
	}
	// obtiene el parametro
	cparam := c.Context.Param(param)
	// Busca el registro en la base de datos
	reg := map[string]interface{}{}
	r := Db.Model(c.Model).Where(query, cparam).First(&reg)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al buscar en base de datos: %v", r.Error),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}

	}
	// Si fue exitos
	return &Response{
		Success:    true,
		Error:      nil,
		Message:    "Registro actualizado",
		HttpStatus: http.StatusOK,
		Payload:    reg,
	}
}

func (c *Crud) Delete(query string, param string) *Response {
	rs, usuario := validaToken(c.Context)
	if rs.Error != nil {
		return rs
	}
	// validar que tenga privilegios
	if !TienePrivilejios(usuario) {
		return &Response{
			Success:    false,
			Error:      errors.New("no tiene privilejios"),
			Message:    "Ocurrio un error usted no tiene privilegios",
			HttpStatus: http.StatusUnauthorized,
			Payload:    nil,
		}
	}
	// obtiene el parametro
	cparam := c.Context.Param(param)
	// Busca el registro en la base de datos
	reg := map[string]interface{}{}
	r := Db.Model(c.Model).Where(query, cparam).Delete(&reg)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al buscar en base de datos: %v", r.Error),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}

	}
	// Si fue exitos
	return &Response{
		Success:    true,
		Error:      nil,
		Message:    "Registro eliminado",
		HttpStatus: http.StatusOK,
		Payload:    nil,
	}
}

func (c *Crud) Get() *Response {
	rs, usuario := validaToken(c.Context)
	if rs.Error != nil {
		return rs
	}
	// validar que tenga privilegios
	if !TienePrivilejios(usuario) {
		return &Response{
			Success:    false,
			Error:      errors.New("no tiene privilejios"),
			Message:    "Ocurrio un error usted no tiene privilegios",
			HttpStatus: http.StatusUnauthorized,
			Payload:    nil,
		}
	}
	// Obtener los registro en la base de datos
	fmt.Println("Obteniendo registros")
	reg := []map[string]interface{}{}
	r := Db.Model(c.Model).Order("id desc").Find(&reg)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al obtener registros en base de datos: %v", r.Error),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}

	}
	// Si fue exitos
	return &Response{
		Success:    true,
		Error:      nil,
		Message:    "Registro actualizado",
		HttpStatus: http.StatusOK,
		Payload:    reg,
	}
}

func (c *Crud) GetQry(query string, param string) *Response {
	rs, usuario := validaToken(c.Context)
	if rs.Error != nil {
		return rs
	}
	// validar que tenga privilegios
	if !TienePrivilejios(usuario) {
		return &Response{
			Success:    false,
			Error:      errors.New("no tiene privilejios"),
			Message:    "Ocurrio un error usted no tiene privilegios",
			HttpStatus: http.StatusUnauthorized,
			Payload:    nil,
		}
	}
	// obtiene el parametro
	cparam := c.Context.Param(param)

	// Obtener los registro en la base de datos
	fmt.Println("Obteniendo registros")
	reg := []map[string]interface{}{}
	r := Db.Model(c.Model).Where(query, cparam).Order("id desc").Find(&reg)
	if r.Error != nil {
		return &Response{
			Success:    false,
			Error:      r.Error,
			Message:    fmt.Sprintf("Ocurrio un error al obtener registros en base de datos: %v", r.Error),
			HttpStatus: http.StatusBadRequest,
			Payload:    nil,
		}

	}
	// Si fue exitos
	return &Response{
		Success:    true,
		Error:      nil,
		Message:    "Registro actualizado",
		HttpStatus: http.StatusOK,
		Payload:    reg,
	}
}

func TienePrivilejios(usr Usuario) bool {
	if usr.TipoUsuario == TipoAdministrador ||
		usr.TipoUsuario == TipoAuditorCrayon ||
		usr.TipoUsuario == TipoAuditorGob {
		return true
	}
	return false
}

func validaToken(c *gin.Context) (*Response, Usuario) {

	aut := c.Request.Header["Authorization"]

	var token string
	var usuario Usuario
	var err error

	if len(aut) > 0 {
		token = aut[0]
	} else {
		return &Response{
			Success:    false,
			Error:      err,
			Message:    "Ocurrio un error usted no esta incluyendo el token en su peticiónn",
			HttpStatus: http.StatusUnauthorized,
		}, usuario
	}

	usuario, err = ValidaToken(token)

	if err != nil {
		return &Response{
			Success:    false,
			Error:      err,
			Message:    "Ocurrio un error al obtener el token",
			HttpStatus: http.StatusUnauthorized,
		}, usuario
	}

	return &Response{
		Success:    true,
		Error:      err,
		Message:    "Obteniendo token correctamente",
		HttpStatus: http.StatusOK,
	}, usuario

}

func MigrarModelos() {
	MigrarUsuarios()
	MigararMaquina()
	MigrarCasino()
	MigrarEntidad()
	MigrarOperadora()
	MigrarPermisionaria()
}
