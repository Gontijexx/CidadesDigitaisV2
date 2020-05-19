package models

import (
	"github.com/jinzhu/gorm"
)

/*	=========================
	STRUCT FATURA
=========================	*/

type Fatura struct {
	NumNF         uint32 `gorm:"primary_key;not null" json:"num_nf"`
	CodIbge       uint32 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	NomeMunicipio string `gorm:"default:null" json:"nome_municipio"`
	Uf            string `gorm:"default:null" json:"uf"`
	DtNf          string `gorm:"default:null" json:"dt_nf"`
}

/*  =========================
	FUNCAO SALVAR FATURA
=========================  */

func (fatura *Fatura) SaveFatura(db *gorm.DB) error {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&fatura).Error

	return err
}

/*  =========================
	FUNCAO LISTAR FATURA POR ID
=========================  */

func (fatura *Fatura) FindFaturaByID(db *gorm.DB, numNF, codIbge uint32) error {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Fatura{}).Where("num_nf = ? AND cod_ibge = ?", numNF, codIbge).Take(&fatura).Error

	return err
}

/*  =========================
	FUNCAO LISTAR TODAS FATURA
=========================  */

func (fatura *Fatura) FindAllFatura(db *gorm.DB) (*[]Fatura, error) {

	allFatura := []Fatura{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Table("fatura").Select("municipio.nome_municipio, municipio.uf, fatura.*").
		Joins("JOIN municipio ON fatura.cod_ibge = municipio.cod_ibge ORDER BY fatura.num_nf ASC ").Scan(&allFatura).Error
	if err != nil {
		return &[]Fatura{}, err
	}

	return &allFatura, err
}

/*  =========================
	FUNCAO EDITAR FATURA
========================   */

func (fatura *Fatura) UpdateFatura(db *gorm.DB, numNF, codIbge uint32) error {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE fatura SET dt_nf = ? WHERE num_nf = ? AND cod_ibge = ?", fatura.DtNf, numNF, codIbge)

	return db.Error
}

/*  =========================
	FUNCAO DELETAR FATURA
=========================  */

func (fatura *Fatura) DeleteFatura(db *gorm.DB, numNF, codIbge uint32) error {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Fatura{}).Where("num_nf = ? AND cod_ibge = ?", numNF, codIbge).Take(&Fatura{}).Delete(&Fatura{})

	return db.Error
}
