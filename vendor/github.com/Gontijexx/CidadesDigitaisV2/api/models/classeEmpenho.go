package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR CLASSE EMPENHO NO BANCO DE DADOS
=========================  */

func (classeEmpenho *ClasseEmpenho) SaveClasseEmpenho(db *gorm.DB) (*ClasseEmpenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&classeEmpenho).Error
	if err != nil {
		return &ClasseEmpenho{}, err
	}
	return classeEmpenho, nil

}

/*  =========================
	FUNCAO LISTAR CLASSE EMPENHO POR ID
=========================  */

func (classeEmpenho *ClasseEmpenho) FindClasseEmpenhoByID(db *gorm.DB, codClasseEmpenho uint64) (*ClasseEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(ClasseEmpenho{}).Where("cod_classe_empenho = ?", codClasseEmpenho).Take(&classeEmpenho).Error

	if err != nil {
		return &ClasseEmpenho{}, err
	}

	return classeEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS CLASSE EMPENHO
=========================  */

func (classeEmpenho *ClasseEmpenho) FindAllClasseEmpenho(db *gorm.DB) (*[]ClasseEmpenho, error) {

	allClasseEmpenho := []ClasseEmpenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&ClasseEmpenho{}).Find(&allClasseEmpenho).Error
	if err != nil {
		return &[]ClasseEmpenho{}, err
	}
	return &allClasseEmpenho, err
}

/*  =========================
	FUNCAO EDITAR CLASSE EMPENHO
=========================  */

func (classeEmpenho *ClasseEmpenho) UpdateClasseEmpenho(db *gorm.DB, codClasseEmpenho uint64) (*ClasseEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE classe_empenho SET descricao = ? WHERE cod_classe_empenho = ?", classeEmpenho.Descricao, codClasseEmpenho).Error

	if db.Error != nil {
		return &ClasseEmpenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&ClasseEmpenho{}).Where("cod_classe_empenho = ?", codClasseEmpenho).Take(&classeEmpenho).Error
	if err != nil {
		return &ClasseEmpenho{}, err
	}

	// retorna o elemento que foi alterado
	return classeEmpenho, err
}

/*  =========================
	FUNCAO DELETAR CLASSE EMPENHO POR ID
=========================  */

func (classeEmpenho *ClasseEmpenho) DeleteClasseEmpenho(db *gorm.DB, codClasseEmpenho uint64) (int64, error) {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&ClasseEmpenho{}).Where("cod_classe_empenho = ?", codClasseEmpenho).Take(&ClasseEmpenho{}).Delete(&ClasseEmpenho{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Classe Empenho not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
