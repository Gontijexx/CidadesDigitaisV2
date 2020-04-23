package models

/*  =========================
	TABELA UACOM ASSUNTO (CD)
=========================  */

type UacomAssunto struct {
	CodIbge    uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Data       string `gorm:"primary_key;foreign_key:Data;not null" json:"data"`
	CodAssunto uint64 `gorm:"primary_key;foreign_key:CodAssunto;not null" json:"cod_assunto"`
}

/*  =========================
	TABELA PONTO (CD)
=========================  */

type Ponto struct {
	CodPonto     uint64 `gorm:"primary_key;not null" json:"cod_ponto"`
	CodCategoria uint64 `gorm:"primary_key;foreign_key:CodCategoria;not null" json:"cod_categoria"`
	CodIbge      uint64 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	CodPid       uint64 `gorm:"foreign_key:CodPid;not null" json:"cod_pid"`
	Endereco     string `gorm:"default:null" json:"endereco"`
	Numero       string `gorm:"default:null;size:10" json:"numero"`
	Complemento  string `gorm:"default:null" json:"complemento"`
	Bairro       string `gorm:"default:null" json:"bairro"`
	Cep          string `gorm:"default:null;size:8" json:"cep"`
	Latitude     uint64 `gorm:"default:null" json:"latitude"`
	Longitude    uint64 `gorm:"default:null" json:"longitude"`
}
