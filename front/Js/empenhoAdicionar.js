//  pega o token do login
let meuToken = localStorage.getItem("token");

//  tratamento de erros
function erros(value) {
  if (value == 400) {
    window.location.href="./errors/400.html";
  } else if (value == 401) {
    window.location.href="./errors/401.html";
  } else if (value == 403) {
    window.location.href="./errors/403.html";
  } else if (value == 404) {
    window.location.href="./errors/404.html";
  } else if (value == 409) {
    alert("Erro: Lote já existente.");
  } else if (value == 412) {
    alert("Erro: Informação colocada é incorreta.");
  } else if (value == 422) {
    alert("Erro: Formato de informação não aceito.");
  } else if (value == 500) {
    window.location.href="./errors/500.html";
  } else if (value == 504) {
    window.location.href="./errors/504.html";
  } else {
    alert("ERRO DESCONHECIDO");
  }
}

//  JSON usado para mandar as informações no fetch
let info = {
    "cod_previsao_empenho": "",
    "cod_empenho": "",
    "data": "",
    "contador": "",
};

function previsao(){
  fetch('http://localhost:8080/read/previsaoempenho', {
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
  fetch('http://localhost:8080/read/empenho', {
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