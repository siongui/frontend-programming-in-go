package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")

	f.Call("addEventListener", "click", func(event *js.Object) {
		f.Set("textContent", "I am clicked")
	})
}
