//pega o CNPJ escolhido anteriormente
let meuCodigo = localStorage.getItem("cod_previsao_empenho");
let meuCodigoSec = localStorage.getItem("cod_lote");

window.onload = function () {

  //preenche os campos
  document.getElementById("cod_previsao_empenho").value = meuCodigo;
  document.getElementById("cod_lote").value = meuCodigoSec;
  document.getElementById("ano_referencia").value = localStorage.getItem("ano_referencia");
  document.getElementById("cod_natureza_despesa").value = localStorage.getItem("natureza_despesa");

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
    "data": document.getElementById("data").value,
    "tipo": document.getElementById("tipo").value,
    "ano_referencia": parseInt(document.getElementById("ano_referencia").value),
  };

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

    //tratamento dos erros
    if (response.status == 200 || response.status == 201) {
      window.location.replace("./fiscPrevisao.html");
    } else {
      erros(response.status);
    }
  });
}