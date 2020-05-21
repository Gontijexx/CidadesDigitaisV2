//Fazer Tabela
let listaModulo = [];

//pega o usuario logado
let userLogado = localStorage.getItem("logado");

//organizar os modulos
let userCriado,
  userTotal = [],
  modulo = [],
  valorModulo = [];

//Fazer Usuário
let info = {
  "nome": "",
  "email": "",
  "login": "",
  "senha": "",
};

//sistema de paginação
let contador = 0;
let porPagina = 5;
let totalPaginas = (userTotal.length + (porPagina - 1)) / porPagina;

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
  fetch(servidor + 'read/usuario', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {
    //tratamento dos erros
    if (response.status == 200) {
      let j = 0;
      //console.log(response.statusText);

      //pegar o json que possui a tabela
      response.json().then(function (json) {

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
            <tr>
            <th scope="col">Código</th>
            <th scope="col">Nome</th>
            <th scope="col">E-mail</th>
            <th scope="col">Login</th>
            <th scope="col">Status</th>
            <th scope="col">Opções</th>
            </tr>
            </thead>`);
        tabela += (`<tbody>`);

        // função para checar quem está ativo ou inativo
        let selecao = document.getElementById("status").value;
        if (selecao == 1) {

          //para quando o status for inativo
          let j = 0;
          let jsonDeStatus = [];
          for (let i = 0; i < json.length; i++) {
            if (json[i]["status"] == 1) {
              jsonDeStatus[j] = json[i];
              j++;
            }
          }
          for (let i = comeco; i < fim && i < jsonDeStatus.length; i++) {
            userTotal[i] = jsonDeStatus[i];
            tabela += (`<tr> <td>`);
            tabela += jsonDeStatus[i]["cod_usuario"];
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["nome"]
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["email"]
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["login"]
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["status"]
            tabela += (`</td> <td> 
                  <span class="d-flex">
                  <button onclick="editarUsuario(` + i + `)" class="btn btn-success">
                  <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
                  </button>
                  </span> </td> </tr>`);
            if (jsonDeStatus[i]["login"] == userLogado) {
              localStorage.setItem("codigoLogado", jsonDeStatus[i]["cod_usuario"]);
            }
          }

        } else if (selecao == 2) {

          //para quando o status for inativo
          let j = 0;
          let jsonDeStatus = [];
          for (let i = 0; i < json.length; i++) {
            if (json[i]["status"] == 0) {
              jsonDeStatus[j] = json[i];
              j++;
            }
          }

          for (let i = comeco; i < fim && i < jsonDeStatus.length; i++) {
            userTotal[i] = jsonDeStatus[i];
            tabela += (`<tr> <td>`);
            tabela += jsonDeStatus[i]["cod_usuario"];
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["nome"]
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["email"]
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["login"]
            tabela += (`</td> <td>`);
            tabela += jsonDeStatus[i]["status"]
            tabela += (`</td> <td> 
                  <span class="d-flex">
                  <button onclick="editarUsuario(` + i + `)" class="btn btn-success">
                  <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
                  </button>
                  </span> </td> </tr>`);
            if (jsonDeStatus[i]["login"] == userLogado) {
              localStorage.setItem("codigoLogado", jsonDeStatus[i]["cod_usuario"]);
            }
          }

        } else {
          for (let i = comeco; i < fim && i < json.length; i++) {
            userTotal[i] = json[i];
            tabela += (`<tr> <td>`);
            tabela += json[i]["cod_usuario"];
            tabela += (`</td> <td>`);
            tabela += json[i]["nome"]
            tabela += (`</td> <td>`);
            tabela += json[i]["email"]
            tabela += (`</td> <td>`);
            tabela += json[i]["login"]
            tabela += (`</td> <td>`);
            tabela += json[i]["status"]
            tabela += (`</td> <td> 
                  <span class="d-flex">
                  <button onclick="editarUsuario(` + i + `)" class="btn btn-success">
                  <i class="material-icons"data-toggle="tooltip" title="Edit">&#xE254;</i>
                  </button>
                  </span> </td> </tr>`);
            if (json[i]["login"] == userLogado) {
              localStorage.setItem("codigoLogado", json[i]["cod_usuario"]);
            }
          }
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;

        totalPaginas = json.length / porPagina;

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
          if (totalPaginas < 10) {
            for (i = 0; i < totalPaginas; i++) {
              if (contador == i) {
                paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link btn active">` + (i + 1) + `</a></li>`;
              } else {
                paginas += `<li class="page-item" id="page` + i + `"><a href="#" onclick="pagina(` + i + `)" class="page-link">` + (i + 1) + `</a></li>`;
              }
            }
            //caso sejam mais de 10
          } else {
            //mostrar apenas inicio e fim
            if (contador == 0) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link btn active">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            //opções do começo
            else if (contador == 1) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link btn active">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link">4</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == 2) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link btn active">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link">4</a></li>`;
              paginas += `<li class="page-item" id="page4"><a href="#" onclick="pagina(4)" class="page-link">5</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == 3) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li class="page-item" id="page1"><a href="#" onclick="pagina(1)" class="page-link">2</a></li>`;
              paginas += `<li class="page-item" id="page2"><a href="#" onclick="pagina(2)" class="page-link">3</a></li>`;
              paginas += `<li class="page-item" id="page3"><a href="#" onclick="pagina(3)" class="page-link btn active">4</a></li>`;
              paginas += `<li class="page-item" id="page4"><a href="#" onclick="pagina(4)" class="page-link">5</a></li>`;
              paginas += `<li class="page-item" id="page5"><a href="#" onclick="pagina(5)" class="page-link">6</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
            //opções no final
            else if (contador == Math.floor(totalPaginas - 3)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 5) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 5) + `)" class="page-link">` + Math.floor(totalPaginas - 4) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 4) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 4) + `)" class="page-link">` + Math.floor(totalPaginas - 3) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 3) + `)" class="page-link btn active">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == Math.floor(totalPaginas - 2)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 4) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 4) + `)" class="page-link">` + Math.floor(totalPaginas - 3) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 3) + `)" class="page-link">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link btn active">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == Math.floor(totalPaginas - 1)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 3) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 3) + `)" class="page-link">` + Math.floor(totalPaginas - 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link btn active">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else if (contador == Math.floor(totalPaginas)) {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 2) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 2) + `)" class="page-link">` + Math.floor(totalPaginas - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas - 1) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas - 1) + `)" class="page-link">` + Math.floor(totalPaginas) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link btn active">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            } else {
              paginas += `<li class="page-item" id="page0"><a href="#" onclick="pagina(0)" class="page-link">1</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador - 2) + `"><a href="#" onclick="pagina(` + (contador - 2) + `)" class="page-link">` + (contador - 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador - 1) + `"><a href="#" onclick="pagina(` + (contador - 1) + `)" class="page-link">` + contador + `</a></li>`;
              paginas += `<li class="page-item" id="page` + contador + `"><a href="#" onclick="pagina(` + contador + `)" class="page-link btn active">` + (contador + 1) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador + 1) + `"><a href="#" onclick="pagina(` + (contador + 1) + `)" class="page-link">` + (contador + 2) + `</a></li>`;
              paginas += `<li class="page-item" id="page` + (contador + 2) + `"><a href="#" onclick="pagina(` + (contador + 2) + `)" class="page-link">` + (contador + 3) + `</a></li>`;
              paginas += `<li><a>...</a></li>`;
              paginas += `<li class="page-item" id="page` + Math.floor(totalPaginas) + `"><a href="#" onclick="pagina(` + Math.floor(totalPaginas) + `)" class="page-link">` + Math.floor(totalPaginas + 1) + `</a></li>`;
            }
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
        if (fim < json.length) {
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
  this.paginacao();

  //Fazer Tabela para Modulos

  //função fetch para mandar itens do modulo
  fetch(servidor + 'read/modulo', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //tratamento dos erros
    if (response.status == 200 || response.status == 201) {
      //console.log(response.statusText);

      //pegar o json que possui a tabela
      response.json().then(function (json) {
        //usado para listar os modulos usados na criação do usuario
        listaModulo = json;
        let tabelaMod = (`<thead style="background: #4b5366; color:white; font-size:15px">
                <tr>
                <th> <span class="custom-checkbox">
                <input type="checkbox" id="selectAll" >
                <label for="selectAll"></label>
                </span></th>
                <th scope="col">Cód. Módulo</th>
                <th scope="col">Módulo</th>
                <th scope="col">Sub. Módulo</th>
                <th scope="col">Ação</th>
                </tr>
                </thead>`);
        tabelaMod += (`<tbody>`);

        for (let i = 0; i < json.length; i++) {
          tabelaMod += (`<tr> <td>
                <span class="custom-checkbox">
                <input class="checking" onclick="modulos(` + i + `)" type="checkbox" id="checkbox` + i + `" name="options[]" value="` + json[i]["cod_modulo"] + `">
                <label for="checkbox` + i + `"></label>
                </span>
                </td>`);
          tabelaMod += (`<td>`);
          tabelaMod += json[i]["cod_modulo"];
          tabelaMod += (`</td> <td>`);
          tabelaMod += json[i]["categoria_1"]
          tabelaMod += (`</td> <td>`);
          tabelaMod += json[i]["categoria_2"]
          tabelaMod += (`</td> <td>`);
          tabelaMod += json[i]["categoria_3"]
          tabelaMod += (`</td> </tr>`);
        }

        tabelaMod += (`</tbody>`);
        document.getElementById("tabelaMod").innerHTML = tabelaMod;

        $(document).ready(function () {
          // Select/Deselect checkboxes
          let checkbox = $('table tbody input[type="checkbox"]');
          $("#selectAll").click(function () {
            if (this.checked) {

              checkbox.each(function () {
                this.checked = true;
              });

              for (i = 0; i < json.length; i++) {
                let mods = [];
                mods[i] = document.getElementById("checkbox" + i);
                valorModulo[i] = mods[i].value;
              }
            } else {

              checkbox.each(function () {
                this.checked = false;
              });

              for (i = 0; i < json.length; i++) {
                let mods = [];
                mods[i] = document.getElementById("checkbox" + i);
                valorModulo[i] = null;
              }
            }
          });
          checkbox.click(function () {
            if (!this.checked) {
              $("#selectAll").prop("checked", false);
            }
          });
        });

      });

    } else {
      erros(response.status);
    }
  });
}

