package models

/*  =========================
	TABELA ENTIDADE
=========================  */

type Entidade struct {
	Cnpj          string `gorm:"primary_key;not null;size:14" json:"cnpj"`
	Nome          string `gorm:"size:50;default:null" json:"nome"`
	Endereco      string `gorm:"size:100;default:null" json:"endereco"`
	Numero        string `gorm:"size:10;default:null" json:"numero"`
	Bairro        string `gorm:"size:100;default:null" json:"bairro"`
	Cep           string `gorm:"size:8;default:null" json:"cep"`
	NomeMunicipio string `gorm:"size:50;default:null" json:"nome_municipio"`
	UF            string `gorm:"size:2;default:null" json:"uf"`
	Observacao    string `gorm:"size:1000;default:null" json:"observacao"`
}

/*  =========================
	TABELA CONTATO
=========================  */

type Contato struct {
	CodContato uint64 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_contato"`
	Cnpj       string `gorm:"foreign_key:Cnpj;default:null;size:14" json:"cnpj"`
	CodIbge    uint64 `gorm:"foreign_key:CodIbge;default:null;size:7" json:"cod_ibge"`
	Nome       string `gorm:"default:null;size:50" json:"nome"`
	Email      string `gorm:"default:null;size:100" json:"email"`
	Funcao     string `gorm:"default:null;size:45" json:"funcao"`
}

/*  =========================
	TABELA TELEFONE
=========================  */

type Telefone struct {
	CodTelefone uint64 `gorm:"primary_key;auto_increment;not null;size:11" json:"cod_telefone"`
	CodContato  uint64 `gorm:"foreign_key:CodContato;not null;size:11" json:"cod_contato"`
	Telefone    string `gorm:"default:null;size:11" json:"telefone"`
	Tipo        string `gorm:"default:null;size:10" json:"tipo"`
}

/* =========================
	TABELA CD
=========================  */

type CD struct {
	CodIbge       uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	CodLote       uint64 `gorm:"foreign_key:CodLote;not null" json:"cod_lote"`
	NomeMunicipio string `gorm:"default:null" json:"nome_municipio"`
	OsPe          string `gorm:"size:10;default:null" json:"os_pe"`
	DataPe        string `gorm:"default:null" json:"data_pe"`
	OsImp         string `gorm:"size:10;default:null" json:"os_imp"`
	DataImp       string `gorm:"default:null" json:"data_imp"`
}

/*  =========================
	TABELA CD ITENS
=========================  */

type CDItens struct {
	CodIbge                    uint64  `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	CodItem                    uint64  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem                uint64  `gorm:"primary_key;foreign_key:CodTipo_item;not null" json:"cod_tipo_item"`
	QuantidadePrevisto         uint64  `gorm:"default:null" json:"quantidade_previsto"`
	QuantidadeProjetoExecutivo float64 `gorm:"default:null" json:"quantidade_projeto_executivo"`
	QuantidadeTermoInstalacao  float64 `gorm:"default:null" json:"quantidade_termo_instalacao"`
}

/*  =========================
	TABELA CD PROCESSO
=========================  */

type Processo struct {
	CodProcesso string `gorm:"primary_key;not null;size:17" json:"cod_processo"`
	CodIbge     uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Descricao   string `gorm:"default:null" json:"descricao"`
}

/*  =========================
	TABELA UACOM (CD)
=========================  */

type Uacom struct {
	CodIbge uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Data    string `gorm:"primary_key;not null" json:"data"`
	Titulo  string `gorm:"default:null" json:"titulo"`
	Relato  string `gorm:"default:null" json:"relato"`
}

/*  =========================
	TABELA UACOM ASSUNTO (CD)
=========================  */

type UacomAssunto struct {
	CodIbge    uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Data       string `gorm:"primary_key;foreign_key:Data;not null" json:"data"`
	CodAssunto uint64 `gorm:"primary_key;foreign_key:CodAssunto;not null" json:"cod_assunto"`
}

/*  =========================
	TABELA PONTO (CD)
=========================  */

type Ponto struct {
	CodPonto     uint64 `gorm:"primary_key;not null" json:"cod_ponto"`
	CodCategoria uint64 `gorm:"primary_key;foreign_key:CodCategoria;not null" json:"cod_categoria"`
	CodIbge      uint64 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	CodPID       uint64 `gorm:"foreign_key:CodPID;not null" json:"cod_pid"`
	Endereco     string `gorm:"default:null" json:"endereco"`
	Numero       string `gorm:"default:null;size:10" json:"numero"`
	Complemento  string `gorm:"default:null" json:"complemento"`
	Bairro       string `gorm:"default:null" json:"bairro"`
	Cep          string `gorm:"default:null;size:8" json:"cep"`
	Latitude     uint64 `gorm:"default:null" json:"latitude"`
	Longitude    uint64 `gorm:"default:null" json:"longitude"`
}

/*   =========================
	TABELA PID (CD)
=========================  */

type PID struct {
	CodPID  uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_pid"`
	CodIbge uint64 `gorm:"foreign_key:CodIbge;not null" json:"cod_ibge"`
	Nome    string `gorm:"default:null" json:"nome"`
	Inep    string `gorm:"default:null" json:"inep"`
}

