package models

import "gorm.io/gorm"

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
