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
	CasinoID          uint   `json:"casino_id" gorm:"index:idx_censo_casino,unique"`
	Casino            Casino `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SistemaPrincipal  string `json:"sistema_principal" gorm:"type:character varying(250);"`
	SistemaSecunadrio string `json:"sistema_Secundario" gorm:"type:character varying(250);"`
	NumeroMaquinas    int    `json:"numero_maquinas"`
	NumeroMesas       int    `json:"numero_mesas"`
	SportsBook        bool   `json:"sportsbook"`
	PersonaAtendio    string `json:"persona_atendio" gorm:"type:character varying(250);"`
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

type ResultadoCensoQry struct {
	Id                  uint
	CasinoId            uint
	SistemaPrincipal    string
	NumeroMaquinas      uint
	NumeroMesas         uint
	SportsBook          bool
	PersonaAtendio      string
	UpdateAt            time.Time
	NombreComercial     string
	PermisionariaRfc    string
	PermisionariaNombre string
	Direccion           string
	Colonia             string
	Municipio           string
	CodigoPostal        uint
}

func ObtenerCenso(entidad_id int) ([]ResultadoCensoQry, error) {
	datosQry := []ResultadoCensoQry{}
	r := Db.Raw(`
		select 
			ce.id,
			ce.casino_id,
			ce.sistema_principal,
			ce.numero_maquinas,
			ce.numero_mesas,
			ce.sports_book,
			ce.persona_atendio,
			ce.updated_at,
			ca.nombre_comercial,
			ca.permisionaria_rfc,
			ca.permisionaria_nombre,
			ca.direccion,
			ca.colonia,
			ca.municipio,
			ca.codigo_postal			
		from 
			org.censos ce,
			org.vw_casinos ca
		where 
			entidad_id = ?
			and ce.deleted_at is null
			and ce.casino_id = ca.id 

	`, entidad_id).Find(&datosQry)

	if r.Error != nil {
		fmt.Println(r.Error)
		return nil, errors.New("Error al los obtener casinos")
	}

	return datosQry, nil
}
