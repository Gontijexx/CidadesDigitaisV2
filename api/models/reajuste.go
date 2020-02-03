package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
		COMENTAR
=========================	*/

/*	=========================
		PRECISA DE MANUTENCAO
=========================	*/

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
	FUNCAO LISTAR REAJUSTE POR ID
=========================  */

func (reajuste *Reajuste) FindReajusteByID(db *gorm.DB, reajusteID uint64) (*Reajuste, error) {

	err := db.Debug().Model(Reajuste{}).Where("ano_ref = ?", reajusteID).Take(&reajuste).Error

	if err != nil {
		return &Reajuste{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Reajuste{}, errors.New("Reajuste Not Found")
	}

	return reajuste, err
}

func (r *Reajuste) UpdateReajuste(db *gorm.DB, rId1, rId2 uint64) (*Reajuste, error) {

	db = db.Debug().Model(&Reajuste{}).Where("ano_ref= ? AND cod_lote= ?", rId1, rId2).Take(&Reajuste{}).UpdateColumns(
		map[string]interface{}{
			"percentual": r.Percentual,
		},
	)

	if db.Error != nil {
		return &Reajuste{}, db.Error
	}

	err := db.Debug().Model(&Reajuste{}).Where("ano_ref= ? AND cod_lote= ?", rId1, rId2).Take(&r).Error
	if err != nil {
		return &Reajuste{}, err
	}

	return r, err

}

/*  =========================
	FUNCAO DELETAR ENTIDADE POR ID
=========================  */

func (r *Reajuste) DeleteReajuste(db *gorm.DB, codLoteID, anoRefID uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Reajuste{}).Where("cod_lote = ? AND ano_ref", codLoteID, anoRefID).Take(&Reajuste{}).Delete(&Reajuste{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
