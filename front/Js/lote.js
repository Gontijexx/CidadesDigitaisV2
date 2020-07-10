//capturar chave primaria
let jsonFinal = [];

function paginacao() {
  porPagina = document.getElementById("quantos").value;
  let comeco = contador * porPagina;
  let fim = (contador + 1) * porPagina;

  //função fetch para chamar itens da tabela
  fetch(servidor + 'read/lote', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {
    //console.log(response)
    //tratamento dos erros
    if (response.status == 200) {
      console.log(response.statusText);
      //pegar o json que possui a tabela
      response.json().then(function (json) {

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
            <tr>
            <th style="width:10%" scope="col">Lote</th>
            <th style="width:20%" scope="col">Entidade - CNPJ</th>
            <th style="width:15%" scope="col">Contrato</th>
            <th style="width:15%" scope="col">Data de Inicio</th>
            <th style="width:15%" scope="col">Data Final</th>
            <th style="width:15%" scope="col">Data de Reajuste</th>
            <th style="width:10%" scope="col">Opções</th>
            </tr>
            </thead>`);
        tabela += (`<tbody>`);

        //sistema de filtragem:
        let filtrado = [];
        filtrado = filtro(json,["cnpj","nome","endereco","bairro","cep","uf","nome_municipio","observacao"]);

        //para edição
        jsonFinal=filtrado;

        for (let i = comeco; i < fim && i < filtrado.length; i++) {
          tabela += (`<tr><td>`);
          tabela += filtrado[i]["cod_lote"];
          tabela += (`</td> <td>`);
          tabela += filtrado[i]["nome"] + " - " + filtrado[i]["cnpj"];
          tabela += (`</td> <td>`);
          tabela += filtrado[i]["contrato"];
          tabela += (`</td> <td>`);

          let data1 = new Date(filtrado[i]["dt_inicio_vig"]);
          let dataf1 = String(data1.getDate()).padStart(2, '0') + '/' + String(data1.getMonth() + 1).padStart(2, '0') + '/' + String(data1.getFullYear()).padStart(4, '0');
          tabela += dataf1;
          tabela += (`</td> <td>`);

          let data2 = new Date(filtrado[i]["dt_final_vig"]);
          let dataf2 = String(data2.getDate()).padStart(2, '0') + '/' + String(data2.getMonth() + 1).padStart(2, '0') + '/' + String(data2.getFullYear()).padStart(4, '0');
          tabela += dataf2;
          tabela += (`</td> <td>`);

          //let data3 = new Date(filtrado[i]["dt_reajuste"]);
          //let dataf3 = String(data3.getDate()).padStart(2, '0') + '/' + String(data3.getMonth() + 1).padStart(2, '0');
          //usar mascara para pegar apenas o que precisa
          tabela += filtrado[i]["dt_reajuste"];

          tabela += (`</td><td>
                  <span class="d-flex">
                  <button onclick="editarLote(` + i + `)" class="btn btn-success">
                  <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
                  </button>
                  </td></tr>`);
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

        paginasOrganizadas(filtrado,comeco,fim);
      });
    } else {
      erros(response.status);
    }
  });
}


function pegarCNPJ() {
  //preenche os CNPJs
  fetch(servidor + 'read/entidade', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {
        //cria variaveis
        let i = 0;
        let x = [];
        x[0] += "<option value='00000000000000'>CNPJ</option>";
        for (i = 0; i < json.length; i++) {
          x[i + 1] += "<option>" + json[i].cnpj + "</option>";
        }
        
        document.getElementById("cnpj").innerHTML = x;
        paginacao();
      });
    } else {
      erros(response.status);
    }
  });
}


window.onload = function () {
  pegarCNPJ();
}

function enviar() {

  //estrutura usada para mandar o JSON no fetch
  let info = {
    "cod_lote": "",
    "cnpj": "",
    "contrato": "",
    "dt_inicio_vig": "",
    "dt_final_vig": "",
    "dt_reajuste": ""
  };

  let a = document.getElementById("cod_lote");
  info.cod_lote = parseFloat(a.value);
  let b = document.getElementById("cnpj");
  info.cnpj = b.value;
  let c = document.getElementById("contrato");
  info.contrato = c.value;
  let d = document.getElementById("dt_inicio_vig");
  info.dt_inicio_vig = d.value;
  let e = document.getElementById("dt_final_vig");
  info.dt_final_vig = e.value;
  let f = document.getElementById("dt_reajuste");
  info.dt_reajuste = f.value;

  //transforma as informações do token em json
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch(servidor + 'read/lote', {
    method: 'POST',
    body: corpo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar o status do pedido
    //console.log(response);

    //tratamento dos erros
    if (response.status == 201) {
      location.reload();
    } else {
      erros(response.status);
    }
  });
}

//leva para o editor do campo selecionado
function editarLote(valor) {
  localStorage.setItem("cod_lote", jsonFinal[valor].cod_lote);
  localStorage.setItem("cnpj", jsonFinal[valor].cnpj);
  localStorage.setItem("contrato", jsonFinal[valor].contrato);
  localStorage.setItem("dt_inicio_vig", jsonFinal[valor].dt_inicio_vig);
  localStorage.setItem("dt_final_vig", jsonFinal[valor].dt_final_vig);
  localStorage.setItem("dt_reajuste", jsonFinal[valor].dt_reajuste);
  console.log(jsonFinal[valor].dt_inicio_vig + " " + jsonFinal[valor].dt_final_vig + " " + jsonFinal[valor].dt_reajuste);
  window.location.href = "./gerenciaLote.html";
}