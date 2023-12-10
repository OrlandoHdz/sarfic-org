package models

import "gorm.io/gorm"

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
