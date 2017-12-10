package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")
	f.Set("innerHTML", "Hello World")
}
