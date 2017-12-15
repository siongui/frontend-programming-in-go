package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")

	// get value
	println(f.Get("dataset").Get("demoValue").String())

	// set value
	f.Get("dataset").Set("demoValue", "world hello")
}