/*  =========================
	TABELA PID TIPOLOGIA (CD)
=========================  */

type PIDTipologia struct {
	CodPID       uint64 `gorm:"primary_key;foreign_key:CodPID;not null" json:"cod_pid"`
	CodTipologia uint64 `gorm:"primary_key;foreign_key:CodTipologia;not null" json:"cod_tipologia"`
}

/*  =========================
	TABELA LOTE
=========================  */

type Lote struct {
	CodLote     uint64 `gorm:"primary_key;not null;size:11" json:"cod_lote"`
	Cnpj        string `gorm:"foreign_key:Cnpj;not null;size:14" json:"cnpj"`
	Contrato    string `gorm:"default:null;size:10" json:"contrato"`
	DtInicioVig string `gorm:"default:null;size:10" json:"dt_inicio_vig" `
	DtFinalVig  string `gorm:"default:null;size:10" json:"dt_final_vig" `
	DtReajuste  string `gorm:"default:null;size:10" json:"dt_reajuste" `
	Nome        string `gorm:"default:null" json:"nome"`
}

/*  =========================
	TABELA REAJUSTE
=========================  */

type Reajuste struct {
	AnoRef     uint64  `gorm:"primary_key;not null;size:11" json:"ano_ref"`
	CodLote    uint64  `gorm:"primary key;foreign_key:CodLote;not null;size:11" json:"cod_lote"`
	Percentual float64 `gorm:"default:null" json:"percentual"`
}

/*  =========================
	TABELA LOTE ITENS
========================= */

type LoteItens struct {
	CodLote     uint64  `gorm:"primary_key;foreign_key:CodLote;not null;size:11" json:"cod_lote"`
	CodItem     uint64  `gorm:"primary_key;foreign_key:CodItem;not null;size:11" json:"cod_item"`
	CodTipoItem uint64  `gorm:"primary_key;foreign_key:CodTipoItem;not null;size:11" json:"cod_tipo_item"`
	Preco       float64 `gorm:"default:null;size:12" json:"preco"`
	Descricao   string `gorm:"default:null" json:"descricao"`
}

/*  =========================
	TABELA EMPENHO
=========================  */

type Empenho struct {
	IDEmpenho          uint64 `gorm:"primary_key;auto_increment;default:not null" json:""id_empenho`
	CodPrevisaoEmpenho uint64 `gorm:"foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho"`
	CodEmpenho         string `gorm:"default:null;size:13" json:"cod_empenho"`
	Data               string `gorm:"default:null" json:"data"`
	Contador           uint64 `gorm:"default:null" json:"contador"`
}

/*  =========================
	TABELA ITENS EMPENHO
=========================  */

type ItensEmpenho struct {
	IDEmpenho          uint64  `gorm:"primary_key;foreign_key:IDEmpenho;not null;size:13" json:"id_empenho"`
	CodItem            uint64  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem        uint64  `gorm:"primary_key;foreign_key:CodTipoItem;not null" json:"cod_tipo_item"`
	CodPrevisaoEmpenho uint64  `gorm:"foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho"`
	Valor              float64 `gorm:"default:null" json:"valor"`
	Quantidade         uint64  `gorm:"default:null" json:"quantidade"`
}

/*	=========================
		TABELA OTB (PAGAMENTO)
=========================	*/

