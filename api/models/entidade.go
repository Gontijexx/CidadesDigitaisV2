package models

import (
	"github.com/jinzhu/gorm"
)

func (entidade *Entidade) SaveEntidade(db *gorm.DB) (*Entidade, error) {

	err := db.Debug().Create(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}
	return entidade, nil

}

/*
func (entidade *Entidade) FindEntidadeByID(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	err := db.Debug().Model(Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error

	if err != nil {
		return &Entidade{}, err
	}

	return entidade, err
}
*/
func (entidade *Entidade) UpdateEntidade(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	db = db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).UpdateColumns(
		map[string]interface{}{
			"nome":           entidade.Nome,
			"endereco":       entidade.Endereco,
			"numero":         entidade.Numero,
			"bairro":         entidade.Bairro,
			"cep":            entidade.Cep,
			"nome_municipio": entidade.Nome_municipio,
			"uf":             entidade.Uf,
			"observacao":     entidade.Observacao,
		},
	)

	if db.Error != nil {
		return &Entidade{}, db.Error
	}

	err := db.Debug().Model(&Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}

	return entidade, err
}

/*
func (entidade *Entidade) DeleteEntidade(db *gorm.DB, entidadeID uint64) (int64, error) {

	db = db.Debug().Model(&Usuario{}).Where("cnpj = ?", entidadeID).Take(&Entidade{}).Delete(&Entidade{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Entidade not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
*/
