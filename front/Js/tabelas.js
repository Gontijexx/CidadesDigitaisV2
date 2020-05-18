//até agora só foi adaptado para lote
//adapte para CD João. Use administracao.js como exemplo

function itens() {

  //cria o botão para editar
  document.getElementById("editar").innerHTML = (`<button id="editar" onclick="editarItem()" class="btn btn-success">Salvar Alterações em Itens</button>`);
  document.getElementById("editar2").innerHTML = (`<button id="editar" onclick="editarItem()" class="btn btn-success">Salvar Alterações em Itens</button>`);

  //função fetch para chamar itens da tabela
  fetch(servidor + 'read/loteitens', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      //console.log(response.statusText);

      //pegar o json que possui a tabela
      response.json().then(function (json) {

        let j = 0;
        //cria uma lista apenas com os itens do lote selecionado
        for (let i = 0; i < json.length; i++) {
          if (json[i]["cod_lote"] == meuLote) {
            listaItem[j] = json[i];
            j++;
          }
        }

        //cria a tabela para 
        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
                <tr>
                <th scope="col">Código do Item, Tipo de Item e Descrição do item</th>
                <th scope="col">Valor</th>
                </tr>
                </thead>`);
        tabela += (`<tbody>`);

        for (i = 0; i < listaItem.length; i++) {

          //salva os valores para edição
          meuItem[i] = listaItem[i]["cod_item"];
          meuTipo[i] = listaItem[i]["cod_tipo_item"];

          //cria json para edição
          edicaoItem[i] = {
            "preco": "",
          };

          //captura itens para tabela
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += listaItem[i]["cod_item"] + "." + listaItem[i]["cod_tipo_item"] + " - " + listaItem[i]["descricao"];
          tabela += (`</td> <td>`);
          tabela += "R$ " + (`<input value="` + listaItem[i]["preco"] + `" onchange="mudaItem(` + i + `)" id="preco` + i + `" type="text" class="preco">`);
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

function mudaItem(valor) {
  edicaoItem[valor].preco = parseFloat(document.getElementById("preco" + valor).value);
  itemMudado[valor] = valor;
}

function editarItem() {
  for (i = 0; i < listaItem.length; i++) {
    if (itemMudado[i] != null) {
      //transforma as informações em string para mandar
      let corpo = JSON.stringify(edicaoItem[i]);
      //função fetch para mandar
      fetch(servidor + 'read/loteitens/' + meuLote + '/' + meuItem[i] + '/' + meuTipo[i], {
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
          //checar a resposta do pedido
          //console.log(json);
          window.location.replace("./gerenciaLote.html");
        } else {
          //erros(response.status);
        }
      });
    }
  }
}







//lote reajustes

function reajuste() {

  document.getElementById("cod_lote1").value = meuLote;
  document.getElementById("cod_lote1").disabled = true;

  document.getElementById("editar").innerHTML = (`<button onclick="editarReajuste()" class="btn btn-success" >Salvar Alterações em Reajustes</button>
                                                  <button class="btn btn-success" data-toggle="modal" data-target="#adicionarReajuste">Novo Reajuste</button>`);
  document.getElementById("editar2").innerHTML = (`<button onclick="editarReajuste()" class="btn btn-success">Salvar Alterações em Reajustes</button>`);

  //função fetch para chamar reajustes da tabela
  fetch(servidor + 'read/reajuste', {
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
        <th style="width:46%" scope="col">Ano de Referência</th>
        <th style="width:46%" scope="col">Percentual de Reajuste</th>
        <th style="width:8%" scope="col">Apagar</th>
        </tr>
        </thead>`);
        tabela += (`<tbody>`);

        let j = 0;
        for (let i = 0; i < json.length; i++) {
          if (json[i].cod_lote == meuLote) {
            listaReajuste[j] = json[i];
            j++;
          }
        }

        for (i = 0; i < listaReajuste.length; i++) {

          //salva os valores para edição
          meuAno[i] = listaReajuste[i]["ano_ref"];

          //cria json para edição
          edicaoReajuste[i] = {
            "percentual": "",
          };

          //captura itens para tabela
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += listaReajuste[i]["ano_ref"];
          tabela += (`</td>`);
          tabela += (`<td>`);
          tabela += (`<input value="` + listaReajuste[i]["percentual"] + `" onchange="mudaReajuste(` + i + `)" id="percentual` + i + `" type="number">`) + "%";
          tabela += (`</td>`);
          tabela += (`<td>
          <button onclick="apagarReajuste(` + listaReajuste[i]["ano_ref"] + `)" class="btn btn-danger">
          <i class="material-icons"data-toggle="tooltip" title="Delete">&#xE872;</i>
          </button>
          </td>`);
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

function mudaReajuste(valor) {
  edicaoReajuste[valor].percentual = parseFloat(document.getElementById("percentual" + valor).value);
  reajusteMudado[valor] = valor;
}

function editarReajuste() {
  for (i = 0; i < listaReajuste.length; i++) {
    if (reajusteMudado[i] != null) {
      //transforma as informações em string para mandar
      let corpo = JSON.stringify(edicaoReajuste[i]);
      //função fetch para mandar
      fetch(servidor + 'read/reajuste/' + listaReajuste[i]["ano_ref"] + '/' + meuLote, {
        method: 'PUT',
        body: corpo,
        headers: {
          'Authorization': 'Bearer ' + meuToken
        },
      }).then(function (response) {

        //tratamento dos erros
        if (response.status == 200 || response.status == 201) {
          //checar a resposta do pedido
          //console.log(json);
          window.location.replace("./gerenciaLote.html");
        } else {
          erros(response.status);
        }
      });
    }
  }
}

function novoReajuste() {

  let infoReajuste = {
    "cod_lote": parseInt(meuLote),
    "ano_ref": "",
    "percentual": "",
  };

  infoReajuste.ano_ref = parseInt(document.getElementById("ano_ref").value);
  infoReajuste.percentual = parseFloat(document.getElementById("percentual").value);

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(infoReajuste);
  //função fetch para mandar
  fetch(servidor + 'read/reajuste', {
    method: 'POST',
    body: corpo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200 || response.status == 201) {
      //checar a resposta do pedido
      //console.log(json);

      location.reload();
    } else {
      erros(response.status);
    }
  });
}

function apagarReajuste(valor) {

  //função fetch para deletar
  fetch(servidor + 'read/reajuste/' + valor + "/" + meuLote, {
    method: 'DELETE',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 204) {
      alert("Apagado com sucesso.");
      window.location.replace("./gerenciaLote.html");
    } else {
      erros(response.status);
    }
    return response.json().then(function (json) {
      console.log(json);
    });
  });
}







//lote previsao:

function previsao() {

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
        for (let i = 0; i < json.length; i++) {
          console.log(json[i].cod_lote)
          if (json[i].cod_lote == meuLote) {
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