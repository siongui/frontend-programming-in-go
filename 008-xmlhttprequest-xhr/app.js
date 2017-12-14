var URL = "https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json";

function GetJSON() {
  console.log(this.responseText);
}

var req = new XMLHttpRequest();
req.addEventListener("load", GetJSON);
req.open("GET", URL);
req.send();
