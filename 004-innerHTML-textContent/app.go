package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	f := js.Global.Get("document").Call("querySelector", "#foo")
	f.Set("innerHTML", "<strong>Hello World</strong>")

	println(f.Get("innerHTML").String())

	f.Set("textContent", "Hello World2")

	println(f.Get("textContent").String())
}