type OTB struct {
	CodOtb uint64 `gorm:"primary_key;not null" json:"cod_otb"`
	DtPgto string `gorm:"default:null" json:"dt_pgto" `
}

/*	=========================
		TABELA FATURA_OTB (PAGAMENTO FATURA/FATURA PAGAMENTO)
=========================	*/

type FaturaOTB struct {
	CodOtb  uint64 `gorm:"primary_key;foreign_key:CodOtb;not null" json:"cod_otb"`
	NumNF   uint64 `gorm:"primary_key;foreign_key:NumNF;not null" json:"num_nf"`
	CodIbge uint64 `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
}

/*	=========================
		TABELA ITENS_OTB (PAGAMENTO)
=========================	*/

type ItensOTB struct {
	CodOtb      uint64  `gorm:"primary_key;foreign_key:CodOtb;not null" json:"cod_otb"`
	NumNF       uint64  `gorm:"primary_key;foreign_key:NumNF;not null" json:"num_nf"`
	CodIbge     uint64  `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	IDEmpenho   uint64  `gorm:"primary_key;foreign_key:IDEmpenho;not null" json:"id_empenho"`
	CodItem     uint64  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem uint64  `gorm:"primary_key;foreign_key:CodTipoIte,;not null" json:"cod_tipo_item"`
	Valor       float64 `gorm:"default:null" json:"valor"`
	Quantidade  uint64  `gorm:"default:null" json:"quantidade"`
}

/*	=========================
		TABELA FATURA
=========================	*/

type Fatura struct {
	NumNF   uint64 `gorm:"primary_key;not null" json:"num_nf"`
	CodIbge uint64 `gorm:"primary_key;foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	DtNf    string `gorm:"default:null" json:"dt_nf"`
}

/*	=========================
		TABELA ITENS FATURA
=========================	*/

type ItensFatura struct {
	NumNF       uint64  `gorm:"primary_key;foreign_key:NumNF;not null" json:"num_nf"`
	CodIbge     uint64  `gorm:"primary_key;foreign_key:CodIbge;not null" json:"cod_ibge"`
	IDEmpenho   uint64  `gorm:"primary_key;foreign_key:IDEmpenho;not null" json:"id_empenho"`
	CodItem     uint64  `gorm:"primary_key;foreign_key:CodItem;not null" json:"cod_item"`
	CodTipoItem uint64  `gorm:"primary_key;foreign_key:CodTipoIte,;not null" json:"cod_tipo_item"`
	Valor       float64 `gorm:"default:null" json:"valor"`
	Quantidade  uint64  `gorm:"default:null" json:"quantidade"`
}

/*  =========================
	TABELA PREVISAO EMPENHO
=========================  */

type PrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint64 `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;auto_increment;not null" json:"cod_previsao_empenho"`
	CodLote            uint64 `gorm:"foreign_key:CodLote;not null;size:11" json:"cod_lote"`
	CodNaturezaDespesa uint64 `gorm:"foreign_key:CodNaturezaDespesa;not null;size:11" json:"cod_natureza_despesa"`
	Data               string `gorm:"default:null" json:"data"`
	Tipo               string `gorm:"default:null;size:1" json:"tipo"`
	Ano_referencia     uint64 `gorm:"default:null;size:11" json:"ano_referencia"`
}

/*  =========================
	TABELA ITENS PREVISAO EMPENHO
=========================  */

type ItensPrevisaoEmpenho struct {
	CodPrevisaoEmpenho uint64  `gorm:"primary_key;foreign_key:CodPrevisaoEmpenho;not null" json:"cod_previsao_empenho"`
	CodItem            uint64  `gorm:"primary_key;foreign_key:CodItem;not null;size:11" json:"cod_item"`
	CodTipoItem        uint64  `gorm:"primary_key;foreign_key:CodTipo_item;not null;size:11" json:"cod_tipo_item"`
	CodLote            uint64  `gorm:"foreign_key:CodLote;not null;size:11" json:"cod_lote"`
	Valor              float64 `gorm:"default:null;size:12" json:"valor"`
	Quantidade         uint64  `gorm:"default:null;size:11" json:"quantidade"`
}

/*	=========================
		TABELA ASSUNTO
=========================	*/

type Assunto struct {
	CodAssunto uint64 `gorm:"primary_key;auto_increment;not null;" json:"cod_assunto"`
	Descricao  string `gorm:"default:null" json:"descricao"`
}

