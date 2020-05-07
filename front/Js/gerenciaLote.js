//pega o CNPJ escolhido anteriormente
let meuLote = localStorage.getItem("cod_lote");
document.getElementById("cod_lote").value = meuLote;

//cuida de lote_itens
let edicaoItem = [];
let listaItem = [];
let meuItem = [];
let meuTipo = [];
let itemMudado = [];

//cuida de reajuste
let listaReajuste = [];
let edicaoReajuste = [];
let meuAno = [];
let reajusteMudado = [];

//cuida de previsao de empenho
let listaPrevisao = [];

//JSON usado para mandar as informações no fetch
let info = {
  "cod_lote": "",
  "cnpj": "",
  "contrato": "",
  "dt_inicio_vig": "",
  "dt_final_vig": "",
  "dt_reajuste": "",
};

function pegarEntidade(){
  fetch(servidor + 'read/entidadeget', {
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
          // o valor pego é o cnpj, mas o campo mostra o nome da entidade
          x[i] += "<option value=" + json[i].cnpj + ">" + json[i].nome + "</option>";
        }
        x.sort();

        document.getElementById("cnpj").innerHTML = x;

        let cnpj1 = document.getElementById("cnpj");
        cnpj1.value = localStorage.getItem("cnpj");
      });
    } else {
      erros(response.status);
    }
  });
}


window.onload = function () {

  //preenche os campos
  let contrato1 = document.getElementById("contrato");
  contrato1.value = localStorage.getItem("contrato");

  //estes campos precisam de adaptações para serem aceitos, como yyyy-MM-dd

  let data1 = new Date(localStorage.getItem("dt_inicio_vig"));
  let data2 = new Date(localStorage.getItem("dt_final_vig"));
  let data3 = new Date(localStorage.getItem("dt_reajuste"));

  let dataFinal1 = String(data1.getFullYear()).padStart(4, '0') + "-" + String(data1.getMonth() + 1).padStart(2, '0') + "-" + String(data1.getDate()).padStart(2, '0');
  let dataFinal2 = String(data2.getFullYear()).padStart(4, '0') + "-" + String(data2.getMonth() + 1).padStart(2, '0') + "-" + String(data2.getDate()).padStart(2, '0');
  let dataFinal3 = String(data3.getFullYear()).padStart(4, '0') + "-" + String(data3.getMonth() + 1).padStart(2, '0') + "-" + String(data3.getDate()).padStart(2, '0');

  document.getElementById("dt_inicio_vig").innerHTML = dataFinal1;
  document.getElementById("dt_final_vig").innerHTML = dataFinal2;
  document.getElementById("dt_reajuste").innerHTML = dataFinal3;

  //esta função preenche o campo de lote
  pegarEntidade();
}

function enviar() {

  info.cod_lote = parseFloat(meuLote);
  info.cnpj = document.getElementById("cnpj").value;
  info.contrato = document.getElementById("contrato").value;
  info.dt_inicio_vig = document.getElementById("dt_inicio_vig").value;
  info.dt_final_vig = document.getElementById("dt_final_vig").value;
  info.dt_reajuste = document.getElementById("dt_reajuste").value;

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch(servidor + 'read/lote/' + meuLote, {
    method: 'PUT',
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
      window.location.replace("./lote.html");
    } else {
      erros(response.status);
    }
  });
}