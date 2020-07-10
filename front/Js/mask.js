// Uma máscara utilizada para campos de porcentagem

$(document).ready(function(){
    $('.percentual').inputmask('999.99%', {reverse: true, numericInput:true, placeholder:"0"});
    $('.data').mask('00/00/0000');
    // $("#percentual").inputmask("999.99%",{reverse: true,numericInput:true, placeholder:"0"});
  });