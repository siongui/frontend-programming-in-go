package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")

	// set the color of element
	f.Get("style").Set("color", "red")

	// get the color of element
	println(f.Get("style").Get("color").String())
}
