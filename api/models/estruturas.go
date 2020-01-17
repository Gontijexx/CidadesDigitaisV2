package models

import "time"

type Entidade struct {
	Cnpj           float64 `gorm:"primary_key;not null;size:14" json:"cnpj"`
	Nome           string  `gorm:"size:50;default:null" json:"nome" validate: "alphanum":`
	Endereco       string  `gorm:"size:100;default:null" json:"endereco" validate: "alphanum":`
	Numero         string  `gorm:"size:10;default:null" json:"numero" validate: "alphanum":`
	Bairro         string  `gorm:"size:100;default:null" json:"bairro" validate: "alphanum":`
	Cep            string  `gorm:"size:8;default:null" json:"cep" validate: "alphanum":`
	Nome_municipio string  `gorm:"size:50;default:null" json:"nome_municipio" validate: "alphanum":`
	Uf             string  `gorm:"size:2;default:null" json:"uf" validate: "alphanum":`
	Observacao     string  `gorm:"size:1000;default:null" json:"observacao" validate: "alphanum":`
}

type Lote struct {
	Cod_lote      uint32    `gorm:"primary_key;not null;size:11" json:"cod_lote" validate: "required":`
	Cnpj          string    `gorm:"foreing_key;not null;size:14" json:"cnpj" validate: "required":`
	Contrato      string    `gorm:"size:10;default:null" json:"contrato" validate: "alphanum":`
	Dt_inicio_vig time.Date `gorm:"size:10;default:null" json:"dt_inicio_vig" validate: "alphanum":`
	Dt_final_vig  time.Date `gorm:"size:10;default:null" json:"dt_final_vig" validate: "alphanum":`
	Dt_reajuste   time.Date `gorm:"size:10;default:null" json:"dt_reajuste" validate: "alphanum":`
}

type Reajuste struct {
	Ano_ref    uint32  `gorm:"primary_key;not null;size:11" json:"ano_ref" validate: "required":`
	Cod_lote   uint32  `gorm:"foreing_key;not null;size:11" json:"cod_lote" validate: "required":`
	Percentual float64 `gorm:"default:null" json:"percentual" validate: "alphanum":`
}

type Cd struct {
	Cod_ibge uint32    `gorm:"primary_key;not null;size:7" json:"cod_ibge" validate: "required":`
	Cod_lote uint32    `gorm:"foreing_key;not null;size:11" json:"cod_lote" validate: "required":`
	Os_pe    string    `gorm:"size:10;default:null" json:"os_pe" validate: "alphanum":`
	Data_pe  time.Date `gorm:"default:null" json:"data_pe" validate: "alphanum":`
	Os_imp   time.Date `gorm:"size:10;default:null" json:"os_imp" validate: "alphanum":`
	Data_imp time.Date `gorm:"sdefault:null" json:"data_imp" validate: "alphanum":`
}
