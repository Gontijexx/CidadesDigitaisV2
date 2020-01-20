package models

//	Estrutura referente a tabela Entidade

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
