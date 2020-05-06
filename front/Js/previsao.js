//cuida de previsao de empenho
let listaPrevisao = [];
 
function previsao() {

    document.getElementById("editar").innerHTML = (`<br>`);
    document.getElementById("editar2").innerHTML = (`<br>`);
  
    //função fetch para chamar os itens de previsão da tabela
    fetch(servidor + 'read/previsaoempenho', {
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
          <tr>
          <th style="width:15%" scope="col">Código de Previsão de Empenho</th>
          <th style="width:40%" scope="col">Natureza da despesa</th>
          <th style="width:10%" scope="col">Tipo</th>
          <th style="width:20%" scope="col">Data</th>
          <th style="width:15%" scope="col">Ano de Referência</th>
          </tr>
          </thead>`);
          tabela += (`<tbody>`);

          let j = 0;
          for (let i = 0; i < json.length; i++) {
            console.log(json[i].cod_lote)
            if (json[i].cod_lote == meuLote) {
              listaPrevisao[j] = json[i];
              j++;
            }
          }
          
          for (i = 0; i < listaPrevisao.length; i++) {
            //captura itens para tabela
            tabela += (`<tr>`);
            tabela += (`<td>`);
            tabela += listaPrevisao[i]["cod_previsao_empenho"];
            tabela += (`</td><td>`);
            tabela += listaPrevisao[i]["cod_natureza_despesa"];
            tabela += (`</td><td>`);
            tabela += listaPrevisao[i]["tipo"];
            tabela += (`</td><td>`);
            let data1 = new Date(listaPrevisao[i]["data"]);
            let dataFinal1 = String(data1.getDate()).padStart(2, '0') + "/" + String(data1.getMonth() + 1).padStart(2, '0') + "/" + String(data1.getFullYear()).padStart(4, '0');
            tabela += dataFinal1;
            tabela += (`</td><td>`);
            tabela += listaPrevisao[i]["ano_referencia"];
            tabela += (`</td>`);
            tabela += (`</tr>`);
          }
          tabela += (`</tbody>`);
          document.getElementById("tabela").innerHTML = tabela;
        });
      } else {
        erros(response.status);
      }
    });
}