function active(num) {
    var btns = document.querySelectorAll(".request-params button")
    var params = document.querySelectorAll(".request-params-expansion > *")
    btns.forEach(function(v,k) {
        if (k === num-1) {
            if (v.classList.contains("active")) {
                v.classList.remove("active")
                params[k].classList.remove("visible")
            } else {
                 v.classList.add("active")
                 params[k].classList.add("visible")
            }
        } else {
            v.classList.remove("active")
            params[k].classList.remove("visible")
        }
        
    })
}