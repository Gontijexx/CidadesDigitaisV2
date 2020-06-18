const masks = {
    porcentagem(value) {
        console.log(value)        
        return value
            .replace(/\D/g, '')
            .replace(/(\d+?$)/, '$1%')
            .replace(/(%\d)\d+?$/, '$1')
    },
    numberOnly(value) {
        console.log(value)
        return value
        .replace(/\D/g,'')
    }

    
}



document.querySelectorAll('input').forEach(($input) => {
    const field = $input.dataset.js
    console.log(field)

    $input.addEventListener('input', (e) => {
        e.target.value = masks[field](e.target.value)
        console.log(e.target.value)
    }, false)
})