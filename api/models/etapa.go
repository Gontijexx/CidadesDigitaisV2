package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ETAPA NO BANCO DE DADOS
=========================  */

func (etapa *Etapa) SaveEtapa(db *gorm.DB) (*Etapa, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&etapa).Error
	if err != nil {
		return &Etapa{}, err
	}
	return etapa, nil

}

/*  =========================
	FUNCAO LISTAR ETAPA POR ID
=========================  */

func (etapa *Etapa) FindEtapaByID(db *gorm.DB, codEtapa uint64) (*Etapa, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Etapa{}).Where("cod_etapa = ?", codEtapa).Take(&etapa).Error

	if err != nil {
		return &Etapa{}, err
	}

	return etapa, err
}

/*  =========================
	FUNCAO LISTAR TODAS ETAPA
=========================  */

func (etapa *Etapa) FindAllEtapa(db *gorm.DB) (*[]Etapa, error) {

	allEtapa := []Etapa{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Etapa{}).Limit(100).Find(&allEtapa).Error
	if err != nil {
		return &[]Etapa{}, err
	}
	return &allEtapa, err
}

/*  =========================
	FUNCAO EDITAR ETAPA
=========================  */

func (etapa *Etapa) UpdateEtapa(db *gorm.DB, codEtapa uint64) (*Etapa, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Etapa{}).Where("cod_etapa = ?", codEtapa).Updates(
		Etapa{
			Descricao: etapa.Descricao,
			Duracao:   etapa.Duracao,
			Depende:   etapa.Depende,
			Delay:     etapa.Delay,
			SetorResp: etapa.SetorResp}).Error

	if db.Error != nil {
		return &Etapa{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Etapa{}).Where("cod_etapa = ?", codEtapa).Take(&etapa).Error
	if err != nil {
		return &Etapa{}, err
	}

	// retorna o elemento que foi alterado
	return etapa, err
}

/*  =========================
	FUNCAO DELETAR ETAPA POR ID
=========================  */

func (etapa *Etapa) DeleteEtapa(db *gorm.DB, codEtapa uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Etapa{}).Where("cod_etapa = ?", codEtapa).Take(&Etapa{}).Delete(&Etapa{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Etapa not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
