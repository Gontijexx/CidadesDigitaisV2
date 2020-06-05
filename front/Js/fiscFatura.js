//variaveis globais
let jsonFinal = [];
let cidades = [];









































//nova subjanela para pagamento com todos os pagamentos relacionados e faz hiperlinks
//remodelar fatura

window.onload = function () {
  paginacao();
  pegarCD();
}

function paginacao() {
  porPagina = document.getElementById("quantos").value;
  let comeco = contador * porPagina;
  let fim = (contador + 1) * porPagina;

  //função fetch para chamar os itens de previsão da tabela
  fetch(servidor + 'read/fatura', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar os status de pedidos
    //console.log(response);

    //tratamento dos erros
    if (response.status == 200) {

      response.json().then(function (json) {
        
        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
        <tr>
        <th style="width:20%" scope="col">Número da Fatura</th>
        <th style="width:50%" scope="col">Município</th>
        <th style="width:20%" scope="col">Data</th>
        <th style="width:10%" scope="col">Opções</th>
        </tr>
        </thead>`);
        tabela += (`<tbody>`);

        
        for (let i = comeco; i < fim && i < json.length; i++) {
          //captura itens para tabela
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += json[i]["num_nf"]; //está sendo enviado assim por algum motivo
          tabela += (`</td><td>`);
          tabela += json[i]["nome_municipio"] + " - " + json[i]["uf"] + " - " + json[i]["cod_ibge"];
          tabela += (`</td><td>`);
          let data1 = new Date(json[i]["dt_nf"]);
          let dataFinal1 = String(data1.getDate()).padStart(2, '0') + "/" + String(data1.getMonth() + 1).padStart(2, '0') + "/" + String(data1.getFullYear()).padStart(4, '0');
          tabela += dataFinal1;
          tabela += (`</td>`);
          tabela += (`<td>
          <span class="d-flex">
          <button onclick="editarFatura(` + i + `)" class="btn btn-success">
          <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
          </button>
          </td>`);
          tabela += (`</tr>`);
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

        paginasOrganizadas(json,comeco,fim);
      });
    } else {
      erros(response.status);
    }
  });
}



function pegarCD() {
  fetch(servidor + 'read/cd', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {

        //variaveis
        let i, j=0;
        let x = [];
        let ufCD = [];

        //para usar em enabler()
        cidades=json;

        //para tirar repetições
        for(i=0; i<json.length;i++){
          if(i != 0 && json[i].uf != json[i-1].uf){
            ufCD[j] = json[i];
            j++;
          }
        }

        //preenche "uf"
        x[0] = "<option value='A'>Estado</option>";
        for (i = 0; i < ufCD.length; i++) {
          x[i+1] += "<option>" + ufCD[i].uf + "</option>";
        }
        
        document.getElementById("uf").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
  });
}



function enabler() {
  document.getElementById("cod_ibge").disabled = false;

  //variaveis
  let uf = document.getElementById("uf");
  let i, j = 0;
  let x = [], cidadesFinal = [];

  //para tirar repetições
  for (i = 0; i < cidades.length; i++) {
    if (cidades[i].uf == uf.value) {
      cidadesFinal[j] = cidades[i];
      j++;
    }
  }

  //preenche "cod_ibge"
  x[0] = "<option value='A'>Cidade</option>";
  for (i = 0; i < cidadesFinal.length; i++) {
    x[i+1] = "<option value=" + cidadesFinal[i].cod_ibge + ">" + cidadesFinal[i].nome_municipio + "</option>";
  }
  
  document.getElementById("cod_ibge").innerHTML = x;
}



function enviar() {

  //  JSON usado para mandar as informações no fetch
  let info = {
    "num_nf": "",
    "cod_ibge": "",
    "dt_nf": "",
  };

  info.num_nf = parseInt(document.getElementById("num_nf").value);
  info.cod_ibge = parseInt(document.getElementById("cod_ibge").value);
  info.dt_nf = document.getElementById("dt_nf").value;

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch(servidor + 'read/fatura', {
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
      //checar o json
      //response.json().then(function (json) {
      //console.log(json);
      //});
      location.reload();
    } else {
      erros(response.status);
    }
  });
}



//leva para o editor do campo selecionado
function editarFatura(valor) {
  localStorage.setItem("num_nf", jsonFinal[valor]["num_nf"]);
  localStorage.setItem("cod_ibge", jsonFinal[valor]["cod_ibge"]);
  localStorage.setItem("dt_nf", jsonFinal[valor]["dt_nf"]);
  localStorage.setItem("uf", jsonFinal[valor]["uf"]);
  localStorage.setItem("nome_municipio", jsonFinal[valor]["nome_municipio"]);
  window.location.href = "./gerenciaFatura.html";
}