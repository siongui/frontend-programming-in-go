package main

import (
	. "github.com/siongui/godom"
)

func main() {
	f := Document.QuerySelector("#foo")

	// create element
	s := Document.CreateElement("span")
	// append element
	f.AppendChild(s)

	// create text node
	t := Document.CreateTextNode("Hello World")
	// append text node
	s.AppendChild(t)
}
