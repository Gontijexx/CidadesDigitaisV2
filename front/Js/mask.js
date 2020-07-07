$(document).ready(function(){
    $('.data').mask('00/00/0000');
    $('.cep').mask('00000-000');
    $('.numero').mask('00000-0000');
    $('.numero_ddd').mask('(00) 00000-0000');
    $('.cpf').mask('000.000.000-00', {reverse: true});
    $('.cnpj').mask('00.000.000/0000-00', {reverse: true});
    $('.preco').mask('R$ 000.000.000.000.000,00', {reverse: true});
});