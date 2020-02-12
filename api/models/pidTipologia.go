package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO DELETAR PID_TIPOLOGIA POR ID
=========================  */

func (pidTipologia *PIDTipologia) DeletePIDTipologia(db *gorm.DB, codPonto, codCategoria, codIbge, codTipologia uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&PIDTipologia{}).Where("cod_ponto = ? AND cod_categoria = ? AND cod_ibge = ? AND cod_tipologia = ?", codPonto, codCategoria, codIbge, codTipologia).Take(&PIDTipologia{}).Delete(&PIDTipologia{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("PIDTipologia not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
