package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PONTO NO BANCO DE DADOS
=========================  */

func (ponto *Ponto) SavePonto(db *gorm.DB) (*Ponto, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&ponto).Error
	if err != nil {
		return &Ponto{}, err
	}

	return ponto, nil
}

/*  =========================
	FUNCAO LISTAR PONTO POR ID
=========================  */

func (ponto *Ponto) FindPontoByID(db *gorm.DB, codPonto, codCategoria, codIbge uint64) (*Ponto, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Ponto{}).Where("cod_ponto = ? AND cod_categoria =? AND cod_ibge = ?", codPonto, codCategoria, codIbge).Take(&ponto).Error

	if err != nil {
		return &Ponto{}, err
	}

	return ponto, err
}

/*  =========================
	FUNCAO LISTAR TODAS PONTO
=========================  */

func (ponto *Ponto) FindAllPonto(db *gorm.DB) (*[]Ponto, error) {

	allPonto := []Ponto{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Ponto{}).Find(&allPonto).Error
	if err != nil {
		return &[]Ponto{}, err
	}
	return &allPonto, err
}

/*  =========================
	FUNCAO EDITAR PONTO
=========================  */

func (ponto *Ponto) UpdatePonto(db *gorm.DB, codPonto, codCategoria, codIbge uint64) (*Ponto, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Ponto{}).Where("cod_ponto = ? AND cod_categoria =? AND cod_ibge = ?", codPonto, codCategoria, codIbge).Updates(
		Ponto{
			Endereco:    ponto.Endereco,
			Numero:      ponto.Numero,
			Complemento: ponto.Complemento,
			Bairro:      ponto.Bairro,
			Cep:         ponto.Cep,
			Latitude:    ponto.Latitude,
			Longitude:   ponto.Longitude}).Error

	if db.Error != nil {
		return &Ponto{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Ponto{}).Where("cod_ponto = ? AND cod_categoria =? AND cod_ibge = ?", codPonto, codCategoria, codIbge).Take(&ponto).Error
	if err != nil {
		return &Ponto{}, err
	}

	// retorna o elemento que foi alterado
	return ponto, err
}

/*  =========================
	FUNCAO DELETAR PONTO POR ID
=========================  */

func (ponto *Ponto) DeletePonto(db *gorm.DB, codPonto, codCategoria, codIbge uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Ponto{}).Where("cod_ponto = ? AND cod_categoria =? AND cod_ibge = ?", codPonto, codCategoria, codIbge).Take(&Ponto{}).Delete(&Ponto{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Ponto not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