/*	=========================
		TABELA CATEGORIA
=========================	*/

type Categoria struct {
	CodCategoria uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_categoria"`
	Descricao    string `gorm:"default:null" json:"descricao"`
}

/*	=========================
		TABELA CLASSE EMPENHO
=========================	*/

type ClasseEmpenho struct {
	CodClasseEmpenho uint64 `gorm:"primary_key;not null" json:"cod_classe_empenho"`
	Descricao        string `gorm:"default:null" json:"descricao"`
}

/*	=========================
		TABELA ETAPA
=========================	*/

type Etapa struct {
	CodEtapa  uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_etapa"`
	Descricao string `gorm:"default:null" json:"descricao"`
	Duracao   uint64 `gorm:"default:null" json:"duracao"`
	Depende   uint64 `gorm:"default:null" json:"depende"`
	Delay     uint64 `gorm:"default:null" json:"delay"`
	SetorResp string `gorm:"default:null" json:"setor_resp"`
}

/*	=========================
		TABELA ITENS
=========================	*/

type Itens struct {
	CodItem            uint64 `gorm:"primary_key;not null" json:"cod_item"`
	CodTipoItem        uint64 `gorm:"primary_key;foreign_key:CodTipoItem;not null" json:"cod_tipo_item"`
	CodNaturezaDespesa uint64 `gorm:"foreign_key:CodNaturezaDespesa;not null" json:"cod_natureza_despesa"`
	CodClasseEmpenho   uint64 `gorm:"foreign_key:CodClasseEmpenho;not null" json:"cod_classe_empenho"`
	Descricao          string `gorm:"default:null" json:"descricao"`
	Unidade            string `gorm:"default:null" json:"unidade"`
}

/*  =========================
	TABELA MUNICIPIO
=========================  */

type Municipio struct {
	CodIbge       uint64  `gorm:"primary_key;not null;size:7" json:"cod_ibge"`
	NomeMunicipio string  `gorm:"default:null" json:"nome_municipio"`
	Populacao     uint64  `gorm:"default:null" json:"populacao"`
	UF            string  `gorm:"default:null;size:2" json:"uf"`
	Regiao        string  `gorm:"default:null;size:15" json:"regiao"`
	Cnpj          string  `gorm:"default:null;size:14" json:"cnpj"`
	DistCapital   uint64  `gorm:"default:null" json:"dist_capital"`
	Endereco      string  `gorm:"default:null" json:"endereco"`
	Numero        string  `gorm:"default:null;size:10" json:"numero"`
	Complemento   string  `gorm:"default:null" json:"complemento"`
	Bairro        string  `gorm:"default:null" json:"bairro"`
	Idhm          float64 `gorm:"default:null" json:"idhm"`
	Latitude      float64 `gorm:"default:null" json:"latitude"`
	Longitude     float64 `gorm:"default:null" json:"longitude"`
}

/*	=========================
	TABELA NATUREZA DE DESPESA
=========================	*/

type NaturezaDespesa struct {
	CodNaturezaDespesa uint64 `gorm:"primary_key;not null" json:"cod_natureza_despesa"`
	Descricao          string `gorm:"default:null" json:"descricao"`
}

/*  =========================
	TABELA PREFEITOS
=========================  */

type Prefeitos struct {
	CodPrefeito uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_prefeito"`
	CodIbge     uint64 `gorm:"foreign_key:CodIbge;not null;size:7" json:"cod_ibge"`
	Nome        string `gorm:"default:null" json:"nome"`
	Cpf         string `gorm:"default:null" json:"cpf"`
	RG          string `gorm:"default:null" json:"rg"`
	Partido     string `gorm:"default:null" json:"partido"`
	Exercicio   string `gorm:"default:null" json:"exercicio"`
}

/*	=========================
		TABELA TIPOLOGIAS
=========================	*/

type Tipologia struct {
	CodTipologia uint64 `gorm:"primary_key;auto_increment;not null" json:"cod_tipologia"`
	Descricao    string `gorm:"default:null" json:"descricao"`
}

/*	=========================
		TABELA TIPO ITEM
=========================	*/

type TipoItem struct {
	CodTipoItem uint64 `gorm:"primary_key;not null" json:"cod_tipo_item"`
	Descricao   string `gorm:"default:null" json:"descricao"`
}
