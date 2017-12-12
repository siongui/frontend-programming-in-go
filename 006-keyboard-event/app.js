var Key = {
  LEFT:   37,
  UP:     38,
  RIGHT:  39,
  DOWN:   40
};

function handleKeyboardEvent(evt) {
  var keycode = evt.keyCode;

  var info = document.querySelector("#info");
  switch (keycode) {
    case Key.LEFT:
      info.innerHTML = "LEFT ";
      break;
    case Key.UP:
      info.innerHTML = "UP ";
      break;
    case Key.RIGHT:
      info.innerHTML = "RIGHT ";
      break;
    case Key.DOWN:
      info.innerHTML = "DOWN ";
      break;
  }
}

window.addEventListener("keyup", handleKeyboardEvent, false);
