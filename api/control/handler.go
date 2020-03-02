package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) CreateHandler() (r *mux.Router) {

	//	CRIA UM ROTEADOR

	r = s.Router

	//	HOME
	r.HandleFunc("/{id}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.Home))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM USUARIO
	=========================	*/

	//LISTA USUARIOS
	r.HandleFunc(config.USER_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllUsuario))).Methods(http.MethodGet)

	//EDITA O USUARIO {cod_usuario}
	r.HandleFunc(config.USER_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateUsuario))).Methods(http.MethodPut)

	//ROTA DE LOGIN
	r.HandleFunc(config.USER_PATH_LOGIN, middlewares.SetMiddleJSON(s.Login)).Methods(http.MethodPost)

	//SALVA USUARIO
	r.HandleFunc(config.USER_PATH_CREATEUSER, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateUsuario))).Methods(http.MethodPost)

	//APAGA O USUARIO
	//	r.HandleFunc(config.USER_PATH_DELETEUSER, middlewares.SetMiddleAuth(s.DeleteUsuario)).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM MODULOS
	=========================	*/

	//	LISTA MODULOS
	r.HandleFunc(config.MODULO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllModulo))).Methods(http.MethodGet)

	//	ADICIONAR MODULO
	r.HandleFunc(config.MODULO_USERLIST_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.AddModulo))).Methods(http.MethodPost)

	/*	=========================
		ROTAS EM ENTIDADE OK!
	=========================	*/

	//	LISTA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllEntidade))).Methods(http.MethodGet)

	//	SALVA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateEntidade))).Methods(http.MethodPost)

	//	EDITA ENTIDADE
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEntidade))).Methods(http.MethodPut)

	//	LISTA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEntidadeByID))).Methods(http.MethodGet)

	//	APAGA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEntidade))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS DE CONTATO OK!
	=========================	*/

	//	LISTA CONTATO
	r.HandleFunc(config.CONTATO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllContato))).Methods(http.MethodGet)

	//	SALVA CONTATO
	r.HandleFunc(config.CONTATO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateContato))).Methods(http.MethodPost)

	//	EDITA CONTATO (cod_contato)
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateContato))).Methods(http.MethodPut)

	//	APAGA LOTE (cod_contato)
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteContato))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS DE TELEFONE OK!
	=========================	*/

	//	LISTA TELEFONE
	r.HandleFunc(config.TELEFONE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllTelefone))).Methods(http.MethodGet)

	//	SALVA TELEFONE
	r.HandleFunc(config.TELEFONE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateTelefone))).Methods(http.MethodPost)

	//	APAGA TELEFONE (cod_telefone)
	r.HandleFunc(config.TELEFONE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteTelefone))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM CIDADE DIGITAL OK!
	=========================	*/

	//	LISTA CD
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCD))).Methods(http.MethodGet)

	//	SALVA CD
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateCD))).Methods(http.MethodPost)

	//	EDITA CD (cod_ibge)
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCD))).Methods(http.MethodPut)

	//	LISTA CD (cod_ibge)
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDByID))).Methods(http.MethodGet)

	//	APAGA CD (cod_ibge)
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteCD))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM CIDADE DIGITAL_ITENS OK!
	=========================	*/

	//	LISTA CD_ITENS
	r.HandleFunc(config.CD_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCDItens))).Methods(http.MethodGet)

	//	EDITA CD_ITENS (cod_ibge, cod_item, cod_tipo_item)
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCDItens))).Methods(http.MethodPut)

	//	LISTA CD_ITENS POR ID (cod_ibge, cod_item, cod_tipo_item)
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDItensByID))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM CIDADE DIGITAL PROCESSO OK!
	=========================	*/

	//	LISTA PROCESSO
	r.HandleFunc(config.PROCESSO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllProcesso))).Methods(http.MethodGet)

	//	SALVA PROCESSO
	r.HandleFunc(config.PROCESSO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateProcesso))).Methods(http.MethodPost)

	//	EDITA PROCESSO (cod_processo)
	r.HandleFunc(config.PROCESSO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateProcesso))).Methods(http.MethodPut)

	//	LISTA PROCESSO (cod_processo)
	r.HandleFunc(config.PROCESSO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetProcessoByID))).Methods(http.MethodGet)

	//	APAGA PROCESSO (cod_processo)
	r.HandleFunc(config.PROCESSO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteProcesso))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM CIDADE DIGITAL UACOM OK!
	=========================	*/

	//	LISTA UACOM
	r.HandleFunc(config.UACOM_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllUacom))).Methods(http.MethodGet)

	//	SALVA UACOM
	r.HandleFunc(config.UACOM_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateUacom))).Methods(http.MethodPost)

	//	LISTA UACOM (cod_ibge, data)
	r.HandleFunc(config.UACOM_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUacomByID))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM CIDADE DIGITAL UACOM_ASSUNTO OK!
	=========================	*/

	//	SALVA UACOM_ASSUNTO
	r.HandleFunc(config.UACOM_ASSUNTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateUacomAssunto))).Methods(http.MethodPost)

	/*	=========================
		ROTAS EM CIDADE DIGITAL PONTO OK!
	=========================	*/

	//	LISTA PONTO
	r.HandleFunc(config.PONTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPonto))).Methods(http.MethodGet)

	//	SALVA PONTO
	r.HandleFunc(config.PONTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePonto))).Methods(http.MethodPost)

	//	EDITA PONTO (cod_ponto, cod_categoria, cod_ibge)
	r.HandleFunc(config.PONTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePonto))).Methods(http.MethodPut)

	//	LISTA PONTO (cod_ponto, cod_categoria, cod_ibge)
	r.HandleFunc(config.PONTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPontoByID))).Methods(http.MethodGet)

	//	APAGA PONTO (cod_ponto, cod_categoria, cod_ibge)
	r.HandleFunc(config.PONTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePonto))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM CIDADE DIGITAL PID_TIPOLOGIA OK!
	=========================	*/

	//	APAGA PID_TIPOLOGIA (cod_ponto, cod_categoria, cod_ibge, cod_tipologia)
	r.HandleFunc(config.PONTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePIDTipologia))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM LOTE OK!
	=========================	*/

	//	LISTA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllLote))).Methods(http.MethodGet)

	//	SALVA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateLote))).Methods(http.MethodPost)

	//	EDITA LOTE
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLote))).Methods(http.MethodPut)

	//	LISTA LOTE (cod_lote)
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteByID))).Methods(http.MethodGet)

	//	APAGA LOTE (cod_lote)
	//	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteLote))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM REAJUSTE OK!
	=========================	*/

	//	SALVA REAJUSTE
	r.HandleFunc(config.REAJUSTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateReajuste))).Methods(http.MethodPost)

	//	LISTA REAJUSTE
	r.HandleFunc(config.REAJUSTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllReajuste))).Methods(http.MethodGet)

	//	APAGA REAJUSTE (lote_cod_lote, ano_ref)
	r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteReajuste))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM LOTE ITENS OK!
	=========================	*/

	//	LISTA LOTE
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllLoteItens))).Methods(http.MethodGet)

	//	EDITA LOTE ITENS
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLoteItens))).Methods(http.MethodPut)

	//	LISTA LOTE ITENS POR ID
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteItensByID))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM EMPENHO OK!
	=========================	*/

	//	LISTA EMPENHO
	r.HandleFunc(config.EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllEmpenho))).Methods(http.MethodGet)

	//	SALVA EMPENHO
	r.HandleFunc(config.EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateEmpenho))).Methods(http.MethodPost)

	//	EDITA EMPENHO (cod_empenho)
	r.HandleFunc(config.EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEmpenho))).Methods(http.MethodPut)

	//	LISTA EMPENHO (cod_empenho)
	r.HandleFunc(config.EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEmpenhoByID))).Methods(http.MethodGet)

	//	APAGA EMPENHO (cod_empenho)
	r.HandleFunc(config.EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEmpenho))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM ITENS EMPENHO OK!
	=========================	*/

	//	EDITA ITENS EMPENHO (cod_empenho)
	r.HandleFunc(config.ITENS_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS EMPENHO (cod_empenho)
	r.HandleFunc(config.ITENS_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensEmpenhoByID))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM OTB
	=========================	*/

	//	LISTA OTB
	r.HandleFunc(config.OTB_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllOTB))).Methods(http.MethodGet)

	//	SALVA OTB
	r.HandleFunc(config.OTB_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateOTB))).Methods(http.MethodPost)

	//	EDITA OTB (cod_otb)
	r.HandleFunc(config.OTB_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateOTB))).Methods(http.MethodPut)

	//	LISTA OTB (cod_otb)
	r.HandleFunc(config.OTB_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetOTBByID))).Methods(http.MethodGet)

	//	APAGA OTB (cod_otb)
	r.HandleFunc(config.OTB_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteOTB))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM OTB FATURA
	=========================	*/

	//	SALVA FATURA OTB
	r.HandleFunc(config.OTB_FATURA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateOTBFatura))).Methods(http.MethodPost)

	//	LISTA FATURA OTB (cod_otb)
	r.HandleFunc(config.OTB_FATURA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetOTBFatura))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM OTB ITENS (ITENS OTB)
	=========================	*/

	//	EDITA ITENS OTB
	r.HandleFunc(config.OTB_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensOTB))).Methods(http.MethodPut)

	//	LISTA ITENS OTB (cod_otb)
	r.HandleFunc(config.OTB_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensOTB))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM FATURA
	=========================	*/

	//	LISTA FATURA
	r.HandleFunc(config.FATURA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllFatura))).Methods(http.MethodGet)

	//	SALVA FATURA
	r.HandleFunc(config.FATURA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateFatura))).Methods(http.MethodPost)

	//	LISTA FATURA (num_nf)
	r.HandleFunc(config.FATURA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetFaturaByID))).Methods(http.MethodGet)

	//	APAGA FATURA (num_nf)
	r.HandleFunc(config.FATURA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteFatura))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM FATURA ITENS (ITENS FATURA)
	=========================	*/

	//	SALVA FATURA ITENS
	//r.HandleFunc(config.OTB_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateOTB))).Methods(http.MethodPost)

	//	EDITA FATURA ITENS
	r.HandleFunc(config.FATURA_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensFatura))).Methods(http.MethodPut)

	//	LISTA FATURA ITENS (num_nf)
	r.HandleFunc(config.FATURA_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllItensFatura))).Methods(http.MethodGet)

	//	APAGA FATURA ITENS (num_nf, id_empenho, cod_item, cod_tipo_item)
	r.HandleFunc(config.FATURA_ITENS_DELETE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteItensFatura))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM FATURA OTB
	=========================	*/

	//	LISTA FATURA (num_nf, cod_ibge)
	r.HandleFunc(config.FATURA_OTB_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetFaturaOTB))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM PREVISAO EMPENHO OK!
	=========================	*/

	//	LISTA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPrevisaoEmpenho))).Methods(http.MethodGet)

	//	SALVA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrevisaoEmpenho))).Methods(http.MethodPost)

	//	EDITA PREVISAO EMPENHO (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA PREVISAO EMPENHO (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	//	APAGA PREVISAO EMPENHO (cod_previsao_empenho)
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePrevisaoEmpenho))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM ITENS PREVISAO EMPENHO OK!
	=========================	*/

	//	EDITA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensPrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	/*	=========================
		ROTAS EM ASSUNTO OK!
	=========================	*/

	//	LISTA ASSUNTO
	r.HandleFunc(config.ASSUNTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllAssunto))).Methods(http.MethodGet)

	//	SALVA ASSUNTO
	r.HandleFunc(config.ASSUNTO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateAssunto))).Methods(http.MethodPost)

	//	EDITA ASSUNTO (cod_assunto)
	r.HandleFunc(config.ASSUNTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateAssunto))).Methods(http.MethodPut)

	//	LISTA ASSUNTO (cod_assunto)
	r.HandleFunc(config.ASSUNTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAssuntoByID))).Methods(http.MethodGet)

	//	APAGA ASSUNTO (cod_assunto)
	r.HandleFunc(config.ASSUNTO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteAssunto))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM CATEGORIA OK!
	=========================	*/
	//	LISTA CATEGORIA
	r.HandleFunc(config.CATEGORIA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCategoria))).Methods(http.MethodGet)

	//	SALVA CATEGORIA
	r.HandleFunc(config.CATEGORIA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateCategoria))).Methods(http.MethodPost)

	//	EDITA CATEGORIA (cod_categoria)
	r.HandleFunc(config.CATEGORIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCategoria))).Methods(http.MethodPut)

	//	LISTA CATEGORIA (cod_categoria)
	r.HandleFunc(config.CATEGORIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCategoriaByID))).Methods(http.MethodGet)

	//	APAGA CATEGORIA (cod_categoria)
	r.HandleFunc(config.CATEGORIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteCategoria))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM CLASSE EMPENHO OK!
	=========================	*/

	//	LISTA CLASSE EMPENHO
	r.HandleFunc(config.CLASSE_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllClasseEmpenho))).Methods(http.MethodGet)

	//	SALVA CLASSE EMPENHO
	r.HandleFunc(config.CLASSE_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateClasseEmpenho))).Methods(http.MethodPost)

	//	EDITA CLASSE EMPENHO (cod_classe_empenho)
	r.HandleFunc(config.CLASSE_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateClasseEmpenho))).Methods(http.MethodPut)

	//	LISTA CLASSE EMPENHO (cod_classe_empenho)
	r.HandleFunc(config.CLASSE_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetClasseEmpenhoByID))).Methods(http.MethodGet)

	//	APAGA CLASSE EMPENHO (cod_classe_empenho)
	r.HandleFunc(config.CLASSE_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteClasseEmpenho))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM ETAPA OK!
	=========================	*/

	//	LISTA ETAPA
	r.HandleFunc(config.ETAPA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllEtapa))).Methods(http.MethodGet)

	//	SALVA ETAPA
	r.HandleFunc(config.ETAPA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateEtapa))).Methods(http.MethodPost)

	//	EDITA ETAPA (cod_etapa)
	r.HandleFunc(config.ETAPA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEtapa))).Methods(http.MethodPut)

	//	LISTA ETAPA (cod_etapa)
	r.HandleFunc(config.ETAPA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEtapaByID))).Methods(http.MethodGet)

	//	APAGA ETAPA (cod_etapa)
	r.HandleFunc(config.ETAPA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEtapa))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM ITENS OK!
	=========================	*/

	//	LISTA ITENS
	r.HandleFunc(config.ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllItens))).Methods(http.MethodGet)

	//	SALVA ITENS
	r.HandleFunc(config.ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateItens))).Methods(http.MethodPost)

	//	EDITA ITENS (cod_itens)
	r.HandleFunc(config.ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItens))).Methods(http.MethodPut)

	//	LISTA ITENS (cod_itens)
	r.HandleFunc(config.ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensByID))).Methods(http.MethodGet)

	//	APAGA ITENS (cod_itens)
	r.HandleFunc(config.ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteItens))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM MUNICIPIOS OK!
	=========================	*/

	//	LISTA MUNICIPIOS
	r.HandleFunc(config.MUNICIPIO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllMunicipio))).Methods(http.MethodGet)

	//	SALVA MUNICIPIOS
	r.HandleFunc(config.MUNICIPIO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateMunicipio))).Methods(http.MethodPost)

	//	EDITA MUNICIPIOS (cod_ibge)
	r.HandleFunc(config.MUNICIPIO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateMunicipio))).Methods(http.MethodPut)

	//	LISTA MUNICIPIOS (cod_ibge)
	r.HandleFunc(config.MUNICIPIO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetMunicipioByID))).Methods(http.MethodGet)

	//	APAGA MUNICIPIOS (cod_ibge)
	r.HandleFunc(config.MUNICIPIO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteMunicipio))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM NATUREZA DESPESA OK!
	=========================	*/

	//	LISTA NATUREZA_DESPESA
	r.HandleFunc(config.NATUREZA_DESPESA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllNaturezaDespesa))).Methods(http.MethodGet)

	//	SALVA NATUREZA_DESPESA
	r.HandleFunc(config.NATUREZA_DESPESA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateNaturezaDespesa))).Methods(http.MethodPost)

	//	EDITA NATUREZA_DESPESA (cod_natureza_despesa)
	r.HandleFunc(config.NATUREZA_DESPESA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateNaturezaDespesa))).Methods(http.MethodPut)

	//	LISTA NATUREZA_DESPESA (cod_natureza_despesa)
	r.HandleFunc(config.NATUREZA_DESPESA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetNaturezaDespesaByID))).Methods(http.MethodGet)

	//	APAGA NATUREZA_DESPESA (cod_natureza_despesa)
	r.HandleFunc(config.NATUREZA_DESPESA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteNaturezaDespesa))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM PREFEITOS OK!
	=========================	*/

	//	LISTA PREFEITOS
	r.HandleFunc(config.PREFEITOS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPrefeitos))).Methods(http.MethodGet)

	//	SALVA PREFEITOS
	r.HandleFunc(config.PREFEITOS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrefeitos))).Methods(http.MethodPost)

	//	EDITA PREFEITOS (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrefeitos))).Methods(http.MethodPut)

	//	LISTA PREFEITOS (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrefeitosByID))).Methods(http.MethodGet)

	//	APAGA PREFEITOS (cod_prefeito)
	r.HandleFunc(config.PREFEITOS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeletePrefeitos))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM TIPOLOGIAS OK!
	=========================	*/

	//	LISTA TIPOLOGIA
	r.HandleFunc(config.TIPOLOGIA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllTipologia))).Methods(http.MethodGet)

	//	SALVA TIPOLOGIA
	r.HandleFunc(config.TIPOLOGIA_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateTipologia))).Methods(http.MethodPost)

	//	EDITA TIPOLOGIA (cod_tipologia)
	r.HandleFunc(config.TIPOLOGIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateTipologia))).Methods(http.MethodPut)

	//	LISTA TIPOLOGIA (cod_tipologia)
	r.HandleFunc(config.TIPOLOGIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetTipologiaByID))).Methods(http.MethodGet)

	//	APAGA TIPOLOGIA (cod_tipologia)
	r.HandleFunc(config.TIPOLOGIA_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteTipologia))).Methods(http.MethodDelete)

	/*	=========================
		ROTAS EM TIPO ITEM OK!
	=========================	*/

	//	LISTA TIPO_ITEM
	r.HandleFunc(config.TIPO_ITEM_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllTipoItem))).Methods(http.MethodGet)

	//	SALVA TIPO_ITEM
	r.HandleFunc(config.TIPO_ITEM_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateTipoItem))).Methods(http.MethodPost)

	//	EDITA TIPO_ITEM (cod_tipo_item)
	r.HandleFunc(config.TIPO_ITEM_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateTipoItem))).Methods(http.MethodPut)

	//	LISTA TIPO_ITEM (cod_tipo_item)
	r.HandleFunc(config.TIPO_ITEM_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetTipoItemByID))).Methods(http.MethodGet)

	//	APAGA TIPO_ITEM (cod_tipo_item)
	r.HandleFunc(config.TIPO_ITEM_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteTipoItem))).Methods(http.MethodDelete)

	return
}
