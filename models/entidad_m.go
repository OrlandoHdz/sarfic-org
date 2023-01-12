package models

import (
	"gorm.io/gorm"
)

// Entidad ...
type Entidad struct {
	gorm.Model
	Nombre string `json:"nombre" gorm:"type:character varying(250);"`
}

// MigrarEntidad ...
func MigrarEntidad() {
	Db.AutoMigrate(&Entidad{})
}

// CrearEntidad ...
func CrearEntidad() {
	Db.Where(Entidad{Nombre: "BAJA CALIFORNIA"}).Assign(Entidad{Nombre: "BAJA CALIFORNIA"}).FirstOrCreate(&Entidad{})
}
