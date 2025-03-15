btn = document.getElementById("play")
btn.onclick = function() {
    const xhttp = new XMLHttpRequest();
    // xhttp.onload = function() {
    //   document.getElementById("demo").innerHTML = this.responseText;
    //   }
    xhttp.open("GET", "/play", true);
    xhttp.send();
}
