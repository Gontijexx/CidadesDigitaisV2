package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCOES QUE FAZEM CONEXAO DIRETA COM O BANCO DE DADOS
=========================  */

/*  =========================
	FUNCAO PREPARE ENTIDADE
=========================  */

//	Prepara as informacoes a serem escritas no banco de dados
func (entidade *Entidade) Prepare() {
	entidade.Cnpj = 0
	entidade.Nome = html.EscapeString(strings.TrimSpace(entidade.Nome))
	entidade.Endereco = html.EscapeString(strings.TrimSpace(entidade.Endereco))
	entidade.Numero = html.EscapeString(strings.TrimSpace(entidade.Numero))
	entidade.Bairro = html.EscapeString(strings.TrimSpace(entidade.Bairro))
	entidade.Cep = html.EscapeString(strings.TrimSpace(entidade.Cep))
	entidade.Nome_municipio = html.EscapeString(strings.TrimSpace(entidade.Nome_municipio))
	entidade.Uf = html.EscapeString(strings.TrimSpace(entidade.Uf))
	entidade.Endereco = html.EscapeString(strings.TrimSpace(entidade.Endereco))
	entidade.Observacao = html.EscapeString(strings.TrimSpace(entidade.Observacao))

}

/*  =========================
	FUNCAO SALVAR ENTIDADE NO BANCO DE DADOS
=========================  */

func (entidade *Entidade) SaveEntidade(db *gorm.DB) (*Entidade, error) {

	err := db.Debug().Create(&entidade).Error
	if err != nil {
		return &Entidade{}, err
	}
	return entidade, nil

}

/*  =========================
	FUNCAO LISTAR ENTIDADE POR ID
=========================  */

func (entidade *Entidade) FindEntidadeByID(db *gorm.DB, entidadeID uint64) (*Entidade, error) {

	err := db.Debug().Model(Entidade{}).Where("cnpj = ?", entidadeID).Take(&entidade).Error

	if err != nil {
		return &Entidade{}, err
	}

	return entidade, err
}

/*  =========================
	FUNCAO LISTAR ENTIDADES
=========================  */

func (entidade *Entidade) FindEntidades(db *gorm.DB) (*[]Entidade, error) {

	entity := []Entidade{}
	err := db.Debug().Model(&Entidade{}).Limit(100).Find(&entity).Error
	if err != nil {
		return &[]Entidade{}, err
	}
	return &entity, err
}

/*  =========================
	FUNCAO EDITAR ENTIDADE
=========================  */

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

/*  =========================
	FUNCAO DELETER ENTIDADE POR ID
=========================  */

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
