//pega o token do login
let meuToken = localStorage.getItem("token");

//tratamento de erros
function erros(value) {
    if (value == 400) {
      window.location.replace("./errors/400.html");
    } else if (value == 401) {
      window.location.replace("./errors/401.html");
    } else if (value == 403) {
      window.location.replace("./errors/403.html");
    } else if (value == 404) {
      window.location.replace("./errors/404.html");
    } else if (value == 409) {
      alert("Erro: Lote já existente.");
    } else if (value == 412) {
      alert("Erro: Informação colocada é incorreta.");
    } else if (value == 422) {
      alert("Erro: Formato de informação não aceito.");
    } else if (value == 500) {
      window.location.replace("./errors/500.html");
    } else if (value == 504) {
      window.location.replace("./errors/504.html");
    } else {
      alert("ERRO DESCONHECIDO");
    }
}

window.onload = function () {
  paginacao();
}


//sistema de paginação
let contador = 0;
let porPagina = 5;
let totalPaginas;


function antes() {
  contador--;
  paginacao();
}

function depois() {
  contador++;
  paginacao();
}

//garantindo o limite de paginação

function pagina(valor) {
  contador = valor;
  paginacao();
}

function paginacao() {
  porPagina = document.getElementById("quantos").value;
  let comeco = contador * porPagina;
  let fim = (contador + 1) * porPagina;

    //função fetch para chamar os itens de previsão da tabela
  fetch('http://localhost:8080/read/previsaoempenho', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar os status de pedidos
    //console.log(response)

    //tratamento dos erros
    if (response.status == 200) {
      //console.log(response.statusText);

      //pegar o json que possui a tabela
      response.json().then(function (json) {

        totalPaginas = json.length / porPagina;
        
        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
        <tr>
        <th style="width:10%" scope="col">Código do Lote</th>
        <th style="width:10%" scope="col">Código de Previsão de Empenho</th>
        <th style="width:40%" scope="col">Natureza da despesa</th>
        <th style="width:10%" scope="col">Tipo</th>
        <th style="width:20%" scope="col">Data</th>
        <th style="width:10%" scope="col">Ano de Referência</th>
        </tr>
        </thead>`);
        tabela += (`<tbody>`);

        json.sort();
        for (let i = comeco; i < fim && i < json.length; i++) {
          //captura itens para tabela
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += json[i]["cod_lote"];
          tabela += (`</td><td>`);
          tabela += json[i]["cod_previsao_empenho"];
          tabela += (`</td><td>`);
          tabela += json[i]["cod_natureza_despesa"];
          tabela += (`</td><td>`);
          tabela += json[i]["tipo"];
          tabela += (`</td><td>`);
          let data1 = new Date(json[i]["data"]);
          let dataFinal1 = String(data1.getDate()).padStart(2, '0') + "/" + String(data1.getMonth() + 1).padStart(2, '0') + "/" + String(data1.getFullYear()).padStart(4, '0');
          tabela += dataFinal1;
          tabela += (`</td><td>`);
          tabela += json[i]["ano_referencia"];
          tabela += (`</td>`);
          tabela += (`</tr>`);
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

         //mostra quanto do total aparece na tela
         document.getElementById("mostrando").innerHTML = "Mostrando " + (comeco + 1) + " a " + fim + " de " + json.length;
         if (porPagina > json.length - comeco) {
           document.getElementById("mostrando").innerHTML = "Mostrando " + (comeco + 1) + " a " + json.length + " de " + json.length;
         }

         //conta quantas paginas é necessário
         let paginas = `<li id="anterior" class="page-item" ><a href="#" class="page-link" onclick="antes()">Anterior</a></li>`;
         if (json.length > porPagina) {
           for (i = 0; i < totalPaginas; i++) {
             paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link">` + (i + 1) + `</a></li>`;
           }
         }
         paginas += `<li id="proximo" class="page-item" ><a href="#" class="page-link" onclick="depois()">Próximo</a></li>`;
         document.getElementById("paginacao").innerHTML = paginas;

         //limite das paginas
         if (contador > 0) {
           document.getElementById("anterior").style.visibility = "visible";
         } else {
           document.getElementById("anterior").style.visibility = "hidden";
         }
         if (fim<json.length) {
           document.getElementById("proximo").style.visibility = "visible";
         } else {
           document.getElementById("proximo").style.visibility = "hidden";
         }

      });
    } else {
      erros(response.status);
    }
  });
}