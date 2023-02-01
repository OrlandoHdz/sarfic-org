package models

const (
	DEstatusPendiente = 1
	DEstatusCapturada = 2
	DEstatusRechazada = 3
	DEstatusAceptada  = 4

	DTipoNormal         = 1
	DTipoComplementaria = 2

	DPagoPendiente  = 1
	DPagoEfuectuado = 2
)

// Declaracion Contiene la declaración definitiva del contribuyente
type Declaracion struct {
	PermisionariaID         uint          `json:"permisionaria_id"`
	Permisionaria           Permisionaria `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EntidadID               uint          `json:"entidad_id"`
	Entidad                 Entidad       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DepositosPrecargado     float32       `json:"depositos_precargado" gorm:"type:numeric(15,4);"`
	RetirosPrecargado       float32       `json:"retiros_precargado" gorm:"type:numeric(15,4);"`
	PremiosPrecargado       float32       `json:"premios_precargado" gorm:"type:numeric(15,4);"`
	DevolucionesPrecargado  float32       `json:"devoluciones_precargado" gorm:"type:numeric(15,4);"`
	Depositos               float32       `json:"depositos" gorm:"type:numeric(15,4);"`
	Retiros                 float32       `json:"retiros" gorm:"type:numeric(15,4);"`
	Premios                 float32       `json:"premios" gorm:"type:numeric(15,4);"`
	Devoluciones            float32       `json:"devoluciones" gorm:"type:numeric(15,4);"`
	Impuesto788             float32       `json:"impuesto788" gorm:"type:numeric(15,4);"`
	Impuesto789             float32       `json:"impuesto789" gorm:"type:numeric(15,4);"`
	Impuesto790             float32       `json:"impuesto790" gorm:"type:numeric(15,4);"`
	TotalApagar             float32       `json:"total_apagar" gorm:"type:numeric(15,4);"`
	Año                     uint          `json:"año"`
	Mes                     uint          `json:"mes"`
	Tipo                    uint          `json:"tipo"`
	Estatus                 uint          `json:"estatus"`
	EstatusPago             uint          `json:"estatus_pago"`
	ComentarioAdministrador string        `json:"comentario_administrador" gorm:"type:character varying(250);"`
	ComentarioContribuyente string        `json:"comentario_contribuyente" gorm:"type:character varying(250);"`
	UsuarioActualizoID      uint          `json:"usuario_actualizo_id"`
	UsuarioActualizo        Usuario       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

/*
788 - IMPUESTO A LA REALIZACION DE JUEGOS CON APUESTAS Y SORTEOS Y OBTENCIÓN DE PREMIOS
789 - IMPUESTO A LAS EROGACIONES EN JUEGOS CON APUESTAS
790 - IMPUESTO POR LA OBTENCIÓN DE PREMIOS
*/

func MigrarDeclaracion() {
	Db.AutoMigrate(&Declaracion{})
}
