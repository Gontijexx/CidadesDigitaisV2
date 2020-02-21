package config

/*  =========================
	Server HTTP
=========================  */

const (
	// defines ip and port address for server instance
	SERVER_ADDR = "localhost:8080"
)

/*  =========================
	USER PATHS HTTP
=========================  */

const (
	USER_PATH            = "/read/usuario"
	USER_ID_PATH         = "/read/usuario/{cod_usuario}"
	USER_PATH_LOGIN      = "/read/usuario/login"
	USER_PATH_CREATEUSER = "/read/usuario/createuser"
	USER_PATH_DELETEUSER = "/read/usuario/deleteuser"
)

/*  =========================
	MODULOS PATHS HTTP
=========================  */

const (
	MODULO_PATH          = "/read/modulo"
	MODULO_USERLIST_PATH = "/read/usuario/{cod_usuario}/modulo"
)

/*  =========================
	ENTIDADE PATHS HTTP OK!
=========================  */

const (
	ENTIDADE_PATH    = "/read/entidade"
	ENTIDADE_ID_PATH = "/read/entidade/{cnpj}"
)

/*  =========================
	CONTATO PATHS HTTP OK!
=========================  */

const (
	CONTATO_PATH    = "/read/contato"
	CONTATO_ID_PATH = "/read/contato/{cod_contato}"
)

/*  =========================
	TELEFONE PATHS HTTP OK!
=========================  */

const (
	TELEFONE_PATH    = "/read/telefone"
	TELEFONE_ID_PATH = "/read/telefone/{cod_telefone}"
)

/*  =========================
	CD PATHS HTTP OK!
=========================  */

const (
	CD_PATH    = "/read/cd"
	CD_ID_PATH = "/read/cd/{cod_ibge}"
)

/*  =========================
	CD ITENS PATHS HTTP OK!
=========================  */

