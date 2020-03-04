package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR REAJUSTE NO BANCO DE DADOS
=========================  */

func (reajuste *Reajuste) SaveReajuste(db *gorm.DB) (*Reajuste, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&reajuste).Error
	if err != nil {
		return &Reajuste{}, err
	}
	return reajuste, nil

}

/*  =========================
	FUNCAO LISTAR TODOS REAJUSTE
=========================  */

func (reajuste *Reajuste) FindAllReajuste(db *gorm.DB) (*[]Reajuste, error) {

	allReajuste := []Reajuste{}

	//	Busca todos os elementos no banco de dados
	err := db.Debug().Model(&Entidade{}).Find(&allReajuste).Error
	if err != nil {
		return &[]Reajuste{}, err
	}

	return &allReajuste, err
}

/*  =========================
	FUNCAO DELETAR REAJUSTE POR ID
=========================  */

func (r *Reajuste) DeleteReajuste(db *gorm.DB, anoRef, codLote uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Reajuste{}).Where("ano_ref = ? AND cod_lote = ?", anoRef, codLote).Take(&Reajuste{}).Delete(&Reajuste{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
