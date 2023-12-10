package models

import "gorm.io/gorm"

type Recaudacion struct {
	gorm.Model
	EntidadID   uint    `json:"entidad_id"`
	Entidad     Entidad `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Año         uint    `json:"año" gorm:"index:idx_casino_principal,priority:1"`
	Mes         uint    `json:"mes" gorm:"index:idx_casino_principal,priority:2"`
	Rfc         string  `json:"rfc" gorm:"type:character varying(50)"`
	Nombre      string  `json:"nombre" gorm:"type:character varying(100)"`
	Impuesto788 float32 `json:"impuesto_788"`
	Impuesto789 float32 `json:"impuesto_789"`
	Impuesto790 float32 `json:"impuesto_790"`
	Recargos    float32 `json:"recargos"`
	Redondeo    float32 `json:"redondeo"`
	Tipo        uint    `json:"yipo"`
}

func MigrarRecaudacion() {
	Db.AutoMigrate(&Recaudacion{})
}

func (r *Recaudacion) CamposObligatorios() []string {
	var campos []string
	campos = append(campos, "Año")
	campos = append(campos, "Mes")
	campos = append(campos, "Rfc")
	campos = append(campos, "Nombre")
	campos = append(campos, "Impuesto788")
	campos = append(campos, "Impuesto789")
	campos = append(campos, "Impuesto790")
	campos = append(campos, "Tipo")

	return campos
}
