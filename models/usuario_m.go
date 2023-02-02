package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/OrlandoHdz/sarfic-org/encrypt"
	"gorm.io/gorm"
)

const (
	TipoAdministrador = 1
	TipoAuditorGob    = 2
	TipoAuditorSat    = 3
	TipoAuditorCrayon = 4
	TipoContribuyente = 5
	TipoConsultas     = 6

	KeyEncrypt = "rfq74564bdb6fc123706eea85ec98431"

	AuttSistema = 1
	AuttGoogle  = 2
)

var (
	DescTipoUsuario = [6]string{"Administrador", "Auditor_Gob", "Auditor_Sat", "Auditor_Crayon",
		"Contribuyente", "Consultas"}

	DescTipoAutt = [2]string{"Sistema", "Google"}
)

// Usuario tabla de usuarios de la plataforma, el campo de Usuario se manejara el RFC para el contribuyente
type Usuario struct {
	gorm.Model
	Usuario         string  `json:"usuario" gorm:"type:character varying(250);"`
	Email           string  `json:"email" gorm:"type:character varying(250);unique;"`
	Nombre          string  `json:"nombre" gorm:"type:character varying(250);"`
	Password        string  `json:"password" gorm:"type:character varying(250);"`
	TipoAutt        uint    `json:"tipo_autt"`
	TipoAuttDesc    string  `gorm:"type:character varying(250);"`
	TipoUsuario     uint    `json:"tipo_usuario"`
	TipoUsuarioDesc string  `gorm:"type:character varying(250);"`
	EntidadID       uint    `json:"entidad_id"`
	Entidad         Entidad `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token           string  `gorm:"type:character varying(2000);"`
	FechaToken      time.Time
}

// MigrarUsuarios ...
func MigrarUsuarios() {
	Db.AutoMigrate(&Usuario{})
}

// CrearUsuario ...
func CrearUsuario(usr Usuario) error {
	var emp Entidad
	r := Db.First(&emp, usr.EntidadID)
	if r.Error == nil {
		r = Db.Create(&usr)
		if r.Error != nil {
			return r.Error
		} else {
			fmt.Println("Usuario creado...")
			return nil
		}
	} else {
		return errors.New("no se encontro la entidad")
	}
}

// ActualizaUsuario ...
func ActualizaUsuario(usr Usuario) error {
	var emp Entidad
	r := Db.First(&emp, usr.EntidadID)
	if r.Error != nil {
		return errors.New("no se encontro la entidad")
	} else {
		r = Db.Save(usr)
		if r.Error != nil {
			return r.Error
		} else {
			fmt.Println("Usuario actualizado...")
			return nil
		}
	}
}

// ObtenerUsuario ...
func ObtenerUsuario(email string) (Usuario, error) {
	var usr Usuario
	//r := Db.Where(Usuario{Email: strings.ToLower(email)}).First(&usr)
	r := Db.Joins("Entidad").First(&usr, Usuario{Email: strings.ToLower(email)})
	if r.Error != nil {
		fmt.Println(r.Error)
		return usr, errors.New("no encontro usuario")
	}

	usr.Password = "**************"
	usr.Token = "**************"

	return usr, nil
}

// ObtenerTodosUsuario ...
func ObtenerTodosUsuario(entidadID uint) ([]Usuario, error) {
	var usrs []Usuario
	var rusrs []Usuario

	r := Db.Joins("Entidad").Find(&usrs, Usuario{EntidadID: entidadID})

	if r.Error != nil {
		fmt.Println(r.Error)
		return usrs, errors.New("no encontro usuario")
	}

	for _, usr := range usrs {
		usr.Password = "**************"
		usr.Token = "**************"
		rusrs = append(rusrs, usr)
	}

	return rusrs, nil
}

// EliminarUsuario ...
func EliminarUsuario(email string) (bool, error) {
	var usr Usuario
	r := Db.First(&usr, Usuario{Email: strings.ToLower(email)})
	if r.Error != nil {
		fmt.Println(r.Error)
		return false, errors.New("no encontro usuario")
	}

	r = Db.Delete(&usr)

	if r.Error != nil {
		fmt.Println(r.Error)
		return false, errors.New("no se pudo eliminar usuario")
	}

	return true, nil
}

// ActualizaTokenUsuario ...
func ActualizaTokenUsuario(id uint, token string) error {
	var usr Usuario
	zonaMty, _ := time.LoadLocation("America/Monterrey")
	r := Db.Model(&usr).
		Where("id = ?", id).
		Updates(Usuario{Token: token, FechaToken: time.Now().In(zonaMty)})
	if r.Error != nil {
		return errors.New("error al actualizar tocken")
	}
	return nil
}

// ValidaToken ...
func ValidaToken(token string) (Usuario, error) {
	ln := len(token)
	var usr Usuario
	if ln > 0 {
		mtoken := token[7:ln]
		r := Db.Where(Usuario{Token: mtoken}).First(&usr)
		if r.Error != nil {
			return usr, errors.New("el usuario no existe en la plataforma o el token expiro")
		} else {
			mfecha := usr.FechaToken.Add(time.Hour * time.Duration(8))
			zonaMty, _ := time.LoadLocation("America/Monterrey")
			if mfecha.After(time.Now().In(zonaMty)) {
				return usr, nil
			} else {
				return usr, errors.New("la fecha del token expiro")
			}
		}
	} else {
		return usr, errors.New("no se encontro el token")
	}
}

// ValidaPassword ...
func ValidaPassword(email string, password string) (Usuario, error) {

	key := "rfq74564bdb6fc123706eea85ec98431"

	esMail := strings.Contains(email, "@")

	usr := Usuario{}
	var r *gorm.DB

	if esMail {
		r = Db.Where(Usuario{Email: email}).First(&usr)
	} else {
		r = Db.Where(Usuario{Usuario: email}).First(&usr)
	}

	if r.Error != nil {
		return usr, errors.New("el usuario no existe en la plataforma")
	} else {
		// Si es por RFC es un contribuyente valida que este dado de alta la permicionaria
		if !esMail {
			per := Permisionaria{}
			r = Db.Where(Permisionaria{Rfc: email}).First(per)
			if r.Error != nil {
				return usr, errors.New("el RFC no existe en la plataforma")
			}
		}

		pass, err := encrypt.Sdecrypt(usr.Password, key)
		if err != nil {
			return usr, err
		} else {
			if pass != password {
				return usr, errors.New("password es incorrecto")
			} else {
				return usr, nil
			}
		}
	}

}

// DescripcionTipoUsuario ...
func DescripcionTipoUsuario(tipo int) string {
	switch tipo {
	case TipoAdministrador:
		return "Administrador"
	case TipoAuditorGob:
		return "AuditorGob"
	case TipoAuditorSat:
		return "AuditorSat"
	case TipoAuditorCrayon:
		return "AuditorCrayon"
	case TipoContribuyente:
		return "Contribuyente"
	case TipoConsultas:
		return "Consultas"
	}

	return "NA"
}

// AccesosUsuario ...
func AccesosUsuario(tipo int) []string {
	var accesos []string

	accesos = append(accesos, "lista_embarques")

	switch tipo {
	case TipoAdministrador:
		accesos = append(accesos, "dashboard")
		accesos = append(accesos, "recoleccion")
		accesos = append(accesos, "impuestos")
		accesos = append(accesos, "alertas")
		accesos = append(accesos, "catalogos")
		accesos = append(accesos, "reportes")
	case TipoAuditorGob:
		accesos = append(accesos, "dashboard")
		accesos = append(accesos, "recoleccion")
		accesos = append(accesos, "impuestos")
		accesos = append(accesos, "alertas")
		accesos = append(accesos, "catalogos")
		accesos = append(accesos, "reportes")
	case TipoAuditorSat:
		accesos = append(accesos, "dashboard")
		accesos = append(accesos, "impuestos")
		accesos = append(accesos, "lineas_captura")
		accesos = append(accesos, "reportes")
	case TipoAuditorCrayon:
		accesos = append(accesos, "dashboard")
		accesos = append(accesos, "recoleccion")
		accesos = append(accesos, "impuestos")
		accesos = append(accesos, "alertas")
		accesos = append(accesos, "catalogos")
		accesos = append(accesos, "reportes")
	case TipoContribuyente:
		accesos = append(accesos, "declaraciones")
		accesos = append(accesos, "reportes")
	case TipoConsultas:
		accesos = append(accesos, "dashboard")
		accesos = append(accesos, "reportes")
	}

	return accesos

}

/*
{
    "email": "admin@sarfic.com.mx",
    "password": "S4rf1cP"
}
*/
