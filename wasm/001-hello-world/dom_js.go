package main

import (
	"syscall/js"
)

type window struct {
	js.Value
}

var Window = &window{js.Global()}
var Document = js.Global().Get("document")

// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (w *window) Alert(s string) {
	w.Call("alert", s)
}
