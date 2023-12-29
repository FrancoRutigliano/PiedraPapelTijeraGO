function choose(x) {
    fetch("/play?c=" + x)
    .then(Response => Response.json())
    .then(data => {
        console.log(data)
    })
}

