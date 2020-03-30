package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Log struct {
	CodLog      uint64    `gorm:"primary_key;auto_increment;not null" json:"cod_log"`
	CodUsuario  uint64    `gorm:"primary_key;foreign_key:CodUsuario;not null" json:"cod_usuario"`
	Data        time.Time `gorm:"not null" json:"data"`
	NomeTabela  string    `gorm:"not null" json:"nome_tabela"`
	Operacao    string    `gorm:"not null" json:"operacao"`
	Espelho     string    `gorm:"default:null" json:"espelho"`
	CodInt1     uint64    `gorm:"default:null" json:"cod_int_1"`
	CodInt2     uint64    `gorm:"default:null" json:"cod_int_2"`
	CodInt3     uint64    `gorm:"default:null" json:"cod_int_3"`
	CodInt4     uint64    `gorm:"default:null" json:"cod_int_4"`
	CodInt5     uint64    `gorm:"default:null" json:"cod_int_5"`
	CodData     string    `gorm:"default:null" json:"cod_data"`
	CodProcesso string    `gorm:"default:null;size:17" json:"cod_processo"`
	CodCnpj     string    `gorm:"default:null;size:14" json:"cod_cnpj"`
	CodEmpenho  string    `gorm:"default:null;size:13" json:"cod_empenho"`
}

func (log *Log) ConcatEntidade(db *gorm.DB, cnpj string) (*Log, error) {

	err := db.Table("entidade").Select("CONCAT(IFNULL(cnpj, 'NULL'), ' ; ', IFNULL(nome, 'NULL'), ' ; ', IFNULL(endereco, 'NULL'), ' ; ', IFNULL(numero, 'NULL'), ' ; ', IFNULL(bairro, 'NULL'), ' ; ', IFNULL(cep, 'NULL'), ' ; ', IFNULL(nome_municipio, 'NULL'), ' ; ', IFNULL(uf, 'NULL'), ' ; ', IFNULL(observacao, 'NULL'))").Where("cnpj = ?", cnpj).Row().Scan(&log.Espelho)

	return &Log{}, err
}

func (log *Log) LogEntidade(db *gorm.DB, cnpj, nomeTabela, operacao string, codUsuario uint32) error {

	log.Data = time.Now()

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, data, nome_tabela, operacao, espelho, cod_cnpj) VALUES (?, ?, ?, ?, ?, ?)", codUsuario, log.Data, nomeTabela, operacao, log.Espelho, cnpj).Error

	return err
}
