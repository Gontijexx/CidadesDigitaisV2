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

function itensFinanceamento(caminho,estrutura) {

  //cria o botão para editar
  document.getElementById("editar").innerHTML = (`<button id="editar" onclick="editarItem()" class="btn btn-success">Salvar Alterações em Itens</button>`);
  document.getElementById("editar2").innerHTML = (`<button id="editar" onclick="editarItem()" class="btn btn-success">Salvar Alterações em Itens</button>`);

  //função fetch para chamar itens da tabela
  fetch(servidor + 'read/' + caminho, {
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

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
        <tr>
        <th scope="col">Descrição</th>
        <th scope="col">Quantidade prevista</th>
        <th scope="col">Quantidade do projeto executivo</th>
        <th scope="col">Quantidade de termo de instalação </th>
        </tr>
        </thead>`);
        tabela += (`<tbody>`);

        //cria uma lista apenas com os itens do grupo selecionado
        let j = 0;
        for (let i = 0; i < json.length; i++) {
          if (json[i]["cod_ibge"] == meuCodigo) {
            listaItem[j] = json[i];
            j++;
          }
        }

        //estrutura

        /*for (i = 0; i < json.length; i++) {
          //captura itens para tabela
          tabela += (`<tr>`);
          for (j = 1; j < estrutura.length; j++) {
            tabela += (`<td>`);
            tabela += json[i][estrutura[j]];
            tabela += (`</td>`);
          }
          tabela += (`</tr>`);
        }*/

        for (i = 0; i < listaItem.length; i++) {

          //salva os valores para edição
          meuItem[i] = listaItem[i]["cod_item"];
          meuTipo[i] = listaItem[i]["cod_tipo_item"];

          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += listaItem[i]["descricao"];
          tabela += (`</td> <td>`);
          tabela += (`<input value="` + listaItem[i]["quantidade_previsto"] + `" onchange="mudaItem(` + i + `)" id="quantidade_previsto` + i + `" type="text">`);
          tabela += (`</td> <td>`);
          tabela += (`<input value="` + listaItem[i]["quantidade_projeto_executivo"] + `" onchange="mudaItem(` + i + `)" id="quantidade_projeto_executivo` + i + `" type="text">`);
          tabela += (`</td> <td>`);
          tabela += (`<input value="` + listaItem[i]["quantidade_termo_instalacao"] + `" onchange="mudaItem(` + i + `)" id="quantidade_termo_instalacao` + i + `" type="text">`);
          tabela += (`</td>`);
          tabela += (`</tr>`);

          edicaoItem[i] = {
            "quantidade_previsto": listaItem[i]["quantidade_previsto"],
            "quantidade_projeto_executivo": listaItem[i]["quantidade_projeto_executivo"],
            "quantidade_termo_instalacao": listaItem[i]["quantidade_termo_instalacao"],
          };
        }


        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

      });
    } else {
      erros(response.status);
    }
  });
}

function mudaItem(valor) {
  edicaoItem[valor].quantidade_previsto = parseInt(document.getElementById("quantidade_previsto" + valor).value);
  edicaoItem[valor].quantidade_projeto_executivo = parseInt(document.getElementById("quantidade_projeto_executivo" + valor).value);
  edicaoItem[valor].quantidade_termo_instalacao = parseInt(document.getElementById("quantidade_termo_instalacao" + valor).value);
  itemMudado[valor] = valor;
}

function editarItem() {

  for (let i = 0; i < listaItem.length; i++) {

    if (itemMudado[i] != null) {
      //transforma as informações do token em json
      let corpo = JSON.stringify(edicaoItem[i]);
      //função fetch para mandar
      fetch(servidor + 'read/cditens/' + meuCD + '/' + meuItem[i] + '/' + meuTipo[i], {
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
          //erros(response.status);
        }
        window.location.replace("./gerenciaCd.html");
      });
    }
  }
}