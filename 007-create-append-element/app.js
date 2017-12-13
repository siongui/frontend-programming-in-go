var f = document.querySelector("#foo");

var s = document.createElement("span");
f.appendChild(s);

var t = document.createTextNode("Hello World");
s.appendChild(t);
