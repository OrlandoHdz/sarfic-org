package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Casino definicion de tabla en la base de datos
type Casino struct {
	gorm.Model
	EntidadID          uint      `json:"entidad_id"`
	Entidad            Entidad   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OperadoraID        uint      `json:"operadora_id"`
	Operadora          Operadora `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	NombreComercial    string    `json:"nombre_comercial" gorm:"type:character varying(250);"`
	Direccion          string    `json:"direccion" gorm:"type:character varying(250);"`
	Colonia            string    `json:"colonia" gorm:"type:character varying(250);"`
	Municipio          string    `json:"municipio" gorm:"type:character varying(250);"`
	CodigoPostal       uint      `json:"codigo_postal"`
	SistemaPrincipal   string    `json:"sistema_principal" gorm:"type:character varying(250);"`
	SistemaSecunadrio  string    `json:"sistema_Secundario" gorm:"type:character varying(250);"`
	NumeroMaquinas     int       `json:"numero_maquinas"`
	NumeroMesas        int       `json:"numero_mesas"`
	SportsBook         bool      `json:"sportsbook"`
	ContactoNombre     string    `json:"contacto_nombre" gorm:"type:character varying(250);"`
	ContactoPuesto     string    `json:"contacto_puesto" gorm:"type:character varying(250);"`
	ContactoEmail      string    `json:"contacto_email" gorm:"type:character varying(250);"`
	ContactoTelefono   string    `json:"contacto_telefono" gorm:"type:character varying(20);"`
	ContactoMovil      string    `json:"contacto_movil" gorm:"type:character varying(20);"`
	ContactoHoario     string    `json:"contacto_horario" gorm:"type:character varying(20);"`
	UsuarioActualizoID uint      `json:"usuario_actualizo_id"`
	UsuarioActualizo   Usuario   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// MigrarCasino migra la tabla
func MigrarCasino() {
	Db.AutoMigrate(&Casino{})
}

func (c *Casino) CamposObligatorios() []string {
	var campos []string
	campos = append(campos, "OperadoraID")
	campos = append(campos, "NombreComercial")
	campos = append(campos, "Direccion")
	campos = append(campos, "Colonia")
	campos = append(campos, "Municipio")
	campos = append(campos, "NumeroMaquinas")
	campos = append(campos, "NumeroMesas")
	campos = append(campos, "SportsBook")

	return campos
}

type ResultadoQry struct {
	EntidadId           uint
	Entidad             string
	PermisionariaId     uint
	PermisionariaRfc    string
	PermisionariaNombre string
	NombreComercial     string
	Direccion           string
	Colonia             string
	CodigoPostal        uint
	SistemaPrincipal    string
	NumeroMaquinas      uint
	NumeroMesas         uint
	ContactoNombre      string
	ContactoEmail       string
	ContactoTelefono    string
	ContactoMivil       string
}

func ObtenerCasinos(entidad_id int) ([]ResultadoQry, error) {
	datosQry := []ResultadoQry{}

	// r := Db.Table("casinos").
	// 	Select(`
	// 		entidads.id entidad_id,
	// 		entidads.nombre entidad,
	// 		operadoras.permisionaria_id,
	// 		permisionaria.rfc permisionaria_rfc,
	// 		permisionaria.descripcion permisionaria_nombre,
	// 		casinos.nombre_comercial,
	// 		casinos.direccion,
	// 		casinos.colonia,
	// 		casinos.municipio,
	// 		casinos.codigo_postal,
	// 		casinos.sistema_principal,
	// 		casinos.numero_maquinas,
	// 		casinos.numero_mesas,
	// 		casinos.contacto_nombre,
	// 		casinos.contacto_email,
	// 		casinos.contacto_telefono,
	// 		casinos.contacto_movil
	// 	`).Joins("lef join entidads on entidads.id = casinos.entidad_id").
	// 	Joins("left join  operadoras.id = casinos.operadora_id").
	// 	Joins("left join permisionaria.id = operadoras.permisionaria_id").
	// 	Where("casinos.entidad_id = ?", entidad_id).
	// 	Order("permisionaria.rfc, casinos.nombre_comercial").
	// 	Find(&datosQry)

	r := Db.Raw(`
		select 
			entidad_id,
			entidad,
			permisionaria_id,
			permisionaria_rfc,
			permisionaria,
			nombre_comercial,
			direccion,
			colonia,
			municipio,
			codigo_postal,
			sistema_principal,
			numero_maquinas,
			numero_mesas,
			contacto_nombre,
			contacto_email,
			contacto_telefono,
			contacto_movil
		from vw_casinos
		where entidad_id = ?
	`, entidad_id).Find(&datosQry)

	if r.Error != nil {
		fmt.Println(r.Error)
		return nil, errors.New("Error al los obtener casinos")
	}

	return datosQry, nil

}
