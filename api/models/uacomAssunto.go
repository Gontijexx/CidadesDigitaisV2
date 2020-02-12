package models

import "github.com/jinzhu/gorm"

/*  =========================
	FUNCAO SALVAR UACOM_ASSUNTO NO BANCO DE DADOS
=========================  */

func (uacomAssunto *UacomAssunto) SaveUacomAssunto(db *gorm.DB) (*UacomAssunto, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&uacomAssunto).Error
	if err != nil {
		return &UacomAssunto{}, err
	}

	return uacomAssunto, nil
}