function editarUsuario(valor) {
  localStorage.setItem("cod_usuario", userTotal[valor].cod_usuario);
  localStorage.setItem("nome", userTotal[valor].nome);
  localStorage.setItem("email", userTotal[valor].email);
  localStorage.setItem("login", userTotal[valor].login);
  localStorage.setItem("status", userTotal[valor].status);
  localStorage.setItem("senha", userTotal[valor].senha);
  window.location.href = "./gerenciaUsuario.html";
}

//pega os valores de modulo dos checkboxes e coloca na estrutura valorModulo
function modulos(numCod) {
  let mods = [];
  mods[numCod] = document.getElementById("checkbox" + numCod);
  if (mods[numCod].checked) {
    valorModulo[numCod] = mods[numCod].value;
  } else {
    valorModulo[numCod] = null;
  }
}

function enviar() {

  let a = document.getElementById("nome");
  info.nome = a.value;
  let b = document.getElementById("email");
  info.email = b.value;
  let c = document.getElementById("login");
  info.login = c.value;
  let d = document.getElementById("senha");
  info.senha = d.value;

  //transforma as informações do token em json
  let corpo = JSON.stringify(info);

  //função fetch para mandar
  fetch(servidor + 'read/usuario/createuser', {
    method: 'POST',
    body: corpo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar o status do pedido
    //console.log(response);

    //tratamento dos erros
    if (response.status == 201) {
      //alert("Usuário criado com sucesso");
      response.json().then(function (json) {
        userCriado = json.cod_usuario;
      });
    } else {
      erros(response.status);
      window.location.reload();
    }
  });
}

function enviarModulo() {

  let i, j = 0;
  for (i = 0; i < listaModulo.length; i++) {
    if (valorModulo[i] != null) {
      modulo[j] = {
        "cod_usuario": parseFloat(userCriado),
        "cod_modulo": parseFloat(valorModulo[i])
      }
      j++;
    }
  }

  //transforma as informações do token em json
  let infoModulo = JSON.stringify(modulo);

  //função fetch para mandar
  fetch(servidor + 'read/usuario/' + 1 + '/modulo', {
    method: 'POST',
    body: infoModulo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar o status do pedido
    console.log(response);

    //tratamento dos erros
    if (response.status == 200 || response.status == 201) {
      alert("Módulos inseridos com sucesso");
      location.reload();
    } else {
      //erros(response.status);
    }
  });
}