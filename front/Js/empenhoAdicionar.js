//  JSON usado para mandar as informações no fetch
let info = {
  "cod_previsao_empenho": "",
  "cod_empenho": "",
  "data": "",
  "contador": "",
};

function previsao() {
  fetch(servidor + 'read/previsaoempenho', {
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
          x[i] += "<option >" + json[i].cod_previsao_empenho + "</option>";
        }
        document.getElementById("cod_previsao_empenho").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
  });
}

window.onload = function () {
  previsao();
}

function enviar() {

  info.cod_previsao_empenho = parseInt(document.getElementById("cod_previsao_empenho").value);
  info.cod_empenho = document.getElementById("cod_empenho").value;
  info.data = document.getElementById("data").value;
  info.contador = parseInt(document.getElementById("contador").value);

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  console.log(corpo);
  //função fetch para mandar
  fetch(servidor + 'read/empenho', {
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
      window.location.replace("./fiscalizacao.html");
    } else {
      erros(response.status);
    }
  });
}