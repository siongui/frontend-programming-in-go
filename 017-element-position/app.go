package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func GetPosition(elm *js.Object) (x, y float64) {
	x = elm.Call("getBoundingClientRect").Get("left").Float() +
		js.Global.Get("pageXOffset").Float()
	y = elm.Call("getBoundingClientRect").Get("top").Float() +
		js.Global.Get("pageYOffset").Float()
	return
}

func main() {
	b1 := js.Global.Get("document").Call("querySelector", "#btn1")
	b2 := js.Global.Get("document").Call("querySelector", "#btn2")
	b3 := js.Global.Get("document").Call("querySelector", "#btn3")

	println(GetPosition(b1))
	println(GetPosition(b2))
	println(GetPosition(b3))
}
