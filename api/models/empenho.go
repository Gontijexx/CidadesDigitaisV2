package models

import "github.com/jinzhu/gorm"

/*  =========================
	STRUCT EMPENHO
=========================  */

type Empenho struct {
	IDEmpenho          uint32 `gorm:"primary_key;auto_increment;default:not null" json:"id_empenho"`
	CodPrevisaoEmpenho uint32 `gorm:"foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho"`
	CodEmpenho         string `gorm:"not null;size:13" json:"cod_empenho"`
	CodLote            uint32 `gorm:"default:null" json:"cod_lote"`
	Data               string `gorm:"default:null" json:"data"`
	Contador           uint32 `gorm:"default:null" json:"contador"`
}

/*  =========================
	FUNCAO SALVAR EMPENHO
=========================  */

func (empenho *Empenho) SaveEmpenho(db *gorm.DB) (*Empenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&empenho).Error
	if err != nil {
		return &Empenho{}, err
	}

	return empenho, err
}

/*  =========================
	FUNCAO LISTAR EMPENHO POR ID
=========================  */

func (empenho *Empenho) FindEmpenhoByID(db *gorm.DB, idEmpenho uint32) (*Empenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Empenho{}).Where("id_empenho = ?", idEmpenho).Take(&empenho).Error
	if err != nil {
		return &Empenho{}, err
	}

	return empenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS EMPENHO
=========================  */

func (empenho *Empenho) FindAllEmpenho(db *gorm.DB) (*[]Empenho, error) {

	allEmpenho := []Empenho{}

	//	Busca todos elementos contidos no banco de dados
	err := db.Debug().Table("empenho").Select("previsao_empenho.cod_lote, empenho.*").
		Joins("JOIN previsao_empenho ON empenho.cod_previsao_empenho = previsao_empenho.cod_previsao_empenho ORDER BY empenho.id_empenho ASC").Scan(&allEmpenho).Error
	if err != nil {
		return &[]Empenho{}, err
	}
	return &allEmpenho, err
}

/*  =========================
	FUNCAO EDITAR EMPENHO
=========================  */

func (empenho *Empenho) UpdateEmpenho(db *gorm.DB, idEmpenho uint32) (*Empenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Exec("UPDATE empenho SET cod_previsao_empenho = ?, cod_empenho = ?, data = ?, contador = ? WHERE id_empenho = ?", empenho.CodPrevisaoEmpenho, empenho.CodEmpenho, empenho.Data, empenho.Contador, idEmpenho)
	if db.Error != nil {
		return &Empenho{}, db.Error
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&Empenho{}).Where("id_empenho = ?", idEmpenho).Take(&empenho).Error
	if err != nil {
		return &Empenho{}, err
	}

	return empenho, err
}

/*  =========================
	FUNCAO DELETAR EMPENHO
=========================

func (empenho *Empenho) DeleteEmpenho(db *gorm.DB, idEmpenho uint) (error) {

	//	Deleta  um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Empenho{}).Where("id_empenho = ?", idEmpenho).Take(&Empenho{}).Delete(&Empenho{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("Empenho not found")
		}
		return db.Error
	}

	return db.Error
}

*/
