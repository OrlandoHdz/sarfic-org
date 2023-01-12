package models

import "gorm.io/gorm"

const (
	EstatusEnLinea       = 1
	EstatusMantenimiento = 2
	EstatusReparacion    = 3
	EstatusFueraLinea    = 4

	PropiedadCasino = 1
	PropiedadRenta  = 2
	PropiedadOtro   = 3

	ProtocoloSas      = 1
	ProtocoloSasPlus  = 2
	ProtocoloEmulador = 3
	ProtocoloManual   = 4
	ProtocoloOtro     = 5

	ComunicacionEnLinea    = 1
	ComunicacionFueraLinea = 2
)

var (
	DescEstatus      = [4]string{"En_Linea", "Mantenimiento", "Reparacion", "Fuera_Linea"}
	DescPropiedad    = [3]string{"Casino", "Renta", "Otro"}
	DescProtocolo    = [5]string{"Sas", "Sas_Plus", "Emulador", "Manual", "Otro"}
	DescComunicacion = [2]string{"En_Linea", "Fuera_Linea"}
)

// Maquina tabla de maquinas
type Maquina struct {
	gorm.Model
	CasinoID           uint    `json:"casino_id"`
	Casino             Casino  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Sala               string  `json:"sala" gorm:"type:character varying(250);"`
	Nombre             string  `json:"nombre" gorm:"type:character varying(250);"`
	Marca              string  `json:"marca" gorm:"type:character varying(250);"`
	Modelo             string  `json:"modelo" gorm:"type:character varying(250);"`
	NumeroSerie        string  `json:"numero_serie" gorm:"type:character varying(250);"`
	Estatus            uint    `json:"estatus_id"`
	DescEstatus        string  `json:"desc_estatus" gorm:"type:character varying(30);"`
	Propiedad          uint    `json:"propiedad_id"`
	DescPropiedad      string  `json:"desc_propiedad" gorm:"type:character varying(30);"`
	Protocolo          uint    `json:"protocolo_id"`
	Comunicacion       uint    `json:"comunicacion_id"`
	DescComunicacion   string  `json:"desc_comunicacion" gorm:"type:character varying(30);"`
	IpPrimaria         string  `json:"ip_primaria" gorm:"type:character varying(30);"`
	IpSecundaria       string  `json:"ip_secundaria" gorm:"type:character varying(30);"`
	IpPrimaria6        string  `json:"ip_primaria6" gorm:"type:character varying(30);"`
	IpSecundaria6      string  `json:"ip_secundaria6" gorm:"type:character varying(30);"`
	UsuarioActualizoID uint    `json:"usuario_actualizo_id"`
	UsuarioActualizo   Usuario `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// MigrarMaquina migra la tabla a la base de datos
func MigararMaquina() {
	Db.AutoMigrate(&Maquina{})
}
