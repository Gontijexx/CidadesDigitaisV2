//pega o token do login
let meuToken = localStorage.getItem("token");

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

let info = {
  "cod_usuario": "",
  "nome": "",
  "email": "",
  "login": "",
  "status": "",
  "senha": "",
};

//pega o valor do usuario logado
let codigoLogado = localStorage.getItem("codigoLogado");

//pega os valores corretos das variaveis
let meuCodigo = localStorage.getItem("cod_usuario");
let meuNome = localStorage.getItem("nome");
let meuEmail = localStorage.getItem("email");
let meuLogin = localStorage.getItem("login");
let meuStatus = localStorage.getItem("status");
let meuSenha = localStorage.getItem("senha");


window.onload = function () {
// inserindo os valores no html
document.getElementById("cod_usuario").value = meuCodigo;
document.getElementById("nome").value = meuNome;
document.getElementById("email").value = meuEmail;
document.getElementById("login").value = meuLogin;
document.getElementById("status").value = meuStatus;
document.getElementById("senha").value = meuSenha;
}


function enviar() {

  info.cod_usuario = parseInt(document.getElementById("cod_usuario").value);
  info.nome = document.getElementById("nome").value;
  info.email = document.getElementById("email").value;
  info.login = document.getElementById("login").value;
  info.status = document.getElementById("status").value;
  info.senha = document.getElementById("senha").value;

  //transforma as informações do token em json
  let corpo = JSON.stringify(info);
  console.log(corpo);
  //função fetch para mandar
  fetch('http://localhost:8080/read/usuario/' + parseInt(codigoLogado), {
    method: 'PUT',
    body: corpo,
    headers: {
      'Authorization': 'Bearer ' + meuToken
    },
  }).then(function (response) {

    //checar o status do pedido
    console.log(response.status);

    //tratamento dos erros
    if (response.status == 200 || response.status == 202) {
      response.json().then(function (json) {
        console.log(json);
      });
      window.location.replace("./usuario.html");
    } else {
      erros(response.status);
    }
  });
}