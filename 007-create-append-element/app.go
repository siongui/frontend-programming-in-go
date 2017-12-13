package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")

	// create element
	s := js.Global.Get("document").Call("createElement", "span")
	// append element
	f.Call("appendChild", s)

	// create text node
	t := js.Global.Get("document").Call("createTextNode", "Hello World")
	// append text node
	s.Call("appendChild", t)
}
