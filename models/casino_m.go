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
	Id                  uint
	EntidadId           uint
	Entidad             string
	PermisionariaId     uint
	PermisionariaRfc    string
	PermisionariaNombre string
	NombreComercial     string
	Direccion           string
	Colonia             string
	Municipio           string
	CodigoPostal        uint
	SistemaPrincipal    string
	NumeroMaquinas      uint
	NumeroMesas         uint
	SportsBook          bool
	ContactoNombre      string
	ContactoEmail       string
	ContactoTelefono    string
	ContactoMivil       string
}

func ObtenerCasinos(entidad_id int) ([]ResultadoQry, error) {
	datosQry := []ResultadoQry{}

	r := Db.Raw(`
		select 
			id,
			entidad_id,
			entidad,
			permisionaria_id,
			permisionaria_rfc,
			permisionaria_nombre,
			nombre_comercial,
			direccion,
			colonia,
			municipio,
			codigo_postal,
			sistema_principal,
			numero_maquinas,
			numero_mesas,
			sports_book,
			contacto_nombre,
			contacto_email,
			contacto_telefono,
			contacto_movil
		from org.vw_casinos
		where entidad_id = ?
	`, entidad_id).Find(&datosQry)

	if r.Error != nil {
		fmt.Println(r.Error)
		return nil, errors.New("Error al los obtener casinos")
	}

	return datosQry, nil

}

type ResultadoCiudadQry struct {
	Municipio  string
	Casinos    uint
	Maquinas   uint
	Mesas      uint
	Sportsbook uint
}

func ObtenerCasinosCiudad(entidad_id int) ([]ResultadoCiudadQry, error) {
	datosQry := []ResultadoCiudadQry{}

	r := Db.Raw(`
		select 
			municipio, count(*) as casinos, 
			sum(numero_maquinas) as maquinas, 
			sum(numero_mesas) as mesas,
			count(case when sports_book = true then 1 else null end) as sportsbook		
		from org.vw_casinos
		where entidad_id = ?
		group by municipio
	`, entidad_id).Find(&datosQry)

	if r.Error != nil {
		fmt.Println(r.Error)
		return nil, errors.New("Error al los obtener casinos ciudad")
	}

	return datosQry, nil

}
