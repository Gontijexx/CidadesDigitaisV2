package models

import (
	"github.com/jinzhu/gorm"
)

/*  =========================
	STRUCT LOG
=========================  */

type Log struct {
	CodLog      uint32 `gorm:"primary_key;auto_increment;not null" json:"cod_log"`
	CodUsuario  uint32 `gorm:"primary_key;foreign_key:CodUsuario;not null" json:"cod_usuario"`
	Data        string `gorm:"not null" json:"data"`
	NomeTabela  string `gorm:"not null" json:"nome_tabela"`
	Operacao    string `gorm:"not null" json:"operacao"`
	Espelho     string `gorm:"default:null" json:"espelho"`
	CodInt1     uint32 `gorm:"default:null" json:"cod_int_1"`
	CodInt2     uint32 `gorm:"default:null" json:"cod_int_2"`
	CodInt3     uint32 `gorm:"default:null" json:"cod_int_3"`
	CodInt4     uint32 `gorm:"default:null" json:"cod_int_4"`
	CodInt5     uint32 `gorm:"default:null" json:"cod_int_5"`
	CodInt6     uint32 `gorm:"default:null" json:"cod_int_6"`
	CodData     string `gorm:"default:null" json:"cod_data"`
	CodProcesso string `gorm:"default:null;size:17" json:"cod_processo"`
	CodCnpj     string `gorm:"default:null;size:14" json:"cod_cnpj"`
	CodEmpenho  string `gorm:"default:null;size:13" json:"cod_empenho"`
}

/*  =========================
	LOG ASSUNTO
=========================  */

func (log *Log) LogAssunto(db *gorm.DB, codAssunto uint32, nomeTabela, operacao string, codUsuario uint32) error {

	db.Table("assunto").Select("CONCAT(IFNULL(cod_assunto, ''), ';', IFNULL(descricao, ''))").Where("cod_assunto = ?", codAssunto).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codAssunto).Error

	return err
}

/*  =========================
	LOG CATEGORIA
=========================  */

func (log *Log) LogCategoria(db *gorm.DB, codCategoria uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("categoria").Select("CONCAT(IFNULL(cod_categoria, ''), ';', IFNULL(descricao, ''))").Where("cod_categoria = ?", codCategoria).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codCategoria).Error

	return err
}

/*  =========================
	LOG CD
=========================  */

func (log *Log) LogCD(db *gorm.DB, codIbge uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("cd").Select("CONCAT(IFNULL(cod_ibge, ''), ';', IFNULL(cod_lote, ''), ';', IFNULL(os_pe, ''), ';', IFNULL(data_pe, ''), ';', IFNULL(os_imp, ''), ';', IFNULL(data_imp, ''))").Where("cod_ibge = ?", codIbge).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codIbge).Error

	return err
}

/*  =========================
	LOG CD ITENS
=========================  */

func (log *Log) LogCDItens(db *gorm.DB, codIbge uint32, codItem uint32, codTipoItem uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("cd_itens").Select("CONCAT(IFNULL(cod_ibge, ''), ';', IFNULL(cod_item, ''), ';', IFNULL(cod_tipo_item, ''), ';', IFNULL(quantidade_previsto, ''), ';', IFNULL(quantidade_projeto_executivo, ''), ';', IFNULL(quantidade_termo_instalacao, ''))").Where("cod_ibge = ? AND cod_item = ? AND cod_tipo_item = ?", codIbge, codItem, codTipoItem).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1, cod_int_2, cod_int_3) VALUES (?, ?, ?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codIbge, codItem, codTipoItem).Error

	return err
}

/*  =========================
	LOG CLASSE EMPENHO
=========================  */

func (log *Log) LogClasseEmpenho(db *gorm.DB, codClasseEmpenho uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("classe_empenho").Select("CONCAT(IFNULL(cod_classe_empenho, ''), ';', IFNULL(descricao, ''))").Where("cod_classe_empenho = ?", codClasseEmpenho).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codClasseEmpenho).Error

	return err
}

/*  =========================
	LOG CONTATO
=========================  */

func (log *Log) LogContato(db *gorm.DB, codContato uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("contato").Select("CONCAT(IFNULL(cod_contato, ''), ';', IFNULL(cnpj, ''), ';', IFNULL(cod_ibge, ''), ';', IFNULL(nome, ''), ';', IFNULL(email, ''), ';', IFNULL(funcao, ''))").Where("cod_contato = ?", codContato).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log (cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codContato).Error

	return err
}

/*  =========================
	LOG EMPENHO
=========================  */

func (log *Log) LogEmpenho(db *gorm.DB, idEmpenho uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("empenho").Select("CONCAT(IFNULL(id_empenho, ''), ';', IFNULL(cod_previsao_empenho, ''), ';', IFNULL(cod_empenho, ''), ';', IFNULL(data, ''), ';', IFNULL(contador, ''))").Where("id_empenho = ?", idEmpenho).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, idEmpenho).Error

	return err
}

