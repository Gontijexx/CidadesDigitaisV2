//pega o CNPJ escolhido anteriormente
let meuCodigo = localStorage.getItem("cod_previsao_empenho");
let meuCodigoSec = localStorage.getItem("cod_lote");

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
          x[i] += "<option value=" + json[i].cod_natureza_despesa + ">" + json[i].cod_natureza_despesa + ' - ' + json[i].descricao + "</option>";
        }
        document.getElementById("cod_natureza_despesa").innerHTML = x;
        //deixar com o valor inicial
        document.getElementById("cod_natureza_despesa").value = localStorage.getItem("cod_natureza_despesa");
      });
    } else {
      erros(response.status);
    }
  });
}

window.onload = function () {

  //esta função preenche o campo de natureza de despesa
  pegarNaturezaDespesa();

  //preenche os campos
  document.getElementById("cod_previsao_empenho").value = meuCodigo;
  document.getElementById("cod_lote").value = meuCodigoSec;
  document.getElementById("ano_referencia").value = localStorage.getItem("ano_referencia");

  document.getElementById("tipo").innerHTML = "<option value='o'>Original</option><option value='r'>Reajuste</option>";
  document.getElementById("tipo").value = localStorage.getItem("tipo");

  //este campo precisa de adaptação para ser aceito, como yyyy-MM-dd
  let data1 = new Date(localStorage.getItem("data"));
  let dataFinal1 = String(data1.getFullYear()).padStart(4, '0') + "-" + String(data1.getMonth() + 1).padStart(2, '0') + "-" + String(data1.getDate()).padStart(2, '0');
  document.getElementById("data").value = dataFinal1;

}

function enviar() {

  //  JSON usado para mandar as informações no fetch
  let info = {
    "cod_natureza_despesa": "",
    "data": "",
    "tipo": "",
    "ano_referencia": "",
  };

  info.cod_natureza_despesa = parseInt(document.getElementById("cod_natureza_despesa").value);
  info.data = document.getElementById("data").value;
  info.tipo = document.getElementById("tipo").value;
  info.ano_referencia = parseInt(document.getElementById("ano_referencia").value);

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch(servidor + 'read/previsaoempenho/' + parseInt(meuCodigo), {
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
      window.location.replace("./previsao.html");
    } else {
      erros(response.status);
    }
  });
}