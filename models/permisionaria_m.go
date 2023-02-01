package models

import "gorm.io/gorm"

// Permisionaria Contiene el permiso ante la secretaria de gobernación
type Permisionaria struct {
	gorm.Model
	Numero             string  `json:"numero" gorm:"type:character varying(50)"`
	Rfc                string  `json:"rfc" gorm:"type:character varying(50)"`
	Descripcion        string  `json:"descripcion" gorm:"type:character varying(250)"`
	UsuarioActualizoID uint    `json:"usuario_actualizo_id"`
	UsuarioActualizo   Usuario `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// MigrarPermisionaria migra la tabla a la base de datos
func MigrarPermisionaria() {
	Db.AutoMigrate(&Permisionaria{})
}

func (c *Permisionaria) CamposObligatorios() []string {
	var campos []string
	campos = append(campos, "Numero")
	campos = append(campos, "Rfc")
	campos = append(campos, "Descripcion")
	return campos
}
