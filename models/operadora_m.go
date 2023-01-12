package models

import "gorm.io/gorm"

// Operadora tabla de operadora
type Operadora struct {
	gorm.Model
	EntidadID          uint          `json:"entidad_id"`
	Entidad            Entidad       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PermisionariaID    uint          `json:"permisionaria_id"`
	Permisionaria      Permisionaria `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Numero             string        `json:"numero" gorm:"type:character varying(50)"`
	Descripcion        string        `json:"descripcion" gorm:"type:character varying(250)"`
	UsuarioActualizoID uint          `json:"usuario_actualizo_id"`
	UsuarioActualizo   Usuario       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// MigrarOperadora migra a la base de datos
func MigrarOperadora() {
	Db.AutoMigrate(&Operadora{})
}
