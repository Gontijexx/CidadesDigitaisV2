package main

import "CidadesDigitaisV2/api"

func main() {
	api.Run()
}

//join do lote_itens e itens
// select * from lote_itens
// inner join itens on lote_itens.cod_item = itens.cod_item and lote_itens.cod_tipo_item = itens.cod_tipo_item

// -- Código do Item, Tipo de Item e Descrição do item	