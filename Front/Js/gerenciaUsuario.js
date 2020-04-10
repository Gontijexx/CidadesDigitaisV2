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


//o json usado para mandar as informações pelo fetch
let info = {
  "nome": "",
  "email": "",
  "status": "",
  "login": "",
};

//pega os valores corretos das variaveis
let meuCodigo = localStorage.getItem("cod_usuario");
let meuNome = localStorage.getItem("nome");
let meuEmail = localStorage.getItem("email");
let meuLogin = localStorage.getItem("login");
let meuStatus = localStorage.getItem("status");

// inserindo os valores no html
document.getElementById("nome").value = meuNome;
document.getElementById("email").value = meuEmail;
document.getElementById("login").value = meuLogin;
document.getElementById("status").value = meuStatus;

function enviar() {

  let a = document.getElementById("nome");
  info.nome = a.value;
  let b = document.getElementById("email");
  info.email =b.value;
  let c = document.getElementById("login");
  info.login = c.value;
  let d = document.getElementById("status");
  info.status =d.value;

  //transforma as informações do token em json
  let corpo = JSON.stringify(info);

  //função fetch para mandar
  fetch('http://localhost:8080/read/usuario/', {
    method: 'POST',
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
      window.location.replace("./Usuario.html");
    } else {
      erros(response.status);
    }
  });
}