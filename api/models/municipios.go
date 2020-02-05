package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR MUNICIPIOS NO BANCO DE DADOS
=========================  */

func (municipios *Municipios) SaveMunicipios(db *gorm.DB) (*Municipios, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&municipios).Error
	if err != nil {
		return &Municipios{}, err
	}
	return municipios, nil

}

/*  =========================
	FUNCAO LISTAR MUNICIPIOS POR ID
=========================  */

func (municipios *Municipios) FindMunicipiosByID(db *gorm.DB, codIbge uint64) (*Municipios, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Municipios{}).Where("cod_ibge = ?", codIbge).Take(&municipios).Error

	if err != nil {
		return &Municipios{}, err
	}

	return municipios, err
}

/*  =========================
	FUNCAO LISTAR TODOS MUNICIPIOS
=========================  */

func (municipios *Municipios) FindAllMunicipios(db *gorm.DB) (*[]Municipios, error) {

	allMunicipios := []Municipios{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Municipios{}).Limit(100).Find(&allMunicipios).Error
	if err != nil {
		return &[]Municipios{}, err
	}
	return &allMunicipios, err
}

/*  =========================
	FUNCAO EDITAR MUNICIPIOS
=========================  */

func (municipios *Municipios) UpdateMunicipios(db *gorm.DB, codIbge uint64) (*Municipios, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Model(&Municipios{}).Where("cod_ibge = ?", codIbge).Updates(
		Municipios{
			NomeMunicipio: municipios.NomeMunicipio,
			Populacao:     municipios.Populacao,
			UF:            municipios.UF,
			Regiao:        municipios.Regiao,
			Cnpj:          municipios.Cnpj,
			DistCapital:   municipios.DistCapital,
			Endereco:      municipios.Endereco,
			Numero:        municipios.Numero,
			Complemento:   municipios.Complemento,
			Bairro:        municipios.Bairro,
			Idhm:          municipios.Idhm,
			Latitude:      municipios.Latitude,
			Longitude:     municipios.Longitude}).Error

	if db.Error != nil {
		return &Municipios{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Municipios{}).Where("cod_ibge = ?", codIbge).Take(&municipios).Error
	if err != nil {
		return &Municipios{}, err
	}

	// retorna o elemento que foi alterado
	return municipios, err
}

/*  =========================
	FUNCAO DELETAR MUNICIPIOS POR ID
=========================  */

func (municipios *Municipios) DeleteMunicipios(db *gorm.DB, codIbge uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Municipios{}).Where("cod_ibge = ?", codIbge).Take(&Municipios{}).Delete(&Municipios{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Municipios not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