/*  =========================
	LOG ENTIDADE
=========================  */

func (log *Log) LogEntidade(db *gorm.DB, cnpj string, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("entidade").Select("CONCAT(IFNULL(cnpj, ''), ';', IFNULL(nome, ''), ';', IFNULL(endereco, ''), ';', IFNULL(numero, ''), ';', IFNULL(bairro, ''), ';', IFNULL(cep, ''), ';', IFNULL(nome_municipio, ''), ';', IFNULL(uf, ''), ';', IFNULL(observacao, ''))").Where("cnpj = ?", cnpj).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_cnpj) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, cnpj).Error

	return err
}

/*  =========================
	LOG ETAPA
=========================  */

func (log *Log) LogEtapa(db *gorm.DB, codEtapa uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("etapa").Select("CONCAT(IFNULL(cod_etapa, ''), ';', IFNULL(descricao, ''), ';', IFNULL(duracao, ''), ';', IFNULL(depende, ''), ';', IFNULL(delay, ''), ';', IFNULL(setor_resp, ''))").Where("cod_etapa = ?", codEtapa).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1) VALUES (?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codEtapa).Error

	return err
}

/*  =========================
	LOG ETAPAS CD
=========================  */

func (log *Log) LogEtapasCD(db *gorm.DB, codIbge uint32, codEtapa uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("etapas_cd").Select("CONCAT(IFNULL(cod_ibge, ''), ';', IFNULL(cod_etapa, ''), ';', IFNULL(dt_inicio, ''), ';', IFNULL(dt_fim, ''), ';', IFNULL(responsavel, ''))").Where("cod_ibge = ? AND cod_etapa = ?", codIbge, codEtapa).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1, cod_int_2) VALUES (?, ?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codIbge, codEtapa).Error

	return err
}

/*  =========================
	LOG FATURA
=========================  */

func (log *Log) LogFatura(db *gorm.DB, numNF uint32, codIbge uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("fatura").Select("CONCAT(IFNULL(num_nf, ''), ';', IFNULL(cod_ibge, ''), ';', IFNULL(dt_nf, ''))").Where("num_nf = ? AND cod_ibge = ?", numNF, codIbge).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1, cod_int_2) VALUES (?, ?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, numNF, codIbge).Error

	return err
}

/*  =========================
	LOG FATURA OTB
=========================  */

func (log *Log) LogFaturaOTB(db *gorm.DB, codOtb uint32, numNF uint32, codIbge uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("fatura_otb").Select("CONCAT(IFNULL(cod_otb, ''), ';', IFNULL(num_nf, ''), ';', IFNULL(cod_ibge, ''))").Where("cod_otb = ? AND num_nf = ? AND cod_ibge =?", codOtb, numNF, codIbge).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1, cod_int_2, cod_int_3) VALUES (?, ?, ?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codOtb, numNF, codIbge).Error

	return err
}

/*  =========================
	LOG ITENS
=========================  */

func (log *Log) LogItens(db *gorm.DB, codItem uint32, codTipoItem uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("itens").Select("CONCAT(IFNULL(cod_item, ''), ';', IFNULL(cod_tipo_item, ''), ';', IFNULL(cod_natureza_despesa, ''), ';', IFNULL(cod_classe_empenho, ''), ';', IFNULL(descricao, ''), ';', IFNULL(unidade, ''))").Where("cod_item = ? AND cod_tipo_item = ?", codItem, codTipoItem).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1, cod_int_2) VALUES (?, ?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, codItem, codTipoItem).Error

	return err
}

/*  =========================
	LOG ITENS EMPENHO
=========================  */

func (log *Log) LogItensEmpenho(db *gorm.DB, idEmpenho uint32, codItem uint32, codTipoItem uint32, nomeTabela string, operacao string, codUsuario uint32) error {

	db.Table("itens_empenho").Select("CONCAT(IFNULL(id_empenho, ''), ';', IFNULL(cod_item, ''), ';', IFNULL(cod_tipo_item, ''), ';', IFNULL(cod_previsao_empenho, ''), ';', IFNULL(valor, ''), ';', IFNULL(quantidade, ''))").Where("id_empenho = ? AND cod_item = ? AND cod_tipo_item = ?", idEmpenho, codItem, codTipoItem).Row().Scan(&log.Espelho)

	err := db.Debug().Exec("INSERT INTO log(cod_usuario, nome_tabela, operacao, espelho, cod_int_1, cod_int_2, cod_int_3) VALUES (?, ?, ?, ?, ?, ?, ?)", codUsuario, nomeTabela, operacao, log.Espelho, idEmpenho, codItem, codTipoItem).Error

	return err
}
