package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR ENTIDADE NO BANCO DE DADOS
=========================  */

func (entidade *Entidade) SaveEntidade(db *gorm.DB) (*Entidade, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}

	return entidade, nil
}

/*  =========================
	FUNCAO LISTAR ENTIDADE POR ID
=========================  */

func (entidade *Entidade) FindEntidadeByID(db *gorm.DB, cnpj uint64) (*Entidade, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Entidade{}).Where("cnpj = ?", cnpj).Take(&entidade).Error

	if err != nil {
		return &Entidade{}, err
	}

	return entidade, err
}

/*  =========================
	FUNCAO LISTAR TODAS ENTIDADE
=========================  */

func (entidade *Entidade) FindAllEntidade(db *gorm.DB) (*[]Entidade, error) {

	allEntidade := []Entidade{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Entidade{}).Find(&allEntidade).Error
	if err != nil {
		return &[]Entidade{}, err
	}
	return &allEntidade, err
}

/*  =========================
	FUNCAO EDITAR ENTIDADE
=========================  */

func (entidade *Entidade) UpdateEntidade(db *gorm.DB, cnpj uint64) (*Entidade, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE entidade SET nome = ?, endereco = ?, numero = ?, bairro = ?, cep = ?, nome_municipio = ?, uf = ?, observacao = ? WHERE cnpj = ?", entidade.Nome, entidade.Endereco, entidade.Numero, entidade.Bairro, entidade.Cep, entidade.NomeMunicipio, entidade.UF, entidade.Observacao, cnpj).Error

	if db.Error != nil {
		return &Entidade{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Entidade{}).Where("cnpj = ?", cnpj).Take(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}

	// retorna o elemento que foi alterado
	return entidade, err
}

/*  =========================
	FUNCAO DELETAR ENTIDADE POR ID
=========================  */

func (entidade *Entidade) DeleteEntidade(db *gorm.DB, cnpj uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Entidade{}).Where("cnpj = ?", cnpj).Take(&Entidade{}).Delete(&Entidade{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Entidade not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

/*  =========================
	FUNCAO LISTAR ENTIDADE.CNPJ E ENTIDADE.NOME
=========================  */

func (entidade *Entidade) GetEntidadeIDAndName(db *gorm.DB) (*[]Entidade, error) {

	allEntidade := []Entidade{}

	err := db.Debug().Select("cnpj, nome").Find(&allEntidade).Error
	if err != nil {
		return &[]Entidade{}, err
	}

	return &allEntidade, err
}
