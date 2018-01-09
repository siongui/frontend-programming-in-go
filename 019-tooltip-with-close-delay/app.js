// indicate if the mouse cursor is in the tooltip
var isCursorInTooltip = false;

// when cursor leaves the text, the delay time to close the tooltip if the
// cursor is not in the tooltip. (milisecond)
var delay = 250;

// create and append invisible tooltip to DOM tree
var tooltip = document.createElement("div");
tooltip.classList.add("tooltip");
tooltip.classList.add("invisible");
tooltip.addEventListener("mouseenter", function() {
  isCursorInTooltip = true;
});
tooltip.addEventListener("mouseleave", function() {
  isCursorInTooltip = false;
  tooltip.classList.add("invisible");
});
document.querySelector("body").appendChild(tooltip);

// event handler for mouseenter event of elements with data-descr attribute
function ShowTooltip(e) {
  var elm = e.target;
  tooltip.textContent = elm.dataset.descr;
  tooltip.style.left = elm.getBoundingClientRect().left + window.pageXOffset + 'px';
  tooltip.style.top = (elm.getBoundingClientRect().top + window.pageYOffset + elm.offsetHeight + 5) + 'px';
  tooltip.classList.remove("invisible");
}

// event handler for mouseleave event of elements with data-descr attribute
function HideTooltip(e) {
  setTimeout(function() {
    if (!isCursorInTooltip) {
      tooltip.classList.add("invisible");
    }
  }, delay);
}

// select all elements with data-descr attribute and attach mouseenter and mouseleave event handler
var els = document.querySelectorAll("*[data-descr]");
for (var i = 0; i < els.length; ++i) {
  var el = els[i];
  el.addEventListener("mouseenter", ShowTooltip);
  el.addEventListener("mouseleave", HideTooltip);
}
