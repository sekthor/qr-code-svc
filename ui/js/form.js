
var baseURL = "http://localhost:8080/qr"

function generate(type) {
    imgurl = baseURL + "/" + type + "/" + document.getElementById(type).value
    document.getElementById("qr").src = imgurl
}
