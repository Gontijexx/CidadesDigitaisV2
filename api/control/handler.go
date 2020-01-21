package control

import (
	"CidadesDigitaisV2/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) CreateHandler() (r *mux.Router) {

	//cria um roteador

	r = s.Router
	//Home
	r.HandleFunc("/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.Home))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM USUARIO
	=========================	*/

	//ROTA DE LOGIN
	r.HandleFunc("/read/usuario/login", middlewares.SetMiddleJSON(s.Login)).Methods(http.MethodPost)

	//LISTA USUARIOS
	r.HandleFunc("/read/usuario", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetUsers))).Methods(http.MethodGet)

	//SALVA USUARIO
	r.HandleFunc("/read/usuario/createUser", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.CreateUser))).Methods(http.MethodPost)

	//LISTA USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetUser))).Methods(http.MethodGet)

	//EDITA O USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateUser))).Methods(http.MethodPut)

	//APAGA O USUARIO
	//	r.HandleFunc(config.USER_ID_PATH, middlewares.SetMiddleAuth(s.DeleteUser)).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ENTIDADE
	=========================	*/

	//LISTA ENTIDADE
	//r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetEntidades))).Methods(http.MethodGet)

	//SALVA ENTIDADE
	r.HandleFunc("/read/entidade", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.CreateEntidade))).Methods(http.MethodPost)

	//edita entidade
	//r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateEntidade))).Methods(http.MethodPut)

	//lista entidade por ID
	//r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetEntidadeByID))).Methods(http.MethodGet)

	//apaga entidade ID
	//r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(s.DeleteEntidades)).Methods(http.MethodDelete)

	/*




		//**********Rotas em Lotes

		//lista lote
		r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetLote))).Methods(http.MethodGet)

		//salva lote
		r.HandleFunc(config.LOTE_PATH_CREATELOTE, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.CreateLote))).Methods(http.MethodPost)

		//edita lote
		r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateLote))).Methods(http.MethodPut)

		//lista lote por ID
		r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetLoteByID))).Methods(http.MethodGet)

		//apaga lote ID
		//r.HandleFunc().Methods(http.MethodDelete)

		//---------Rotas de Reajuste

		//lista reajuste
		//nao existe

		//salva reajuste
		r.HandleFunc(config.REAJUSTE_PATH_CREATEREAJUSTE, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.CreateReajuste))).Methods(http.MethodPost)

		//lista reajuste por ID
		r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetReajusteByID))).Methods(http.MethodGet)

		//apaga reajuste (lote_cod_lote, ano_ref)
		r.HandleFunc(config.REAJUSTE_DEL, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.DeleteReajuste))).Methods(http.MethodDelete)

		//**********Rotas em Cd

		//lista cd
		r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetCd))).Methods(http.MethodGet)

		//salva cd
		r.HandleFunc(config.CD_PATH_CREATECD, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.CreateCd))).Methods(http.MethodPost)

		//edita cd
		r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateCd))).Methods(http.MethodPut)

		//lista cd por ID
		r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetCdByID))).Methods(http.MethodGet)

		//----------Rotas de Cd Itens
		//edita cd itens
		r.HandleFunc("/{id1}/{id2}/{id3}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateCd_itens))).Methods(http.MethodPut)

		//lista cd itens ID
		r.HandleFunc(config.CD_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetCd_itensByID))).Methods(http.MethodGet)

		//----------Rotas de Lote Itens
		//edita lote itens
		r.HandleFunc("/{id1}/{id2}/{id3}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateLote_itens))).Methods(http.MethodPut)

		//lista lote itens ID
		r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetCd_itensByID))).Methods(http.MethodGet)

		//----------Rotas de PrevisaoEmpenho
		//lista previsao empenho
		r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetPrevisao_empenhos))).Methods(http.MethodGet)

		//salva previsao empenho
		r.HandleFunc(config.PREVISAO_EMPENHO_PATH_CREATEPREVISAO_EMPENHO, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.CreatePrevisao_empenho))).Methods(http.MethodPost)

		//edita previsao empenho
		r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdatePrevisao_empenho))).Methods(http.MethodPut)

		//lista previsao emprenho por ID
		r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetPrevisao_empenhoByID))).Methods(http.MethodGet)

		//----------Rotas de ItensPrevisaoEmpenho
		//edita itens previsao empenho
		r.HandleFunc("/{id1}/{id2}/{id3}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.UpdateItens_previsao_empenho))).Methods(http.MethodPut)

		//lista itens previsao empenho (cod_lote, cod_previsao_empenho)
		r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuthMod(s.GetItens_previsao_empenhoByID))).Methods(http.MethodGet)



	*/
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
