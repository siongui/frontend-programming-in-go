// querySelector
var d = document.querySelector("#foo");
console.log(d);

// querySelectorAll
var nodeList = document.querySelectorAll("span");
for (var i = 0; i < nodeList.length; ++i) {
  var span = nodeList[i];
  console.log(span);
}

// getElementById
var d2 = document.getElementById("foo");
console.log(d2);
