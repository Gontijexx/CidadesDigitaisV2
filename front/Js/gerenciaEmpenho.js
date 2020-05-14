//pega o CNPJ escolhido anteriormente
let meuEmpenho = localStorage.getItem("id_empenho");

function pegarPrevisao(){
  fetch(servidor + 'read/previsaoempenho', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {
        //pegar o json
        //console.log(json);
        let x = [];
        for (i = 0; i < json.length; i++) {
          x[i] += "<option>" + json[i]["cod_previsao_empenho"] + "</option>";
        }
        x.sort();
        document.getElementById("cod_previsao_empenho").innerHTML = x;

        //colocando valor original
        document.getElementById("cod_previsao_empenho").value = localStorage.getItem("cod_previsao_empenho");
      });
    } else {
      erros(response.status);
    }
  });
}


window.onload = function () {

    console.log(meuEmpenho);

  //preenche os campos
  this.document.getElementById("cod_empenho").value = localStorage.getItem("cod_empenho");

  //esta função preenche o campo de lote
  pegarPrevisao();

  //este campo precisa de adaptação para ser aceito, como yyyy-MM-dd

  let data1 = new Date(localStorage.getItem("data"));
  let dataFinal1 = String(data1.getFullYear()).padStart(4, '0') + "-" + String(data1.getMonth() + 1).padStart(2, '0') + "-" + String(data1.getDate()).padStart(2, '0');
  document.getElementById("data").value = dataFinal1;
  
  
}

function enviar() {

  //JSON usado para mandar as informações no fetch
  let info = {
    "cod_empenho": "",
    "cod_previsao_empenho": "",
    "data": "",
  };

  info.cod_empenho = document.getElementById("cod_empenho").value;
  info.cod_previsao_empenho = parseInt(document.getElementById("cod_previsao_empenho").value);
  info.data = document.getElementById("data").value;

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch(servidor + 'read/empenho/' + meuEmpenho, {
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
      window.location.replace("./empenho.html");
    } else {
      erros(response.status);
    }
  });
}