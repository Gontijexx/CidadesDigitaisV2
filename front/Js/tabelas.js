//até agora só foi adaptado para lote
//objetivo de pegar todas as tabelas de
//adapte para CD João. Use administracao.js como exemplo



//tabela pra previsão de empenho:

function previsao(valorCodigo) {

  document.getElementById("editar").innerHTML = (`<br>`);
  document.getElementById("editar2").innerHTML = (`<br>`);

  //função fetch para chamar os itens de previsão da tabela
  fetch(servidor + 'read/previsaoempenho', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {
    //checar os status de pedidos
    //console.log(response)
    //tratamento dos erros
    if (response.status == 200) {
      //console.log(response.statusText);

      //pegar o json que possui a tabela
      response.json().then(function (json) {

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
          <tr>
          <th style="width:15%" scope="col">Código de Previsão de Empenho</th>
          <th style="width:40%" scope="col">Natureza da despesa</th>
          <th style="width:10%" scope="col">Tipo</th>
          <th style="width:20%" scope="col">Data</th>
          <th style="width:15%" scope="col">Ano de Referência</th>
          </tr>
          </thead>`);
        tabela += (`<tbody>`);

        let j = 0;
        let listaPrevisao = [];
        for (let i = 0; i < json.length; i++) {
          if (valorCodigo == meuCodigo) {
            listaPrevisao[j] = json[i];
            j++;
          }
        }

        for (i = 0; i < listaPrevisao.length; i++) {
          //captura itens para tabela
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += listaPrevisao[i]["cod_previsao_empenho"];
          tabela += (`</td><td>`);
          tabela += listaPrevisao[i]["natureza_despesa"];
          tabela += (`</td><td>`);
          tabela += listaPrevisao[i]["tipo"];
          tabela += (`</td><td>`);
          let data1 = new Date(listaPrevisao[i]["data"]);
          let dataFinal1 = String(data1.getDate()).padStart(2, '0') + "/" + String(data1.getMonth() + 1).padStart(2, '0') + "/" + String(data1.getFullYear()).padStart(4, '0');
          tabela += dataFinal1;
          tabela += (`</td><td>`);
          tabela += listaPrevisao[i]["ano_referencia"];
          tabela += (`</td>`);
          tabela += (`</tr>`);
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;
      });
    } else {
      erros(response.status);
    }
  });
}







//Itens de financeamento

let listaItem = [],
  meuItem = [],
  meuTipo = [],
  edicaoItem = [],
  itemMudado = [];

function itensFinanceamento(caminho) {

  //cria o botão para editar
  if(caminho == "itensprevisaoempenho"){
    document.getElementById("editar").innerHTML = (`<button class="btn btn-success" onclick="editarItem('itensprevisaoempenho')">Salvar Alterações em Itens</button>`);
    document.getElementById("editar2").innerHTML = (`<button class="btn btn-success" onclick="editarItem('itensprevisaoempenho')">Salvar Alterações em Itens</button>`);
  }
  else if(caminho == "itensempenho"){
    document.getElementById("editar").innerHTML = (`<button class="btn btn-success" onclick="editarItem('itensempenho')">Salvar Alterações em Itens</button>`);
    document.getElementById("editar2").innerHTML = (`<button class="btn btn-success" onclick="editarItem('itensempenho')">Salvar Alterações em Itens</button>`);
  }
  else if(caminho == "itensfatura"){
    document.getElementById("editar").innerHTML = (`<button class="btn btn-success" onclick="editarItem('itensfatura')">Salvar Alterações em Itens</button>`);
    document.getElementById("editar2").innerHTML = (`<button class="btn btn-success" onclick="editarItem('itensfatura')">Salvar Alterações em Itens</button>`);
  }
 
  //função fetch para chamar itens da tabela
  fetch(servidor + 'read/' + caminho + "/" + meuCodigo + "/" + meuCodigoSec, {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar os status de pedidos
    //console.log(response)

    //tratamento dos erros
    if (response.status == 200) {
      console.log(response.statusText);

      //pegar o json que possui a tabela
      response.json().then(function (json) {

        console.log(meuCodigoSec)

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
        <tr>
        <th style="width:40%" scope="col">Descrição</th>
        <th style="width:20%" scope="col">Quantidade</th>
        <th style="width:20%" scope="col">Valor</th>
        <th style="width:20%" scope="col">Subtotal</th>
        </tr>
        </thead>`);
        tabela += (`<tbody>`);

        //armazenando para edição
        listaItem = json;

        //calculo do total
        let total = 0;

        for (i = 0; i < listaItem.length; i++) {
          //salva os valores para edição
          meuTipo[i] = listaItem[i]["cod_tipo_item"];
          meuItem[i] = listaItem[i]["cod_item"];

          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += listaItem[i]["cod_tipo_item"] + '.' + listaItem[i]["cod_item"] + ' - ' + listaItem[i]["descricao"];
          tabela += (`</td> <td>`);
          tabela += (`<input value="` + listaItem[i]["quantidade"] + `" onchange="mudaItem(` + i + `)" id="quantidade` + i + `" type="number">`);
          tabela += (`</td> <td>`);
          tabela += (`<input value="` + listaItem[i]["valor"] + `" onchange="mudaItem(` + i + `)" id="valor` + i + `" type="number">`);
          tabela += (`</td> <td>`);
          tabela += listaItem[i]["quantidade_disponivel"];
          total = total + listaItem[i]["quantidade_disponivel"];
          tabela += (`</td>`);
          tabela += (`</tr>`);

          edicaoItem[i] = {
            "quantidade": listaItem[i]["quantidade"],
            "valor": listaItem[i]["valor"],
          };
        }
        tabela += (`<tr>`);

        tabela += (`<td>`);
        tabela += (`<p> Total: </p>`);
        tabela += (`</td>`);

        //espaços
        tabela += (`<td>`);
        tabela += (`</td>`);
        tabela += (`<td>`);
        tabela += (`</td>`);

        //valor final
        tabela += (`<td>`);
        tabela += total;
        tabela += (`</td>`);

        tabela += (`</tr>`);

        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;
      });
    } else {
      erros(response.status);
    }
  });
}

function mudaItem(itemPego) {
  edicaoItem[itemPego].quantidade = parseInt(document.getElementById("quantidade" + itemPego).value);
  edicaoItem[itemPego].valor = parseInt(document.getElementById("valor" + itemPego).value);
  itemMudado[itemPego] = itemPego;
}

function editarItem(caminho) {

  for (let i = 0; i < listaItem.length; i++) {

    if (itemMudado[i] != null) {
      //transforma as informações do token em json
      let corpo = JSON.stringify(edicaoItem[i]);
      //função fetch para mandar
      fetch(servidor + 'read/' + caminho + '/' + meuCodigo + '/' + meuItem[i] + '/' + meuTipo[i], {
        method: 'PUT',
        body: corpo,
        headers: {
          'Authorization': 'Bearer ' + meuToken
        },
      }).then(function (response) {
        //checar o status do pedido
        //console.log(response.statusText);

        //tratamento dos erros
        if (response.status == 200 || response.status == 201) {
          location.reload();
        } else {
          erros(response.status);
        }
        location.reload();
      });
    }
  }
}