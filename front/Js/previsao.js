let previsaoTotal = [];

window.onload = function () {
  paginacao();
  //altera o campo tipos como pedido, pode alterar aqui se necessário
  document.getElementById("tipo").innerHTML = "<option value=''>Tipo</option><option value='o'>Original</option><option value='r'>Reajuste</option>";
}

//sistema de paginação
let contador = 0;
let porPagina = 5;
let totalPaginas;

function antes() {
  contador--;
  paginacao();
}

function depois() {
  contador++;
  paginacao();
}

//garantindo o limite de paginação
function pagina(valor) {
  contador = valor;
  paginacao();
}

function paginacao() {
  porPagina = document.getElementById("quantos").value;
  let comeco = contador * porPagina;
  let fim = (contador + 1) * porPagina;

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

      response.json().then(function (json) {
        
        totalPaginas = json.length / porPagina;

        //para edição
        previsaoTotal=json;

        //pegar o json
        //console.log(json);

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
        <tr>
        <th style="width:10%" scope="col">Código de Previsão de Empenho</th>
        <th style="width:10%" scope="col">Código do Lote</th>
        <th style="width:45%" scope="col">Natureza da despesa</th>
        <th style="width:10%" scope="col">Tipo</th>
        <th style="width:10%" scope="col">Data</th>
        <th style="width:10%" scope="col">Ano de Referência</th>
        <th style="width:5%" scope="col">Opções</th>
        </tr>
        </thead>`);
        tabela += (`<tbody>`);

        
        for (let i = comeco; i < fim && i < json.length; i++) {
          //captura itens para tabela
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += json[i]["cod_previsao_empenho"];
          tabela += (`</td>`);
          tabela += (`<td>`);
          tabela += json[i]["cod_lote"];
          tabela += (`</td>`);
          tabela += (`<td>`);
          tabela += json[i]["natureza_despesa"];
          tabela += (`</td>`);
          tabela += (`<td>`);
          if(json[i]["tipo"]=="o"){
            tabela += "Original";
          }
          else if(json[i]["tipo"]=="r"){
            tabela += "Reajuste";
          }
          tabela += (`</td>`);
          tabela += (`<td>`);
          let data1 = new Date(json[i]["data"]);
          let dataFinal1 = String(data1.getDate()).padStart(2, '0') + "/" + String(data1.getMonth() + 1).padStart(2, '0') + "/" + String(data1.getFullYear()).padStart(4, '0');
          tabela += dataFinal1;
          tabela += (`</td>`);
          tabela += (`<td>`);
          tabela += json[i]["ano_referencia"];
          tabela += (`</td>`);
          tabela += (`<td> 
          <span class="d-flex">
          <button onclick="editarPrevisao(` + i + `)" class="btn btn-success">
          <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
          </button>
          </td>`);
          tabela += (`</tr>`);
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

        //mostra quanto do total aparece na tela
        document.getElementById("mostrando").innerHTML = "Mostrando " + (comeco + 1) + " a " + fim + " de " + json.length;
        if (porPagina > json.length - comeco) {
          document.getElementById("mostrando").innerHTML = "Mostrando " + (comeco + 1) + " a " + json.length + " de " + json.length;
        }

        //conta quantas paginas é necessário
        let paginas = `<li id="anterior" class="page-item" ><a href="#" class="page-link" onclick="antes()">Anterior</a></li>`;
        //apenas aciona se precisar de paginação
        if (json.length > porPagina) {
          //caso seja apenas 10 paginas
          if (totalPaginas < 10) {
            for (i = 0; i < totalPaginas; i++) {
              if (contador == i) {
                paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link btn active">` + (i + 1) + `</a></li>`;
              } else {
                paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link">` + (i + 1) + `</a></li>`;
              }
            }
            //caso sejam mais de 10
          } else {
            //mostrar apenas inicio e fim
            if (contador == 0) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link btn active">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            //opções do começo
            else if (contador == 1) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link btn active">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link">4</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == 2) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link btn active">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link">4</a></li>`;
              paginas += `<li class="page-item" id="page4"><a href="#" onclick="pagina(4)" class="page-link">5</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == 3) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link btn active">4</a></li>`;
              paginas += `<li class="page-item" id="page4"><a href="#" onclick="pagina(4)" class="page-link">5</a></li>`;
              paginas += `<li class="page-item" id="page5"><a href="#" onclick="pagina(5)" class="page-link">6</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            //opções no final
            else if (contador == Math.floor(totalPaginas - 3)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 5) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 5) + `)" class="page-link">` + Math.floor(totalPaginas - 4) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 4) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 4) + `)" class="page-link">` + Math.floor(totalPaginas - 3) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 3) + `)" class="page-link btn active">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == Math.floor(totalPaginas - 2)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 4) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 4) + `)" class="page-link">` + Math.floor(totalPaginas - 3) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 3) + `)" class="page-link">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link btn active">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == Math.floor(totalPaginas - 1)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 3) + `)" class="page-link">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link btn active">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == Math.floor(totalPaginas)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link btn active">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador - 2) + `"><a href="#" onclick="pagina(` + (contador - 2) + `)" class="page-link">` + (contador - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador - 1) + `"><a href="#" onclick="pagina(` + (contador - 1) + `)" class="page-link">` + contador + `</a></li>`;
              paginas += `<li class="page-item" id="page` + contador + `"><a href="#" onclick="pagina(` + contador + `)" class="page-link btn active">` + (contador + 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador + 1) + `"><a href="#" onclick="pagina(` + (contador + 1) + `)" class="page-link">` + (contador + 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador + 2) + `"><a href="#" onclick="pagina(` + (contador + 2) + `)" class="page-link">` + (contador + 3) + `</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
          }
        }
        paginas += `<li id="proximo" class="page-item" ><a href="#" class="page-link" onclick="depois()">Próximo</a></li>`;
        document.getElementById("paginacao").innerHTML = paginas;

        //limite das paginas
        if (contador > 0) {
          document.getElementById("anterior").style.visibility = "visible";
        } else {
          document.getElementById("anterior").style.visibility = "hidden";
        }
        if (fim < json.length) {
          document.getElementById("proximo").style.visibility = "visible";
        } else {
          document.getElementById("proximo").style.visibility = "hidden";
        }

      });
    } else {
      erros(response.status);
    }
  });
}



//funções para enviar

function pegarLote() {
  fetch(servidor + 'read/lote', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {
        //console.log(json);
        let x = [];
        x[0] = "<option value=''>Lote</option>";
        for (i = 0; i < json.length; i++) {
          x[i+1] += "<option >" + json[i].cod_lote + "</option>";
        }
        document.getElementById("cod_lote").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
  });
}

function pegarNaturezaDespesa() {
  fetch(servidor + 'read/naturezadespesa', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {
        //console.log(json);
        let x = [];
        for (i = 0; i < json.length; i++) {
          // o valor pego é o codigo, mas o campo mostra a descrição
          x[i] += "<option value=" + json[i].cod_natureza_despesa + ">" + json[i].descricao + "</option>";
        }
        document.getElementById("cod_natureza_despesa").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
  });
}


function enviar() {

  //  JSON usado para mandar as informações no fetch
  let info = {
    "cod_lote": "",
    "cod_natureza_despesa": "",
    "data": "",
    "tipo": "",
    "ano_referencia": "",
  };

  info.cod_lote = parseInt(document.getElementById("cod_lote").value);
  info.cod_natureza_despesa = parseInt(document.getElementById("cod_natureza_despesa").value);
  info.data = document.getElementById("data").value;
  info.tipo = document.getElementById("tipo").value;
  info.ano_referencia = parseInt(document.getElementById("ano_referencia").value);

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  console.log(corpo);
  //função fetch para mandar
  fetch(servidor + 'read/previsaoempenho', {
    method: 'POST',
    body: corpo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar o status do pedido
    //console.log(response);

    //tratamento dos erros
    if (response.status == 200 || response.status == 201) {
      location.reload();
    } else {
      erros(response.status);
    }
  });
}



//leva para o editor do campo selecionado
function editarPrevisao(valor) {
  localStorage.setItem("cod_previsao_empenho", previsaoTotal[valor].cod_previsao_empenho);
  localStorage.setItem("cod_lote", previsaoTotal[valor].cod_lote);
  localStorage.setItem("cod_natureza_despesa", previsaoTotal[valor].cod_natureza_despesa);
  localStorage.setItem("data", previsaoTotal[valor].data);
  localStorage.setItem("tipo", previsaoTotal[valor].tipo);
  localStorage.setItem("ano_referencia", previsaoTotal[valor].ano_referencia);
  window.location.href = "./gerenciaPrevisao.html";
}