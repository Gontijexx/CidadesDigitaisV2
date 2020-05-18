//pega o CNPJ escolhido anteriormente
let meuFatura = localStorage.getItem("id_fatura");

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
        cdTotal=json;
      });
      pegarMunicipio();
    } else {
      erros(response.status);
    }
  });
}

function pegarMunicipio() {

  //preenche os campos para estado e municipio
  fetch(servidor + 'read/municipio', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {
        //cria variaveis
        let i, j = 0;
        let x = [],
          valorUF = [];

        //impede que haja repetição
        for (i = 0; i < json.length; i++) {
          if (i>0 && json[i].uf != json[i-1].uf) {
            valorUF[j] = json[i];
            j++;
          }
        }
        console.log(cdTotal);
        //faz a ligação entre variaveis e valores do banco
        let k = 0;
        for (i = 0; i < valorUF.length; i++) {
          for (j = 0; j < cdTotal.length; j++) {
            if (valorUF[i].uf == cdTotal[j].uf) {
              cidades[k] = cdTotal[j].uf;
              k++;
            }
          }
        }

        x[0] += "<option value=''>Estado</option>";
        for (i = 0; i < j; i++) {
          x[i + 1] += "<option>" + cidades[i] + "</option>";
        }
        x.sort();
        document.getElementById("uf").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
  });
}


function enabler() {

  document.getElementById("cod_ibge").disabled = false;
  let uf = document.getElementById("uf");
  let i, j = 0,
    x = [],
    cidadesFinal = [];
  for (i = 0; i < cidades.length; i++) {
    if (cidades[i].uf == uf.value) {
      cidadesFinal[j] = cidades[i];
      j++;
    }
  }
  for (i = 0; i < cidadesFinal.length; i++) {
    x[i] = "<option value=" + cidadesFinal[i].cod_ibge + ">" + cidadesFinal[i].nome_municipio + "</option>";
  }
  x.sort();
  document.getElementById("cod_ibge").innerHTML = x;
}


window.onload = function () {

  //preenche os campos
  this.document.getElementById("num_nf").value = localStorage.getItem("num_nf");

  //esta função preenche os campos de municipio
  pegarCD();

  //este campo precisa de adaptação para ser aceito, como yyyy-MM-dd

  let data = new Date(localStorage.getItem("dt_nf"));
  let dataFinal = String(data.getFullYear()).padStart(4, '0') + "-" + String(data.getMonth() + 1).padStart(2, '0') + "-" + String(data.getDate()).padStart(2, '0');
  document.getElementById("dt_nf").value = dataFinal;
  
}

function enviar() {

  //JSON usado para mandar as informações no fetch
  let info = {
    "num_nf": "",
    "cod_ibge": "",
    "dt_nf": "",
  };

  info.num_nf = document.getElementById("num_nf").value;
  info.cod_ibge = parseInt(document.getElementById("cod_ibge").value);
  info.dt_nf = document.getElementById("dt_nf").value;

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch(servidor + 'read/fatura/' + meuFatura, {
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
      window.location.replace("./fatura.html");
    } else {
      erros(response.status);
    }
  });
}