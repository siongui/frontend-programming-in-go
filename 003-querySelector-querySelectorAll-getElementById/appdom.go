package main

import (
	. "github.com/siongui/godom"
)

func main() {
	// querySelector
	d := Document.QuerySelector("#foo")
	println(d)

	// querySelectorAll
	nodeList := Document.QuerySelectorAll("span")
	for _, span := range nodeList {
		println(span)
	}

	// getElementById
	d2 := Document.GetElementById("foo")
	println(d2)
}
