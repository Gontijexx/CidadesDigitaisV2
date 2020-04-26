//pega o token do login
let meuToken = localStorage.getItem("token");

//Fazer Tabela
let cdTotal = [];
let meuLote;

//pega o JSON de municípios para uso na tabela e para adcionar "CD"s
let cidades = [];
document.getElementById("cod_ibge").disabled = true;

//tratamento de erros
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
    alert("Erro: Cidade Digital já existente.");
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


//Fazer Entidade
let info = {
  "cod_ibge": " ",
  "cod_lote": " ",
  "os_pe": " ",
  "data_pe": " ",
  "os_imp": " ",
  "data_imp": " "
};


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

  //função fetch para mandar
  fetch('http://localhost:8080/read/cd', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {
    //tratamento dos erros
    if (response.status == 200) {
      console.log(response.statusText);
      //pegar o json que possui a tabela
      response.json().then(function (json) {

        cdTotal = json;
        totalPaginas = Math.floor(json.length / porPagina);
        console.log(totalPaginas);
        
        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
            <tr>
            <th scope="col">Código IBGE do Município</th>
            <th scope="col">Município</th>
            <th scope="col">Código Lote</th>
            <th scope="col">O.S. Projeto Executivo</th>
            <th scope="col">Data de Projeto Executivo</th>
            <th scope="col">O.S. Implementação</th>
            <th scope="col">Data de Implementação</th>
            <th scope="col">Opções</th>
            </tr>
            </thead>`);
        tabela += (`<tbody>`);

        for (let i = comeco; i < fim && i < json.length; i++) {
          tabela += (`<tr>`);
          tabela += (`<td>`);
          tabela += json[i]["cod_ibge"];
          tabela += (`</td> <td>`);
          tabela += cidades[i]["nome_municipio"] + " - " + cidades[i]["uf"];
          tabela += (`</td> <td>`);
          tabela += json[i]["cod_lote"]
          tabela += (`</td> <td>`);
          tabela += json[i]["os_pe"]
          tabela += (`</td> <td>`);
          let data1 = new Date(json[i]["data_pe"]);
          let dataf1 = String(data1.getDate()).padStart(2, '0') + '/' + String(data1.getMonth()+1).padStart(2, '0') + '/' + String(data1.getFullYear()).padStart(4, '0');
          tabela += dataf1;
          tabela += (`</td> <td>`);
          tabela += json[i]["os_imp"]
          tabela += (`</td> <td>`);
          let data2 = new Date(json[i]["data_imp"]);
          let dataf2 = String(data2.getDate()).padStart(2, '0') + '/' + String(data2.getMonth()+1).padStart(2, '0') + '/' + String(data2.getFullYear()).padStart(4, '0');
          tabela += dataf2;
          tabela += (`</td> <td> 
                <span class="d-flex">
                <button onclick="editarCd(` + i + `)" class="btn btn-success">
                <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
                </button>
                </span> </td>`);
          tabela += (`</tr>`);
        }
        tabela += (`</tr> </tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

        //mostra quanto do total aparece na tela
        document.getElementById("mostrando").innerHTML = "Mostrando " + (comeco + 1) + " a " + fim + " de " + json.length;
        if (porPagina > json.length - comeco) {
          document.getElementById("mostrando").innerHTML = "Mostrando " + (comeco + 1) + " a " + json.length + " de " + json.length;
        }

        //conta quantas paginas é necessário
        let paginas = `<li id="anterior" class="page-item" ><a href="#" class="page-link" onclick="antes()">Anterior</a></li>`;
        //apenas aciona se precisar de paginação
        if (json.length > porPagina) {
          //caso seja apenas 10 paginas
          if(totalPaginas<10){
            for (i = 0; i < totalPaginas; i++) {
              if(contador==i){
                paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link btn active">` + (i + 1) + `</a></li>`;
              }
              else{
                paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link">` + (i + 1) + `</a></li>`;
              }
            }
            //caso sejam mais de 10
          } else{
            //mostrar apenas inicio e fim
            if(contador==0){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link btn active">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            //opções do começo
            else if(contador==1){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link btn active">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link">4</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            else if(contador==2){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link btn active">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link">4</a></li>`;
              paginas += `<li class="page-item" id="page4"><a href="#" onclick="pagina(4)" class="page-link">5</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            else if(contador==3){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link btn active">4</a></li>`;
              paginas += `<li class="page-item" id="page4"><a href="#" onclick="pagina(4)" class="page-link">5</a></li>`;
              paginas += `<li class="page-item" id="page5"><a href="#" onclick="pagina(5)" class="page-link">6</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            //opções no final
            else if(contador==Math.floor(totalPaginas-3)){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-5) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-5) + `)" class="page-link">` + Math.floor(totalPaginas - 4) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-4) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-4) + `)" class="page-link">` + Math.floor(totalPaginas - 3) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-3) + `)" class="page-link btn active">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            else if(contador==Math.floor(totalPaginas-2)){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-4) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-4) + `)" class="page-link">` + Math.floor(totalPaginas - 3) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-3) + `)" class="page-link">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-2) + `)" class="page-link btn active">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            else if(contador==Math.floor(totalPaginas-1)){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-3) + `)" class="page-link">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-1) + `)" class="page-link btn active">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            else if(contador==Math.floor(totalPaginas)){
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas-1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas-1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link btn active">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            else{
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador-2) + `"><a href="#" onclick="pagina(` + (contador-2) + `)" class="page-link">` + (contador-1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador-1) + `"><a href="#" onclick="pagina(` + (contador-1) + `)" class="page-link">` + contador + `</a></li>`;
              paginas += `<li class="page-item" id="page` + contador + `"><a href="#" onclick="pagina(` + contador + `)" class="page-link btn active">` + (contador+ 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador+1) + `"><a href="#" onclick="pagina(` + (contador+1) + `)" class="page-link">` + (contador+2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador+2) + `"><a href="#" onclick="pagina(` + (contador+2) + `)" class="page-link">` + (contador+3) + `</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` +  Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
          }
        }
        paginas += `<li id="proximo" class="page-item" ><a href="#" class="page-link" onclick="depois()">Próximo</a></li>`;
        document.getElementById("paginacao").innerHTML = paginas;

        //controla quando aparece o botão de antes e depois
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



window.onload = function () {

  fetch('http://localhost:8080/read/municipio', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200) {
      response.json().then(function (json) {
        //pegando valores para usar em municipios
        cidades = json;
        //cria variaveis
        let i, j = 0;
        let x = [],
          valorUF = [],
          valorFinalUF = [];

        //faz a ligação entre variaveis e valores do banco
        for (i = 0; i < json.length; i++) {
          valorUF[i] = json[i].uf;
          if (i!=0 && valorUF[i] != valorUF[i - 1]) {
            valorFinalUF[j] = valorUF[i];
            j++;
          }
        }
        for (i = 0; i < j; i++) {
          x[i] += "<option>" + valorFinalUF[i] + "</option>";
        }
        x.sort();
        document.getElementById("uf").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
    this.paginacao();
  });

  //preenche os cod_lotes
  fetch('http://localhost:8080/read/lote', {
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
        x[0] += "<option value='00000000000000'>Código do Lote</option>";
        for (i = 0; i < json.length; i++) {
          x[i + 1] += "<option>" + json[i].cod_lote + "</option>";
        }
        x.sort();
        document.getElementById("cod_lote").innerHTML = x;
      });
    } else {
      erros(response.status);
    }
  });

}


function enabler() {
  document.getElementById("cod_ibge").disabled = false;
  let uf = document.getElementById("uf");
  let i,j=0, x = [], cidadesFinal=[];
  for (i = 0; i < cidades.length; i++) {
    if (cidades[i].uf == uf.value) {
      cidadesFinal[j]=cidades[i];
      j++;
    }
  }
  for (i = 0; i < cidadesFinal.length; i++) {
    x[i] = "<option value='"+cidadesFinal[i].cod_ibge+"'>" + cidadesFinal[i].nome_municipio + "</option>";
  }
  x.sort();
  document.getElementById("cod_ibge").innerHTML = x;
}



function editarCd(valor) {
  localStorage.setItem("cod_ibge", cdTotal[valor].cod_ibge);
  localStorage.setItem("cod_lote", cdTotal[valor].cod_lote);
  localStorage.setItem("os_pe", cdTotal[valor].os_pe);
  localStorage.setItem("data_pe", cdTotal[valor].data_pe);
  localStorage.setItem("os_imp", cdTotal[valor].os_imp);
  localStorage.setItem("data_imp", cdTotal[valor].data_imp);
  localStorage.setItem("nome_municipio", cidades[valor].nome_municipio);
  localStorage.setItem("uf", cidades[valor].uf);
  window.location.href = "./gerenciaCd.html";
}




function enviar() {

  info.cod_ibge = parseInt(document.getElementById("cod_ibge").value);
  info.cod_lote = parseInt(document.getElementById("cod_lote").value);
  info.os_pe = document.getElementById("os_pe").value;
  info.data_pe = document.getElementById("data_pe").value;
  info.os_imp  = document.getElementById("os_imp").value;
  info.data_imp = document.getElementById("data_imp").value;

  //transforma as informações do token em json
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch('http://localhost:8080/read/cd', {
    method: 'POST',
    body: corpo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar o status do pedido
    //console.log(response);

    //tratamento dos erros
    if (response.status == 200 || response.status == 201 || response.status == 202) {
      response.json().then(function (json) {
        //console.log(json);
      });
      window.location.replace("./cd.html");
    } else {
      erros(response.status);
    }
  });
}