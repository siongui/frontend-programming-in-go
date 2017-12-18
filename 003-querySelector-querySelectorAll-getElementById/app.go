package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// querySelector
	d := js.Global.Get("document").Call("querySelector", "#foo")
	println(d)

	// querySelectorAll
	nodeList := js.Global.Get("document").Call("querySelectorAll", "span")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		span := nodeList.Call("item", i)
		println(span)
	}

	// getElementById
	d2 := js.Global.Get("document").Call("getElementById", "foo")
	println(d2)
}
