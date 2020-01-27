package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) CreateHandler() (r *mux.Router) {

	//cria um roteador

	r = s.Router
	//Home
	r.HandleFunc("/{id}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.Home))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM USUARIO
	=========================	*/

	//ROTA DE LOGIN
	r.HandleFunc("/read/usuario/login", middlewares.SetMiddleJSON(s.Login)).Methods(http.MethodPost)

	//LISTA USUARIOS
	r.HandleFunc("/read/usuario", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUsers))).Methods(http.MethodGet)

	//SALVA USUARIO
	r.HandleFunc("/read/usuario/createuser", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateUser))).Methods(http.MethodPost)

	//LISTA USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUser))).Methods(http.MethodGet)

	//EDITA O USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateUser))).Methods(http.MethodPut)

	//APAGA O USUARIO
	//	r.HandleFunc(config.USER_ID_PATH, middlewares.SetMiddleAuth(s.DeleteUser)).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ENTIDADE
	=========================	*/

	//	LISTA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEntidades))).Methods(http.MethodGet)

	//	SALVA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.AddEntidade))).Methods(http.MethodPost)

	//	EDITA ENTIDADE
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEntidade))).Methods(http.MethodPut)

	//	LISTA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEntidadeByID))).Methods(http.MethodGet)

	//	APAGA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEntidade))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM LOTE
	=========================	*/

	//	LISTA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLotes))).Methods(http.MethodGet)

	//	SALVA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.AddLote))).Methods(http.MethodPost)

	//	EDITA LOTE
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLote))).Methods(http.MethodPut)

	//	LISTA LOTE POR ID
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteByID))).Methods(http.MethodGet)

	//	APAGA LOTE POR ID
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteLote))).Methods(http.MethodDelete)

	//---------Rotas de Reajuste

	//lista reajuste
	//nao existe

	//salva reajuste
	r.HandleFunc(config.REAJUSTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.AddReajuste))).Methods(http.MethodPost)

	//lista reajuste
	r.HandleFunc(config.REAJUSTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetReajustes))).Methods(http.MethodGet)

	//lista reajuste por ID
	r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetReajuste))).Methods(http.MethodGet)

	//edita reajuste por ID
	r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateReajuste))).Methods(http.MethodGet)

	//apaga reajuste (lote_cod_lote, ano_ref)
	r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteReajuste))).Methods(http.MethodDelete)

	//**********Rotas em Cd

	//lista cd
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCD))).Methods(http.MethodGet)

	//salva cd
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateCd))).Methods(http.MethodPost)

	//edita cd
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCd))).Methods(http.MethodPut)

	//lista cd por ID
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDByID))).Methods(http.MethodGet)

	//----------Rotas de Cd Itens
	//edita cd itens
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCdItens))).Methods(http.MethodPut)

	//lista cd itens
	r.HandleFunc(config.CD_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCdItens))).Methods(http.MethodGet)

	//lista cd itens ID
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCdItem))).Methods(http.MethodGet)

	//----------Rotas de Lote Itens
	//edita lote itens
	r.HandleFunc("/{id1}/{id2}/{id3}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLote_itens))).Methods(http.MethodPut)

	//lista lote itens ID
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLote_itensByID))).Methods(http.MethodGet)

	//----------Rotas de PrevisaoEmpenho
	//lista previsao empenho
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrevisaoEmpenhos))).Methods(http.MethodGet)

	//salva previsao empenho
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrevisaoEmpenho))).Methods(http.MethodPost)

	//edita previsao empenho
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrevisaoEmpenho))).Methods(http.MethodPut)

	//lista previsao emprenho por ID
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrevisaoEmpenho))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM ITENS PREVISAO EMPENHO
	=========================	*/

	//	EDITA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensPrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS PREVISAO EMPENDO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	/*
		//**********Rotas em Assunto

	/*
		//**********Rotas em Assunto
		//----------Rotas de Assunto
		//lista assunto
		r.HandleFunc().Methods(http.MethodGet)

		//salva assunto
		r.HandleFunc().Methods(http.MethodPost)

		//lista assunto por ID
		r.HandleFunc().Methods(http.MethodGet)

		//edita assunto ID
		r.HandleFunc().Methods(http.MethodPut)

		//apaga assunto ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Categoria
		//Rotas de Categoria
		//lista categotia
		r.HandleFunc().Methods(http.MethodGet)

		//salva categoria
		r.HandleFunc().Methods(http.MethodPost)

		//edita categoria
		r.HandleFunc().Methods(http.MethodPut)

		//lista categoria por ID
		r.HandleFunc().Methods(http.MethodGet)

		//deleta categoria ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Cidades Digitais
			//---------Rotas de Estado---------//
			//---------Rotas de Municipios da Cidade Digital---------//
			//---------Rotas de Cidade Digital---------//
			//---------Rotas de Itens da Cidade Digital---------//
			//---------Rotas de Processo da Cidade Digital---------//
			//---------Rotas de Acompanhamento da Cidade Digital---------//
			//---------Rotas de Ponto da Cidade Digital---------//
			//---------Rotas de Pagamento da Cidade Digital---------//

		//**********Rotas em ClasseEmpenho
		//Rotas de ClaseEmpenho
		//lista classe empenho
		r.HandleFunc().Methods(http.MethodGet)

		//salva classe empenho
		r.HandleFunc().Methods(http.MethodPost)

		//edita classe empenho
		r.HandleFunc().Methods(http.MethodPut)

		//lista classe empenho por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga classe empenho ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Contato
			//----------Rotas De Contato
				//ROTAS EM CONTATO_PATH
					//lista contato
						r.HandleFunc().Methods(http.MethodGet)

					//salva contato
						r.HandleFunc().Methods(http.MethodPost)

					//edita contato
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM CONTATO_ID_PATH
					//apaga contato ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas De Telefone
				//ROTAS EM TEL_PATH
					//lista telefone
						r.HandleFunc().Methods(http.MethodGet)

					//salva telefone
						r.HandleFunc().Methods(http.MethodPost)
				ROTAS EM TEL_ID
					//apaga telefone ID
						r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Empenho
			//---------Rotas de Empenho
				//ROTAS EM EMPENHO_PATH
					//lista empenho
						r.HandleFunc().Methods(http.MethodGet)

					//salva empenho
						r.HandleFunc().Methods(http.MethodPost)

					//edita empenho
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM EMPENHO_ID_PATH
					//lista empenho por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga emepnho ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Empenho Itens----------//
				//ROTAS EM EMPENHO_ITENS
					//edita empenho itens
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM EMPENHO_ITENS_ID
					//lista empenho itens ID
						r.HandleFunc().Methods(http.MethodGet)


		//**********Rotas em Etapa
		//----------Rotas de Etapa
		//lista etapa
		r.HandleFunc().Methods(http.MethodGet)

		//salva etapa
		r.HandleFunc().Methods(http.MethodPost)

		//edita etapa
		r.HandleFunc().Methods(http.MethodPut)

		//lista etapa por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga etapa ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Fatura
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//

		//**********Rotas em Itens
		//----------Rotas de Itens
		//lista itens
		r.HandleFunc().Methods(http.MethodGet)

		//salva itens
		r.HandleFunc().Methods(http.MethodPost)

		//edita itens
		r.HandleFunc().Methods(http.MethodPut)

		//lista itens por ID (cod_item, cod_tipo_item)
		r.HandleFunc().Methods(http.MethodGet)

		//apaga itens (cod_item, cod_tipo_item)
		r.HandleFunc().Methods(http.MethodDelete)





		//**********Rotas em Modulos
		//----------Rotas de Modulos
		//lista modulos
		r.HandleFunc().Methods(http.MethodGet)

		//lista Modulos Usuario (cod_usuario)
		r.HandleFunc().Methods(http.MethodGet)

		//**********Rotas em Municipios
		//----------Rotas de Municipio
		//lista minicipio
		r.HandleFunc().Methods(http.MethodGet)

		//salva municipio
		r.HandleFunc().Methods(http.MethodPost)

		//edita municipio
		r.HandleFunc().Methods(http.MethodPut)

		//lista municipio por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga municipio ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em NaturezaDespesa
		//----------Rotas de NaturezaDespesa
		//lista natureza despesa
		r.HandleFunc().Methods(http.MethodGet)

		//salva natureza despesa
		r.HandleFunc().Methods(http.MethodPost)

		//edita natureza despesa
		r.HandleFunc().Methods(http.MethodPut)

		//lista natureza despesa por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga natureza despesa ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Pagamento (otb)
			//----------Rotas de Pagamento
				//ROTAS EM PAG_PATH
					//lista pagamento
						r.HandleFunc().Methods(http.MethodGet)

					//salva pagamento
						r.HandleFunc().Methods(http.MethodPost)

					//edita pagamento
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM PAG_ID_PATH
					//lista pagamento por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga pagamento ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Municipio Fatura
				//ROTAS EM PAG_FAT_MUNICIPIO
					//lista municipio fatura (cd_municipio_cod_ibge)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_FAT_LIST
					//lista pagamento fatura (cod_otb)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_FAT_SAVE
					//salva pagamento fatura
						r.HandleFunc().Methods(http.MethodPost)

			//----------Rotas de Pagamento Itens
				//ROTAS EM PAG_LIST_ITENS
					//lista pagamento itens	(cod_otb)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_EDIT_ITENS
					//edita pagamento itens
						r.HandleFunc().Methods(http.MethodPut)

		//**********Rotas em Prefeitos
		//----------Rotas de Prefeitos
		//lista prefeitos
		r.HandleFunc().Methods(http.MethodGet)

		//salva prefeitos
		r.HandleFunc().Methods(http.MethodPost)

		//edita prefeitos
		r.HandleFunc().Methods(http.MethodPut)

		//lista prefeitos por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga prefeitos ID
		r.HandleFunc().Methods(http.MethodDelete)





		//**********Rotas em TipoItem

		//**********Rotas em Tipologia



	*/
	return
}

//Methods: OPTIONS, GET, POST, PUT, DELETE
