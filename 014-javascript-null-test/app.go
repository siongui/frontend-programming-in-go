package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")

	if f == nil {
		println("querySelector #foo returns null")
	} else {
		println("querySelector #foo returns element")
	}
}
