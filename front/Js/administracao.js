//pega o token do login
let meuToken = localStorage.getItem("token");

//guarda info pra ser usado assim depois
let info = {};

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
    alert("Erro: Adição já existente.");
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



//informaçãoes para assunto
function addAssunto() {

  let formulario = (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  document.getElementById("modalAdicao").innerHTML = formulario;

  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioAssunto()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioAssunto() {
  info = {
    "descricao": document.getElementById("descricao").value,
  };
  mandar("assunto");
}

function visAssunto() {
  visualizar("assunto", [`<th style="width:20%" scope="col">Código do Assunto</th>
  <th style="width:80%" scope="col">Descrição</th>`, "cod_assunto", "descricao"]);
}



//informaçãoes para categoria
function addCategoria() {

  let formulario = (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  document.getElementById("modalAdicao").innerHTML = formulario;
  
  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioCategoria()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioCategoria() {
  info = {
    "descricao": document.getElementById("descricao").value,
  };
  mandar("categoria");
}

function visCategoria() {
  visualizar("categoria", [`<th style="width:20%" scope="col">Código da Categoria</th>
  <th style="width:80%" scope="col">Descrição</th>`, "cod_categoria", "descricao"]);
}



//informaçãoes para classe de empenho
function addClasseEmpenho() {

  let formulario = (`<label for="cod_classe_empenho">Código da Classe</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_classe_empenho" id="cod_classe_empenho"></input>`);
  formulario += (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  document.getElementById("modalAdicao").innerHTML = formulario;

  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioClasseEmpenho()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioClasseEmpenho() {
  info = {
    "cod_classe_empenho" : parseInt(document.getElementById("cod_classe_empenho").value),
    "descricao" : document.getElementById("descricao").value,
  };
  mandar("classeempenho");
}

function visClasseEmpenho() {
  visualizar("classeempenho", [`<th style="width:20%" scope="col">Código de Classe de Empenho</th>
  <th style="width:80%" scope="col">Descrição</th>`, "cod_classe_empenho", "descricao"]);
}



//informaçãoes para etapas
function addEtapa() {

  let formulario = (`<label for="descricao">Descrição</label>`);
  formulario +=(`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  formulario += (`<label for="duracao">Duração</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="duracao" id="duracao"></input>`);
  formulario += (`<label for="depende">Depende</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="depende" id="depende"></input>`);
  formulario += (`<label for="delay">Delay</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="delay" id="delay"></input>`);
  formulario += (`<label for="setor_resp">Setor Responsável</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="setor_resp" id="setor_resp" maxlength="45"></input>`);
  document.getElementById("modalAdicao").innerHTML = formulario;

  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioEtapa()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioEtapa() {
  info = {
    "descricao": document.getElementById("descricao").value,
    "duracao": parseInt(document.getElementById("duracao").value),
    "depende": parseInt(document.getElementById("depende").value),
    "delay": parseInt(document.getElementById("delay").value),
    "setor_resp": document.getElementById("setor_resp").value,
  };
  mandar("etapa");
}

function visEtapa() {
  visualizar("etapa", [`<th style="width:10%" scope="col">Código de Etapas</th>
  <th style="width:40%" scope="col">Descrição</th>
  <th style="width:10%" scope="col">Duração</th>
  <th style="width:10%" scope="col">Depende</th>
  <th style="width:10%" scope="col">Delay</th>
  <th style="width:20%" scope="col">Setor Responsável</th>`, "cod_etapa", "descricao", "duracao", "depende", "delay", "setor_resp"]);
}



//informaçãoes para itens

function selectNatureza(){

}
function selectClasse(){

}

function addItem() {

  let formulario = (`<label for="cod_item">Código do Item</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_item" id="cod_item"></input>`);
  formulario += (`<label for="cod_tipo_item">Código do Tipo de Item</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_tipo_item" id="cod_tipo_item"></input>`);
  formulario += (`<label for="cod_natureza_despeza">Código da Natureza</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_natureza_despeza" id="cod_natureza_despeza" maxlength="45"></input>`);
  formulario += (`<label for="cod_classe_empenho">Código da Classe</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_classe_empenho" id="cod_classe_empenho" maxlength="45"></input>`);
  formulario += (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="200"></textarea>`);
  formulario += (`<label for="unidade">Unidade</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="unidade" id="unidade" maxlength="45"></input>`);

  document.getElementById("modalAdicao").innerHTML = formulario;
  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioItem()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;

  selectNatureza();
  selectClasse();
}

function envioItem() {
  info = {
    "cod_item": parseInt(document.getElementById("cod_item").value),
    "categoria_1": document.getElementById("categoria_1").value,
    "categoria_2": document.getElementById("categoria_2").value,
    "descricao": document.getElementById("descricao").value,
    "categoria_3": document.getElementById("categoria_3").value,
  };
  mandar("itens");
}

function visItem() {
  visualizar("itens", [`<th style="width:20%" scope="col">Código do Item</th>
  <th style="width:10%" scope="col">Código do Tipo de Item</th>
  <th style="width:10%" scope="col">Código da Natureza</th>
  <th style="width:10%" scope="col">Código da Classe</th>
  <th style="width:50%" scope="col">Descrição</th>
  <th style="width:10%" scope="col">Unidade</th>`, "cod_item", "cod_tipo_item", "cod_natureza_despesa", "cod_classe_empenho", "descricao", "unidade"]);
}



//informaçãoes para natreza de despesas
function addNaturezaDespesa() {

  let formulario = (`<label for="cod_natureza_despesa">Código da Natureza</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_natureza_despesa" id="cod_natureza_despesa"></input>`);
  formulario += (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  document.getElementById("modalAdicao").innerHTML = formulario;

  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioNaturezaDespesa()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioNaturezaDespesa() {
  info = {
    "cod_natureza_despesa": parseInt(document.getElementById("cod_natureza_despesa").value),
    "descricao": document.getElementById("descricao").value,
  };
  mandar("naturezadespesa");
}

function visNaturezaDespesa() {
  visualizar("naturezadespesa", [`<th style="width:20%" scope="col">Código da Natureza da Despesa</th>
  <th style="width:80%" scope="col">Descrição</th>`, "cod_natureza_despesa", "descricao"]);
}


//PRECISO CHECAR SE É PARA SER ESSE O FETCH
//informaçãoes para modulo
function addModulo() {

  let formulario = (`<label for="cod_modulo">Código do Módulo</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="cod_modulo" id="cod_modulo"></input>`);
  formulario += (`<label for="categoria_1">Categoria 1</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="categoria_1" id="categoria_1" maxlength="45"></input>`);
  formulario += (`<label for="categoria_2">Categoria 2</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="categoria_2" id="categoria_2" maxlength="45"></input>`);
  formulario += (`<label for="categoria_3">Categoria 3</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="categoria_3" id="categoria_3" maxlength="45"></input>`);
  formulario += (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="200"></textarea>`);
  
  document.getElementById("modalAdicao").innerHTML = formulario;
  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioModulo()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioModulo() {
  info = {
    "cod_modulo": parseInt(document.getElementById("cod_modulo").value),
    "categoria_1": document.getElementById("categoria_1").value,
    "categoria_2": document.getElementById("categoria_2").value,
    "categoria_3": document.getElementById("categoria_3").value,
    "descricao": document.getElementById("descricao").value,
  };
  mandar("modulo");
}

function visModulo() {
  visualizar("modulo", [`<th style="width:20%" scope="col">Código de Modulo</th>
  <th style="width:10%" scope="col">Categoria 1</th>
  <th style="width:10%" scope="col">Categoria 2</th>
  <th style="width:10%" scope="col">Categoria 3</th>
  <th style="width:50%" scope="col">Descrição</th>`, "cod_modulo", "categoria_1", "categoria_2", "categoria_3", "descricao"]);
}



//informaçãoes para municipio
function addMunicipio() {

  let formulario = (`<label for="descricao">Descrição</label>`);
  formulario +=(`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  formulario += (`<label for="duracao">Duração</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="duracao" id="duracao"></input>`);
  formulario += (`<label for="depende">Depende</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="depende" id="depende"></input>`);
  formulario += (`<label for="delay">Delay</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="delay" id="delay"></input>`);
  formulario += (`<label for="setor_resp">Setor Responsável</label>`);
  formulario += (`<input class="multisteps-form__input form-control" name="setor_resp" id="setor_resp" maxlength="45"></input>`);
  document.getElementById("modalAdicao").innerHTML = formulario;

  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioMunicipio()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioMunicipio() {
  info = {
    "descricao": document.getElementById("descricao").value,
    "duracao": parseInt(document.getElementById("duracao").value),
    "depende": parseInt(document.getElementById("depende").value),
    "delay": parseInt(document.getElementById("delay").value),
    "setor_resp": document.getElementById("setor_resp").value,
  };
  mandar("municipio");
}

function visMunicipio() {
  visualizar("municipio", [`<th style="width:10%" scope="col">Código do Municipio</th>
  <th style="width:40%" scope="col">Descrição</th>
  <th style="width:10%" scope="col">Duração</th>
  <th style="width:10%" scope="col">Depende</th>
  <th style="width:10%" scope="col">Delay</th>
  <th style="width:20%" scope="col">Setor Responsável</th>`, "cod_ibge", "descricao", "duracao", "depende", "delay", "setor_resp"]);
}



//informaçãoes para Tipologia
function addTipologia() {

  let formulario = (`<label for="descricao">Descrição</label>`);
  formulario += (`<textarea class="multisteps-form__input form-control" name="descricao" id="descricao" maxlength="1000"></textarea>`);
  document.getElementById("modalAdicao").innerHTML = formulario;

  let botao = (`<button class="btn btn-primary multi-button ml-auto js-btn-next" type="button" onclick="envioTipologia()" title="Next">Adicionar</button>`);
  document.getElementById("botaoEnvio").innerHTML = botao;
}

function envioTipologia() {
  info = {
    "descricao": document.getElementById("descricao").value,
  };
  mandar("tipologia");
}

function visTipologia() {
  visualizar("tipologia", [`<th style="width:20%" scope="col">Código de Tipologia</th>
  <th style="width:80%" scope="col">Descrição</th>`, "cod_tipologia", "descricao"]);
}



function mandar(caminho) {
  //transforma as informações do token em json
  let corpo = JSON.stringify(info);
  //função fetch para mandar
  fetch('http://localhost:8080/read/' + caminho, {
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
      window.location.replace("./administracao.html");
    } else {
      erros(response.status);
    }
  });
}



function visualizar(caminho, estrutura) {
  //função fetch para chamar os itens de previsão da tabela
  fetch('http://localhost:8080/read/' + caminho, {
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

        let tabela = (`<thead style="background: #4b5366; color:white; font-size:15px">
        <tr>` + estrutura[0] + `</tr>
        </thead>`);
        tabela += (`<tbody>`);

        console.log(json);

        for (i = 0; i < json.length; i++) {
          //captura itens para tabela
          tabela += (`<tr>`);
          for (j = 1; j < estrutura.length; j++) {
            tabela += (`<td>`);
            tabela += json[i][estrutura[j]];
            tabela += (`</td>`);
          }
          tabela += (`</tr>`);
        }
        tabela += (`</tbody>`);
        document.getElementById("tabela").innerHTML = tabela;
      });
    } else {
      //erros(response.status);
    }
  });
}