//  JSON usado para mandar as informações no fetch
let info = {
  "num_nf": "",
  "cod_ibge": "",
  "dt_nf": "",
};

function enviar() {

  info.num_nf = parseInt(document.getElementById("num_nf").value);
  info.cod_ibge = parseInt(document.getElementById("cod_ibge").value);
  info.dt_nf = document.getElementById("dt_nf").value;
  

  //transforma as informações em string para mandar
  let corpo = JSON.stringify(info);
  console.log(corpo);
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
      window.location.replace("./fiscalizacao.html");
    } else {
      erros(response.status);
    }
  });
}