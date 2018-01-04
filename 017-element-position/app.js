function GetPosition(elm) {
  var x = elm.getBoundingClientRect().left + window.pageXOffset;
  var y = elm.getBoundingClientRect().top + window.pageYOffset;
  return [x, y];
}

var b1 = document.querySelector("#btn1");
var b2 = document.querySelector("#btn2");
var b3 = document.querySelector("#btn3");

console.log(GetPosition(b1));
console.log(GetPosition(b2));
console.log(GetPosition(b3));
