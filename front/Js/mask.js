// Uma máscara utilizada para campos de porcentagem
document.addEventListener("DOMContentLoaded", function(){
    let els = document.querySelectorAll(".pecent"),
    calc = (i) => {
       let ttl = 0,
       thisval = i.target.value;
 
       for(let x=0; x<els.length; x++){
          ttl += Number(els[x].value.replace("%","").replace(",","."));
          if(ttl > 999){
             i.target.value = '';
          }
       }
       
       if(thisval.indexOf("%") == -1) i.target.value += "%";
 
       i.target.selectionStart = i.target.selectionEnd = thisval.replace("%","").length;
    }
 
    for(let x=0; x<els.length; x++){
       els[x].addEventListener("input", calc);
    }
 });

$('.pecent').mask('P', {
    translation: {
        'P': {
            pattern: /[\d\.,]/,
            recursive: true
        }
    },
    onKeyPress: function(val, e, field, options) {
        var old_value = $(field).data('oldValue') || '';

        val = val.trim();
        val = val.replace(',', '.');
        val = val.length > 0 ? val : '0';

        // Transformando múltiplos pontos em um único ponto
        val = val.replace(/[\.]+/, '.');

        // Verificando se o valor contém mais de uma ocorrência de ponto
        var dot_occurrences = (val.match(/\./g) || []).length > 1;

        // Verificando se o valor está de acordo com a sintaxe do float
        var is_float = /[-+]?[\d]*\.?[\d]+/.test(val);

        if (dot_occurrences || !is_float) {
            val = old_value;
        }

        // Força o valor a ficar no intervalo de 0 à 999
        val = parseFloat(val) >= 999 ? '999' : val;
        val = parseFloat(val) <  0   ? '0'   : val;

        $(field)
            .val(val)
            .data('oldValue', val);
    }
});