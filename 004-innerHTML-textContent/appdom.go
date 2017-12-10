package main

import (
	. "github.com/siongui/godom"
)

func main() {
	f := Document.QuerySelector("#foo")
	f.SetInnerHTML("<strong>Hello World</strong>")
	println(f.InnerHTML())

	f.SetTextContent("Hello World2")

	println(f.TextContent())
}
