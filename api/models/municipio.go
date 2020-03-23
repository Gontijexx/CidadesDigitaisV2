package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR MUNICIPIO NO BANCO DE DADOS
=========================  */

func (municipio *Municipio) SaveMunicipio(db *gorm.DB) (*Municipio, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&municipio).Error
	if err != nil {
		return &Municipio{}, err
	}
	return municipio, nil

}

/*  =========================
	FUNCAO LISTAR MUNICIPIO POR ID
=========================  */

func (municipio *Municipio) FindMunicipioByID(db *gorm.DB, codIbge uint64) (*Municipio, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Municipio{}).Where("cod_ibge = ?", codIbge).Take(&municipio).Error

	if err != nil {
		return &Municipio{}, err
	}

	return municipio, err
}

/*  =========================
	FUNCAO LISTAR TODOS MUNICIPIO
=========================  */

func (municipio *Municipio) FindAllMunicipio(db *gorm.DB) (*[]Municipio, error) {

	allMunicipio := []Municipio{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Municipio{}).Find(&allMunicipio).Error
	if err != nil {
		return &[]Municipio{}, err
	}
	return &allMunicipio, err
}

/*  =========================
	FUNCAO EDITAR MUNICIPIO
=========================  */

func (municipio *Municipio) UpdateMunicipio(db *gorm.DB, codIbge uint64) (*Municipio, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE municipio SET nome_municipio = ?, populacao = ?, uf = ?, regiao = ?, cnpj = ?, dist_capital = ?, endereco = ?, numero = ?, complemento = ?, bairro = ?, idhm = ?, latitude = ?, longitude = ? WHERE cod_ibge = ?", municipio.NomeMunicipio, municipio.Populacao, municipio.UF, municipio.Regiao, municipio.Cnpj, municipio.DistCapital, municipio.Endereco, municipio.Numero, municipio.Complemento, municipio.Bairro, municipio.Idhm, municipio.Latitude, municipio.Longitude, codIbge).Error

	if db.Error != nil {
		return &Municipio{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Municipio{}).Where("cod_ibge = ?", codIbge).Take(&municipio).Error
	if err != nil {
		return &Municipio{}, err
	}

	// retorna o elemento que foi alterado
	return municipio, err
}

/*  =========================
	FUNCAO DELETAR MUNICIPIO POR ID
=========================  */

func (municipio *Municipio) DeleteMunicipio(db *gorm.DB, codIbge uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Municipio{}).Where("cod_ibge = ?", codIbge).Take(&Municipio{}).Delete(&Municipio{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Municipio not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

/*  =========================
	FUNCAO LISTAR MUNICIPIO.COD_IBGE E MUNICIPIO.NOME_MUNICIPIO
=========================  */

func (municipio *Municipio) FindMunicipioIDandNomeMunicipio(db *gorm.DB) (*[]Municipio, error) {

	allMunicipio := []Municipio{}

	err := db.Debug().Select("cod_ibge, nome_municipio").Find(&allMunicipio).Error
	if err != nil {
		return &[]Municipio{}, err
	}

	return &allMunicipio, err
}