const (
	CD_ITENS_PATH    = "/read/cditens"
	CD_ITENS_ID_PATH = "/read/cditens/{cod_ibge}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	CD PROCESSO PATHS HTTP
=========================  */

const (
	PROCESSO_PATH    = "/read/processo"
	PROCESSO_ID_PATH = "/read/processo/{cod_processo}/{cod_ibge}"
)

/*  =========================
	CD UACOM PATHS HTTP
=========================  */

const (
	UACOM_PATH    = "/read/uacom"
	UACOM_ID_PATH = "/read/uacom/{cod_ibge}/{data}"
)

/*  =========================
	CD UACOM_ASSUNTO PATHS HTTP
=========================  */

const (
	UACOM_ASSUNTO_PATH = "/read/uacomassunto"
)

/*  =========================
	CD PONTO PATHS HTTP
=========================  */

const (
	PONTO_PATH    = "/read/ponto"
	PONTO_ID_PATH = "/read/ponto/{cod_ponto}/{cod_categoria}/{cod_ibge}"
)

/*  =========================
	CD PID_TIPOOGIA PATHS HTTP
=========================  */

const (
	PID_TOPOLOGIA_ID_PATH = "/read/pontotipologia/{cod_ponto}/{cod_categoria}/{cod_ibge}/{cod_tipologia}"
)

/*  =========================
	LOTE PATHS HTTP OK!
=========================  */

const (
	LOTE_PATH    = "/read/lote"
	LOTE_ID_PATH = "/read/lote/{cod_lote}"
)

/*  =========================
	REAJUSTE PATHS HTTP OK!
=========================  */
const (
	REAJUSTE_PATH    = "/read/reajuste"
	REAJUSTE_ID_PATH = "/read/reajuste/{cod_lote}/{ano_ref}"
)

/*  =========================
	LOTE ITENS PATHS HTTP OK!
=========================  */

const (
	LOTE_ITENS_PATH    = "/read/loteitens"
	LOTE_ITENS_ID_PATH = "/read/loteitens/{cod_lote}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	EMPENHO PATHS HTTP
=========================  */

const (
	EMPENHO_PATH    = "/read/empenho"
	EMPENHO_ID_PATH = "/read/empenho/{id_empenho}"
)

/*  =========================
	ITENS EMPENHO PATHS HTTP
=========================  */

const (
	ITENS_EMPENHO_PATH    = "/read/itensempenho"
	ITENS_EMPENHO_ID_PATH = "/read/itensempenho/{id_empenho}/{cod_item}/{cod_tipo_item}"
)

/*	=========================
	OTB PATH HTTP (PAGAMENTO)
=========================	*/

const (
	OTB_PATH    = "/read/otb"
	OTB_ID_PATH = "/read/otb/{cod_otb}"
)

/*	=========================
	OTB FATURA PATH HTTP (PAGAMENTO FATURA)
=========================	*/
const (
	OTB_FATURA_PATH    = "/read/otbfatura"
	OTB_FATURA_ID_PATH = "/read/otbfatura/{cod_otb}"
)

/*	=========================
	OTB ITENS PATH HTTP (PAGAMENTO ITENS)
=========================	*/

const (
	OTB_ITENS_PATH    = "/read/otbitens"
	OTB_ITENS_ID_PATH = "/read/otbitens/{cod_otb}"
)

/*	=========================
	FATURA PATH HTTP
=========================	*/

const (
	FATURA_PATH    = "/read/fatura"
	FATURA_ID_PATH = "/read/fatura/{num_nf}"
)

/*	=========================
	FATURA OTB PATH HTTP (FATURA PAGAMENTO)
=========================	*/
const (
	FATURA_OTB_ID_PATH = "/read/faturaotb/{num_nf}/{cod_ibge}"
)

/*	=========================
	FATURA ITENS PATH HTTP (ITENS FATURA)
=========================	*/

const (
	FATURA_ITENS_PATH        = "/read/faturaitens"
	FATURA_ITENS_ID_PATH     = "/read/faturaitens/{num_nf}"
	FATURA_ITENS_DELETE_PATH = "/read/faturaitens/{num_nf}/{id_empenho}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	PREVISAO EMPENHO PATHS HTTP OK!
=========================  */

const (
	PREVISAO_EMPENHO_PATH    = "/read/previsaoempenho"
	PREVISAO_EMPENHO_ID_PATH = "/read/previsaoempenho/{cod_previsao_empenho}"
)

/*  =========================
	ITENS PREVISAO EMPENHO PATHS HTTP
=========================  */

const (
	ITENS_PREVISAO_EMPENHO_PATH    = "/read/itensprevisaoempenho"
	ITENS_PREVISAO_EMPENHO_ID_PATH = "/read/itensprevisaoempenho/{cod_previsao_empenho}/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	ASSUNTO PATHS HTTP OK!
=========================  */

const (
	ASSUNTO_PATH    = "/read/assunto"
	ASSUNTO_ID_PATH = "/read/assunto/{cod_assunto}"
)

/*	=========================
	CATEGORIA PATHS HTTP OK!
=========================	*/
const (
	CATEGORIA_PATH    = "/read/categoria"
	CATEGORIA_ID_PATH = "/read/categoria/{cod_categoria}"
)

/*	=========================
	CLASSE EMPENHO PATHS HTTP OK!
=========================	*/

const (
	CLASSE_EMPENHO_PATH    = "/read/classeempenho"
	CLASSE_EMPENHO_ID_PATH = "/read/classeempenho/{cod_classe_empenho}"
)

/*  =========================
	ETAPA PATHS HTTP
=========================  */

const (
	ETAPA_PATH    = "/read/etapa"
	ETAPA_ID_PATH = "/read/etapa/{cod_etapa}"
)

/*	=========================
	ITENS PATHS HTTP OK!
=========================	*/

const (
	ITENS_PATH    = "/read/itens"
	ITENS_ID_PATH = "/read/itens/{cod_item}/{cod_tipo_item}"
)

/*  =========================
	MUNICIPIO PATHS HTTP OK!
=========================  */

const (
	MUNICIPIO_PATH    = "/read/municipio"
	MUNICIPIO_ID_PATH = "/read/municipio/{cod_ibge}"
)

/*	=========================
	NATUREZA DE DESPESA PATHS HTTP OK!
=========================	*/

const (
	NATUREZA_DESPESA_PATH    = "/read/naturezadespesa"
	NATUREZA_DESPESA_ID_PATH = "/read/naturezadespesa/{cod_natureza_despesa}"
)

/*  =========================
	PREFEITOS PATHS HTTP OK!
=========================  */

const (
	PREFEITOS_PATH    = "/read/prefeitos"
	PREFEITOS_ID_PATH = "/read/prefeitos/{cod_prefeito}"
)

/*	=========================
	TIPOLOGIAS PATHS HTTP OK!
=========================	*/

const (
	TIPOLOGIA_PATH    = "/read/tipologia"
	TIPOLOGIA_ID_PATH = "/read/tipologia/{cod_tipologia}"
)

/*	=========================
	TIPO ITEM PATHS HTTP OK!
=========================	*/

const (
	TIPO_ITEM_PATH    = "/read/tipoitem"
	TIPO_ITEM_ID_PATH = "/read/tipoitem/{cod_tipo_item}"
)
