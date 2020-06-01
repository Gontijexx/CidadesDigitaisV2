const masks = {
    reajusteMask(value){
        console.log(value)
        VMasker.toPattern(value, "999.999.999-99");
        return value
    }
}

document.querySelectorAll('input').forEach(($input) => {
    const field = $input.dataset.js
        
    $input.addEventListener('input', (e) => {
        e.target.value =  masks[field](e.target.value)
    }, false)
})