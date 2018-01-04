package main

import (
	. "github.com/siongui/godom"
)

func GetPosition(elm *Object) (x, y float64) {
	x = elm.GetBoundingClientRect().Left() + Window.PageXOffset()
	y = elm.GetBoundingClientRect().Top() + Window.PageYOffset()
	return
}

func main() {
	b1 := Document.QuerySelector("#btn1")
	b2 := Document.QuerySelector("#btn2")
	b3 := Document.QuerySelector("#btn3")

	println(GetPosition(b1))
	println(GetPosition(b2))
	println(GetPosition(b3))
}
