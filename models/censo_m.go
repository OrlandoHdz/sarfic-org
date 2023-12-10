package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Censo definicion de la tabla
type Censo struct {
	gorm.Model
	CasinoID          uint      `json:"casino_id" gorm:"index:idx_censo_casino,unique"`
	Casino            Casino    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SistemaPrincipal  string    `json:"sistema_principal" gorm:"type:character varying(250);"`
	SistemaSecunadrio string    `json:"sistema_Secundario" gorm:"type:character varying(250);"`
	NumeroMaquinas    int       `json:"numero_maquinas"`
	NumeroMesas       int       `json:"numero_mesas"`
	SportsBook        bool      `json:"sportsbook"`
	PersonaAtendio    string    `json:"persona_atendio" gorm:"type:character varying(250);"`
	Fecha             time.Time `gorm:"type:datetime;"`
	FechaAct          time.Time `json:"fecha_act"`
}

// MigrarCenso migrar la tabla
func MigrarCenso() {
	Db.AutoMigrate(&Censo{})
}

func (c *Censo) CamposObligatoriosCenso() []string {
	var campos []string
	campos = append(campos, "CasinoID")
	campos = append(campos, "SistemaPrincipal")
	campos = append(campos, "NumeroMaquinas")
	return campos
}

// type ResultadoCensoQry struct {
// 	Id                     uint      `json:"id"`
// 	CasinoId               uint      `json:"casino_id"`
// 	SistemaPrincipal       string    `json:"sistema_principal"`
// 	NumeroMaquinas         uint      `json:"numero_maquinas"`
// 	NumeroMesas            uint      `json:"numero_mesas"`
// 	SportsBook             bool      `json:"sports_book"`
// 	PersonaAtendio         string    `json:"persona_atendio"`
// 	UpdateAt               time.Time `json:"update_at"`
// 	UpdateAtStr            string    `json:"update_at_str"`
// 	NombreComercial        string    `json:"nombre_comercial"`
// 	PermisionariaRfc       string    `json:"permisionaria_rfc"`
// 	PermisionariaNombre    string    `json:"permisionaria_nombre"`
// 	Direccion              string    `json:"direccion"`
// 	Colonia                string    `json:"colonia"`
// 	Municipio              string    `json:"municipio"`
// 	CodigoPostal           uint      `json:"codigo_postal"`
// 	SistemaPrincipalCasino string    `json:"sistema_principal_casino"`
// 	NumeroMaquinasCasino   uint      `json:"numero_maquinas_casino"`
// 	NumeroMesasCasino      uint      `json:"numero_mesas_casino"`
// 	SportsBookCasino       bool      `json:"sports_book_casino"`
// }

type ResultadoCensoQry struct {
	Id                     uint
	CasinoId               uint
	SistemaPrincipal       string
	NumeroMaquinas         uint
	NumeroMesas            uint
	SportsBook             bool
	PersonaAtendio         string
	FechaAct               time.Time
	NombreComercial        string
	PermisionariaRfc       string
	PermisionariaNombre    string
	Direccion              string
	Colonia                string
	Municipio              string
	CodigoPostal           uint
	SistemaPrincipalCasino string
	NumeroMaquinasCasino   uint
	NumeroMesasCasino      uint
	SportsBookCasino       bool
}

func ObtenerCenso(entidad_id int) ([]ResultadoCensoQry, error) {

	datosQry := []ResultadoCensoQry{}

	r := Db.Table("org.censos").
		Select(`
		censos.id,
		censos.casino_id,
		censos.sistema_principal,
		censos.numero_maquinas,
		censos.numero_mesas,
		censos.sports_book,
		censos.persona_atendio,
		censos.fecha_act AT TIME ZONE 'CST',
		vw_casinos.nombre_comercial,
		vw_casinos.permisionaria_rfc,
		vw_casinos.permisionaria_nombre,
		vw_casinos.direccion,
		vw_casinos.colonia,
		vw_casinos.municipio,
		vw_casinos.codigo_postal,
		vw_casinos.sistema_principal sistema_principal_casino,
		vw_casinos.numero_maquinas numero_maquinas_casino,
		vw_casinos.numero_mesas numero_mesas_casino,
		vw_casinos.sports_book sports_book_casino
	`).Joins("left join org.vw_casinos on org.vw_casinos.id = org.censos.casino_id").
		Where("org.vw_casinos.entidad_id = ? ", entidad_id).
		Find(&datosQry)

	if r.Error != nil {
		fmt.Println(r.Error)
		return nil, errors.New("Error al los obtener casinos")
	}

	return datosQry, nil
}
