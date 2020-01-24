package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO ADICIONA REAJUSTE
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
	FUNCAO LISTAR REAJUSTE
=========================  */

func (e *Reajuste) FindReajustes(db *gorm.DB) (*[]Reajuste, error) {
	reajuste := []Reajuste{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Reajuste{}).Find(&reajuste).Error
	if err != nil {
		return &[]Reajuste{}, err
	}
	return &reajuste, err
}

/*  =========================
FUNCAO LISTAR REAJUSTE POR ID
=========================  */

func (reajuste *Reajuste) FindReajuste(db *gorm.DB, reajusteID uint64) (*Reajuste, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Reajuste{}).Where("ano_ref = ?", reajusteID).Take(&reajuste).Error

	if err != nil {
		return &Reajuste{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Reajuste{}, errors.New("Reajuste Not Found")
	}

	return reajuste, err
}

/*  =========================
	FUNCAO ATUALIZAR CD_ITENS
=========================  */

func (reajuste *Reajuste) UpdateReajuste(db *gorm.DB, rId1, rId2 uint64) (*Reajuste, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&Reajuste{}).Where("ano_ref= ? AND cod_lote= ?", rId1, rId2).Take(&Reajuste{}).UpdateColumns(
		map[string]interface{}{
			"percentual": reajuste.Percentual,
		},
	)

	if db.Error != nil {
		return &Reajuste{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Reajuste{}).Where("ano_ref= ? AND cod_lote= ?", rId1, rId2).Take(&reajuste).Error
	if err != nil {
		return &Reajuste{}, err
	}

	// retorna o elemento que foi alterado
	return reajuste, err

}

func (r *Reajuste) DeleteReajuste(db *gorm.DB, rID1 uint64, rID2 uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Reajuste{}).Where("ano_ref = ? AND cod_lote", rID1, rID2).Take(&Reajuste{}).Delete(&Reajuste{})

	if db.Error != nil {
		return 0, db.Error
	}

	//	Retorna um valor nulo caso seja deletada a row
	return db.RowsAffected, nil
}
