package main

import (
	. "github.com/siongui/godom"
)

func main() {
	f := Document.QuerySelector("#foo")

	// set the color of element
	f.Style().SetColor("red")

	// get the color of element
	println(f.Style().Color())
}
