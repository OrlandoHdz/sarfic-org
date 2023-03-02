package models

import (
	"time"

	"gorm.io/gorm"
)

// Oficio definicion de la tabla en la base de datos
type Oficio struct {
	gorm.Model
	CasinoNombre string    `json:"casino_nombre" gorm:"type:character varying(250);"`
	Numero       string    `json:"numero" gorm:"type:character varying(250);"`
	FechaLLegada time.Time `json:"fecha_llegada"`
	FechaCierre  time.Time `json:"fecha_cierre"`
	Archivo      string    `json:"archivo" gorm:"type:character varying(250);"`
}

// MigrarOficio migrar tabla
func MigrarOficio() {
	Db.AutoMigrate(&Oficio{})
}

func (c *Oficio) CamposObligatorios() []string {
	var campos []string
	campos = append(campos, "CasinoNombre")
	campos = append(campos, "Numero")
	campos = append(campos, "FechaLLegada")
	campos = append(campos, "FechaCierre")
	campos = append(campos, "Archivo")
	return campos
}
