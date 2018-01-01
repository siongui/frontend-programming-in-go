package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Call("addEventListener", "load", func(event *js.Object) {
		s := js.Global.Get("document").Call("querySelector", ".loader")
		s.Get("classList").Call("add", "invisible")
		println("All resources finished loading!")
	})
}
